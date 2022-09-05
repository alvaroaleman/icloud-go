package icloudgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/net/publicsuffix"
	"k8s.io/apimachinery/pkg/util/sets"
)

type User struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type LoginData struct {
	User        `json:",inline"`
	RememberMe  bool     `json:"rememberMe"`
	TrustTokens []string `json:"trustTokens"`
}

const (
	AuthEndpoint  = "https://idmsa.apple.com/appleauth/auth"
	SetupEndpoint = "https://setup.icloud.com/setup/ws/1"
)

type Session struct {
	client         *retryablehttp.Client `json:"-"`
	ClientID       string                `json:"clientID"`
	AccountCountry string                `json:"accountCountry"`
	SessionID      string                `json:"sessionID"`
	SessionToken   string                `json:"sessionToken"`
	TrustToken     string                `json:"trustToken"`
	SCNT           string                `json:"scnt"`
	Data           AccountLoginResponse  `json:"data"`
}

type requestParameters struct {
	method              string
	url                 string
	body                interface{}
	other               func(r *retryablehttp.Request)
	expectedStatusCodes sets.Int
	into                interface{}
}

func (s *Session) do(ctx context.Context, params requestParameters) error {
	var body []byte
	if params.body != nil {
		var err error
		body, err = json.Marshal(params.body)
		if err != nil {
			return fmt.Errorf("failed to serialize request body: %w", err)
		}
	}
	req, err := retryablehttp.NewRequestWithContext(ctx, params.method, params.url, body)
	if err != nil {
		return fmt.Errorf("failed to construct request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	if strings.HasPrefix(params.url, AuthEndpoint) {
		s.addAuthHeaders(req)
	}

	if params.other != nil {
		params.other(req)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if val := resp.Header.Get("X-Apple-ID-Account-Country"); val != "" {
		s.AccountCountry = val
	}
	if val := resp.Header.Get("X-Apple-ID-Session-Id"); val != "" {
		s.SessionID = val
	}
	if val := resp.Header.Get("X-Apple-Session-Token"); val != "" {
		s.SessionToken = val

	}
	if val := resp.Header.Get("X-Apple-TwoSV-Trust-Token"); val != "" {
		s.TrustToken = val
	}
	if val := resp.Header.Get("scnt"); val != "" {
		s.SCNT = val
	}

	if params.expectedStatusCodes != nil && !params.expectedStatusCodes.Has(resp.StatusCode) {
		bodyRead, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("failed to read body from request with unexpected response code: %v\n", err)
		}
		return fmt.Errorf("unexpected response code: %d, expected %v. Body: %s", resp.StatusCode, params.expectedStatusCodes.List(), string(bodyRead))
	} else if params.expectedStatusCodes == nil {
		fmt.Printf("request for %s returned status code %d\n", req.URL.String(), resp.StatusCode)
	}

	if params.into != nil {
		if err := json.NewDecoder(resp.Body).Decode(params.into); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}

func (s *Session) addAuthHeaders(r *retryablehttp.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Apple-OAuth-Client-Id", "d39ba9916b7251055b22c7f910e2ea796ee65e98b2ddecea8f5dde8d9d1a815d")
	r.Header.Set("X-Apple-OAuth-Client-Type", "firstPartyAuth")
	r.Header.Set("X-Apple-OAuth-Redirect-URI", "https://www.icloud.com")
	r.Header.Set("X-Apple-OAuth-Require-Grant-Code", "true")
	r.Header.Set("X-Apple-OAuth-Response-Mode", "web_message")
	r.Header.Set("X-Apple-OAuth-Response-Type", "code")
	r.Header.Set("X-Apple-OAuth-State", s.ClientID)
	r.Header.Set("X-Apple-Widget-Key", "d39ba9916b7251055b22c7f910e2ea796ee65e98b2ddecea8f5dde8d9d1a815d")
}

func (s *Session) requires2FA() bool {
	return s.Data.DsInfo.HsaVersion == 2 && s.Data.HsaChallengeRequired
}

type twoFactorAuth struct {
	SecurityCode securityCode `json:"securityCode"`
}

type securityCode struct {
	Code string `json:"code"`
}

func (s *Session) validate2FACode(ctx context.Context, code string) error {
	return s.do(ctx, requestParameters{
		method: http.MethodPost,
		url:    AuthEndpoint + "/verify/trusteddevice/securitycode",
		other: func(r *retryablehttp.Request) {
			r.Header.Set("X-Apple-Id-Session-Id", s.SessionID)
			if s.SCNT != "" {
				r.Header.Set("scnt", s.SCNT)
			}
		},
		body:                &twoFactorAuth{securityCode{Code: code}},
		expectedStatusCodes: sets.NewInt(204),
	})
}

func (s *Session) trustSession(ctx context.Context) error {
	if err := s.do(ctx, requestParameters{
		method: http.MethodGet,
		url:    AuthEndpoint + "/2sv/trust",
		other: func(r *retryablehttp.Request) {
			r.Header.Set("X-Apple-Id-Session-Id", s.SessionID)
			if s.SCNT != "" {
				r.Header.Set("scnt", s.SCNT)
			}
		},
	}); err != nil {
		return fmt.Errorf("failed to trust session: %w", err)
	}

	// TODO: Can this request be avoided?
	if err := s.authenticateWithToken(ctx); err != nil {
		return fmt.Errorf("failed to authenticate with token: %w", err)
	}

	return nil
}

func (s *Session) authenticateWithToken(ctx context.Context) error {
	return s.do(ctx, requestParameters{
		method: http.MethodPost,
		body: &authenticationData{
			AccountCountryCode: s.AccountCountry,
			DSWebAuthToken:     s.SessionToken,
			ExtendedLogin:      true,
			TrustToken:         s.TrustToken,
		},
		url: SetupEndpoint + "/accountLogin",
		other: func(r *retryablehttp.Request) {
			r.Header.Set("Origin", "https://www.icloud.com")
			r.Header.Set("Referer", "https://www.icloud.com/")

		},
		expectedStatusCodes: sets.NewInt(200),
		into:                &s.Data,
	})
}

type authenticationData struct {
	AccountCountryCode string `json:"accountCountryCode"`
	DSWebAuthToken     string `json:"dsWebAuthToken"`
	ExtendedLogin      bool   `json:"extendedLogin"`
	TrustToken         string `json:"trustToken"`
}

type CredentialFetcher func() (string, error)

func Login(ctx context.Context, user string, password, twoFactorCode CredentialFetcher) error {
	sessionUUID, err := uuid.DefaultGenerator.NewV1()
	if err != nil {
		return fmt.Errorf("failed to generate session UUID: %w", err)
	}
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return fmt.Errorf("failed to create cookie jar: %w", err)
	}
	client := retryablehttp.NewClient()
	client.HTTPClient.Jar = cookieJar

	s := Session{
		client:   client,
		ClientID: "auth-" + strings.ToLower(sessionUUID.String()),
	}

	defer func() {
		serialized, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			fmt.Printf("failed to serialize session: %v\n", err)
			return
		}
		println(string(serialized))
	}()

	passcode, err := password()
	if err != nil {
		return fmt.Errorf("failed to get password: %w", err)
	}

	params := requestParameters{
		method: http.MethodPost,
		url:    AuthEndpoint + "/signin?isRememberMeEnabled=true",
		body: &LoginData{
			User: User{
				AccountName: user,
				Password:    passcode,
			},
			RememberMe: true,
		},
		// TODO: This is what happens with hsa2 - What happens if that is not enabled?
		expectedStatusCodes: sets.NewInt(409),
		into:                &signInResponse{},
	}
	if err := s.do(ctx, params); err != nil {
		return fmt.Errorf("failed to sign in: %w", err)
	}

	if err := s.authenticateWithToken(ctx); err != nil {
		return fmt.Errorf("failed to authenticate with token: %w", err)
	}

	if s.requires2FA() {
		twoFactorCodeRaw, err := twoFactorCode()
		if err != nil {
			return fmt.Errorf("failed to get 2fa code: %w", err)
		}

		if err := s.validate2FACode(ctx, twoFactorCodeRaw); err != nil {
			return fmt.Errorf("failed to validate 2fa code: %w", err)
		}
	}

	if err := s.trustSession(ctx); err != nil {
		return fmt.Errorf("failed to trust session: %w", err)
	}

	return nil
}

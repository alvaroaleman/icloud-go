package cookiejar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"sync"

	"golang.org/x/net/publicsuffix"
)

type SerializableCookieJar struct {
	http.CookieJar
	lock sync.Mutex
	data map[string][]*http.Cookie
}

func (j *SerializableCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.lock.Lock()
	defer j.lock.Unlock()

	if j.data == nil {
		j.data = map[string][]*http.Cookie{}

	}
	j.data[u.String()] = append(j.data[u.String()], cookies...)
	j.CookieJar.SetCookies(u, cookies)
}

func (j *SerializableCookieJar) Save(dest io.Writer) error {
	j.lock.Lock()
	defer j.lock.Unlock()

	return json.NewEncoder(dest).Encode(j.data)
}

func (j *SerializableCookieJar) Load(src io.Reader) error {
	j.lock.Lock()
	defer j.lock.Unlock()

	if err := json.NewDecoder(src).Decode(&j.data); err != nil {
		return err
	}

	for urlRaw, cookies := range j.data {
		u, err := url.Parse(urlRaw)
		if err != nil {
			return err
		}
		j.CookieJar.SetCookies(u, cookies)
	}

	return nil
}

func New() (*SerializableCookieJar, error) {
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	return &SerializableCookieJar{
		CookieJar: cookieJar,
	}, nil
}

func NewFromPath(path string) (*SerializableCookieJar, error) {
	jar, err := New()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer file.Close()

	if err := jar.Load(file); err != nil {
		return nil, fmt.Errorf("failed to load cookie jar from %s: %w", path, err)
	}

	return jar, nil
}

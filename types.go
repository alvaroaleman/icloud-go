package icloudgo

type AccountLoginResponse struct {
	DsInfo                       DsInfo      `json:"dsInfo"`
	HasMinimumDeviceForPhotosWeb bool        `json:"hasMinimumDeviceForPhotosWeb"`
	ICDPEnabled                  bool        `json:"iCDPEnabled"`
	Webservices                  Webservices `json:"webservices"`
	PcsEnabled                   bool        `json:"pcsEnabled"`
	ConfigBag                    ConfigBag   `json:"configBag"`
	HsaTrustedBrowser            bool        `json:"hsaTrustedBrowser"`
	AppsOrder                    []string    `json:"appsOrder"`
	Version                      int         `json:"version"`
	IsExtendedLogin              bool        `json:"isExtendedLogin"`
	PcsServiceIdentitiesIncluded bool        `json:"pcsServiceIdentitiesIncluded"`
	HsaChallengeRequired         bool        `json:"hsaChallengeRequired"`
	RequestInfo                  RequestInfo `json:"requestInfo"`
	PcsDeleted                   bool        `json:"pcsDeleted"`
	ICloudInfo                   ICloudInfo  `json:"iCloudInfo"`
	Apps                         Apps        `json:"apps"`
}

type AppleIDEntries struct {
	IsPrimary bool   `json:"isPrimary"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type MailFlags struct {
	IsThreadingAvailable  bool `json:"isThreadingAvailable"`
	IsSearchV2Provisioned bool `json:"isSearchV2Provisioned"`
	IsCKMail              bool `json:"isCKMail"`
}

type BeneficiaryInfo struct {
	IsBeneficiary bool `json:"isBeneficiary"`
}

type DsInfo struct {
	LastName                        string           `json:"lastName"`
	ICDPEnabled                     bool             `json:"iCDPEnabled"`
	TantorMigrated                  bool             `json:"tantorMigrated"`
	Dsid                            string           `json:"dsid"`
	HsaEnabled                      bool             `json:"hsaEnabled"`
	IroncadeMigrated                bool             `json:"ironcadeMigrated"`
	Locale                          string           `json:"locale"`
	BrZoneConsolidated              bool             `json:"brZoneConsolidated"`
	IsManagedAppleID                bool             `json:"isManagedAppleID"`
	IsCustomDomainsFeatureAvailable bool             `json:"isCustomDomainsFeatureAvailable"`
	IsHideMyEmailFeatureAvailable   bool             `json:"isHideMyEmailFeatureAvailable"`
	GilliganInvited                 bool             `json:"gilligan-invited"`
	AppleIDAliases                  []interface{}    `json:"appleIdAliases"`
	HsaVersion                      int              `json:"hsaVersion"`
	UbiquityEOLEnabled              bool             `json:"ubiquityEOLEnabled"`
	IsPaidDeveloper                 bool             `json:"isPaidDeveloper"`
	CountryCode                     string           `json:"countryCode"`
	NotificationID                  string           `json:"notificationId"`
	PrimaryEmailVerified            bool             `json:"primaryEmailVerified"`
	ADsID                           string           `json:"aDsID"`
	Locked                          bool             `json:"locked"`
	HasICloudQualifyingDevice       bool             `json:"hasICloudQualifyingDevice"`
	PrimaryEmail                    string           `json:"primaryEmail"`
	AppleIDEntries                  []AppleIDEntries `json:"appleIdEntries"`
	GilliganEnabled                 bool             `json:"gilligan-enabled"`
	FullName                        string           `json:"fullName"`
	MailFlags                       MailFlags        `json:"mailFlags"`
	LanguageCode                    string           `json:"languageCode"`
	AppleID                         string           `json:"appleId"`
	AnalyticsOptInStatus            bool             `json:"analyticsOptInStatus"`
	FirstName                       string           `json:"firstName"`
	ICloudAppleIDAlias              string           `json:"iCloudAppleIdAlias"`
	NotesMigrated                   bool             `json:"notesMigrated"`
	BeneficiaryInfo                 BeneficiaryInfo  `json:"beneficiaryInfo"`
	HasPaymentInfo                  bool             `json:"hasPaymentInfo"`
	PcsDeleted                      bool             `json:"pcsDeleted"`
	AppleIDAlias                    string           `json:"appleIdAlias"`
	BrMigrated                      bool             `json:"brMigrated"`
	StatusCode                      int              `json:"statusCode"`
	FamilyEligible                  bool             `json:"familyEligible"`
}

type WebserviceReminders struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Ckdatabasews struct {
	PcsRequired bool   `json:"pcsRequired"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type Photosupload struct {
	PcsRequired bool   `json:"pcsRequired"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type WebservicePhotos struct {
	PcsRequired bool   `json:"pcsRequired"`
	UploadURL   string `json:"uploadUrl"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type Drivews struct {
	PcsRequired bool   `json:"pcsRequired"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type Uploadimagews struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Schoolwork struct {
}

type Cksharews struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Findme struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Ckdeviceservice struct {
	URL string `json:"url"`
}

type Iworkthumbnailws struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type WebserviceCalendar struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Docws struct {
	PcsRequired bool   `json:"pcsRequired"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type WebserviceSettings struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Premiummailsettings struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Ubiquity struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Streams struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Keyvalue struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Archivews struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Push struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Iwmb struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Iworkexportws struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Sharedlibrary struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Geows struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type ICloudEnv struct {
	ShortID   string `json:"shortId"`
	VipSuffix string `json:"vipSuffix"`
}

type Account struct {
	ICloudEnv ICloudEnv `json:"iCloudEnv"`
	URL       string    `json:"url"`
	Status    string    `json:"status"`
}

type WebserviceContacts struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Webservices struct {
	Reminders           WebserviceReminders `json:"reminders"`
	Ckdatabasews        Ckdatabasews        `json:"ckdatabasews"`
	Photosupload        Photosupload        `json:"photosupload"`
	Photos              WebservicePhotos    `json:"photos"`
	Drivews             Drivews             `json:"drivews"`
	Uploadimagews       Uploadimagews       `json:"uploadimagews"`
	Schoolwork          Schoolwork          `json:"schoolwork"`
	Cksharews           Cksharews           `json:"cksharews"`
	Findme              Findme              `json:"findme"`
	Ckdeviceservice     Ckdeviceservice     `json:"ckdeviceservice"`
	Iworkthumbnailws    Iworkthumbnailws    `json:"iworkthumbnailws"`
	Calendar            WebserviceCalendar  `json:"calendar"`
	Docws               Docws               `json:"docws"`
	Settings            WebserviceSettings  `json:"settings"`
	Premiummailsettings Premiummailsettings `json:"premiummailsettings"`
	Ubiquity            Ubiquity            `json:"ubiquity"`
	Streams             Streams             `json:"streams"`
	Keyvalue            Keyvalue            `json:"keyvalue"`
	Archivews           Archivews           `json:"archivews"`
	Push                Push                `json:"push"`
	Iwmb                Iwmb                `json:"iwmb"`
	Iworkexportws       Iworkexportws       `json:"iworkexportws"`
	Sharedlibrary       Sharedlibrary       `json:"sharedlibrary"`
	Geows               Geows               `json:"geows"`
	Account             Account             `json:"account"`
	Contacts            WebserviceContacts  `json:"contacts"`
}

type Urls struct {
	AccountCreateUI     string `json:"accountCreateUI"`
	AccountLoginUI      string `json:"accountLoginUI"`
	AccountLogin        string `json:"accountLogin"`
	AccountRepairUI     string `json:"accountRepairUI"`
	DownloadICloudTerms string `json:"downloadICloudTerms"`
	RepairDone          string `json:"repairDone"`
	AccountAuthorizeUI  string `json:"accountAuthorizeUI"`
	VettingURLForEmail  string `json:"vettingUrlForEmail"`
	AccountCreate       string `json:"accountCreate"`
	GetICloudTerms      string `json:"getICloudTerms"`
	VettingURLForPhone  string `json:"vettingUrlForPhone"`
}

type ConfigBag struct {
	Urls                 Urls `json:"urls"`
	AccountCreateEnabled bool `json:"accountCreateEnabled"`
}

type RequestInfo struct {
	Country  string `json:"country"`
	TimeZone string `json:"timeZone"`
	Region   string `json:"region"`
}

type ICloudInfo struct {
	SafariBookmarksHasMigratedToCloudKit bool `json:"SafariBookmarksHasMigratedToCloudKit"`
}

type Calendar struct {
}

type Reminders struct {
}

type Keynote struct {
	IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
}

type Settings struct {
	CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
}

type Mail struct {
}

type Numbers struct {
	IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
}

type Photos struct {
}

type Pages struct {
	IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
}

type Notes3 struct {
}

type Find struct {
	CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
}

type Iclouddrive struct {
}

type Newspublisher struct {
	IsHidden bool `json:"isHidden"`
}

type Contacts struct {
}

type Apps struct {
	Calendar      Calendar      `json:"calendar"`
	Reminders     Reminders     `json:"reminders"`
	Keynote       Keynote       `json:"keynote"`
	Settings      Settings      `json:"settings"`
	Mail          Mail          `json:"mail"`
	Numbers       Numbers       `json:"numbers"`
	Photos        Photos        `json:"photos"`
	Pages         Pages         `json:"pages"`
	Notes3        Notes3        `json:"notes3"`
	Find          Find          `json:"find"`
	Iclouddrive   Iclouddrive   `json:"iclouddrive"`
	Newspublisher Newspublisher `json:"newspublisher"`
	Contacts      Contacts      `json:"contacts"`
}

type signInResponse struct {
	AuthType string `json:"authType"`
}

const (
	TwoFactorAuthType = "hsa2"
)

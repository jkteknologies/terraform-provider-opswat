package opswatClient

// globalSyncTimeout API /admin/config/file/sync
type globalSyncTimeout struct {
	Timeout int `json:"timeout"`
}

// Session API /admin/config/session
type Session struct {
	AbsoluteSessionTimeout int  `json:"absoluteSessionTimeout"`
	AllowCrossIpSessions   bool `json:"allowCrossIpSessions"`
	AllowDuplicateSession  bool `json:"allowDuplicateSession"`
	SessionTimeout         int  `json:"sessionTimeout"`
}

// Quarantine API /admin/config/quarantine
type Quarantine struct {
	Cleanuprange int `json:"cleanuprange"`
}

// ConfigUpdates API /admin/config/update
type Updates struct {
	Autoupdateperiod     int           `json:"autoupdateperiod"`
	Deleteafterimport    bool          `json:"deleteafterimport"`
	Disabledupdate       []interface{} `json:"disabledupdate"`
	Pickupfolder         string        `json:"pickupfolder"`
	Skipenginedependency bool          `json:"skipenginedependency"`
	Source               string        `json:"source"`
}

// API /admin/license
type License struct {
	ActivationKey   string `json:"activation_key"`
	DaysLeft        int    `json:"days_left"`
	Deployment      string `json:"deployment"`
	Expiration      string `json:"expiration"`
	LicensedEngines string `json:"licensed_engines"`
	LicensedTo      string `json:"licensed_to"`
	MatchStrategy   string `json:"match_strategy"`
	MatchingLicense string `json:"matchingLicense"`
	MatchingMachine string `json:"matchingMachine"`
	MaxAgentCount   string `json:"max_agent_count"`
	OnlineActivated bool   `json:"online_activated"`
	ProductId       string `json:"product_id"`
	ProductName     string `json:"product_name"`
	Serial          string `json:"serial"`
}

// API /admin/config/scan
type scanQueue struct {
	MaxQueuePerAgent int `json:"max_queue_per_agent"`
}

// API /admin/config/rule
type Workflow struct {
	AllowCert                                   bool          `json:"allow_cert"`
	AllowCertCert                               string        `json:"allow_cert.cert"`
	AllowCertCertValidity                       int           `json:"allow_cert.cert_validity"`
	AllowLocalFiles                             bool          `json:"allow_local_files"`
	AllowLocalFilesWhiteList                    bool          `json:"allow_local_files.local_files_white_list"`
	AllowLocalFilesLocalPaths                   []string      `json:"allow_local_files.local_paths"`
	Description                                 string        `json:"description"`
	Id                                          int           `json:"id"`
	IncludeWebhookSignature                     bool          `json:"include_webhook_signature"`
	IncludeWebhookSignatureWebhookCertificateId int           `json:"include_webhook_signature.webhook_certificate_id"`
	LastModified                                int64         `json:"last_modified"`
	Mutable                                     bool          `json:"mutable"`
	Name                                        string        `json:"name"`
	ScanAllowed                                 []interface{} `json:"scan_allowed"`
	WorkflowId                                  int           `json:"workflow_id"`
	ZoneId                                      int           `json:"zone_id"`
	PrefHashes                                  PrefHash      `json:"pref_hashes"`
}

// PrefHashes
type PrefHash struct {
	DSADVANCEDSETTINGHASH string `json:"DS_ADVANCED_SETTING_HASH"`
}

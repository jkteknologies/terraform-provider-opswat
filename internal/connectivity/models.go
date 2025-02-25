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
type Queue struct {
	MaxQueuePerAgent int `json:"max_queue_per_agent"`
}

// API /admin/config/rule
type Workflow struct {
	AllowCert                                   bool            `json:"allow_cert"`
	AllowCertCert                               string          `json:"allow_cert.cert"`
	AllowCertCertValidity                       int             `json:"allow_cert.cert_validity"`
	AllowLocalFiles                             bool            `json:"allow_local_files"`
	AllowLocalFilesWhiteList                    bool            `json:"allow_local_files.local_files_white_list"`
	AllowLocalFilesLocalPaths                   []string        `json:"allow_local_files.local_paths"`
	Description                                 string          `json:"description"`
	Id                                          int             `json:"id,omitempty"`
	IncludeWebhookSignature                     bool            `json:"include_webhook_signature"`
	IncludeWebhookSignatureWebhookCertificateId int             `json:"include_webhook_signature.webhook_certificate_id"`
	LastModified                                int             `json:"last_modified,omitempty"`
	Mutable                                     bool            `json:"mutable,omitempty"`
	Name                                        string          `json:"name"`
	ScanAllowed                                 []int           `json:"scan_allowed"`
	WorkflowId                                  int             `json:"workflow_id"`
	ZoneId                                      int             `json:"zone_id"`
	ResultAllowed                               []ResultAllowed `json:"result_allowed"`
	OptionValues                                OptionValues    `json:"option_values"`
	UserAgents                                  []string        `json:"user_agents"`
}

type OptionValues struct {
	ArchiveHandlingMaxNumberFiles           int  `json:"archive.archive_handling.max_number_files"`
	ArchiveHandlingMaxRecursionLevel        int  `json:"archive.archive_handling.max_recursion_level"`
	ArchiveHandlingMaxSizeFiles             int  `json:"archive.archive_handling.max_size_files"`
	ArchiveHandlingTimeout                  int  `json:"archive.archive_handling.timeout"`
	FiletypeAnalysisTimeout                 int  `json:"filetype_analysis.timeout"`
	ProcessInfoGlobalTimeout                bool `json:"process_info.global_timeout"`
	ProcessInfoGlobalTimeoutValue           int  `json:"process_info.global_timeout.value"`
	ProcessInfoMaxDownloadSize              int  `json:"process_info.max_download_size"`
	ProcessInfoMaxFileSize                  int  `json:"process_info.max_file_size"`
	ProcessInfoQuarantine                   bool `json:"process_info.quarantine"`
	ProcessInfoSkipHash                     bool `json:"process_info.skip_hash"`
	ProcessInfoSkipProcessingFastSymlink    bool `json:"process_info.skip_processing_fast_symlink"`
	ProcessInfoWorkflowPriority             int  `json:"process_info.workflow_priority"`
	ScanFilescanCheckAvEngine               bool `json:"scan.filescan.check_av_engine"`
	ScanFilescanDownloadTimeout             int  `json:"scan.filescan.download_timeout"`
	ScanFilescanGlobalScanTimeout           int  `json:"scan.filescan.global_scan_timeout"`
	ScanFilescanPerEngineScanTimeout        int  `json:"scan.filescan.per_engine_scan_timeout"`
	VulFilescanTimeoutVulnerabilityScanning int  `json:"vul.filescan.timeout_vulnerability_scanning"`
}

// ResultAllowed
type ResultAllowed struct {
	Role       int `json:"role"`
	Visibility int `json:"visibility"`
}

// UserDirectory API /admin/userdirectory
type UserDirectory struct {
	ID               int    `json:"id,omitempty"`
	Type             string `json:"type"`
	Enabled          bool   `json:"enabled"`
	Name             string `json:"name"`
	UserIdentifiedBy string `json:"user_identified_by"`
	Sp               Sp     `json:"sp"`
	Role             Role   `json:"role"`
	Version          string `json:"version"`
	Idp              Idp    `json:"idp"`
}

type Sp struct {
	LoginUrl           string `json:"login_url"`
	SupportLogoutUrl   bool   `json:"support_logout_url"`
	SupportPrivateKey  bool   `json:"support_private_key"`
	SupportEntityId    bool   `json:"support_entity_id"`
	EnableIdpInitiated bool   `json:"enable_idp_initiated"`
	EntityId           string `json:"entity_id"`
}

type Role struct {
	Option  string    `json:"option"`
	Details []Details `json:"details"`
}

type Details struct {
	Key    string   `json:"key"`
	Values []Values `json:"values"`
}

type Values struct {
	Condition string   `json:"condition"`
	RoleIds   []string `json:"role_ids"`
	Type      string   `json:"type"`
}

type Idp struct {
	AuthnRequestSigned bool         `json:"authn_request_signed"`
	EntityId           string       `json:"entity_id"`
	LoginMethod        LoginMethod  `json:"login_method"`
	LogoutMethod       LogoutMethod `json:"logout_method"`
	ValidUntil         string       `json:"valid_until"`
	X509Cert           []string     `json:"x509_cert"`
}

type LoginMethod struct {
	Post     string `json:"post"`
	Redirect string `json:"redirect"`
}

type LogoutMethod struct {
	Redirect string `json:"redirect"`
}

// User API /admin/user
type User struct {
	ApiKey      string   `json:"api_key,omitempty"`
	DirectoryId int      `json:"directory_id,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Email       string   `json:"email,omitempty"`
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Password    string   `json:"password,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

// Role API /admin/role
type UserRole struct {
	Name        string     `json:"name,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	ID          int        `json:"id,omitempty"`
	UserRights  UserRights `json:"rights,omitempty"`
}

type UserRights struct {
	//Scanlog     []string `json:"scanlog,omitempty"`
	//Statistics  []string `json:"statistics,omitempty"`
	//Quarantine  []string `json:"quarantine,omitempty"`
	//Updatelog   []string `json:"updatelog,omitempty"`
	//Configlog   []string `json:"configlog,omitempty"`
	//Rule        []string `json:"rule,omitempty"`
	//Workflow    []string `json:"workflow,omitempty"`
	//Zone        []string `json:"zone,omitempty"`
	//Agents      []string `json:"agents,omitempty"`
	//Engines     []string `json:"engines,omitempty"`
	//External    []string `json:"external,omitempty"`
	//Skip        []string `json:"skip,omitempty"`
	//Cert        []string `json:"cert,omitempty"`
	//WebhookAuth []string `json:"webhook_auth,omitempty"`
	//Retention   []string `json:"retention,omitempty"`
	//Users       []string `json:"users,omitempty"`
	//License     []string `json:"license,omitempty"`
	//Update      []string `json:"update,omitempty"`
	//Scan        []string `json:"scan,omitempty"`
	//Healthcheck []string `json:"healthcheck,omitempty"`
	Fetch    []string `json:"fetch"`
	Download []string `json:"download"`
}

// globalSyncTimeout API /admin/config/file/sync
type scanHistory struct {
	Cleanuprange int `json:"cleanuprange"`
}

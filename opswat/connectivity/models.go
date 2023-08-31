package opswatClient

// API /admin/config/file/sync
type globalSyncTimeout struct {
	Timeout int `json:"timeout"`
}

// API /admin/config/session
type ConfigSession struct {
	AbsoluteSessionTimeout int  `json:"absoluteSessionTimeout"`
	AllowCrossIpSessions   bool `json:"allowCrossIpSessions"`
	AllowDuplicateSession  bool `json:"allowDuplicateSession"`
	SessionTimeout         int  `json:"sessionTimeout"`
}

// API /admin/license
type license struct {
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
type workflow struct {
	AllowCert                                   bool          `json:"allow_cert"`
	AllowCertCert                               string        `json:"allow_cert.cert"`
	AllowCertCertValidity                       int           `json:"allow_cert.cert_validity"`
	AllowLocalFiles                             bool          `json:"allow_local_files"`
	AllowLocalFilesLocalFilesWhiteList          bool          `json:"allow_local_files.local_files_white_list"`
	AllowLocalFilesLocalPaths                   []interface{} `json:"allow_local_files.local_paths"`
	Description                                 string        `json:"description"`
	Id                                          int           `json:"id"`
	IncludeWebhookSignature                     bool          `json:"include_webhook_signature"`
	IncludeWebhookSignatureWebhookCertificateId int           `json:"include_webhook_signature.webhook_certificate_id"`
	LastModified                                int64         `json:"last_modified"`
	Mutable                                     bool          `json:"mutable"`
	Name                                        string        `json:"name"`
	OptionValues                                struct {
		ArchiveArchiveHandlingExtractTaskFailed                      bool   `json:"archive.archive_handling.extract_task_failed,omitempty"`
		ArchiveArchiveHandlingExtractTaskFailedExtractedPartially    bool   `json:"archive.archive_handling.extract_task_failed.extracted_partially,omitempty"`
		ArchiveArchiveHandlingExtractTaskFailedExtractionOtherErrors bool   `json:"archive.archive_handling.extract_task_failed.extraction_other_errors,omitempty"`
		ArchiveArchiveHandlingExtractTaskFailedInvalidFileStructure  bool   `json:"archive.archive_handling.extract_task_failed.invalid_file_structure,omitempty"`
		ProcessInfoMaxFileSize                                       int    `json:"process_info.max_file_size,omitempty"`
		Compression7ZGroupCompress                                   bool   `json:"compression.7z.group.compress,omitempty"`
		Compression7ZGroupCompressTo                                 string `json:"compression.7z.group.compress.to,omitempty"`
		CompressionDeepCdr                                           bool   `json:"compression.deep_cdr,omitempty"`
		CompressionDeepCdrAllowedOriginal                            bool   `json:"compression.deep_cdr.allowed_original,omitempty"`
		CompressionGzGroupCompress                                   bool   `json:"compression.gz.group.compress,omitempty"`
		CompressionGzGroupCompressTo                                 string `json:"compression.gz.group.compress.to,omitempty"`
		CompressionRarGroupCompress                                  bool   `json:"compression.rar.group.compress,omitempty"`
		CompressionRarGroupCompressTo                                string `json:"compression.rar.group.compress.to,omitempty"`
		CompressionXzGroupCompress                                   bool   `json:"compression.xz.group.compress,omitempty"`
		CompressionXzGroupCompressTo                                 string `json:"compression.xz.group.compress.to,omitempty"`
		CompressionZipGroupCompress                                  bool   `json:"compression.zip.group.compress,omitempty"`
		CompressionZipGroupCompressTo                                string `json:"compression.zip.group.compress.to,omitempty"`
		DsFilescan                                                   bool   `json:"ds.filescan,omitempty"`
		DsFilescanBmpGroupConvert                                    bool   `json:"ds.filescan.bmp.group.convert,omitempty"`
		DsFilescanBmpGroupConvertTo                                  string `json:"ds.filescan.bmp.group.convert.to,omitempty"`
		DsFilescanCsvGroupConvert                                    bool   `json:"ds.filescan.csv.group.convert,omitempty"`
		DsFilescanCsvGroupConvertTo                                  string `json:"ds.filescan.csv.group.convert.to,omitempty"`
		DsFilescanDocGroupConvert                                    bool   `json:"ds.filescan.doc.group.convert,omitempty"`
		DsFilescanDocGroupConvertTo                                  string `json:"ds.filescan.doc.group.convert.to,omitempty"`
		DsFilescanDocmGroupConvert                                   bool   `json:"ds.filescan.docm.group.convert,omitempty"`
		DsFilescanDocmGroupConvertTo                                 string `json:"ds.filescan.docm.group.convert.to,omitempty"`
		DsFilescanDocxGroupConvert                                   bool   `json:"ds.filescan.docx.group.convert,omitempty"`
		DsFilescanDocxGroupConvertTo                                 string `json:"ds.filescan.docx.group.convert.to,omitempty"`
		DsFilescanDotGroupConvert                                    bool   `json:"ds.filescan.dot.group.convert,omitempty"`
		DsFilescanDotGroupConvertTo                                  string `json:"ds.filescan.dot.group.convert.to,omitempty"`
		DsFilescanDotmGroupConvert                                   bool   `json:"ds.filescan.dotm.group.convert,omitempty"`
		DsFilescanDotmGroupConvertTo                                 string `json:"ds.filescan.dotm.group.convert.to,omitempty"`
		DsFilescanDotxGroupConvert                                   bool   `json:"ds.filescan.dotx.group.convert,omitempty"`
		DsFilescanDotxGroupConvertTo                                 string `json:"ds.filescan.dotx.group.convert.to,omitempty"`
		DsFilescanDwgGroupConvert                                    bool   `json:"ds.filescan.dwg.group.convert,omitempty"`
		DsFilescanDwgGroupConvertTo                                  string `json:"ds.filescan.dwg.group.convert.to,omitempty"`
		DsFilescanGifGroupConvert                                    bool   `json:"ds.filescan.gif.group.convert,omitempty"`
		DsFilescanGifGroupConvertTo                                  string `json:"ds.filescan.gif.group.convert.to,omitempty"`
		DsFilescanHtmlGroupConvert                                   bool   `json:"ds.filescan.html.group.convert,omitempty"`
		DsFilescanHtmlGroupConvertTo                                 string `json:"ds.filescan.html.group.convert.to,omitempty"`
		DsFilescanHwpGroupConvert                                    bool   `json:"ds.filescan.hwp.group.convert,omitempty"`
		DsFilescanHwpGroupConvertTo                                  string `json:"ds.filescan.hwp.group.convert.to,omitempty"`
		DsFilescanJpgGroupConvert                                    bool   `json:"ds.filescan.jpg.group.convert,omitempty"`
		DsFilescanJpgGroupConvertTo                                  string `json:"ds.filescan.jpg.group.convert.to,omitempty"`
		DsFilescanJtdGroupConvert                                    bool   `json:"ds.filescan.jtd.group.convert,omitempty"`
		DsFilescanJtdGroupConvertTo                                  string `json:"ds.filescan.jtd.group.convert.to,omitempty"`
		DsFilescanOdtGroupConvert                                    bool   `json:"ds.filescan.odt.group.convert,omitempty"`
		DsFilescanOdtGroupConvertTo                                  string `json:"ds.filescan.odt.group.convert.to,omitempty"`
		DsFilescanPdfGroupConvert                                    bool   `json:"ds.filescan.pdf.group.convert,omitempty"`
		DsFilescanPdfGroupConvertTo                                  string `json:"ds.filescan.pdf.group.convert.to,omitempty"`
		DsFilescanPngGroupConvert                                    bool   `json:"ds.filescan.png.group.convert,omitempty"`
		DsFilescanPngGroupConvertTo                                  string `json:"ds.filescan.png.group.convert.to,omitempty"`
		DsFilescanPpsxGroupConvert                                   bool   `json:"ds.filescan.ppsx.group.convert,omitempty"`
		DsFilescanPpsxGroupConvertTo                                 string `json:"ds.filescan.ppsx.group.convert.to,omitempty"`
		DsFilescanPptGroupConvert                                    bool   `json:"ds.filescan.ppt.group.convert,omitempty"`
		DsFilescanPptGroupConvertTo                                  string `json:"ds.filescan.ppt.group.convert.to,omitempty"`
		DsFilescanPptmGroupConvert                                   bool   `json:"ds.filescan.pptm.group.convert,omitempty"`
		DsFilescanPptmGroupConvertTo                                 string `json:"ds.filescan.pptm.group.convert.to,omitempty"`
		DsFilescanPptxGroupConvert                                   bool   `json:"ds.filescan.pptx.group.convert,omitempty"`
		DsFilescanPptxGroupConvertTo                                 string `json:"ds.filescan.pptx.group.convert.to,omitempty"`
		DsFilescanRtfGroupConvert                                    bool   `json:"ds.filescan.rtf.group.convert,omitempty"`
		DsFilescanRtfGroupConvertTo                                  string `json:"ds.filescan.rtf.group.convert.to,omitempty"`
		DsFilescanSvgGroupConvert                                    bool   `json:"ds.filescan.svg.group.convert,omitempty"`
		DsFilescanSvgGroupConvertTo                                  string `json:"ds.filescan.svg.group.convert.to,omitempty"`
		DsFilescanTiffGroupConvert                                   bool   `json:"ds.filescan.tiff.group.convert,omitempty"`
		DsFilescanTiffGroupConvertTo                                 string `json:"ds.filescan.tiff.group.convert.to,omitempty"`
		DsFilescanWmfGroupConvert                                    bool   `json:"ds.filescan.wmf.group.convert,omitempty"`
		DsFilescanWmfGroupConvertTo                                  string `json:"ds.filescan.wmf.group.convert.to,omitempty"`
		DsFilescanXlsGroupConvert                                    bool   `json:"ds.filescan.xls.group.convert,omitempty"`
		DsFilescanXlsGroupConvertTo                                  string `json:"ds.filescan.xls.group.convert.to,omitempty"`
		DsFilescanXlsbGroupConvert                                   bool   `json:"ds.filescan.xlsb.group.convert,omitempty"`
		DsFilescanXlsbGroupConvertTo                                 string `json:"ds.filescan.xlsb.group.convert.to,omitempty"`
		DsFilescanXlsmGroupConvert                                   bool   `json:"ds.filescan.xlsm.group.convert,omitempty"`
		DsFilescanXlsmGroupConvertTo                                 string `json:"ds.filescan.xlsm.group.convert.to,omitempty"`
		DsFilescanXlsxGroupConvert                                   bool   `json:"ds.filescan.xlsx.group.convert,omitempty"`
		DsFilescanXlsxGroupConvertTo                                 string `json:"ds.filescan.xlsx.group.convert.to,omitempty"`
		DsFilescanXmlDocGroupConvert                                 bool   `json:"ds.filescan.xml-doc.group.convert,omitempty"`
		DsFilescanXmlDocGroupConvertTo                               string `json:"ds.filescan.xml-doc.group.convert.to,omitempty"`
		DsFilescanXmlDocxGroupConvert                                bool   `json:"ds.filescan.xml-docx.group.convert,omitempty"`
		DsFilescanXmlDocxGroupConvertTo                              string `json:"ds.filescan.xml-docx.group.convert.to,omitempty"`
		DsFilescanXmlXlsGroupConvert                                 bool   `json:"ds.filescan.xml-xls.group.convert,omitempty"`
		DsFilescanXmlXlsGroupConvertTo                               string `json:"ds.filescan.xml-xls.group.convert.to,omitempty"`
		DsFilescanXmlGroupConvert                                    bool   `json:"ds.filescan.xml.group.convert,omitempty"`
		DsFilescanXmlGroupConvertTo                                  string `json:"ds.filescan.xml.group.convert.to,omitempty"`
		ProcessInfoGlobalTimeoutValue                                int    `json:"process_info.global_timeout.value,omitempty"`
		DlpFilescan                                                  bool   `json:"dlp.filescan,omitempty"`
		DlpFilescanScanCertaintyThreshold                            int    `json:"dlp.filescan.scan.certainty_threshold,omitempty"`
		DlpFilescanScanCidr                                          bool   `json:"dlp.filescan.scan.cidr,omitempty"`
		DlpFilescanScanIpv4                                          bool   `json:"dlp.filescan.scan.ipv4,omitempty"`
		DlpFilescanScanRegex                                         bool   `json:"dlp.filescan.scan.regex,omitempty"`
		DlpFilescanScanRegexRegexlist                                []struct {
			Allow               bool   `json:"allow"`
			Certainty           int    `json:"certainty"`
			Keywords            string `json:"keywords"`
			Name                string `json:"name"`
			Redact              bool   `json:"redact"`
			ValidationBeginning string `json:"validation.beginning"`
			ValidationDelimiter string `json:"validation.delimiter"`
			ValidationDuplicate bool   `json:"validation.duplicate"`
			ValidationEnding    string `json:"validation.ending"`
			ValidationPrefix    string `json:"validation.prefix"`
			ValidationSuffix    string `json:"validation.suffix"`
			Value               string `json:"value"`
		} `json:"dlp.filescan.scan.regex.regexlist,omitempty"`
		ProcessInfoQuarantine                   bool `json:"process_info.quarantine,omitempty"`
		ArchiveArchiveHandling                  bool `json:"archive.archive_handling,omitempty"`
		ArchiveArchiveHandlingMaxNumberFiles    int  `json:"archive.archive_handling.max_number_files,omitempty"`
		ArchiveArchiveHandlingMaxRecursionLevel int  `json:"archive.archive_handling.max_recursion_level,omitempty"`
		ArchiveArchiveHandlingMaxSizeFiles      int  `json:"archive.archive_handling.max_size_files,omitempty"`
		ArchiveArchiveHandlingTimeout           int  `json:"archive.archive_handling.timeout,omitempty"`
		FiletypeAnalysisTimeout                 int  `json:"filetype_analysis.timeout,omitempty"`
		ProcessInfoGlobalTimeout                bool `json:"process_info.global_timeout,omitempty"`
		ProcessInfoMaxDownloadSize              int  `json:"process_info.max_download_size,omitempty"`
		ProcessInfoSkipHash                     bool `json:"process_info.skip_hash,omitempty"`
		ProcessInfoSkipProcessingFastSymlink    bool `json:"process_info.skip_processing_fast_symlink,omitempty"`
		ProcessInfoWorkflowPriority             int  `json:"process_info.workflow_priority,omitempty"`
		ScanFilescanCheckAvEngine               bool `json:"scan.filescan.check_av_engine,omitempty"`
		ScanFilescanDownloadTimeout             int  `json:"scan.filescan.download_timeout,omitempty"`
		ScanFilescanGlobalScanTimeout           int  `json:"scan.filescan.global_scan_timeout,omitempty"`
		ScanFilescanPerEngineScanTimeout        int  `json:"scan.filescan.per_engine_scan_timeout,omitempty"`
		VulFilescanTimeoutVulnerabilityScanning int  `json:"vul.filescan.timeout_vulnerability_scanning,omitempty"`
	} `json:"option_values"`
	PrefHashes struct {
		DSADVANCEDSETTINGHASH string `json:"DS_ADVANCED_SETTING_HASH"`
	} `json:"pref_hashes"`
	ResultAllowed []struct {
		Role       interface{} `json:"role"`
		Visibility int         `json:"visibility"`
	} `json:"result_allowed"`
	ScanAllowed []interface{} `json:"scan_allowed"`
	UserAgents  []string      `json:"user_agents"`
	WorkflowId  int           `json:"workflow_id"`
	ZoneId      int           `json:"zone_id"`
}

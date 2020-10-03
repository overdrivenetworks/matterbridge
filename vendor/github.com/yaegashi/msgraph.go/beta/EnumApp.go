// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AppInstallControlType undocumented
type AppInstallControlType string

const (
	// AppInstallControlTypeVNotConfigured undocumented
	AppInstallControlTypeVNotConfigured AppInstallControlType = "notConfigured"
	// AppInstallControlTypeVAnywhere undocumented
	AppInstallControlTypeVAnywhere AppInstallControlType = "anywhere"
	// AppInstallControlTypeVStoreOnly undocumented
	AppInstallControlTypeVStoreOnly AppInstallControlType = "storeOnly"
	// AppInstallControlTypeVRecommendations undocumented
	AppInstallControlTypeVRecommendations AppInstallControlType = "recommendations"
	// AppInstallControlTypeVPreferStore undocumented
	AppInstallControlTypeVPreferStore AppInstallControlType = "preferStore"
)

var (
	// AppInstallControlTypePNotConfigured is a pointer to AppInstallControlTypeVNotConfigured
	AppInstallControlTypePNotConfigured = &_AppInstallControlTypePNotConfigured
	// AppInstallControlTypePAnywhere is a pointer to AppInstallControlTypeVAnywhere
	AppInstallControlTypePAnywhere = &_AppInstallControlTypePAnywhere
	// AppInstallControlTypePStoreOnly is a pointer to AppInstallControlTypeVStoreOnly
	AppInstallControlTypePStoreOnly = &_AppInstallControlTypePStoreOnly
	// AppInstallControlTypePRecommendations is a pointer to AppInstallControlTypeVRecommendations
	AppInstallControlTypePRecommendations = &_AppInstallControlTypePRecommendations
	// AppInstallControlTypePPreferStore is a pointer to AppInstallControlTypeVPreferStore
	AppInstallControlTypePPreferStore = &_AppInstallControlTypePPreferStore
)

var (
	_AppInstallControlTypePNotConfigured   = AppInstallControlTypeVNotConfigured
	_AppInstallControlTypePAnywhere        = AppInstallControlTypeVAnywhere
	_AppInstallControlTypePStoreOnly       = AppInstallControlTypeVStoreOnly
	_AppInstallControlTypePRecommendations = AppInstallControlTypeVRecommendations
	_AppInstallControlTypePPreferStore     = AppInstallControlTypeVPreferStore
)

// AppListType undocumented
type AppListType string

const (
	// AppListTypeVNone undocumented
	AppListTypeVNone AppListType = "none"
	// AppListTypeVAppsInListCompliant undocumented
	AppListTypeVAppsInListCompliant AppListType = "appsInListCompliant"
	// AppListTypeVAppsNotInListCompliant undocumented
	AppListTypeVAppsNotInListCompliant AppListType = "appsNotInListCompliant"
)

var (
	// AppListTypePNone is a pointer to AppListTypeVNone
	AppListTypePNone = &_AppListTypePNone
	// AppListTypePAppsInListCompliant is a pointer to AppListTypeVAppsInListCompliant
	AppListTypePAppsInListCompliant = &_AppListTypePAppsInListCompliant
	// AppListTypePAppsNotInListCompliant is a pointer to AppListTypeVAppsNotInListCompliant
	AppListTypePAppsNotInListCompliant = &_AppListTypePAppsNotInListCompliant
)

var (
	_AppListTypePNone                   = AppListTypeVNone
	_AppListTypePAppsInListCompliant    = AppListTypeVAppsInListCompliant
	_AppListTypePAppsNotInListCompliant = AppListTypeVAppsNotInListCompliant
)

// AppLockerApplicationControlType undocumented
type AppLockerApplicationControlType string

const (
	// AppLockerApplicationControlTypeVNotConfigured undocumented
	AppLockerApplicationControlTypeVNotConfigured AppLockerApplicationControlType = "notConfigured"
	// AppLockerApplicationControlTypeVEnforceComponentsAndStoreApps undocumented
	AppLockerApplicationControlTypeVEnforceComponentsAndStoreApps AppLockerApplicationControlType = "enforceComponentsAndStoreApps"
	// AppLockerApplicationControlTypeVAuditComponentsAndStoreApps undocumented
	AppLockerApplicationControlTypeVAuditComponentsAndStoreApps AppLockerApplicationControlType = "auditComponentsAndStoreApps"
	// AppLockerApplicationControlTypeVEnforceComponentsStoreAppsAndSmartlocker undocumented
	AppLockerApplicationControlTypeVEnforceComponentsStoreAppsAndSmartlocker AppLockerApplicationControlType = "enforceComponentsStoreAppsAndSmartlocker"
	// AppLockerApplicationControlTypeVAuditComponentsStoreAppsAndSmartlocker undocumented
	AppLockerApplicationControlTypeVAuditComponentsStoreAppsAndSmartlocker AppLockerApplicationControlType = "auditComponentsStoreAppsAndSmartlocker"
)

var (
	// AppLockerApplicationControlTypePNotConfigured is a pointer to AppLockerApplicationControlTypeVNotConfigured
	AppLockerApplicationControlTypePNotConfigured = &_AppLockerApplicationControlTypePNotConfigured
	// AppLockerApplicationControlTypePEnforceComponentsAndStoreApps is a pointer to AppLockerApplicationControlTypeVEnforceComponentsAndStoreApps
	AppLockerApplicationControlTypePEnforceComponentsAndStoreApps = &_AppLockerApplicationControlTypePEnforceComponentsAndStoreApps
	// AppLockerApplicationControlTypePAuditComponentsAndStoreApps is a pointer to AppLockerApplicationControlTypeVAuditComponentsAndStoreApps
	AppLockerApplicationControlTypePAuditComponentsAndStoreApps = &_AppLockerApplicationControlTypePAuditComponentsAndStoreApps
	// AppLockerApplicationControlTypePEnforceComponentsStoreAppsAndSmartlocker is a pointer to AppLockerApplicationControlTypeVEnforceComponentsStoreAppsAndSmartlocker
	AppLockerApplicationControlTypePEnforceComponentsStoreAppsAndSmartlocker = &_AppLockerApplicationControlTypePEnforceComponentsStoreAppsAndSmartlocker
	// AppLockerApplicationControlTypePAuditComponentsStoreAppsAndSmartlocker is a pointer to AppLockerApplicationControlTypeVAuditComponentsStoreAppsAndSmartlocker
	AppLockerApplicationControlTypePAuditComponentsStoreAppsAndSmartlocker = &_AppLockerApplicationControlTypePAuditComponentsStoreAppsAndSmartlocker
)

var (
	_AppLockerApplicationControlTypePNotConfigured                            = AppLockerApplicationControlTypeVNotConfigured
	_AppLockerApplicationControlTypePEnforceComponentsAndStoreApps            = AppLockerApplicationControlTypeVEnforceComponentsAndStoreApps
	_AppLockerApplicationControlTypePAuditComponentsAndStoreApps              = AppLockerApplicationControlTypeVAuditComponentsAndStoreApps
	_AppLockerApplicationControlTypePEnforceComponentsStoreAppsAndSmartlocker = AppLockerApplicationControlTypeVEnforceComponentsStoreAppsAndSmartlocker
	_AppLockerApplicationControlTypePAuditComponentsStoreAppsAndSmartlocker   = AppLockerApplicationControlTypeVAuditComponentsStoreAppsAndSmartlocker
)

// AppLogDecryptionAlgorithm undocumented
type AppLogDecryptionAlgorithm string

const (
	// AppLogDecryptionAlgorithmVAes256 undocumented
	AppLogDecryptionAlgorithmVAes256 AppLogDecryptionAlgorithm = "aes256"
)

var (
	// AppLogDecryptionAlgorithmPAes256 is a pointer to AppLogDecryptionAlgorithmVAes256
	AppLogDecryptionAlgorithmPAes256 = &_AppLogDecryptionAlgorithmPAes256
)

var (
	_AppLogDecryptionAlgorithmPAes256 = AppLogDecryptionAlgorithmVAes256
)

// AppLogUploadState undocumented
type AppLogUploadState string

const (
	// AppLogUploadStateVPending undocumented
	AppLogUploadStateVPending AppLogUploadState = "pending"
	// AppLogUploadStateVCompleted undocumented
	AppLogUploadStateVCompleted AppLogUploadState = "completed"
	// AppLogUploadStateVFailed undocumented
	AppLogUploadStateVFailed AppLogUploadState = "failed"
)

var (
	// AppLogUploadStatePPending is a pointer to AppLogUploadStateVPending
	AppLogUploadStatePPending = &_AppLogUploadStatePPending
	// AppLogUploadStatePCompleted is a pointer to AppLogUploadStateVCompleted
	AppLogUploadStatePCompleted = &_AppLogUploadStatePCompleted
	// AppLogUploadStatePFailed is a pointer to AppLogUploadStateVFailed
	AppLogUploadStatePFailed = &_AppLogUploadStatePFailed
)

var (
	_AppLogUploadStatePPending   = AppLogUploadStateVPending
	_AppLogUploadStatePCompleted = AppLogUploadStateVCompleted
	_AppLogUploadStatePFailed    = AppLogUploadStateVFailed
)

// AppManagementLevel undocumented
type AppManagementLevel string

const (
	// AppManagementLevelVUnspecified undocumented
	AppManagementLevelVUnspecified AppManagementLevel = "unspecified"
	// AppManagementLevelVUnmanaged undocumented
	AppManagementLevelVUnmanaged AppManagementLevel = "unmanaged"
	// AppManagementLevelVMDM undocumented
	AppManagementLevelVMDM AppManagementLevel = "mdm"
	// AppManagementLevelVAndroidEnterprise undocumented
	AppManagementLevelVAndroidEnterprise AppManagementLevel = "androidEnterprise"
)

var (
	// AppManagementLevelPUnspecified is a pointer to AppManagementLevelVUnspecified
	AppManagementLevelPUnspecified = &_AppManagementLevelPUnspecified
	// AppManagementLevelPUnmanaged is a pointer to AppManagementLevelVUnmanaged
	AppManagementLevelPUnmanaged = &_AppManagementLevelPUnmanaged
	// AppManagementLevelPMDM is a pointer to AppManagementLevelVMDM
	AppManagementLevelPMDM = &_AppManagementLevelPMDM
	// AppManagementLevelPAndroidEnterprise is a pointer to AppManagementLevelVAndroidEnterprise
	AppManagementLevelPAndroidEnterprise = &_AppManagementLevelPAndroidEnterprise
)

var (
	_AppManagementLevelPUnspecified       = AppManagementLevelVUnspecified
	_AppManagementLevelPUnmanaged         = AppManagementLevelVUnmanaged
	_AppManagementLevelPMDM               = AppManagementLevelVMDM
	_AppManagementLevelPAndroidEnterprise = AppManagementLevelVAndroidEnterprise
)

// AppVulnerabilityTaskMitigationType undocumented
type AppVulnerabilityTaskMitigationType string

const (
	// AppVulnerabilityTaskMitigationTypeVUnknown undocumented
	AppVulnerabilityTaskMitigationTypeVUnknown AppVulnerabilityTaskMitigationType = "unknown"
	// AppVulnerabilityTaskMitigationTypeVUpdate undocumented
	AppVulnerabilityTaskMitigationTypeVUpdate AppVulnerabilityTaskMitigationType = "update"
	// AppVulnerabilityTaskMitigationTypeVUninstall undocumented
	AppVulnerabilityTaskMitigationTypeVUninstall AppVulnerabilityTaskMitigationType = "uninstall"
	// AppVulnerabilityTaskMitigationTypeVSecurityConfiguration undocumented
	AppVulnerabilityTaskMitigationTypeVSecurityConfiguration AppVulnerabilityTaskMitigationType = "securityConfiguration"
)

var (
	// AppVulnerabilityTaskMitigationTypePUnknown is a pointer to AppVulnerabilityTaskMitigationTypeVUnknown
	AppVulnerabilityTaskMitigationTypePUnknown = &_AppVulnerabilityTaskMitigationTypePUnknown
	// AppVulnerabilityTaskMitigationTypePUpdate is a pointer to AppVulnerabilityTaskMitigationTypeVUpdate
	AppVulnerabilityTaskMitigationTypePUpdate = &_AppVulnerabilityTaskMitigationTypePUpdate
	// AppVulnerabilityTaskMitigationTypePUninstall is a pointer to AppVulnerabilityTaskMitigationTypeVUninstall
	AppVulnerabilityTaskMitigationTypePUninstall = &_AppVulnerabilityTaskMitigationTypePUninstall
	// AppVulnerabilityTaskMitigationTypePSecurityConfiguration is a pointer to AppVulnerabilityTaskMitigationTypeVSecurityConfiguration
	AppVulnerabilityTaskMitigationTypePSecurityConfiguration = &_AppVulnerabilityTaskMitigationTypePSecurityConfiguration
)

var (
	_AppVulnerabilityTaskMitigationTypePUnknown               = AppVulnerabilityTaskMitigationTypeVUnknown
	_AppVulnerabilityTaskMitigationTypePUpdate                = AppVulnerabilityTaskMitigationTypeVUpdate
	_AppVulnerabilityTaskMitigationTypePUninstall             = AppVulnerabilityTaskMitigationTypeVUninstall
	_AppVulnerabilityTaskMitigationTypePSecurityConfiguration = AppVulnerabilityTaskMitigationTypeVSecurityConfiguration
)

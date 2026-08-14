package service

// APIs used to interact with Core
const (
	UserProfileURL             = "/api/v1/users/profile/"                   // get the current user's basic info
	TerminalRegisterURL        = "/api/v1/terminal/terminal-registrations/" // register
	TerminalConfigURL          = "/api/v1/terminal/terminals/config/"       // get config
	TerminalHeartBeatURL       = "/api/v1/terminal/terminals/status/"
	TerminalEncryptedConfigURL = "/api/v1/terminal/encrypted-config/"
)

// APIs used for user login authentication
const (
	UserTokenAuthURL   = "/api/v1/authentication/tokens/" // user login validation
	UserConfirmAuthURL = "/api/v1/authentication/login-confirm-ticket/status/"
	AuthMFASelectURL   = "/api/v1/authentication/mfa/select/" // select MFA method

)

// Session-related APIs
const (
	SessionListURL      = "/api/v1/terminal/sessions/"           // upload the created asset session id
	SessionDetailURL    = "/api/v1/terminal/sessions/%s/"        // sent when finishing the session
	SessionReplayURL    = "/api/v1/terminal/sessions/%s/replay/" // upload the recording
	SessionCommandURL   = "/api/v1/terminal/commands/"           // upload batch commands
	FinishTaskURL       = "/api/v1/terminal/tasks/%s/"
	JoinRoomValidateURL = "/api/v1/terminal/sessions/join/validate/"
	FTPLogListURL       = "/api/v1/audits/ftp-logs/" // upload ftp log
	FTPLogUpdateURL     = "/api/v1/audits/ftp-logs/%s/"
	FTPLogFileURL       = "/api/v1/audits/ftp-logs/%s/upload/"

	SessionLifecycleLogURL = "/api/v1/terminal/sessions/%s/lifecycle_log/"
)

// Permission-related APIs
const (
	UserPermsNodesListURL         = "/api/v1/perms/users/%s/nodes/"
	UserPermsNodeAssetsListURL    = "/api/v1/perms/users/%s/nodes/%s/assets/"
	UserPermsNodeTreeWithAssetURL = "/api/v1/perms/users/%s/nodes/children-with-assets/tree/" // asset tree
)

// Resource detail APIs
const (
	UserListURL      = "/api/v1/users/users/"
	UserDetailURL    = "/api/v1/users/users/%s/"
	AssetPlatFormURL = "/api/v1/assets/assets/%s/platform/"

	DomainDetailWithGateways = "/api/v1/assets/domains/%s/?gateway=1"

	UserSuggestionsURL = "/api/v1/users/users/suggestions/"
)

const (
	NotificationCommandURL = "/api/v1/terminal/commands/insecure-command/"
)

// Command review

const (
	ShareCreateURL        = "/api/v1/terminal/session-sharings/"
	ShareSessionJoinURL   = "/api/v1/terminal/session-join-records/"
	ShareSessionFinishURL = "/api/v1/terminal/session-join-records/%s/finished/"
)

const (
	PublicSettingURL = "/api/v1/settings/public/"
)

const (
	TicketSessionURL = "/api/v1/tickets/ticket-session-relation/"
)

const (
	DBListenPortsURL = "/api/v1/terminal/db-listen-ports/"
	DBPortInfoURL    = "/api/v1/terminal/db-listen-ports/db-info/"

	SuperConnectTokenSecretURL        = "/api/v1/authentication/super-connection-token/secret/"
	SuperConnectTokenInfoURL          = "/api/v1/authentication/super-connection-token/"
	SuperConnectTokenK8sAliasesURL    = "/api/v1/authentication/super-connection-token/k8s-aliases/"
	SuperConnectTokenK8sAliasesSetURL = "/api/v1/authentication/super-connection-token/k8s-aliases/set/"

	UserPermsAssetAccountsURL = "/api/v1/perms/users/%s/assets/%s/"
	AccountSecretURL          = "/api/v1/assets/account-secrets/%s/"
	UserPermsAssetsURL        = "/api/v1/perms/users/%s/assets/"

	AssetLoginConfirmURL = "/api/v1/acls/login-asset/check/"
	AclCommandReviewURL  = "/api/v1/acls/command-filter-acls/command-review/"
)

const (
	UserKoKoPreferenceURL = "/api/v1/users/preference/?category=koko"
)

// applet and virtual app related
const (
	SuperConnectTokenAppletOptionURL        = "/api/v1/authentication/super-connection-token/applet-option/"
	SuperConnectAppletHostAccountReleaseURL = "/api/v1/authentication/super-connection-token/applet-account/release/"
	SuperConnectTokenVirtualAppOptionURL    = "/api/v1/authentication/super-connection-token/virtual-app-option/"

	SuperConnectTokenCheckURL = "/api/v1/authentication/super-connection-token/%s/check/"
	SuperTokenRenewalURL      = "/api/v1/authentication/super-connection-token/renewal/"
)

const (
	FaceRecognitionURL    = "/api/v1/authentication/face/callback/"
	FaceMonitorURL        = "/api/v1/authentication/face-monitor/callback/"
	FaceMonitorContextUrl = "/api/v1/authentication/face-monitor/context/"
)

const (
	AccountsChatURL = "/api/v1/accounts/accounts/chat/"
)

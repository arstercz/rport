package routes

const (
	ParamClientID       = "client_id"
	ParamClientAuthID   = "client_auth_id"
	ParamUserID         = "user_id"
	ParamSessionID      = "session_id"
	ParamJobID          = "job_id"
	ParamGroupID        = "group_id"
	ParamTokenPrefix    = "prefix"
	ParamVaultValueID   = "vault_value_id"
	ParamScriptValueID  = "script_value_id"
	ParamCommandValueID = "command_value_id"
	ParamGraphName      = "graph_name"
	ParamTemplateID     = "template_id"
	ParamProblemID      = "problem_id"
	ParamNotificationID = "notification_id"

	AllRoutesPrefix             = "/api/v1"
	AuthRoutesPrefix            = "/auth"
	AuthProviderRoute           = "/provider"
	AuthSettingsRoute           = "/ext/settings"
	AuthDeviceSettingsRoute     = "/ext/settings/device"
	AlertingServiceRoutesPrefix = "/monitoring"
	ASRuleSetRoute              = "/rules"
	ASTemplatesRoute            = "/notification-templates"
	ASProblemsRoute             = "/problems"
	TotPRoutes                  = "/me/totp-secret"
	Verify2FaRoute              = "/verify-2fa"
	FilesUploadRouteName        = "files"
)

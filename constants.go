package libra

// 通用元数据键
const (
	XLibraPrefix           = "x-libra-"
	XLibraToken            = XLibraPrefix + "token"
	XLibraAppId            = XLibraPrefix + "app-id"
	XLibraAppSecret        = XLibraPrefix + "app-secret"
	XLibraAuthBy           = XLibraPrefix + "auth-by"
	XLibraTrustedPrefix    = XLibraPrefix + "trusted-"
	XLibraTrustedAuthBy    = XLibraTrustedPrefix + "auth-by"
	XLibraTrustedAppId     = XLibraTrustedPrefix + "app-id"
	XLibraTrustedUserId    = XLibraTrustedPrefix + "user-id"
	XLibraTrustedRoleId    = XLibraTrustedPrefix + "role-id"
	XLibraTrustedRoleIndex = XLibraTrustedPrefix + "role-index"
)

// 通用元数据值
const (
	XLibraAuthByToken                  = "token"
	XLibraAuthBySecret                 = "secret"
	XLibraAuthBySecretAndOptionalToken = "secret-and-optional-token"
)

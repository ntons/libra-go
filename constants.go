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

// 通过secret验证将附加元数据
//   x-libra-trusted-app-id
// 通过token验证将附加元数据
//   x-libra-trusted-app-id
//   x-libra-trusted-user-id
//   会话签入后会附加额外元数据
//   x-libra-trusted-role-id
//   x-libra-trusted-role-index

// 基础策略 x-libra-trusted-auth-by列表只包含基础策略
const (
	// 必须提供正确token
	XLibraAuthByToken = "token"
	// 必须提供正确secret
	XLibraAuthBySecret = "secret"
)

// 策略组合 用于网关策略配置
const (
	// 必须提供正确secret，如果同时提供了token必须也是正确的且AppId匹配
	XLibraAuthBySecretAndOptionalToken = "secret-and-optional-token"
	// 必须提供secret和token中的至少1个，如果同时都提供则必须都正确且AppId匹配
	XLibraAuthBySecretOrToken = "secret-or-token"
	XLibraAuthByTokenOrSecret = "token-or-secret"
)

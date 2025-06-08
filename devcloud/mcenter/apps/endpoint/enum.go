package endpoint

type AccessMode int8

const (
	AccessModeRead = iota
	AccessModeReadWrite
)

const (
	MetaRequiredAuthKey      = "required_auth"
	MetaRequiredCodeKey      = "required_code"
	MetaRequiredPermKey      = "required_perm"
	MetaRequiredRoleKey      = "required_role"
	MetaRequiredAuditKey     = "required_audit"
	MetaRequiredNamespaceKey = "required_namespace"
	MetaResourceKey          = "resource"
	MetaActionKey            = "action"
)

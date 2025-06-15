package comptroller

const (
	// MetaAuditKey 定义开启审计的元数据键
	MetaAuditKey = "audit"
)

func Enable(v bool) (string, bool) {
	return MetaAuditKey, v
}

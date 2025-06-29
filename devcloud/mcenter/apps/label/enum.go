package label

type ValueType string

const (
	ValueTypeText     ValueType = "text"
	ValueTypeBoolean  ValueType = "bool"
	ValueTypeEnum     ValueType = "enum"
	ValueTypeHttpEnum ValueType = "http_enum" // 基于 url 的远程选项拉取,仅存储 url 地址,前端自己处理
)

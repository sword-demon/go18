package application

type Type int32

const (
	TypeSourceCode     Type = 0  // 源代码
	TypeContainerImage Type = 1  // 容器镜像
	TypeOther          Type = 15 // 其他类型
)

var (
	TypeName = map[Type]string{
		TypeSourceCode:     "SourceCode",
		TypeContainerImage: "ContainerImage",
		TypeOther:          "Other",
	}
)

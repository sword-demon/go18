// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package label

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
)

type Label struct {
	apps.ResourceMeta
	CreateLabelRequest `bson:",inline" gorm:"embedded"`
}

func (l *Label) TableName() string {
	return "labels"
}

func (l *Label) String() string {
	return pretty.ToJSON(l)
}

func NewLabel(spc *CreateLabelRequest) (*Label, error) {
	if err := spc.Validate(); err != nil {
		return nil, err
	}
	return &Label{
		ResourceMeta:       *apps.NewResourceMeta(),
		CreateLabelRequest: *spc,
	}, nil
}

type CreateLabelRequest struct {
	CreateBy        string `json:"create_by" bson:"create_by" gorm:"column:create_by;type:varchar(255)" description:"创建者"`
	Domain          string `json:"domain" bson:"domain" gorm:"column:domain;type:varchar(100)" description:"标签的键"`
	Namespace       string `json:"namespace" bson:"namespace" gorm:"column:namespace;type:varchar(100)" description:"标签的命名空间"`
	CreateLabelSpec        // 用户参数
}

func NewCreateLabelRequest() *CreateLabelRequest {
	return &CreateLabelRequest{
		CreateLabelSpec: CreateLabelSpec{
			Resources:   []string{},
			EnumOptions: []*EnumOption{},
			Extras:      map[string]string{},
		},
	}
}

func (c *CreateLabelRequest) Validate() error {
	return validator.Validate(c)
}

func (c *CreateLabelRequest) AddEnumOption(enums ...*EnumOption) *CreateLabelRequest {
	if c.EnumOptions == nil {
		c.EnumOptions = make([]*EnumOption, 0)
	}
	c.EnumOptions = append(c.EnumOptions, enums...)
	return c
}

type CreateLabelSpec struct {
	Resources      []string        `json:"resources" bson:"resources" gorm:"column:resources;type:json;serializer:json" description:"适用于哪些资源" optional:"true"`
	Key            string          `json:"key" bson:"key" gorm:"column:key;type:varchar(255)" description:"标签的键" validate:"required"`                  // 标签的键不允许修改
	KeyDesc        string          `json:"key_desc" bson:"key_desc" gorm:"column:key_desc;type:varchar(255)" description:"标签的键描述" validate:"required"` // 标签的键描述
	Color          string          `json:"color" bson:"color" gorm:"column:color;type:varchar(100)" description:"标签的颜色"`
	ValueType      ValueType       `json:"value_type" bson:"value_type" gorm:"column:value_type;type:varchar(20)" description:"标签的值类型"`                          // 标签的值类型
	DefaultValue   string          `json:"default_value" bson:"default_value" gorm:"column:default_value;type:text" description:"标签的默认值"`                        // 标签的默认值
	ValueDesc      string          `json:"value_desc" bson:"value_desc" gorm:"column:value_desc;type:text" description:"标签的值描述"`                                 // 标签的值描述
	Multiple       bool            `json:"multiple" bson:"multiple" gorm:"column:multiple;type:tinyint(1)" description:"标签的值是否允许多选"`                             // 标签的值是否允许多选
	EnumOptions    []*EnumOption   `json:"enum_options,omitempty" bson:"enum_options" gorm:"column:enum_options;type:json;serializer:json" description:"枚举值的选项"` // 枚举选项
	HttpEnumConfig *HttpEnumConfig `json:"http_enum_config" gorm:"embedded" bson:"http_enum_config" description:"基于 HTTP 的枚举配置"`                                 // 基于 HTTP 的枚举配置
	Example        string          `json:"example" bson:"example" gorm:"column:example;type:text" description:"值的样例"`                                            // 标签的示例值

	Extras map[string]string `json:"extras" bson:"extras" gorm:"column:extras;type:json;serializer:json" description:"扩展属性"` // 扩展属性

}

type EnumOption struct {
	Label         string            `json:"label" bson:"label"`                           // 选项的说明
	Input         string            `json:"input" bson:"input" validate:"required"`       // 用户输入
	Value         string            `json:"value" bson:"value"`                           // 选项的值 根据 parent.input + children.input 自动生成
	Disabled      bool              `json:"disabled" bson:"disabled"`                     // 是否禁止选中,配合前端 UI 组件使用
	Color         string            `json:"color" bson:"color"`                           // 标签的颜色
	Deprecate     bool              `json:"deprecate" bson:"deprecate"`                   // 是否废弃
	DeprecateDesc string            `json:"deprecate_desc" bson:"deprecate_desc"`         // 废弃说明
	Children      []*EnumOption     `json:"children,omitempty" bson:"children,omitempty"` // 枚举的子选项
	Extensions    map[string]string `json:"extensions" bson:"extensions"`                 // 扩展属性
}

type HttpEnumConfig struct {
	Url        string `json:"url" bson:"url" gorm:"column:http_enum_config_url;type:text" description:"基于枚举的 URL,注意只支持 GET 方法"`                                 // 基于枚举的 URL,注意只支持 GET 方法
	KeyField   string `json:"enum_label_name" bson:"enum_label_name" gorm:"column:http_enum_config_key_filed;type:varchar(100)" description:"Enum Label映射的字段名"` // Enum Label映射的字段名
	ValueFiled string `json:"enum_label_value" bson:"enum_label_value" gorm:"column:http_enum_config_value_filed;type:varchar(100)"`                            // Enum Value 映射的字段名
}

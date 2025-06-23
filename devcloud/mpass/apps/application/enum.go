// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package application

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

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
	TypeValue = map[string]Type{
		"SourceCode":     TypeSourceCode,
		"ContainerImage": TypeContainerImage,
		"Other":          TypeOther,
	}
)

// ParseTypeFromString 将字符串转换为 Type 枚举
func ParseTypeFromString(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TypeValue[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown type %s", str)
	}

	return Type(v), nil
}

func (t *Type) String() string {
	if name, ok := TypeName[*t]; ok {
		return name
	}
	return fmt.Sprintf("Type(%d)", t)
}

func (t *Type) Equal(target Type) bool {
	return t == &target
}

func (t *Type) IsIn(targets ...Type) bool {
	return slices.ContainsFunc(targets, t.Equal)
}

func (t *Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ScmProvider 源代码仓库类型
type ScmProvider string

const (
	ScmProviderGithub    ScmProvider = "github"    // Github
	ScmProviderGitlab    ScmProvider = "gitlab"    // Gitlab
	ScmProviderBitbucket ScmProvider = "bitbucket" // Bitbucket
)

type Language string

const (
	LanguageJava       Language = "java"
	LanguageJavascript Language = "javascript"
	LanguageGolang     Language = "golang"
	LanguagePython     Language = "python"
	LanguagePhp        Language = "php"
	LanguageCSharp     Language = "csharp"
	LanguageC          Language = "c"
	LanguageCPlusPlus  Language = "cpp"
	LanguageSwift      Language = "swift"
	LanguageObjectC    Language = "objectivec"
	LanguageRust       Language = "rust"
	LanguageRuby       Language = "ruby"
	LanguageDart       Language = "dart"
	LanguageKotlin     Language = "kotlin"
	LanguageShell      Language = "shell"
	LanguagePowerShell Language = "powershell"
)

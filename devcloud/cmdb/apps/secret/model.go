// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package secret

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/infraboard/devops/pkg/model"
	"github.com/infraboard/mcube/v2/crypto/cbc"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/cmdb/apps/resource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
)

type Secret struct {
	model.Meta           // 基本属性
	policy.ResourceLabel // 资源范围 Namespace 是继承的,Scope 是 API 添加的
	CreateSecretRequest  // 资源定义
}

func NewSecret(in *CreateSecretRequest) *Secret {
	uid := uuid.NewMD5(uuid.Nil, fmt.Append(nil, "%s.%s.%s", in.Vendor, in.Address, in.ApiKey)).String()
	return &Secret{
		Meta:                *model.NewMeta().SetId(uid),
		CreateSecretRequest: *in,
	}
}

func NewSecretSet() *types.Set[*Secret] {
	return types.New[*Secret]()
}

func (s *Secret) TableName() string {
	return "secrets"
}

func (s *Secret) SetDefault() *Secret {
	if s.SyncLimit == 0 {
		s.SyncLimit = 10
	}
	return s
}

func (s *Secret) String() string {
	return pretty.ToJSON(s)
}

type CreateSecretRequest struct {
	Enabled      *bool           `json:"enabled" gorm:"column:enabled"`                                       // Enabled 是否启用
	Name         string          `json:"name" gorm:"column:name"`                                             // Name 名称
	Vendor       resource.Vendor `json:"vendor" gorm:"column:vendor"`                                         // Vendor 厂商
	Address      string          `json:"address" gorm:"column:address"`                                       // 地址
	ApiKey       string          `json:"api_key" gorm:"column:api_key"`                                       // ApiKey 需要脱敏
	ApiSecret    string          `json:"api_secret" gorm:"column:api_secret"`                                 // ApiSecret
	Regions      []string        `json:"regions" gorm:"column:regions;serializer:json;type:json"`             // Regions 资源所在区域
	ResourceType []resource.Type `json:"resource_type" gorm:"column:resource_type;serializer:json;type:json"` // ResourceType 资源类型
	SyncLimit    int64           `json:"sync_limit" gorm:"column:sync_limit"`                                 // SyncLimit 同步分页大小

	isEncrypted bool `gorm:"-"`
}

func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		Regions:   []string{},
		SyncLimit: 10,
	}
}

func (r *CreateSecretRequest) SetIsEncrypted(v bool) *CreateSecretRequest {
	r.isEncrypted = v
	return r
}

func (r *CreateSecretRequest) SetEnabled(v bool) *CreateSecretRequest {
	r.Enabled = &v
	return r
}

func (r *CreateSecretRequest) GetEnabled() bool {
	if r.Enabled == nil {
		return false
	}
	return *r.Enabled
}

func (r *CreateSecretRequest) GetSyncLimit() int64 {
	if r.SyncLimit == 0 {
		return 10
	}
	return r.SyncLimit
}

func (r *CreateSecretRequest) EncryptedApiSecret() error {
	if r.isEncrypted {
		return nil
	}
	// Hash, 对称，非对称
	// 对称加密 AES(cbc)
	// @v1,xxxx@xxxxx

	key, err := base64.StdEncoding.DecodeString(SecretKey)
	if err != nil {
		return err
	}

	cipherText, err := cbc.MustNewAESCBCCihper(key).Encrypt([]byte(r.ApiSecret))
	if err != nil {
		return err
	}
	r.ApiSecret = base64.StdEncoding.EncodeToString(cipherText)
	r.SetIsEncrypted(true)
	return nil

}

func (r *CreateSecretRequest) DecryptedApiSecret() error {
	if r.isEncrypted {
		cipherdText, err := base64.StdEncoding.DecodeString(r.ApiSecret)
		if err != nil {
			return err
		}

		key, err := base64.StdEncoding.DecodeString(SecretKey)
		if err != nil {
			return err
		}

		plainText, err := cbc.MustNewAESCBCCihper(key).Decrypt([]byte(cipherdText))
		if err != nil {
			return err
		}
		r.ApiSecret = string(plainText)
		r.SetIsEncrypted(false)
	}
	return nil
}

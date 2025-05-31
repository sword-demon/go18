// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/rs/zerolog"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"time"
)

func init() {
	ioc.Controller().Registry(&TokenServiceImpl{})
}

var _ token.Service = (*TokenServiceImpl)(nil)

type TokenServiceImpl struct {
	ioc.ObjectImpl
	user user.Service
	log  *zerolog.Logger

	// 自动刷新
	AutoRefresh       bool   `json:"auto_refresh" toml:"auto_refresh" yaml:"auto_refresh" env:"AUTO_REFRESH"`
	RefreshTTSecond   uint64 `json:"refresh_ttl" toml:"refresh_ttl" yaml:"refresh_ttl" env:"REFRESH_TTL"`
	MaxActiveApiToken uint8  `json:"max_active_api_token" toml:"max_active_api_token" yaml:"max_active_api_token" env:"MAX_ACTIVE_API_TOKEN"`
	refreshDuration   time.Duration
}

func (i *TokenServiceImpl) Init() error {
	// 自动创建表
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&token.Token{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Name 定义托管到 ioc 的名称
func (i *TokenServiceImpl) Name() string {
	return token.AppName
}

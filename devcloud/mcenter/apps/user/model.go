// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package user

import (
	"fmt"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 存储数据库的对象 PO
type User struct {
	apps.ResourceMeta // 固定字段 id 创建时间,删除时间

	// CreateUserRequest 用户传递过来的请求,大部分也都是模型字段
	CreateUserRequest

	// PwdIntensity 密码强度
	PwdIntensity int8 `json:"pwd_intensity" gorm:"column:pwd_intensity;type:tinyint(1);not null" optional:"true"`
}

func NewUser(req *CreateUserRequest) *User {
	req.PasswordHash()
	return &User{
		CreateUserRequest: *req,
		ResourceMeta:      *apps.NewResourceMeta(),
	}
}

func (u *User) String() string {
	return pretty.ToJSON(u)
}

// CheckPassword 判断该用户的密码是否正确
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// TableName orm 负责调用TableName() 来动态获取你这个对象要存储的表的名称
func (u *User) TableName() string {
	return "users"
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Extras: map[string]string{},
	}
}

type CreateUserRequest struct {
	// 账号提供方
	Provider PROVIDER `json:"provider" gorm:"column:provider;type:tinyint(1);not null;index" description:"账号提供方"`
	// 创建方式
	CreateType CreateType `json:"create_type" gorm:"column:create_type;type:tinyint(1);not null;index" optional:"true"`
	// 用户名
	UserName string `json:"user_name" gorm:"column:user_name;type:varchar(100);not null;uniqueIndex" description:"用户名"`
	// 密码(Hash过后的)
	Password string `json:"password" gorm:"column:password;type:varchar(200);not null" description:"用户密码"`
	// 用户描述
	Description string `json:"description" gorm:"column:description;type:varchar(200);not null" description:"用户描述"`
	// 用户类型
	Type TYPE `json:"type" gorm:"column:type;type:varchar(200);not null" description:"用户类型"`
	// 用户描述
	Domain string `json:"domain" gorm:"column:domain;type:varchar(200);" description:"用户所属域"`

	// 支持接口调用
	EnabledApi bool `json:"enabled_api" gorm:"column:enabled_api;type:tinyint(1)" optional:"true" description:"支持接口调用"`
	// 是不是管理员
	IsAdmin bool `json:"is_admin" gorm:"column:is_admin;type:tinyint(1)" optional:"true" description:"是不是管理员"`
	// 用户状态，01:正常，02:冻结
	Locked bool `json:"stat" gorm:"column:stat;type:tinyint(1)" optional:"true" description:"用户状态, 01:正常, 02:冻结"`
	// 激活，1：激活，0：未激活
	Activate bool `json:"activate" gorm:"column:activate;type:tinyint(1)" optional:"true" description:"激活, 1: 激活, 0: 未激活"`
	// 生日
	Birthday *time.Time `json:"birthday" gorm:"column:birthday;type:varchar(200)" optional:"true" description:"生日"`
	// 昵称
	NickName string `json:"nick_name" gorm:"column:nick_name;type:varchar(200)" optional:"true" description:"昵称"`
	// 头像图片
	UserIcon string `json:"user_icon" gorm:"column:user_icon;type:varchar(500)" optional:"true" description:"头像图片"`
	// 性别, 1:男，2:女，0：保密
	Sex SEX `json:"sex" gorm:"column:sex;type:tinyint(1)" optional:"true" description:"性别, 1:男, 2:女, 0: 保密"`

	// 邮箱
	Email string `json:"email" gorm:"column:email;type:varchar(200);index" description:"邮箱" unique:"true"`
	// 邮箱是否验证ok
	IsEmailConfirmed bool `json:"is_email_confirmed" gorm:"column:is_email_confirmed;type:tinyint(1)" optional:"true" description:"邮箱是否验证ok"`
	// 手机
	Mobile string `json:"mobile" gorm:"column:mobile;type:varchar(200);index" optional:"true" description:"手机" unique:"true"`
	// 手机释放验证ok
	IsMobileConfirmed bool `json:"is_mobile_confirmed" gorm:"column:is_mobile_confirmed;type:tinyint(1)" optional:"true" description:"手机释放验证ok"`
	// 手机登录标识
	MobileTGC string `json:"mobile_tgc" gorm:"column:mobile_tgc;type:char(64)" optional:"true" description:"手机登录标识"`
	// 标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" optional:"true" description:"标签"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" optional:"true" description:"其他扩展信息"`

	isHashed bool `json:"-"`
}

func (req *CreateUserRequest) SetIsHashed() {
	req.isHashed = true
}

func (req *CreateUserRequest) Validate() error {
	if req.UserName == "" || req.Password == "" {
		return fmt.Errorf("用户名或者密码需要填写")
	}
	return nil
}

func (req *CreateUserRequest) PasswordHash() {
	if req.isHashed {
		return
	}

	// 默认的强度因子是 10
	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(b)
	req.isHashed = true
}

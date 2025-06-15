// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package application

import "time"

type Application struct {
	Id       string    `json:"id" bson:"id"`
	CreateAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdateBy string    `json:"update_by" bson:"update_by"`
	CreateApplicationRequest
}

type CreateApplicationRequest struct {
	CreateBy  string `json:"create_by" bson:"create_by" description:"创建人"`
	Namespace string `json:"namespace" bson:"namespace" description:"应用命名空间 所属空间名称 从 token 中取出,不需要用户传递"`
	CreateApplicationSpec
}

type CreateApplicationSpec struct {
	Name        string            `json:"name" bson:"name" description:"应用名称"`
	Description string            `json:"description" bson:"description" description:"应用描述"`
	Icon        string            `json:"icon" bson:"icon" description:"应用图标"`
	Type        Type              `json:"type" bson:"type" description:"应用类型, SourceCode, ContainerImage, Other"`
	GitURL      string            `json:"git_url" bson:"git_url" description:"应用代码仓库地址"`
	ImageURL    string            `json:"image_url" bson:"image_url" description:"应用容器镜像地址"`
	Owner       string            `json:"owner" bson:"owner" description:"应用所有者"`
	Level       uint32            `json:"level" bson:"level" description:"应用等级"`
	Priority    uint32            `json:"priority" bson:"priority" description:"应用优先级,用于控制应用启动的先后顺序"`
	Labels      map[string]string `json:"labels" bson:"labels" description:"应用标签,用于标识应用的特性,如:前端,后端,数据库等"`
}

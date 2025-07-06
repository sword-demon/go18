// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package resource

type Vendor string

const (
	VendorAliYun  Vendor = "ali"
	VendorTencent Vendor = "tx"
	VendorHuawei  Vendor = "hw"
	VendorVMWare  Vendor = "vmware"
	VendorAWS     Vendor = "aws"
)

type Type string

const (
	TypeVM        Type = "vm"
	TypeRDS       Type = "rds"
	TypeRedis     Type = "redis"
	TypeBucket    Type = "bucket"
	TypeDisk      Type = "disk"
	TypeLB        Type = "lb"
	TypeDomain    Type = "domain"
	TypeEIP       Type = "eip"
	TypeMongoDB   Type = "mongodb"
	TypeDatabase  Type = "database"
	TypeAccount   Type = "account"
	TypeOtherType Type = "other" // 其他资源
	TypeBillType  Type = "bill"  // 辅助资源
	TypeOrder     Type = "order"
)

type StatusEnum string

# 定义全局 Makefile 变量方便后面引用
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
PROJ_ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# 构建产物、临时文件存放目录
OUTPUT_DIR := $(PROJ_ROOT_DIR)/_output

PKG := "github.com/sword-demon/go18"
MOD_DIR := $(shell go env GOPATH)/pkg/mod
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v redis)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

GIT_TAG := $(shell git describe --tags --exact-match 2>/dev/null || echo "UnTag")
BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "github.com/infraboard/mcube/v2/ioc/config/application"
OUTPUT_NAME := "devcloud-api"

# 定义版本相关变量

## 指定应用使用的 version 包，会通过 `-ldflags -X` 向该包中指定的变量注入值
VERSION_PACKAGE=github.com/sword-demon/go18/devcloud/pkg/version
## 定义 VERSION 语义化版本号
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

## 检查代码仓库是否是 dirty（默认dirty）
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
    GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS := -X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}' -X '${VERSION_PATH}.GIT_TAG=${GIT_TAG}'


# 定义默认目标为 all
.DEFAULT_GOAL := all

# 定义 Makefile all 伪目标，执行 make 时，会默认执行 all 伪目标
.PHONY: all
all: tidy format add-copyright build

.PHONY: run
run: # run devcloud
	@go run main.go start

.PHONY: build
build: tidy # 编译源码,依赖 tidy 目标自动添加/移除依赖包
	@go build -a -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/$(OUTPUT_NAME) $(PROJ_ROOT_DIR)/main.go

.PHONY: format
format: # 格式化 go 源码
	@gofmt -s -w ./


.PHONY: add-copyright
add-copyright: # 添加版权头信息
	@addlicense -v -f $(PROJ_ROOT_DIR)/scripts/boilerplate.txt $(PROJ_ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)


.PHONY: tidy
tidy: # 自动添加、移除依赖
	@go mod tidy


.PHONY: clean
clean: # 清理构建产物,临时文件等
	@-rm -vrf $(OUTPUT_DIR)
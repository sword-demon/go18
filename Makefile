# 定义全局 Makefile 变量方便后面引用

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# 项目根目录
PROJ_ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))

# 构建产物、临时文件存放目录
# OUTPUT_DIR := $(PROJ_ROOT_DIR)/_output

# 定义默认目标为 all
.DEFAULT_GOAL := all

# 定义 Makefile all 伪目标，执行 make 时，会默认执行 all 伪目标
.PHONY: all
all: tidy format add-copyright


.PHONY: format
format: # 格式化 go 源码
	@gofmt -s -w ./


.PHONY: add-copyright
add-copyright: # 添加版权头信息
	@addlicense -v -f $(PROJ_ROOT_DIR)/scripts/boilerplate.txt $(PROJ_ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)


.PHONY: tidy
tidy: # 自动添加、移除依赖
	@go mod tidy
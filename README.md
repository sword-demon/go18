# go18

## 程序设计

1. 需求手机，具体问题具体分析，给出具体的解决方案
2. 问题抽象：建立这类问题的 通用解决模型(程序设计)
3. 架构与实现

## 项目整体介绍

- Book Api Server
- 用户中心
- 应用中心
- 审计中心
- 资源中心(CMDB)
- 发布中心
- 应用流水线 pipline

## 环境 mac/linux

- go version go1.24.0 darwin/arm64
- vscode / goland
- 命令行操作

### 注意事项

- 代码自己写
- 学会排查问题
- 思维转变：程序一次性写完，能正常运行是巧合，一运行就跑错这才是常态，不怕报错，认知查看报错问题或者借助 AI 工具进行帮忙分析
- 持之以恒，每天都写一点，如果时间不过，写一个函数或者少写，便实践的技术
- 项目里面添加自己的文档目录，添加自己的文档和思考，也会查阅周边资料，都是回顾的途径
- 创建一个自己的代码仓库：github/gitee，要 public

## Gin + GORM 开发简单的 Book API Server

从写脚本开始 与 学会合理使用包来组织你的项目工程

### 初始化 mod

```sh
go mod init "github.com/sword-demon/go18"
```

### 安装 gin

```sh
go get github.com/gin-gonic/gin
```

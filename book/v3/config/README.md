# 程序的配置管理

1. 配置的加载
```go
// 用于加载配置
config.LoadConfigFromYaml(yamlConfigFilePath)
```

2. 程序内部如何使用配置

```go
// Get Config -> ConfigObject
config.C().MySQL.Host
```

## 如何验证config包的业务逻辑是否正确

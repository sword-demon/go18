# 技术点

- gin
- gorm
- 配置
- zerolog 打印日志
- 统一响应

## MySQL 环境

```bash
docker run --name mysql8 \                                            130 ↵
  -e MYSQL_ROOT_PASSWORD=admin888 \
  -p 3306:3306 \
  -v /Users/wxvirus/env/mysql8-data:/var/lib/mysql \
  -d mysql:8.0
```

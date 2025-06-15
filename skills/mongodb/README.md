# mongodb

## 使用 docker 安装使用

```bash
docker run -itd -p 27017:27017 --name mongo mongo
```

启动带认证的 mongo

```bash
docker run -d \
  --name mongo \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=secret \
  mongo
```

```bash
docker exec -it mongo mongosh -u admin -p secret
Current Mongosh Log ID:	684e6fdef9b2c208e91b5ff1
Connecting to:		mongodb://<credentials>@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.2
Using MongoDB:		8.0.10
Using Mongosh:		2.5.2

For mongosh info see: https://www.mongodb.com/docs/mongodb-shell/


To help improve our products, anonymous usage data is collected and sent to MongoDB periodically (https://www.mongodb.com/legal/privacy-policy).
You can opt-out by running the disableTelemetry() command.

------
   The server generated these startup warnings when booting
   2025-06-15T07:01:27.278+00:00: Soft rlimits for open file descriptors too low
   2025-06-15T07:01:27.278+00:00: For customers running the current memory allocator, we suggest changing the contents of the following sysfsFile
   2025-06-15T07:01:27.278+00:00: For customers running the current memory allocator, we suggest changing the contents of the following sysfsFile
   2025-06-15T07:01:27.278+00:00: We suggest setting the contents of sysfsFile to 0.
   2025-06-15T07:01:27.278+00:00: We suggest setting swappiness to 0 or 1, as swapping can cause performance problems.
------

test>
```

>判断可用性


```bash
docker logs mongo | grep -i "waiting for connections"
```
# restdoc-api

## 首次部署
1. 创建数据库
``` SQL
CREATE DATABASE restdoc character set UTF8mb4 collate utf8mb4_bin;
```

2. 启动程序 (自动创建数据库表)
``` bash
# 默认配置文件 (config.json)
nohup ./restdoc-api &

# 指定配置文件
nohup ./restdoc-api config.json & 
```

3. 不停服更新
``` bash
kill -USR2 $MASTER_PID
```


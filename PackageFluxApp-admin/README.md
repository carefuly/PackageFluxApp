## PackageFluxApp-admin

### 1. 热更新系统后台

### 2. NaCos配置详情

```yaml
# 服务器配置
server:
  host: "0.0.0.0"
  port: 8080
# 数据库配置
database:
  type: "mysql"
  host: "192.168.66.102"
  port: 3306
  username: "root"
  password: "123456"
  dbname: "package_hot_app"
  charset: "utf8mb4"
  prefix: "package_"
# Redis配置
redis:
  host: "192.168.66.102"
  port: 6379
  password: ""
  db: 0
# Token
token:
  apiKeyAuth: ""
  expire: 60 * 4
# 邮箱配置
email:
  host: "smtp.qq.com"
  port: 25
  username: ""
  password: ""
  from: ""
  to: ""
# 阿里云
aliyun:
  accessKeyID: ""
  accessKeySecret: ""
```


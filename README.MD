# 基于Gin框架开发的博客

### 配置文件
````
[server]
# debug 开发模式，release 生产模式
AppMode = debug
HttpPort = :3000
JwtKey = 32jz2324sd
BaseUrl = "localhost"

[database]
Db = mysql
DbHost = 
DbPort = 
DbUser = 
DbPassWord = 
DbName = 

[qiniu]
AccessKey = 
SecretKey = 
Bucket = 
QiniuSever = 
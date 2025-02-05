# go-zero 实操记录

## 使用 gen.sh 脚本执行

## 目录生成步骤

### 生成 model

1、 通过连接数据库表生成 model

```javascript

goctl model mysql datasource -url="root:123456@tcp(localhost:3306)/gozero" -table=article -dir="./" -cache=true --style=goZero

```

2、通过.sql 文件生成 model

```javascript

goctl model mysql ddl --src user.sql --dir .

```

### 生成 API

1、生成 api 目录代码结构

user.api 表示 你定义的.api 文件名， 在 api/desc 目录下 执行

```javascript
goctl api go -api user.api -dir ../ --style=goZero

```

2、 生成 swagger.json 接口文档

```javascript

goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .

```

### 生成 RPC

1、 生成 proto 文件

```javascript
sql2pb -go_package ../user -host localhost -package user -password 123456 -port 3306 -schema gozero -service_name user -user root > user.proto

```

2、 生成 rpc 目录结构代码

```javascript
goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=../../app/user/rpc --style=goZero

```

### 安装依赖

```javascript
 go mod tidy

```

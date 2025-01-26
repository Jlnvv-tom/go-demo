# 在当前目录下执行命令

## 生成 proto 文件

sql2pb -go_package ./pb -host localhost -package pb -password 123456 -port 3306 -schema gozero -service_name article -user root > article.proto

## 生成 rpc 目录结构代码

goctl rpc protoc article.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero


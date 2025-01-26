## 生成 api 目录代码结构

goctl api go -api main.api -dir ../ --style=goZero

## 生成 swagger.json 接口文档

goctl api plugin -plugin goctl-swagger="swagger -filename main.json" -api main.api -dir .

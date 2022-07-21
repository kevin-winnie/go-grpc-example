//1、编写.proto文件 语法参考proto规则
//2、执行命令，生成 .pd.go 系统文件 再proto目录下执行
protoc --go_out=plugins=grpc:. *.proto
//3、编写service
//4、编写client

//启动服务端
go run ./service/service.go
//启动客户端
go run ./client/client.go

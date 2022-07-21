#GRPC

##1、什么是 RPC
RPC 代指远程过程调用（Remote Procedure Call），它的调用包含了传输协议和编码（对象序列号）协议等等。允许运行于一台计算机的程序调用另一台计算机的子程序，而开发人员无需额外地为这个交互作用编程

实际场景：
有两台服务器，分别是 A、B。在 A 上的应用 C 想要调用 B 服务器上的应用 D，它们可以直接本地调用吗？
答案是不能的，但走 RPC 的话，十分方便。因此常有人称使用 RPC，就跟本地调用一个函数一样简单

##2、Protobuf
Protocol Buffers 是一种与语言、平台无关，可扩展的序列化结构化数据的方法，常用于通信协议，数据存储等等。相较于 JSON、XML，它更小、更快、更简单，因此也更受开发人员的青眯

##3、安装Protoc Plugin和grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

##4、安装protoc
wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-all-3.5.1.zip
unzip protobuf-all-3.5.1.zip
cd protobuf-3.5.1/
./configure
make
make install
检查是否安装成功
protoc --version

##5、grpc程序编写流程
5.1 编写.proto文件 语法参考proto规则
5.2 执行命令，生成 .pd.go 系统文件 再proto目录下执行
protoc --go_out=plugins=grpc:. *.proto
5.3 编写service
5.4 编写client


//启动服务端
go run ./service/service.go
//启动客户端
go run ./client/client.go

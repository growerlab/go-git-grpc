# go-git-grpc

为go-git支持GRPC的能力

#### 目标

- 通过grpc远程调用go-git读取仓库信息
- 通过grpc远程调用git-receive-pack、git-upload-pack命名完成推拉操作

#### 测试

- 初始化 `test/init.sh`
- 执行 `test/test.go`

#### 性能

待测

#### 生成 proto

打开 https://github.com/protocolbuffers/protobuf/releases
下载protoc编译器：protoc-xxx-osx-x86_64.zip
将 bin/protoc 移动 $GOPATH/bin 目录下。

```
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

$ protoc --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src pb/storer.proto --plugin=grpc
```

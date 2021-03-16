# go-git-grpc

为go-git支持GRPC的能力

#### 目标

- 通过grpc通信
- 通过go-git完成各类git操作

#### 性能

待测


#### 生成 proto

```
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

$ protoc --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src pb/storer.proto --plugin=grpc
```

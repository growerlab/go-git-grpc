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

#### EncodedObject 流程

1. client 调用 NewEncodedObject() 从 server 获取 EncodedObject 对象
    - server 为 EncodedObject 注册一个RW IO 
    - 将 RW IO 返回给 client
2. client 对 RW IO 对象做相关的设置，读写操作

RW IO 对象的行为


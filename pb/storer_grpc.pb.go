// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorerClient is the client API for Storer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorerClient interface {
	// EncodedObjectStorer
	NewEncodedObject(ctx context.Context, in *None, opts ...grpc.CallOption) (*UUID, error)
	SetEncodedObject(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Hash, error)
	SetEncodedObjectType(ctx context.Context, in *Int, opts ...grpc.CallOption) (*None, error)
	SetEncodedObjectSetSize(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*None, error)
	EncodedObjectType(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int, error)
	EncodedObjectHash(ctx context.Context, in *None, opts ...grpc.CallOption) (*Hash, error)
	EncodedObjectSize(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int64, error)
	EncodedObjectRWStream(ctx context.Context, opts ...grpc.CallOption) (Storer_EncodedObjectRWStreamClient, error)
	// ReferenceStorer
	SetReference(ctx context.Context, in *Reference, opts ...grpc.CallOption) (*None, error)
	CheckAndSetReference(ctx context.Context, in *SetReferenceParams, opts ...grpc.CallOption) (*None, error)
	GetReference(ctx context.Context, in *ReferenceName, opts ...grpc.CallOption) (*Reference, error)
	GetReferences(ctx context.Context, in *None, opts ...grpc.CallOption) (*References, error)
	RemoveReference(ctx context.Context, in *ReferenceName, opts ...grpc.CallOption) (*None, error)
	CountLooseRefs(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int64, error)
	PackRefs(ctx context.Context, in *None, opts ...grpc.CallOption) (*None, error)
	// ShallowStorer
	SetShallow(ctx context.Context, in *Hashs, opts ...grpc.CallOption) (*None, error)
	Shallow(ctx context.Context, in *None, opts ...grpc.CallOption) (*Hashs, error)
	// IndexStorer
	SetIndex(ctx context.Context, in *Index, opts ...grpc.CallOption) (*None, error)
	GetIndex(ctx context.Context, in *None, opts ...grpc.CallOption) (*Index, error)
	// ConfigStorer
	GetConfig(ctx context.Context, in *None, opts ...grpc.CallOption) (*Config, error)
	SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*None, error)
	// ModuleStorer
	Modules(ctx context.Context, in *None, opts ...grpc.CallOption) (*ModuleNames, error)
}

type storerClient struct {
	cc grpc.ClientConnInterface
}

func NewStorerClient(cc grpc.ClientConnInterface) StorerClient {
	return &storerClient{cc}
}

func (c *storerClient) NewEncodedObject(ctx context.Context, in *None, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/pb.Storer/NewEncodedObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetEncodedObject(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetEncodedObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetEncodedObjectType(ctx context.Context, in *Int, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetEncodedObjectType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetEncodedObjectSetSize(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetEncodedObjectSetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) EncodedObjectType(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int, error) {
	out := new(Int)
	err := c.cc.Invoke(ctx, "/pb.Storer/EncodedObjectType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) EncodedObjectHash(ctx context.Context, in *None, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/pb.Storer/EncodedObjectHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) EncodedObjectSize(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/pb.Storer/EncodedObjectSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) EncodedObjectRWStream(ctx context.Context, opts ...grpc.CallOption) (Storer_EncodedObjectRWStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storer_ServiceDesc.Streams[0], "/pb.Storer/EncodedObjectRWStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &storerEncodedObjectRWStreamClient{stream}
	return x, nil
}

type Storer_EncodedObjectRWStreamClient interface {
	Send(*RWStream) error
	Recv() (*RWStream, error)
	grpc.ClientStream
}

type storerEncodedObjectRWStreamClient struct {
	grpc.ClientStream
}

func (x *storerEncodedObjectRWStreamClient) Send(m *RWStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storerEncodedObjectRWStreamClient) Recv() (*RWStream, error) {
	m := new(RWStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storerClient) SetReference(ctx context.Context, in *Reference, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) CheckAndSetReference(ctx context.Context, in *SetReferenceParams, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/CheckAndSetReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) GetReference(ctx context.Context, in *ReferenceName, opts ...grpc.CallOption) (*Reference, error) {
	out := new(Reference)
	err := c.cc.Invoke(ctx, "/pb.Storer/GetReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) GetReferences(ctx context.Context, in *None, opts ...grpc.CallOption) (*References, error) {
	out := new(References)
	err := c.cc.Invoke(ctx, "/pb.Storer/GetReferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) RemoveReference(ctx context.Context, in *ReferenceName, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/RemoveReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) CountLooseRefs(ctx context.Context, in *None, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/pb.Storer/CountLooseRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) PackRefs(ctx context.Context, in *None, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/PackRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetShallow(ctx context.Context, in *Hashs, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetShallow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) Shallow(ctx context.Context, in *None, opts ...grpc.CallOption) (*Hashs, error) {
	out := new(Hashs)
	err := c.cc.Invoke(ctx, "/pb.Storer/Shallow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetIndex(ctx context.Context, in *Index, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) GetIndex(ctx context.Context, in *None, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/pb.Storer/GetIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) GetConfig(ctx context.Context, in *None, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/pb.Storer/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/pb.Storer/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storerClient) Modules(ctx context.Context, in *None, opts ...grpc.CallOption) (*ModuleNames, error) {
	out := new(ModuleNames)
	err := c.cc.Invoke(ctx, "/pb.Storer/Modules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorerServer is the server API for Storer service.
// All implementations must embed UnimplementedStorerServer
// for forward compatibility
type StorerServer interface {
	// EncodedObjectStorer
	NewEncodedObject(context.Context, *None) (*UUID, error)
	SetEncodedObject(context.Context, *UUID) (*Hash, error)
	SetEncodedObjectType(context.Context, *Int) (*None, error)
	SetEncodedObjectSetSize(context.Context, *Int64) (*None, error)
	EncodedObjectType(context.Context, *None) (*Int, error)
	EncodedObjectHash(context.Context, *None) (*Hash, error)
	EncodedObjectSize(context.Context, *None) (*Int64, error)
	EncodedObjectRWStream(Storer_EncodedObjectRWStreamServer) error
	// ReferenceStorer
	SetReference(context.Context, *Reference) (*None, error)
	CheckAndSetReference(context.Context, *SetReferenceParams) (*None, error)
	GetReference(context.Context, *ReferenceName) (*Reference, error)
	GetReferences(context.Context, *None) (*References, error)
	RemoveReference(context.Context, *ReferenceName) (*None, error)
	CountLooseRefs(context.Context, *None) (*Int64, error)
	PackRefs(context.Context, *None) (*None, error)
	// ShallowStorer
	SetShallow(context.Context, *Hashs) (*None, error)
	Shallow(context.Context, *None) (*Hashs, error)
	// IndexStorer
	SetIndex(context.Context, *Index) (*None, error)
	GetIndex(context.Context, *None) (*Index, error)
	// ConfigStorer
	GetConfig(context.Context, *None) (*Config, error)
	SetConfig(context.Context, *Config) (*None, error)
	// ModuleStorer
	Modules(context.Context, *None) (*ModuleNames, error)
	mustEmbedUnimplementedStorerServer()
}

// UnimplementedStorerServer must be embedded to have forward compatible implementations.
type UnimplementedStorerServer struct {
}

func (UnimplementedStorerServer) NewEncodedObject(context.Context, *None) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEncodedObject not implemented")
}
func (UnimplementedStorerServer) SetEncodedObject(context.Context, *UUID) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEncodedObject not implemented")
}
func (UnimplementedStorerServer) SetEncodedObjectType(context.Context, *Int) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEncodedObjectType not implemented")
}
func (UnimplementedStorerServer) SetEncodedObjectSetSize(context.Context, *Int64) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEncodedObjectSetSize not implemented")
}
func (UnimplementedStorerServer) EncodedObjectType(context.Context, *None) (*Int, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodedObjectType not implemented")
}
func (UnimplementedStorerServer) EncodedObjectHash(context.Context, *None) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodedObjectHash not implemented")
}
func (UnimplementedStorerServer) EncodedObjectSize(context.Context, *None) (*Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodedObjectSize not implemented")
}
func (UnimplementedStorerServer) EncodedObjectRWStream(Storer_EncodedObjectRWStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EncodedObjectRWStream not implemented")
}
func (UnimplementedStorerServer) SetReference(context.Context, *Reference) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReference not implemented")
}
func (UnimplementedStorerServer) CheckAndSetReference(context.Context, *SetReferenceParams) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndSetReference not implemented")
}
func (UnimplementedStorerServer) GetReference(context.Context, *ReferenceName) (*Reference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReference not implemented")
}
func (UnimplementedStorerServer) GetReferences(context.Context, *None) (*References, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferences not implemented")
}
func (UnimplementedStorerServer) RemoveReference(context.Context, *ReferenceName) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReference not implemented")
}
func (UnimplementedStorerServer) CountLooseRefs(context.Context, *None) (*Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountLooseRefs not implemented")
}
func (UnimplementedStorerServer) PackRefs(context.Context, *None) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PackRefs not implemented")
}
func (UnimplementedStorerServer) SetShallow(context.Context, *Hashs) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetShallow not implemented")
}
func (UnimplementedStorerServer) Shallow(context.Context, *None) (*Hashs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shallow not implemented")
}
func (UnimplementedStorerServer) SetIndex(context.Context, *Index) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIndex not implemented")
}
func (UnimplementedStorerServer) GetIndex(context.Context, *None) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndex not implemented")
}
func (UnimplementedStorerServer) GetConfig(context.Context, *None) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedStorerServer) SetConfig(context.Context, *Config) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedStorerServer) Modules(context.Context, *None) (*ModuleNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modules not implemented")
}
func (UnimplementedStorerServer) mustEmbedUnimplementedStorerServer() {}

// UnsafeStorerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorerServer will
// result in compilation errors.
type UnsafeStorerServer interface {
	mustEmbedUnimplementedStorerServer()
}

func RegisterStorerServer(s grpc.ServiceRegistrar, srv StorerServer) {
	s.RegisterService(&Storer_ServiceDesc, srv)
}

func _Storer_NewEncodedObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).NewEncodedObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/NewEncodedObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).NewEncodedObject(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetEncodedObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetEncodedObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetEncodedObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetEncodedObject(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetEncodedObjectType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetEncodedObjectType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetEncodedObjectType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetEncodedObjectType(ctx, req.(*Int))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetEncodedObjectSetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetEncodedObjectSetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetEncodedObjectSetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetEncodedObjectSetSize(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_EncodedObjectType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).EncodedObjectType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/EncodedObjectType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).EncodedObjectType(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_EncodedObjectHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).EncodedObjectHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/EncodedObjectHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).EncodedObjectHash(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_EncodedObjectSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).EncodedObjectSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/EncodedObjectSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).EncodedObjectSize(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_EncodedObjectRWStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorerServer).EncodedObjectRWStream(&storerEncodedObjectRWStreamServer{stream})
}

type Storer_EncodedObjectRWStreamServer interface {
	Send(*RWStream) error
	Recv() (*RWStream, error)
	grpc.ServerStream
}

type storerEncodedObjectRWStreamServer struct {
	grpc.ServerStream
}

func (x *storerEncodedObjectRWStreamServer) Send(m *RWStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storerEncodedObjectRWStreamServer) Recv() (*RWStream, error) {
	m := new(RWStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storer_SetReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetReference(ctx, req.(*Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_CheckAndSetReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReferenceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).CheckAndSetReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/CheckAndSetReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).CheckAndSetReference(ctx, req.(*SetReferenceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_GetReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferenceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).GetReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/GetReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).GetReference(ctx, req.(*ReferenceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_GetReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).GetReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/GetReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).GetReferences(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_RemoveReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferenceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).RemoveReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/RemoveReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).RemoveReference(ctx, req.(*ReferenceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_CountLooseRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).CountLooseRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/CountLooseRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).CountLooseRefs(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_PackRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).PackRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/PackRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).PackRefs(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetShallow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hashs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetShallow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetShallow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetShallow(ctx, req.(*Hashs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_Shallow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).Shallow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/Shallow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).Shallow(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetIndex(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/GetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).GetIndex(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).GetConfig(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).SetConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storer_Modules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).Modules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storer/Modules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).Modules(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

// Storer_ServiceDesc is the grpc.ServiceDesc for Storer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Storer",
	HandlerType: (*StorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewEncodedObject",
			Handler:    _Storer_NewEncodedObject_Handler,
		},
		{
			MethodName: "SetEncodedObject",
			Handler:    _Storer_SetEncodedObject_Handler,
		},
		{
			MethodName: "SetEncodedObjectType",
			Handler:    _Storer_SetEncodedObjectType_Handler,
		},
		{
			MethodName: "SetEncodedObjectSetSize",
			Handler:    _Storer_SetEncodedObjectSetSize_Handler,
		},
		{
			MethodName: "EncodedObjectType",
			Handler:    _Storer_EncodedObjectType_Handler,
		},
		{
			MethodName: "EncodedObjectHash",
			Handler:    _Storer_EncodedObjectHash_Handler,
		},
		{
			MethodName: "EncodedObjectSize",
			Handler:    _Storer_EncodedObjectSize_Handler,
		},
		{
			MethodName: "SetReference",
			Handler:    _Storer_SetReference_Handler,
		},
		{
			MethodName: "CheckAndSetReference",
			Handler:    _Storer_CheckAndSetReference_Handler,
		},
		{
			MethodName: "GetReference",
			Handler:    _Storer_GetReference_Handler,
		},
		{
			MethodName: "GetReferences",
			Handler:    _Storer_GetReferences_Handler,
		},
		{
			MethodName: "RemoveReference",
			Handler:    _Storer_RemoveReference_Handler,
		},
		{
			MethodName: "CountLooseRefs",
			Handler:    _Storer_CountLooseRefs_Handler,
		},
		{
			MethodName: "PackRefs",
			Handler:    _Storer_PackRefs_Handler,
		},
		{
			MethodName: "SetShallow",
			Handler:    _Storer_SetShallow_Handler,
		},
		{
			MethodName: "Shallow",
			Handler:    _Storer_Shallow_Handler,
		},
		{
			MethodName: "SetIndex",
			Handler:    _Storer_SetIndex_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _Storer_GetIndex_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Storer_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Storer_SetConfig_Handler,
		},
		{
			MethodName: "Modules",
			Handler:    _Storer_Modules_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EncodedObjectRWStream",
			Handler:       _Storer_EncodedObjectRWStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/storer.proto",
}

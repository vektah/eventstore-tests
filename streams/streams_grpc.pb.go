// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package streams

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamsClient is the client API for Streams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamsClient interface {
	Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (Streams_ReadClient, error)
	Append(ctx context.Context, opts ...grpc.CallOption) (Streams_AppendClient, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
	Tombstone(ctx context.Context, in *TombstoneReq, opts ...grpc.CallOption) (*TombstoneResp, error)
}

type streamsClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamsClient(cc grpc.ClientConnInterface) StreamsClient {
	return &streamsClient{cc}
}

func (c *streamsClient) Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (Streams_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streams_serviceDesc.Streams[0], "/event_store.client.streams.Streams/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamsReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streams_ReadClient interface {
	Recv() (*ReadResp, error)
	grpc.ClientStream
}

type streamsReadClient struct {
	grpc.ClientStream
}

func (x *streamsReadClient) Recv() (*ReadResp, error) {
	m := new(ReadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamsClient) Append(ctx context.Context, opts ...grpc.CallOption) (Streams_AppendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streams_serviceDesc.Streams[1], "/event_store.client.streams.Streams/Append", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamsAppendClient{stream}
	return x, nil
}

type Streams_AppendClient interface {
	Send(*AppendReq) error
	CloseAndRecv() (*AppendResp, error)
	grpc.ClientStream
}

type streamsAppendClient struct {
	grpc.ClientStream
}

func (x *streamsAppendClient) Send(m *AppendReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamsAppendClient) CloseAndRecv() (*AppendResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AppendResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamsClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/event_store.client.streams.Streams/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamsClient) Tombstone(ctx context.Context, in *TombstoneReq, opts ...grpc.CallOption) (*TombstoneResp, error) {
	out := new(TombstoneResp)
	err := c.cc.Invoke(ctx, "/event_store.client.streams.Streams/Tombstone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamsServer is the server API for Streams service.
// All implementations must embed UnimplementedStreamsServer
// for forward compatibility
type StreamsServer interface {
	Read(*ReadReq, Streams_ReadServer) error
	Append(Streams_AppendServer) error
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
	Tombstone(context.Context, *TombstoneReq) (*TombstoneResp, error)
	mustEmbedUnimplementedStreamsServer()
}

// UnimplementedStreamsServer must be embedded to have forward compatible implementations.
type UnimplementedStreamsServer struct {
}

func (*UnimplementedStreamsServer) Read(*ReadReq, Streams_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStreamsServer) Append(Streams_AppendServer) error {
	return status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (*UnimplementedStreamsServer) Delete(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedStreamsServer) Tombstone(context.Context, *TombstoneReq) (*TombstoneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tombstone not implemented")
}
func (*UnimplementedStreamsServer) mustEmbedUnimplementedStreamsServer() {}

func RegisterStreamsServer(s *grpc.Server, srv StreamsServer) {
	s.RegisterService(&_Streams_serviceDesc, srv)
}

func _Streams_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamsServer).Read(m, &streamsReadServer{stream})
}

type Streams_ReadServer interface {
	Send(*ReadResp) error
	grpc.ServerStream
}

type streamsReadServer struct {
	grpc.ServerStream
}

func (x *streamsReadServer) Send(m *ReadResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Streams_Append_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamsServer).Append(&streamsAppendServer{stream})
}

type Streams_AppendServer interface {
	SendAndClose(*AppendResp) error
	Recv() (*AppendReq, error)
	grpc.ServerStream
}

type streamsAppendServer struct {
	grpc.ServerStream
}

func (x *streamsAppendServer) SendAndClose(m *AppendResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamsAppendServer) Recv() (*AppendReq, error) {
	m := new(AppendReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streams_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_store.client.streams.Streams/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Streams_Tombstone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TombstoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServer).Tombstone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_store.client.streams.Streams/Tombstone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServer).Tombstone(ctx, req.(*TombstoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Streams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event_store.client.streams.Streams",
	HandlerType: (*StreamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _Streams_Delete_Handler,
		},
		{
			MethodName: "Tombstone",
			Handler:    _Streams_Tombstone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _Streams_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Append",
			Handler:       _Streams_Append_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "streams.proto",
}
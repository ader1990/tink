// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/target/target.proto

package target

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{0}
}

func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UpdateRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type UUID struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{3}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Targets struct {
	JSON                 string   `protobuf:"bytes,1,opt,name=JSON,proto3" json:"JSON,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Targets) Reset()         { *m = Targets{} }
func (m *Targets) String() string { return proto.CompactTextString(m) }
func (*Targets) ProtoMessage()    {}
func (*Targets) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{5}
}

func (m *Targets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Targets.Unmarshal(m, b)
}
func (m *Targets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Targets.Marshal(b, m, deterministic)
}
func (m *Targets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Targets.Merge(m, src)
}
func (m *Targets) XXX_Size() int {
	return xxx_messageInfo_Targets.Size(m)
}
func (m *Targets) XXX_DiscardUnknown() {
	xxx_messageInfo_Targets.DiscardUnknown(m)
}

var xxx_messageInfo_Targets proto.InternalMessageInfo

func (m *Targets) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

type TargetList struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetList) Reset()         { *m = TargetList{} }
func (m *TargetList) String() string { return proto.CompactTextString(m) }
func (*TargetList) ProtoMessage()    {}
func (*TargetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9a68d06d9471ac, []int{6}
}

func (m *TargetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetList.Unmarshal(m, b)
}
func (m *TargetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetList.Marshal(b, m, deterministic)
}
func (m *TargetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetList.Merge(m, src)
}
func (m *TargetList) XXX_Size() int {
	return xxx_messageInfo_TargetList.Size(m)
}
func (m *TargetList) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetList.DiscardUnknown(m)
}

var xxx_messageInfo_TargetList proto.InternalMessageInfo

func (m *TargetList) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TargetList) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "github.com.packethost.tinkerbell.protos.target.PushRequest")
	proto.RegisterType((*GetRequest)(nil), "github.com.packethost.tinkerbell.protos.target.GetRequest")
	proto.RegisterType((*UpdateRequest)(nil), "github.com.packethost.tinkerbell.protos.target.UpdateRequest")
	proto.RegisterType((*UUID)(nil), "github.com.packethost.tinkerbell.protos.target.UUID")
	proto.RegisterType((*Empty)(nil), "github.com.packethost.tinkerbell.protos.target.Empty")
	proto.RegisterType((*Targets)(nil), "github.com.packethost.tinkerbell.protos.target.Targets")
	proto.RegisterType((*TargetList)(nil), "github.com.packethost.tinkerbell.protos.target.TargetList")
}

func init() { proto.RegisterFile("protos/target/target.proto", fileDescriptor_1c9a68d06d9471ac) }

var fileDescriptor_1c9a68d06d9471ac = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xa5, 0xa5, 0xb6, 0x38, 0xa5, 0x52, 0xf6, 0x24, 0x8b, 0x82, 0xee, 0x49, 0x11, 0x36, 0xa5,
	0xa5, 0xe2, 0xb9, 0x46, 0x24, 0x22, 0x2a, 0xd5, 0x5c, 0xbc, 0x6d, 0x9b, 0xa5, 0x09, 0x5a, 0x12,
	0xb3, 0x13, 0x6d, 0x11, 0xfc, 0x68, 0xbf, 0x40, 0xb2, 0xdb, 0xb5, 0xe9, 0x41, 0xe8, 0x7a, 0xca,
	0x24, 0xf3, 0xde, 0xbc, 0x99, 0x79, 0x19, 0xa0, 0x59, 0x9e, 0x62, 0xaa, 0x3c, 0x14, 0xf9, 0x4c,
	0xe2, 0xea, 0xc1, 0xf5, 0x47, 0x72, 0x3a, 0x4b, 0x30, 0x2e, 0x26, 0x7c, 0x9a, 0xce, 0x79, 0x26,
	0xa6, 0x2f, 0x12, 0xe3, 0x54, 0x21, 0xcf, 0xd3, 0x77, 0x99, 0x1b, 0x88, 0xe2, 0x86, 0xc0, 0x8e,
	0xa1, 0xfd, 0x50, 0xa8, 0x78, 0x2c, 0xdf, 0x0a, 0xa9, 0x90, 0x10, 0x68, 0x44, 0x02, 0xc5, 0x7e,
	0xed, 0xa8, 0x76, 0xb2, 0x3b, 0xd6, 0x31, 0x3b, 0x00, 0xb8, 0x96, 0x68, 0x11, 0x7b, 0x50, 0x0f,
	0xfc, 0x55, 0xbe, 0x1e, 0xf8, 0x6c, 0x00, 0x9d, 0x30, 0x8b, 0x04, 0xca, 0x3f, 0x00, 0xbf, 0x25,
	0xeb, 0x95, 0x92, 0x14, 0x1a, 0x61, 0x68, 0x72, 0x45, 0x91, 0x44, 0x56, 0xae, 0x8c, 0x59, 0x0b,
	0x76, 0xae, 0xe6, 0x19, 0x2e, 0xd9, 0x21, 0xb4, 0x9e, 0x74, 0x93, 0xaa, 0xc4, 0xdd, 0x3c, 0xde,
	0xdf, 0x59, 0x5c, 0x19, 0xb3, 0x1e, 0x80, 0x49, 0xdf, 0x26, 0xdb, 0xa9, 0xf6, 0xbf, 0x1b, 0xd0,
	0x34, 0x14, 0xb2, 0x80, 0xce, 0x65, 0x2e, 0x05, 0x4a, 0xab, 0x70, 0xce, 0xb7, 0xde, 0x19, 0xaf,
	0x2c, 0x8c, 0x7a, 0x0e, 0x3c, 0x3d, 0xf2, 0x87, 0x6d, 0x7b, 0xb4, 0x0c, 0x7c, 0x32, 0x74, 0xa0,
	0xaf, 0x4d, 0xa0, 0x7d, 0x07, 0x9a, 0x9d, 0xf0, 0x0b, 0xba, 0xc6, 0xa8, 0x8a, 0xfc, 0x85, 0x4b,
	0xf7, 0x55, 0x97, 0x69, 0xcf, 0x81, 0xa9, 0xed, 0x24, 0x9f, 0xd0, 0xf5, 0xe5, 0xab, 0xdc, 0xd0,
	0xff, 0xe7, 0xf8, 0xee, 0xe2, 0x0b, 0x68, 0x97, 0xbf, 0x89, 0xdd, 0x85, 0x73, 0x01, 0x3a, 0x74,
	0xde, 0x78, 0xa9, 0xd7, 0xab, 0x8d, 0xce, 0x9e, 0x2b, 0xd7, 0xe8, 0xad, 0x99, 0x9e, 0x66, 0x7a,
	0x1b, 0x57, 0x3c, 0x69, 0xea, 0xd7, 0xc1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x72, 0x01,
	0x7d, 0xdd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TargetClient is the client API for Target service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TargetClient interface {
	CreateTargets(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*UUID, error)
	TargetByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Targets, error)
	UpdateTargetByID(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteTargetByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error)
	ListTargets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Target_ListTargetsClient, error)
}

type targetClient struct {
	cc *grpc.ClientConn
}

func NewTargetClient(cc *grpc.ClientConn) TargetClient {
	return &targetClient{cc}
}

func (c *targetClient) CreateTargets(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/github.com.packethost.tinkerbell.protos.target.Target/CreateTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetClient) TargetByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Targets, error) {
	out := new(Targets)
	err := c.cc.Invoke(ctx, "/github.com.packethost.tinkerbell.protos.target.Target/TargetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetClient) UpdateTargetByID(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.packethost.tinkerbell.protos.target.Target/UpdateTargetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetClient) DeleteTargetByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.packethost.tinkerbell.protos.target.Target/DeleteTargetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetClient) ListTargets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Target_ListTargetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Target_serviceDesc.Streams[0], "/github.com.packethost.tinkerbell.protos.target.Target/ListTargets", opts...)
	if err != nil {
		return nil, err
	}
	x := &targetListTargetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Target_ListTargetsClient interface {
	Recv() (*TargetList, error)
	grpc.ClientStream
}

type targetListTargetsClient struct {
	grpc.ClientStream
}

func (x *targetListTargetsClient) Recv() (*TargetList, error) {
	m := new(TargetList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TargetServer is the server API for Target service.
type TargetServer interface {
	CreateTargets(context.Context, *PushRequest) (*UUID, error)
	TargetByID(context.Context, *GetRequest) (*Targets, error)
	UpdateTargetByID(context.Context, *UpdateRequest) (*Empty, error)
	DeleteTargetByID(context.Context, *GetRequest) (*Empty, error)
	ListTargets(*Empty, Target_ListTargetsServer) error
}

// UnimplementedTargetServer can be embedded to have forward compatible implementations.
type UnimplementedTargetServer struct {
}

func (*UnimplementedTargetServer) CreateTargets(ctx context.Context, req *PushRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTargets not implemented")
}
func (*UnimplementedTargetServer) TargetByID(ctx context.Context, req *GetRequest) (*Targets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TargetByID not implemented")
}
func (*UnimplementedTargetServer) UpdateTargetByID(ctx context.Context, req *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTargetByID not implemented")
}
func (*UnimplementedTargetServer) DeleteTargetByID(ctx context.Context, req *GetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTargetByID not implemented")
}
func (*UnimplementedTargetServer) ListTargets(req *Empty, srv Target_ListTargetsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTargets not implemented")
}

func RegisterTargetServer(s *grpc.Server, srv TargetServer) {
	s.RegisterService(&_Target_serviceDesc, srv)
}

func _Target_CreateTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServer).CreateTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.packethost.tinkerbell.protos.target.Target/CreateTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServer).CreateTargets(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Target_TargetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServer).TargetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.packethost.tinkerbell.protos.target.Target/TargetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServer).TargetByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Target_UpdateTargetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServer).UpdateTargetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.packethost.tinkerbell.protos.target.Target/UpdateTargetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServer).UpdateTargetByID(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Target_DeleteTargetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServer).DeleteTargetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.packethost.tinkerbell.protos.target.Target/DeleteTargetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServer).DeleteTargetByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Target_ListTargets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TargetServer).ListTargets(m, &targetListTargetsServer{stream})
}

type Target_ListTargetsServer interface {
	Send(*TargetList) error
	grpc.ServerStream
}

type targetListTargetsServer struct {
	grpc.ServerStream
}

func (x *targetListTargetsServer) Send(m *TargetList) error {
	return x.ServerStream.SendMsg(m)
}

var _Target_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.packethost.tinkerbell.protos.target.Target",
	HandlerType: (*TargetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTargets",
			Handler:    _Target_CreateTargets_Handler,
		},
		{
			MethodName: "TargetByID",
			Handler:    _Target_TargetByID_Handler,
		},
		{
			MethodName: "UpdateTargetByID",
			Handler:    _Target_UpdateTargetByID_Handler,
		},
		{
			MethodName: "DeleteTargetByID",
			Handler:    _Target_DeleteTargetByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTargets",
			Handler:       _Target_ListTargets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/target/target.proto",
}

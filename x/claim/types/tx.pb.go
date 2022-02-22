// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgInitialClaim struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgInitialClaim) Reset()         { *m = MsgInitialClaim{} }
func (m *MsgInitialClaim) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaim) ProtoMessage()    {}
func (*MsgInitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{0}
}
func (m *MsgInitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaim.Merge(m, src)
}
func (m *MsgInitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaim proto.InternalMessageInfo

func (m *MsgInitialClaim) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgInitialClaimResponse struct {
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgInitialClaimResponse) Reset()         { *m = MsgInitialClaimResponse{} }
func (m *MsgInitialClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaimResponse) ProtoMessage()    {}
func (*MsgInitialClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{1}
}
func (m *MsgInitialClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaimResponse.Merge(m, src)
}
func (m *MsgInitialClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaimResponse proto.InternalMessageInfo

func (m *MsgInitialClaimResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

type MsgClaimFor struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action  Action `protobuf:"varint,3,opt,name=action,proto3,enum=publicawesome.stargaze.claim.v1beta1.Action" json:"action,omitempty"`
}

func (m *MsgClaimFor) Reset()         { *m = MsgClaimFor{} }
func (m *MsgClaimFor) String() string { return proto.CompactTextString(m) }
func (*MsgClaimFor) ProtoMessage()    {}
func (*MsgClaimFor) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{2}
}
func (m *MsgClaimFor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimFor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimFor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimFor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimFor.Merge(m, src)
}
func (m *MsgClaimFor) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimFor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimFor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimFor proto.InternalMessageInfo

func (m *MsgClaimFor) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgClaimFor) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimFor) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionInitialClaim
}

type MsgClaimForResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgClaimForResponse) Reset()         { *m = MsgClaimForResponse{} }
func (m *MsgClaimForResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimForResponse) ProtoMessage()    {}
func (*MsgClaimForResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{3}
}
func (m *MsgClaimForResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimForResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimForResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimForResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimForResponse.Merge(m, src)
}
func (m *MsgClaimForResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimForResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimForResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimForResponse proto.InternalMessageInfo

func (m *MsgClaimForResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimForResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInitialClaim)(nil), "publicawesome.stargaze.claim.v1beta1.MsgInitialClaim")
	proto.RegisterType((*MsgInitialClaimResponse)(nil), "publicawesome.stargaze.claim.v1beta1.MsgInitialClaimResponse")
	proto.RegisterType((*MsgClaimFor)(nil), "publicawesome.stargaze.claim.v1beta1.MsgClaimFor")
	proto.RegisterType((*MsgClaimForResponse)(nil), "publicawesome.stargaze.claim.v1beta1.MsgClaimForResponse")
}

func init() { proto.RegisterFile("stargaze/claim/v1beta1/tx.proto", fileDescriptor_9ee4a19153cf6635) }

var fileDescriptor_9ee4a19153cf6635 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x89, 0x14, 0x60, 0x0f, 0x0e, 0xc9, 0xfc, 0x33, 0x29, 0xec, 0xc8, 0xa2, 0xf0,
	0x49, 0xdc, 0xae, 0x92, 0x88, 0x02, 0x24, 0x8a, 0xbb, 0x43, 0x48, 0x57, 0xb8, 0x71, 0x49, 0x73,
	0x5a, 0xdb, 0x2b, 0xb3, 0xc2, 0xf6, 0x5a, 0x9e, 0xcd, 0x91, 0xa3, 0x06, 0x1a, 0x1a, 0xde, 0x02,
	0x89, 0x77, 0xa0, 0xbf, 0xf2, 0x4a, 0xaa, 0x03, 0x25, 0x6f, 0xc0, 0x13, 0xa0, 0xac, 0xd7, 0x56,
	0x72, 0x52, 0xa4, 0x40, 0x75, 0x95, 0xbd, 0x9a, 0xf9, 0x7d, 0x33, 0xf3, 0xed, 0x0e, 0x76, 0x41,
	0xb1, 0x2a, 0x65, 0x1f, 0x38, 0x8d, 0x33, 0x26, 0x72, 0x7a, 0x3a, 0x8a, 0xb8, 0x62, 0x23, 0xaa,
	0x66, 0xa4, 0xac, 0xa4, 0x92, 0xd6, 0x93, 0x72, 0x1a, 0x65, 0x22, 0x66, 0xef, 0x39, 0xc8, 0x9c,
	0x93, 0x26, 0x9d, 0xe8, 0x74, 0x62, 0xd2, 0x07, 0xf7, 0x53, 0x99, 0x4a, 0x0d, 0xd0, 0xe5, 0x5f,
	0xcd, 0x0e, 0x9c, 0x58, 0x42, 0x2e, 0x81, 0x46, 0x0c, 0x78, 0xab, 0x1c, 0x4b, 0x51, 0x98, 0xf8,
	0xde, 0x86, 0xe2, 0xfa, 0x74, 0x52, 0xf1, 0x58, 0x56, 0x49, 0x9d, 0xea, 0xed, 0xe1, 0xbb, 0x01,
	0xa4, 0xc7, 0x85, 0x50, 0x82, 0x65, 0x47, 0xcb, 0xb8, 0xf5, 0x10, 0xf7, 0x81, 0x17, 0x09, 0xaf,
	0x6c, 0x34, 0x44, 0xfe, 0xad, 0xd0, 0x9c, 0xbc, 0x6f, 0x08, 0x3f, 0xba, 0x92, 0x1b, 0x72, 0x28,
	0x65, 0x01, 0xdc, 0xfa, 0x82, 0xf0, 0xae, 0x56, 0xe7, 0xc9, 0x09, 0xcb, 0xe5, 0xb4, 0x50, 0x76,
	0x77, 0xd8, 0xf3, 0x77, 0xc6, 0x8f, 0x49, 0xdd, 0x2b, 0x59, 0xf6, 0xda, 0x8c, 0x45, 0x8e, 0xa4,
	0x28, 0x0e, 0x8f, 0xcf, 0x2f, 0xdd, 0xce, 0x9f, 0x4b, 0xf7, 0xc1, 0x19, 0xcb, 0xb3, 0x17, 0xde,
	0x3a, 0xee, 0x7d, 0xff, 0xe5, 0xfa, 0xa9, 0x50, 0x6f, 0xa7, 0x11, 0x89, 0x65, 0x4e, 0xcd, 0xc4,
	0xf5, 0x67, 0x1f, 0x92, 0x77, 0x54, 0x9d, 0x95, 0x1c, 0xb4, 0x12, 0x84, 0x77, 0x0c, 0x7c, 0x50,
	0xb3, 0x9f, 0x10, 0xde, 0x09, 0x20, 0xd5, 0x2d, 0xbe, 0x96, 0xd5, 0xa6, 0x89, 0x2c, 0x1b, 0xdf,
	0x60, 0x49, 0x52, 0x71, 0x00, 0xbb, 0xab, 0x03, 0xcd, 0xd1, 0x7a, 0x85, 0xfb, 0x2c, 0x56, 0x42,
	0x16, 0x76, 0x6f, 0x88, 0xfc, 0xdd, 0xf1, 0x53, 0xb2, 0xcd, 0x75, 0x91, 0x03, 0xcd, 0x84, 0x86,
	0xf5, 0x7e, 0x20, 0x7c, 0x6f, 0xa5, 0x8f, 0xd6, 0xad, 0x95, 0xba, 0x68, 0xbd, 0xee, 0xb5, 0xf2,
	0x71, 0xfc, 0xb9, 0x8b, 0x7b, 0x01, 0xa4, 0xd6, 0x47, 0x84, 0x6f, 0xaf, 0x3d, 0x91, 0x67, 0xdb,
	0xd9, 0x71, 0xe5, 0xb5, 0x0c, 0x5e, 0xfe, 0x17, 0xd6, 0xda, 0x36, 0xc3, 0x37, 0xdb, 0x2b, 0x1d,
	0x6d, 0x2d, 0xd5, 0x20, 0x83, 0xe7, 0xff, 0x8c, 0x34, 0x95, 0x0f, 0x83, 0xf3, 0xb9, 0x83, 0x2e,
	0xe6, 0x0e, 0xfa, 0x3d, 0x77, 0xd0, 0xd7, 0x85, 0xd3, 0xb9, 0x58, 0x38, 0x9d, 0x9f, 0x0b, 0xa7,
	0xf3, 0x66, 0xb2, 0xe2, 0x6d, 0x2d, 0xbf, 0x6f, 0xf4, 0x69, 0xbb, 0x84, 0xa7, 0x13, 0x3a, 0x33,
	0x9b, 0xa8, 0xcd, 0x8e, 0xfa, 0x7a, 0xf7, 0x26, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x58, 0xe3,
	0x1a, 0xbb, 0x25, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error) {
	out := new(MsgInitialClaimResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.claim.v1beta1.Msg/InitialClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error) {
	out := new(MsgClaimForResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.claim.v1beta1.Msg/ClaimFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InitialClaim(context.Context, *MsgInitialClaim) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	ClaimFor(context.Context, *MsgClaimFor) (*MsgClaimForResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InitialClaim(ctx context.Context, req *MsgInitialClaim) (*MsgInitialClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialClaim not implemented")
}
func (*UnimplementedMsgServer) ClaimFor(ctx context.Context, req *MsgClaimFor) (*MsgClaimForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimFor not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InitialClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitialClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitialClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.claim.v1beta1.Msg/InitialClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitialClaim(ctx, req.(*MsgInitialClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimFor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.claim.v1beta1.Msg/ClaimFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimFor(ctx, req.(*MsgClaimFor))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.claim.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitialClaim",
			Handler:    _Msg_InitialClaim_Handler,
		},
		{
			MethodName: "ClaimFor",
			Handler:    _Msg_ClaimFor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargaze/claim/v1beta1/tx.proto",
}

func (m *MsgInitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitialClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimFor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimFor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimFor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimForResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimForResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimForResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitialClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgClaimFor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovTx(uint64(m.Action))
	}
	return n
}

func (m *MsgClaimForResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitialClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgInitialClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitialClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClaimFor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClaimFor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimFor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClaimForResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClaimForResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimForResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

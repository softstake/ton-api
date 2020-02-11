// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type FetchTransactionsRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Lt                   int64    `protobuf:"varint,2,opt,name=lt,proto3" json:"lt,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchTransactionsRequest) Reset()         { *m = FetchTransactionsRequest{} }
func (m *FetchTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*FetchTransactionsRequest) ProtoMessage()    {}
func (*FetchTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *FetchTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchTransactionsRequest.Unmarshal(m, b)
}
func (m *FetchTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *FetchTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTransactionsRequest.Merge(m, src)
}
func (m *FetchTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_FetchTransactionsRequest.Size(m)
}
func (m *FetchTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTransactionsRequest proto.InternalMessageInfo

func (m *FetchTransactionsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FetchTransactionsRequest) GetLt() int64 {
	if m != nil {
		return m.Lt
	}
	return 0
}

func (m *FetchTransactionsRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type FetchTransactionsResponse struct {
	Items                []*Transaction `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchTransactionsResponse) Reset()         { *m = FetchTransactionsResponse{} }
func (m *FetchTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*FetchTransactionsResponse) ProtoMessage()    {}
func (*FetchTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *FetchTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchTransactionsResponse.Unmarshal(m, b)
}
func (m *FetchTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *FetchTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTransactionsResponse.Merge(m, src)
}
func (m *FetchTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_FetchTransactionsResponse.Size(m)
}
func (m *FetchTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTransactionsResponse proto.InternalMessageInfo

func (m *FetchTransactionsResponse) GetItems() []*Transaction {
	if m != nil {
		return m.Items
	}
	return nil
}

type Transaction struct {
	Data                 string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Fee                  int64                  `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
	InMsg                *RawMessage            `protobuf:"bytes,3,opt,name=in_msg,json=inMsg,proto3" json:"in_msg,omitempty"`
	OtherFee             int64                  `protobuf:"varint,4,opt,name=other_fee,json=otherFee,proto3" json:"other_fee,omitempty"`
	OutMsgs              []*RawMessage          `protobuf:"bytes,5,rep,name=out_msgs,json=outMsgs,proto3" json:"out_msgs,omitempty"`
	StorageFee           int64                  `protobuf:"varint,6,opt,name=storage_fee,json=storageFee,proto3" json:"storage_fee,omitempty"`
	TransactionId        *InternalTransactionId `protobuf:"bytes,7,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Transaction) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Transaction) GetInMsg() *RawMessage {
	if m != nil {
		return m.InMsg
	}
	return nil
}

func (m *Transaction) GetOtherFee() int64 {
	if m != nil {
		return m.OtherFee
	}
	return 0
}

func (m *Transaction) GetOutMsgs() []*RawMessage {
	if m != nil {
		return m.OutMsgs
	}
	return nil
}

func (m *Transaction) GetStorageFee() int64 {
	if m != nil {
		return m.StorageFee
	}
	return 0
}

func (m *Transaction) GetTransactionId() *InternalTransactionId {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

type RawMessage struct {
	BodyHash             string   `protobuf:"bytes,1,opt,name=body_hash,json=bodyHash,proto3" json:"body_hash,omitempty"`
	CreatedLt            int64    `protobuf:"varint,2,opt,name=created_lt,json=createdLt,proto3" json:"created_lt,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	FwdFee               int64    `protobuf:"varint,4,opt,name=fwd_fee,json=fwdFee,proto3" json:"fwd_fee,omitempty"`
	IhrFee               int64    `protobuf:"varint,5,opt,name=ihr_fee,json=ihrFee,proto3" json:"ihr_fee,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Source               string   `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Value                int64    `protobuf:"varint,8,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawMessage) Reset()         { *m = RawMessage{} }
func (m *RawMessage) String() string { return proto.CompactTextString(m) }
func (*RawMessage) ProtoMessage()    {}
func (*RawMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *RawMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawMessage.Unmarshal(m, b)
}
func (m *RawMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawMessage.Marshal(b, m, deterministic)
}
func (m *RawMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawMessage.Merge(m, src)
}
func (m *RawMessage) XXX_Size() int {
	return xxx_messageInfo_RawMessage.Size(m)
}
func (m *RawMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RawMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RawMessage proto.InternalMessageInfo

func (m *RawMessage) GetBodyHash() string {
	if m != nil {
		return m.BodyHash
	}
	return ""
}

func (m *RawMessage) GetCreatedLt() int64 {
	if m != nil {
		return m.CreatedLt
	}
	return 0
}

func (m *RawMessage) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *RawMessage) GetFwdFee() int64 {
	if m != nil {
		return m.FwdFee
	}
	return 0
}

func (m *RawMessage) GetIhrFee() int64 {
	if m != nil {
		return m.IhrFee
	}
	return 0
}

func (m *RawMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RawMessage) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RawMessage) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type InternalTransactionId struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Lt                   int64    `protobuf:"varint,2,opt,name=lt,proto3" json:"lt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTransactionId) Reset()         { *m = InternalTransactionId{} }
func (m *InternalTransactionId) String() string { return proto.CompactTextString(m) }
func (*InternalTransactionId) ProtoMessage()    {}
func (*InternalTransactionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *InternalTransactionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTransactionId.Unmarshal(m, b)
}
func (m *InternalTransactionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTransactionId.Marshal(b, m, deterministic)
}
func (m *InternalTransactionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTransactionId.Merge(m, src)
}
func (m *InternalTransactionId) XXX_Size() int {
	return xxx_messageInfo_InternalTransactionId.Size(m)
}
func (m *InternalTransactionId) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTransactionId.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTransactionId proto.InternalMessageInfo

func (m *InternalTransactionId) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *InternalTransactionId) GetLt() int64 {
	if m != nil {
		return m.Lt
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchTransactionsRequest)(nil), "api.FetchTransactionsRequest")
	proto.RegisterType((*FetchTransactionsResponse)(nil), "api.FetchTransactionsResponse")
	proto.RegisterType((*Transaction)(nil), "api.Transaction")
	proto.RegisterType((*RawMessage)(nil), "api.RawMessage")
	proto.RegisterType((*InternalTransactionId)(nil), "api.InternalTransactionId")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x6d, 0xc7, 0x63, 0x51, 0xca, 0x0a, 0xa8, 0x29, 0x2a, 0x44, 0x3e, 0x54, 0x11,
	0x87, 0x1e, 0xc2, 0x91, 0x53, 0x85, 0x54, 0x51, 0x89, 0x5c, 0xac, 0x1c, 0x38, 0x61, 0x6d, 0xb3,
	0x13, 0x7b, 0xa5, 0xc4, 0x6b, 0x3c, 0x6b, 0x22, 0xfe, 0x80, 0x0f, 0xe5, 0x43, 0xd0, 0x8e, 0xdd,
	0xc6, 0xa2, 0xed, 0x6d, 0xe6, 0x8d, 0xf7, 0xed, 0x9b, 0xb7, 0xcf, 0x10, 0xcb, 0x46, 0x5f, 0x36,
	0xad, 0xb1, 0x46, 0xf8, 0xb2, 0xd1, 0xd9, 0x77, 0x48, 0xaf, 0xd1, 0xae, 0xab, 0x55, 0x2b, 0x6b,
	0x92, 0x6b, 0xab, 0x4d, 0x4d, 0x39, 0xfe, 0xec, 0x90, 0xac, 0x48, 0x21, 0x92, 0x4a, 0xb5, 0x48,
	0x94, 0x7a, 0x33, 0x6f, 0x1e, 0xe7, 0x77, 0xad, 0x38, 0x86, 0xc9, 0xd6, 0xa6, 0x93, 0x99, 0x37,
	0xf7, 0xf3, 0xc9, 0xd6, 0x0a, 0x01, 0x47, 0x95, 0xa4, 0x2a, 0xf5, 0xf9, 0x33, 0xae, 0xb3, 0x2f,
	0xf0, 0xf6, 0x11, 0x66, 0x6a, 0x4c, 0x4d, 0x28, 0x2e, 0x20, 0xd0, 0x16, 0x77, 0x8e, 0xd8, 0x9f,
	0x27, 0x8b, 0x93, 0x4b, 0x27, 0x6b, 0xf4, 0x65, 0xde, 0x8f, 0xb3, 0x3f, 0x13, 0x48, 0x46, 0xb0,
	0xbb, 0x48, 0x49, 0x2b, 0x07, 0x3d, 0x5c, 0x8b, 0x13, 0xf0, 0x37, 0x88, 0x83, 0x1a, 0x57, 0x8a,
	0x0b, 0x08, 0x75, 0x5d, 0xec, 0xa8, 0x64, 0x41, 0xc9, 0xe2, 0x05, 0xd3, 0xe7, 0x72, 0xbf, 0x44,
	0x22, 0x59, 0x62, 0x1e, 0xe8, 0x7a, 0x49, 0xa5, 0x78, 0x07, 0xb1, 0xb1, 0x15, 0xb6, 0x85, 0x3b,
	0x7f, 0xc4, 0xe7, 0xa7, 0x0c, 0x5c, 0x23, 0x8a, 0x8f, 0x30, 0x35, 0x9d, 0x75, 0x2c, 0x94, 0x06,
	0xac, 0xf2, 0x01, 0x4d, 0x64, 0x3a, 0xbb, 0xa4, 0x92, 0xc4, 0x07, 0x48, 0xc8, 0x9a, 0x56, 0x96,
	0xc8, 0x54, 0x21, 0x53, 0xc1, 0x00, 0x39, 0xb2, 0x2b, 0x38, 0xb6, 0x87, 0x35, 0x0a, 0xad, 0xd2,
	0x88, 0x95, 0x9d, 0x31, 0xe5, 0x4d, 0x6d, 0xb1, 0xad, 0xe5, 0x76, 0xb4, 0xe9, 0x8d, 0xca, 0x9f,
	0xdb, 0x71, 0x9b, 0xfd, 0xf5, 0x00, 0x0e, 0x77, 0x3b, 0xed, 0xb7, 0x46, 0xfd, 0x2e, 0xd8, 0xf7,
	0xde, 0x8e, 0xa9, 0x03, 0xbe, 0x4a, 0xaa, 0xc4, 0x39, 0xc0, 0xba, 0x45, 0x69, 0x51, 0x15, 0xf7,
	0xef, 0x14, 0x0f, 0xc8, 0x37, 0x2b, 0x66, 0x90, 0x28, 0x24, 0xab, 0x6b, 0xe9, 0xb8, 0x87, 0x57,
	0x1b, 0x43, 0xe2, 0x14, 0xa2, 0xcd, 0x5e, 0x8d, 0x7c, 0x09, 0x37, 0x7b, 0xe5, 0x16, 0x39, 0x85,
	0x48, 0x57, 0xbd, 0x61, 0x41, 0x3f, 0xd0, 0x15, 0xdb, 0x95, 0x42, 0xb4, 0xeb, 0xa5, 0xf1, 0xfa,
	0x71, 0x7e, 0xd7, 0x8a, 0x37, 0x10, 0x92, 0xe9, 0xda, 0x35, 0xf2, 0xce, 0x71, 0x3e, 0x74, 0xe2,
	0x15, 0x04, 0xbf, 0xe4, 0xb6, 0xc3, 0x74, 0xca, 0x44, 0x7d, 0x93, 0x7d, 0x86, 0xd7, 0x8f, 0xda,
	0x71, 0x9f, 0x31, 0xef, 0x90, 0xb1, 0xff, 0x73, 0xb8, 0xf8, 0x01, 0xe1, 0xca, 0xd4, 0x57, 0x8d,
	0x16, 0x2b, 0x78, 0xf9, 0x20, 0x7d, 0xe2, 0x9c, 0xdd, 0x7e, 0x2a, 0xef, 0x67, 0xef, 0x9f, 0x1a,
	0xf7, 0xa1, 0xcd, 0x9e, 0xdd, 0x86, 0xfc, 0xe7, 0x7c, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x58,
	0x3f, 0x97, 0x3f, 0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TonApiClient is the client API for TonApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TonApiClient interface {
	FetchTransactions(ctx context.Context, in *FetchTransactionsRequest, opts ...grpc.CallOption) (*FetchTransactionsResponse, error)
}

type tonApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTonApiClient(cc grpc.ClientConnInterface) TonApiClient {
	return &tonApiClient{cc}
}

func (c *tonApiClient) FetchTransactions(ctx context.Context, in *FetchTransactionsRequest, opts ...grpc.CallOption) (*FetchTransactionsResponse, error) {
	out := new(FetchTransactionsResponse)
	err := c.cc.Invoke(ctx, "/api.TonApi/FetchTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TonApiServer is the server API for TonApi service.
type TonApiServer interface {
	FetchTransactions(context.Context, *FetchTransactionsRequest) (*FetchTransactionsResponse, error)
}

// UnimplementedTonApiServer can be embedded to have forward compatible implementations.
type UnimplementedTonApiServer struct {
}

func (*UnimplementedTonApiServer) FetchTransactions(ctx context.Context, req *FetchTransactionsRequest) (*FetchTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTransactions not implemented")
}

func RegisterTonApiServer(s *grpc.Server, srv TonApiServer) {
	s.RegisterService(&_TonApi_serviceDesc, srv)
}

func _TonApi_FetchTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonApiServer).FetchTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TonApi/FetchTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonApiServer).FetchTransactions(ctx, req.(*FetchTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TonApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TonApi",
	HandlerType: (*TonApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTransactions",
			Handler:    _TonApi_FetchTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

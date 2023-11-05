// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/v1/wallet_service.proto

package wallet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateWalletRequest struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string             `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Exchange             string             `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Assets               map[string]float64 `protobuf:"bytes,4,rep,name=assets,proto3" json:"assets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateWalletRequest) Reset()         { *m = CreateWalletRequest{} }
func (m *CreateWalletRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWalletRequest) ProtoMessage()    {}
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5cbb926e133f949, []int{0}
}

func (m *CreateWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletRequest.Unmarshal(m, b)
}
func (m *CreateWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletRequest.Marshal(b, m, deterministic)
}
func (m *CreateWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletRequest.Merge(m, src)
}
func (m *CreateWalletRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWalletRequest.Size(m)
}
func (m *CreateWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletRequest proto.InternalMessageInfo

func (m *CreateWalletRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateWalletRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateWalletRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *CreateWalletRequest) GetAssets() map[string]float64 {
	if m != nil {
		return m.Assets
	}
	return nil
}

type GetWalletRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	WalletId             string   `protobuf:"bytes,3,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWalletRequest) Reset()         { *m = GetWalletRequest{} }
func (m *GetWalletRequest) String() string { return proto.CompactTextString(m) }
func (*GetWalletRequest) ProtoMessage()    {}
func (*GetWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5cbb926e133f949, []int{1}
}

func (m *GetWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletRequest.Unmarshal(m, b)
}
func (m *GetWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletRequest.Marshal(b, m, deterministic)
}
func (m *GetWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletRequest.Merge(m, src)
}
func (m *GetWalletRequest) XXX_Size() int {
	return xxx_messageInfo_GetWalletRequest.Size(m)
}
func (m *GetWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletRequest proto.InternalMessageInfo

func (m *GetWalletRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetWalletRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetWalletRequest) GetWalletId() string {
	if m != nil {
		return m.WalletId
	}
	return ""
}

type GetWalletResponse struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Wallet               *SimpleWallet `protobuf:"bytes,3,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetWalletResponse) Reset()         { *m = GetWalletResponse{} }
func (m *GetWalletResponse) String() string { return proto.CompactTextString(m) }
func (*GetWalletResponse) ProtoMessage()    {}
func (*GetWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5cbb926e133f949, []int{2}
}

func (m *GetWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletResponse.Unmarshal(m, b)
}
func (m *GetWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletResponse.Marshal(b, m, deterministic)
}
func (m *GetWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletResponse.Merge(m, src)
}
func (m *GetWalletResponse) XXX_Size() int {
	return xxx_messageInfo_GetWalletResponse.Size(m)
}
func (m *GetWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletResponse proto.InternalMessageInfo

func (m *GetWalletResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetWalletResponse) GetWallet() *SimpleWallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

type GetUserWalletsRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserWalletsRequest) Reset()         { *m = GetUserWalletsRequest{} }
func (m *GetUserWalletsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserWalletsRequest) ProtoMessage()    {}
func (*GetUserWalletsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5cbb926e133f949, []int{3}
}

func (m *GetUserWalletsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserWalletsRequest.Unmarshal(m, b)
}
func (m *GetUserWalletsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserWalletsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserWalletsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserWalletsRequest.Merge(m, src)
}
func (m *GetUserWalletsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserWalletsRequest.Size(m)
}
func (m *GetUserWalletsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserWalletsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserWalletsRequest proto.InternalMessageInfo

func (m *GetUserWalletsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetUserWalletsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserWalletsResponse struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Wallet               *UserWallets `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserWalletsResponse) Reset()         { *m = GetUserWalletsResponse{} }
func (m *GetUserWalletsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserWalletsResponse) ProtoMessage()    {}
func (*GetUserWalletsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5cbb926e133f949, []int{4}
}

func (m *GetUserWalletsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserWalletsResponse.Unmarshal(m, b)
}
func (m *GetUserWalletsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserWalletsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserWalletsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserWalletsResponse.Merge(m, src)
}
func (m *GetUserWalletsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserWalletsResponse.Size(m)
}
func (m *GetUserWalletsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserWalletsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserWalletsResponse proto.InternalMessageInfo

func (m *GetUserWalletsResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetUserWalletsResponse) GetWallet() *UserWallets {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateWalletRequest)(nil), "wallet.v1.CreateWalletRequest")
	proto.RegisterMapType((map[string]float64)(nil), "wallet.v1.CreateWalletRequest.AssetsEntry")
	proto.RegisterType((*GetWalletRequest)(nil), "wallet.v1.GetWalletRequest")
	proto.RegisterType((*GetWalletResponse)(nil), "wallet.v1.GetWalletResponse")
	proto.RegisterType((*GetUserWalletsRequest)(nil), "wallet.v1.GetUserWalletsRequest")
	proto.RegisterType((*GetUserWalletsResponse)(nil), "wallet.v1.GetUserWalletsResponse")
}

func init() {
	proto.RegisterFile("wallet/v1/wallet_service.proto", fileDescriptor_f5cbb926e133f949)
}

var fileDescriptor_f5cbb926e133f949 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x0e, 0x18, 0x32, 0x29, 0x55, 0x59, 0x20, 0xb5, 0x5c, 0x54, 0x05, 0x9f, 0x2a, 0x0e,
	0x5e, 0x35, 0x5c, 0xa0, 0x70, 0x80, 0x02, 0x2a, 0x15, 0x12, 0x07, 0x17, 0x54, 0x04, 0x87, 0xc8,
	0xb1, 0x47, 0xae, 0x85, 0xed, 0x35, 0xbb, 0x6b, 0x97, 0xfe, 0x4a, 0xce, 0xfc, 0x1b, 0x94, 0x5d,
	0xc7, 0xda, 0x7c, 0x10, 0x81, 0x7a, 0x9b, 0x79, 0xf3, 0xe6, 0xf9, 0xcd, 0x8c, 0x17, 0xf6, 0x2f,
	0xa3, 0x3c, 0x47, 0x49, 0x9b, 0x43, 0xaa, 0xa3, 0x89, 0x40, 0xde, 0x64, 0x31, 0x06, 0x15, 0x67,
	0x92, 0x91, 0xbe, 0x46, 0x83, 0xe6, 0xd0, 0x1b, 0x2e, 0x53, 0x35, 0xc5, 0xff, 0x6d, 0xc1, 0xfd,
	0x37, 0x1c, 0x23, 0x89, 0xe7, 0x0a, 0x0e, 0xf1, 0x47, 0x8d, 0x42, 0x92, 0x6d, 0xb0, 0xb3, 0xc4,
	0xb5, 0x46, 0xd6, 0x41, 0x3f, 0xb4, 0xb3, 0x84, 0xec, 0xc2, 0xed, 0x5a, 0x20, 0x9f, 0x64, 0x89,
	0x6b, 0x2b, 0xd0, 0x99, 0xa5, 0xa7, 0x09, 0xf1, 0xe0, 0x0e, 0xfe, 0x8c, 0x2f, 0xa2, 0x32, 0x45,
	0xb7, 0xa7, 0x2a, 0x5d, 0x4e, 0x8e, 0xc1, 0x89, 0x84, 0x40, 0x29, 0xdc, 0x9b, 0xa3, 0xde, 0xc1,
	0x60, 0xfc, 0x24, 0xe8, 0x0c, 0x05, 0x6b, 0x3e, 0x1a, 0xbc, 0x56, 0xe4, 0x77, 0xa5, 0xe4, 0x57,
	0x61, 0xdb, 0xe9, 0x3d, 0x87, 0x81, 0x01, 0x93, 0x1d, 0xe8, 0x7d, 0xc7, 0xab, 0xd6, 0xd8, 0x2c,
	0x24, 0x0f, 0xe0, 0x56, 0x13, 0xe5, 0x35, 0x2a, 0x5f, 0x56, 0xa8, 0x93, 0x23, 0xfb, 0x99, 0xe5,
	0x7f, 0x83, 0x9d, 0x13, 0x94, 0x9b, 0xe7, 0x32, 0xed, 0xdb, 0x4b, 0xf6, 0xf7, 0xa0, 0x5d, 0xe0,
	0x6c, 0xea, 0x76, 0x36, 0x0d, 0x9c, 0x26, 0xfe, 0x27, 0xb8, 0x67, 0x88, 0x8b, 0x8a, 0x95, 0x02,
	0x57, 0xd4, 0x29, 0x38, 0xba, 0x41, 0xb5, 0x0f, 0xc6, 0xbb, 0xc6, 0x02, 0xce, 0xb2, 0xa2, 0xca,
	0xe7, 0x0b, 0x68, 0x69, 0xfe, 0x2b, 0x78, 0x78, 0x82, 0xf2, 0xb3, 0x40, 0xae, 0x0b, 0xe2, 0x7f,
	0xef, 0xe1, 0x7f, 0x81, 0xe1, 0xb2, 0xc2, 0x5f, 0xcc, 0x05, 0x9d, 0x39, 0x5b, 0x99, 0x1b, 0x1a,
	0xe6, 0xcc, 0xfe, 0x96, 0x35, 0xfe, 0x65, 0xc3, 0x5d, 0x8d, 0x9d, 0xe9, 0xbf, 0x8c, 0x7c, 0x84,
	0x2d, 0xf3, 0x8c, 0x64, 0x7f, 0xf3, 0x7d, 0xbd, 0x47, 0x46, 0x7d, 0x65, 0x79, 0xfe, 0x0d, 0xf2,
	0x01, 0xb6, 0xde, 0x62, 0x8e, 0x9d, 0xde, 0xde, 0x7a, 0xfe, 0xbf, 0x89, 0xbd, 0x87, 0x7e, 0x07,
	0x5f, 0x4f, 0xe9, 0x1c, 0xb6, 0x17, 0x57, 0x4a, 0x46, 0x8b, 0x1d, 0xab, 0xf7, 0xf2, 0x1e, 0x6f,
	0x60, 0xcc, 0x85, 0x8f, 0x5f, 0x7e, 0x3d, 0x4a, 0x33, 0x79, 0x51, 0x4f, 0x83, 0x98, 0x15, 0xf4,
	0x32, 0xe2, 0x55, 0xca, 0xa9, 0x4c, 0xa6, 0xd1, 0x24, 0x66, 0x45, 0xc1, 0x4a, 0x9a, 0x32, 0xaa,
	0x9e, 0x69, 0xcc, 0x72, 0xda, 0x3d, 0xe0, 0x17, 0x3a, 0x9a, 0x3a, 0xaa, 0xf4, 0xf4, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x08, 0x19, 0x0c, 0x31, 0x06, 0x04, 0x00, 0x00,
}

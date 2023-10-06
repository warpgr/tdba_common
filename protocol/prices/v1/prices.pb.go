// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prices/v1/prices.proto

package prices

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

type GetPriceForRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceForRequest) Reset()         { *m = GetPriceForRequest{} }
func (m *GetPriceForRequest) String() string { return proto.CompactTextString(m) }
func (*GetPriceForRequest) ProtoMessage()    {}
func (*GetPriceForRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{0}
}

func (m *GetPriceForRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceForRequest.Unmarshal(m, b)
}
func (m *GetPriceForRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceForRequest.Marshal(b, m, deterministic)
}
func (m *GetPriceForRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceForRequest.Merge(m, src)
}
func (m *GetPriceForRequest) XXX_Size() int {
	return xxx_messageInfo_GetPriceForRequest.Size(m)
}
func (m *GetPriceForRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceForRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceForRequest proto.InternalMessageInfo

func (m *GetPriceForRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPriceForRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPriceForRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type GetPriceForResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceForResponse) Reset()         { *m = GetPriceForResponse{} }
func (m *GetPriceForResponse) String() string { return proto.CompactTextString(m) }
func (*GetPriceForResponse) ProtoMessage()    {}
func (*GetPriceForResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{1}
}

func (m *GetPriceForResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceForResponse.Unmarshal(m, b)
}
func (m *GetPriceForResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceForResponse.Marshal(b, m, deterministic)
}
func (m *GetPriceForResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceForResponse.Merge(m, src)
}
func (m *GetPriceForResponse) XXX_Size() int {
	return xxx_messageInfo_GetPriceForResponse.Size(m)
}
func (m *GetPriceForResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceForResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceForResponse proto.InternalMessageInfo

func (m *GetPriceForResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPriceForResponse) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPriceForResponse) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetPriceForResponse) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type GetPricesForRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbols              []string `protobuf:"bytes,3,rep,name=symbols,proto3" json:"symbols,omitempty"`
	SuspendMili          int64    `protobuf:"varint,4,opt,name=suspend_mili,json=suspendMili,proto3" json:"suspend_mili,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesForRequest) Reset()         { *m = GetPricesForRequest{} }
func (m *GetPricesForRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricesForRequest) ProtoMessage()    {}
func (*GetPricesForRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{2}
}

func (m *GetPricesForRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesForRequest.Unmarshal(m, b)
}
func (m *GetPricesForRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesForRequest.Marshal(b, m, deterministic)
}
func (m *GetPricesForRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesForRequest.Merge(m, src)
}
func (m *GetPricesForRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricesForRequest.Size(m)
}
func (m *GetPricesForRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesForRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesForRequest proto.InternalMessageInfo

func (m *GetPricesForRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricesForRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPricesForRequest) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *GetPricesForRequest) GetSuspendMili() int64 {
	if m != nil {
		return m.SuspendMili
	}
	return 0
}

type GetPricesForResponse struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string             `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Prices               map[string]float64 `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPricesForResponse) Reset()         { *m = GetPricesForResponse{} }
func (m *GetPricesForResponse) String() string { return proto.CompactTextString(m) }
func (*GetPricesForResponse) ProtoMessage()    {}
func (*GetPricesForResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{3}
}

func (m *GetPricesForResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesForResponse.Unmarshal(m, b)
}
func (m *GetPricesForResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesForResponse.Marshal(b, m, deterministic)
}
func (m *GetPricesForResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesForResponse.Merge(m, src)
}
func (m *GetPricesForResponse) XXX_Size() int {
	return xxx_messageInfo_GetPricesForResponse.Size(m)
}
func (m *GetPricesForResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesForResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesForResponse proto.InternalMessageInfo

func (m *GetPricesForResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricesForResponse) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPricesForResponse) GetPrices() map[string]float64 {
	if m != nil {
		return m.Prices
	}
	return nil
}

type GetPriceFromToRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	From                 int64    `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceFromToRequest) Reset()         { *m = GetPriceFromToRequest{} }
func (m *GetPriceFromToRequest) String() string { return proto.CompactTextString(m) }
func (*GetPriceFromToRequest) ProtoMessage()    {}
func (*GetPriceFromToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{4}
}

func (m *GetPriceFromToRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceFromToRequest.Unmarshal(m, b)
}
func (m *GetPriceFromToRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceFromToRequest.Marshal(b, m, deterministic)
}
func (m *GetPriceFromToRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceFromToRequest.Merge(m, src)
}
func (m *GetPriceFromToRequest) XXX_Size() int {
	return xxx_messageInfo_GetPriceFromToRequest.Size(m)
}
func (m *GetPriceFromToRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceFromToRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceFromToRequest proto.InternalMessageInfo

func (m *GetPriceFromToRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPriceFromToRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPriceFromToRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetPriceFromToRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *GetPriceFromToRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type GetPriceFromToResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceFromToResponse) Reset()         { *m = GetPriceFromToResponse{} }
func (m *GetPriceFromToResponse) String() string { return proto.CompactTextString(m) }
func (*GetPriceFromToResponse) ProtoMessage()    {}
func (*GetPriceFromToResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{5}
}

func (m *GetPriceFromToResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceFromToResponse.Unmarshal(m, b)
}
func (m *GetPriceFromToResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceFromToResponse.Marshal(b, m, deterministic)
}
func (m *GetPriceFromToResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceFromToResponse.Merge(m, src)
}
func (m *GetPriceFromToResponse) XXX_Size() int {
	return xxx_messageInfo_GetPriceFromToResponse.Size(m)
}
func (m *GetPriceFromToResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceFromToResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceFromToResponse proto.InternalMessageInfo

func (m *GetPriceFromToResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPriceFromToResponse) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPriceFromToResponse) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type GetPricesFromToRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbols              []string `protobuf:"bytes,3,rep,name=symbols,proto3" json:"symbols,omitempty"`
	From                 int64    `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesFromToRequest) Reset()         { *m = GetPricesFromToRequest{} }
func (m *GetPricesFromToRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricesFromToRequest) ProtoMessage()    {}
func (*GetPricesFromToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{6}
}

func (m *GetPricesFromToRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesFromToRequest.Unmarshal(m, b)
}
func (m *GetPricesFromToRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesFromToRequest.Marshal(b, m, deterministic)
}
func (m *GetPricesFromToRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesFromToRequest.Merge(m, src)
}
func (m *GetPricesFromToRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricesFromToRequest.Size(m)
}
func (m *GetPricesFromToRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesFromToRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesFromToRequest proto.InternalMessageInfo

func (m *GetPricesFromToRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricesFromToRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPricesFromToRequest) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *GetPricesFromToRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *GetPricesFromToRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type GetPricesFromToResponse struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string             `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Prices               map[string]float64 `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPricesFromToResponse) Reset()         { *m = GetPricesFromToResponse{} }
func (m *GetPricesFromToResponse) String() string { return proto.CompactTextString(m) }
func (*GetPricesFromToResponse) ProtoMessage()    {}
func (*GetPricesFromToResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70da17ac568b99be, []int{7}
}

func (m *GetPricesFromToResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesFromToResponse.Unmarshal(m, b)
}
func (m *GetPricesFromToResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesFromToResponse.Marshal(b, m, deterministic)
}
func (m *GetPricesFromToResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesFromToResponse.Merge(m, src)
}
func (m *GetPricesFromToResponse) XXX_Size() int {
	return xxx_messageInfo_GetPricesFromToResponse.Size(m)
}
func (m *GetPricesFromToResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesFromToResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesFromToResponse proto.InternalMessageInfo

func (m *GetPricesFromToResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricesFromToResponse) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetPricesFromToResponse) GetPrices() map[string]float64 {
	if m != nil {
		return m.Prices
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPriceForRequest)(nil), "prices.v1.GetPriceForRequest")
	proto.RegisterType((*GetPriceForResponse)(nil), "prices.v1.GetPriceForResponse")
	proto.RegisterType((*GetPricesForRequest)(nil), "prices.v1.GetPricesForRequest")
	proto.RegisterType((*GetPricesForResponse)(nil), "prices.v1.GetPricesForResponse")
	proto.RegisterMapType((map[string]float64)(nil), "prices.v1.GetPricesForResponse.PricesEntry")
	proto.RegisterType((*GetPriceFromToRequest)(nil), "prices.v1.GetPriceFromToRequest")
	proto.RegisterType((*GetPriceFromToResponse)(nil), "prices.v1.GetPriceFromToResponse")
	proto.RegisterType((*GetPricesFromToRequest)(nil), "prices.v1.GetPricesFromToRequest")
	proto.RegisterType((*GetPricesFromToResponse)(nil), "prices.v1.GetPricesFromToResponse")
	proto.RegisterMapType((map[string]float64)(nil), "prices.v1.GetPricesFromToResponse.PricesEntry")
}

func init() {
	proto.RegisterFile("prices/v1/prices.proto", fileDescriptor_70da17ac568b99be)
}

var fileDescriptor_70da17ac568b99be = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x65, 0x25, 0xdb, 0xad, 0xc7, 0xa5, 0x94, 0xad, 0xeb, 0x0a, 0x9f, 0x5c, 0x9d, 0x0c, 0x05,
	0x09, 0xb7, 0x87, 0xd6, 0xed, 0xad, 0x25, 0xce, 0x29, 0x10, 0x44, 0x0e, 0xc1, 0x17, 0xa3, 0x8f,
	0x8d, 0xbd, 0x44, 0xd2, 0x2a, 0xbb, 0x2b, 0x27, 0xbe, 0x24, 0x90, 0x9f, 0x96, 0x53, 0x7e, 0x56,
	0xd0, 0xae, 0x22, 0x7f, 0xc4, 0x81, 0x20, 0x7c, 0x9b, 0x37, 0x42, 0x6f, 0xde, 0x9b, 0x79, 0x2c,
	0xf4, 0x32, 0x4e, 0x43, 0x22, 0xdc, 0xe5, 0xc8, 0xd5, 0x95, 0x93, 0x71, 0x26, 0x19, 0x6e, 0x97,
	0x68, 0x39, 0xb2, 0xcf, 0x01, 0x1f, 0x13, 0x79, 0x5a, 0xe0, 0x09, 0xe3, 0x1e, 0xb9, 0xca, 0x89,
	0x90, 0xf8, 0x23, 0x18, 0x34, 0xb2, 0xd0, 0x00, 0x0d, 0xdb, 0x9e, 0x41, 0x23, 0xdc, 0x87, 0xf7,
	0xe4, 0x26, 0x5c, 0xf8, 0xe9, 0x9c, 0x58, 0x86, 0xea, 0x56, 0x18, 0xf7, 0xa0, 0x25, 0x56, 0x49,
	0xc0, 0x62, 0xcb, 0x54, 0x5f, 0x4a, 0x64, 0x33, 0xf8, 0xbc, 0xc5, 0x2c, 0x32, 0x96, 0x0a, 0x72,
	0x08, 0x6a, 0xdc, 0x85, 0xa6, 0x72, 0x60, 0x35, 0x06, 0x68, 0x88, 0x3c, 0x0d, 0xec, 0xdb, 0xf5,
	0x40, 0x51, 0xd3, 0x8b, 0x05, 0xef, 0xf4, 0x08, 0x61, 0x99, 0x03, 0x73, 0xd8, 0xf6, 0x9e, 0x21,
	0xfe, 0x06, 0x1f, 0x44, 0x2e, 0x32, 0x92, 0x46, 0xb3, 0x84, 0xc6, 0x54, 0x4d, 0x36, 0xbd, 0x4e,
	0xd9, 0x3b, 0xa1, 0x31, 0xb5, 0x1f, 0x10, 0x74, 0xb7, 0x05, 0xd4, 0xb0, 0xfc, 0x1f, 0x5a, 0xfa,
	0x38, 0x4a, 0x40, 0xe7, 0xc7, 0x77, 0xa7, 0xba, 0x95, 0xb3, 0x8f, 0xdc, 0xd1, 0x9d, 0xa3, 0x54,
	0xf2, 0x95, 0x57, 0xfe, 0xda, 0x1f, 0x43, 0x67, 0xa3, 0x8d, 0x3f, 0x81, 0x79, 0x49, 0x56, 0xa5,
	0x80, 0xa2, 0x2c, 0x16, 0xb8, 0xf4, 0xe3, 0x5c, 0x8f, 0x47, 0x9e, 0x06, 0x7f, 0x8c, 0xdf, 0xc8,
	0xbe, 0x83, 0x2f, 0xd5, 0xd5, 0x38, 0x4b, 0xce, 0xd8, 0x01, 0x23, 0x81, 0x31, 0x34, 0x2e, 0x38,
	0x4b, 0xca, 0xe5, 0xa9, 0xba, 0xe0, 0x95, 0xcc, 0x6a, 0xaa, 0x8e, 0x21, 0x99, 0x3d, 0x85, 0xde,
	0xae, 0x80, 0x1a, 0x6b, 0xac, 0x12, 0x62, 0x6e, 0x26, 0xe4, 0x1e, 0xad, 0xc9, 0x45, 0x7d, 0x7b,
	0xaf, 0xa7, 0xe4, 0x2d, 0x06, 0x1f, 0x11, 0x7c, 0x7d, 0x21, 0xa2, 0x86, 0xc5, 0xc9, 0x4e, 0x52,
	0x9c, 0xbd, 0x49, 0xd9, 0xe2, 0x3f, 0x70, 0x58, 0xfe, 0x8d, 0xa7, 0xbf, 0xe6, 0x54, 0x2e, 0xf2,
	0xc0, 0x09, 0x59, 0xe2, 0x5e, 0xfb, 0x3c, 0x9b, 0x73, 0x57, 0x46, 0x81, 0x3f, 0x0b, 0x59, 0x92,
	0xb0, 0xd4, 0x55, 0xcf, 0x4d, 0xc8, 0x62, 0xb7, 0x7a, 0x88, 0xfe, 0xea, 0x2a, 0x68, 0xa9, 0x4f,
	0x3f, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x1b, 0xed, 0xcc, 0xa3, 0x04, 0x00, 0x00,
}

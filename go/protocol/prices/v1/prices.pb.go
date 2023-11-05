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
	SuspendMili          int64    `protobuf:"varint,4,opt,name=suspend_mili,json=suspendMili,proto3" json:"suspend_mili,omitempty"`
	Additional           string   `protobuf:"bytes,5,opt,name=additional,proto3" json:"additional,omitempty"`
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

func (m *GetPriceForRequest) GetSuspendMili() int64 {
	if m != nil {
		return m.SuspendMili
	}
	return 0
}

func (m *GetPriceForRequest) GetAdditional() string {
	if m != nil {
		return m.Additional
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
	Additional           string   `protobuf:"bytes,5,opt,name=additional,proto3" json:"additional,omitempty"`
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

func (m *GetPricesForRequest) GetAdditional() string {
	if m != nil {
		return m.Additional
	}
	return ""
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
	Additional           string   `protobuf:"bytes,6,opt,name=additional,proto3" json:"additional,omitempty"`
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

func (m *GetPriceFromToRequest) GetAdditional() string {
	if m != nil {
		return m.Additional
	}
	return ""
}

type GetPriceFromToResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	TimePoint            int64    `protobuf:"varint,4,opt,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
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

func (m *GetPriceFromToResponse) GetTimePoint() int64 {
	if m != nil {
		return m.TimePoint
	}
	return 0
}

type GetPricesFromToRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbols              []string `protobuf:"bytes,3,rep,name=symbols,proto3" json:"symbols,omitempty"`
	From                 int64    `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
	Additional           string   `protobuf:"bytes,6,opt,name=additional,proto3" json:"additional,omitempty"`
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

func (m *GetPricesFromToRequest) GetAdditional() string {
	if m != nil {
		return m.Additional
	}
	return ""
}

type GetPricesFromToResponse struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string             `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Prices               map[string]float64 `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	TimePoint            int64              `protobuf:"varint,4,opt,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
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

func (m *GetPricesFromToResponse) GetTimePoint() int64 {
	if m != nil {
		return m.TimePoint
	}
	return 0
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
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x99, 0x64, 0xb7, 0xda, 0x53, 0x11, 0x19, 0xd7, 0x1a, 0x16, 0x94, 0x9a, 0xab, 0x82,
	0x90, 0xb0, 0x7a, 0xa3, 0xab, 0x57, 0x8a, 0xeb, 0x95, 0xb0, 0x0c, 0x5e, 0x79, 0x53, 0xd2, 0x64,
	0xec, 0x0e, 0x66, 0x72, 0xe2, 0xcc, 0xa4, 0x9a, 0x67, 0x11, 0x14, 0x7c, 0x14, 0x5f, 0xc6, 0xd7,
	0x90, 0xcc, 0xc4, 0x6c, 0xdb, 0xad, 0x08, 0xd9, 0xdc, 0x9d, 0xff, 0x0f, 0x33, 0xe7, 0xe7, 0xcb,
	0xcf, 0xc0, 0xb4, 0x54, 0x22, 0xe5, 0x3a, 0x5e, 0x9f, 0xc4, 0x6e, 0x8a, 0x4a, 0x85, 0x06, 0xe9,
	0xb8, 0x55, 0xeb, 0x93, 0xf0, 0x1b, 0x01, 0xfa, 0x96, 0x9b, 0xf3, 0xc6, 0x38, 0x43, 0xc5, 0xf8,
	0xe7, 0x8a, 0x6b, 0x43, 0x6f, 0x83, 0x27, 0xb2, 0x80, 0xcc, 0xc8, 0x7c, 0xcc, 0x3c, 0x91, 0xd1,
	0x63, 0xb8, 0xc9, 0xbf, 0xa6, 0x17, 0x49, 0xb1, 0xe2, 0x81, 0x67, 0xdd, 0x4e, 0xd3, 0x29, 0x8c,
	0x74, 0x2d, 0x97, 0x98, 0x07, 0xbe, 0xfd, 0xd2, 0x2a, 0xfa, 0x08, 0x6e, 0xe9, 0x4a, 0x97, 0xbc,
	0xc8, 0x16, 0x52, 0xe4, 0x22, 0x38, 0x98, 0x91, 0xb9, 0xcf, 0x26, 0xad, 0xf7, 0x4e, 0xe4, 0x82,
	0x3e, 0x04, 0x48, 0xb2, 0x4c, 0x18, 0x81, 0x45, 0x92, 0x07, 0x87, 0xf6, 0xf8, 0x86, 0x13, 0x22,
	0xdc, 0xdd, 0x0a, 0xa7, 0x4b, 0x2c, 0x34, 0x1f, 0x24, 0xdd, 0x11, 0x1c, 0x5a, 0x0a, 0x36, 0x16,
	0x61, 0x4e, 0x84, 0xdf, 0xc9, 0xe5, 0x46, 0xdd, 0x93, 0x47, 0x00, 0x37, 0xdc, 0x0e, 0x1d, 0xf8,
	0x33, 0x7f, 0x3e, 0x66, 0x7f, 0xe5, 0x10, 0x44, 0x7e, 0x11, 0x38, 0xda, 0x0e, 0xd8, 0x83, 0xc9,
	0x6b, 0x18, 0xb9, 0x06, 0xd8, 0x80, 0x93, 0x27, 0x8f, 0xa3, 0xae, 0x10, 0xd1, 0xbe, 0xcb, 0x23,
	0xe7, 0xbc, 0x29, 0x8c, 0xaa, 0x59, 0x7b, 0xf4, 0xf8, 0x39, 0x4c, 0x36, 0x6c, 0x7a, 0x07, 0xfc,
	0x4f, 0xbc, 0x6e, 0x03, 0x34, 0x63, 0x43, 0x78, 0x9d, 0xe4, 0x95, 0x5b, 0x4f, 0x98, 0x13, 0xa7,
	0xde, 0x33, 0x12, 0xfe, 0x20, 0x70, 0xaf, 0xfb, 0xaf, 0x0a, 0xe5, 0x7b, 0x1c, 0xb2, 0x77, 0x14,
	0x0e, 0x3e, 0x2a, 0x94, 0x2d, 0x5d, 0x3b, 0x37, 0xf7, 0x1a, 0xb4, 0x38, 0x7d, 0xe6, 0x19, 0xdc,
	0xc1, 0x3c, 0xba, 0x82, 0xb9, 0x86, 0xe9, 0x6e, 0xc0, 0x1e, 0x9c, 0xbb, 0x8e, 0xf9, 0x1b, 0x1d,
	0xa3, 0x0f, 0x00, 0x8c, 0x90, 0x7c, 0x51, 0xa2, 0x28, 0x4c, 0x9b, 0x72, 0xdc, 0x38, 0xe7, 0x8d,
	0x11, 0xfe, 0x24, 0x97, 0xbb, 0x75, 0x7f, 0x3a, 0xff, 0x6e, 0xe1, 0x10, 0x7c, 0x7e, 0x13, 0xb8,
	0x7f, 0x25, 0x64, 0x0f, 0x42, 0x67, 0x3b, 0x4d, 0x8c, 0xf6, 0x36, 0x71, 0xeb, 0xfe, 0x7d, 0x65,
	0xfc, 0x0f, 0xd3, 0x6b, 0x74, 0xf5, 0xd5, 0xcb, 0x0f, 0xa7, 0x2b, 0x61, 0x2e, 0xaa, 0x65, 0x94,
	0xa2, 0x8c, 0xbf, 0x24, 0xaa, 0x5c, 0xa9, 0xd8, 0x64, 0xcb, 0x64, 0x91, 0xa2, 0x94, 0x58, 0xc4,
	0x2b, 0x8c, 0xed, 0xab, 0x9a, 0x62, 0x1e, 0x77, 0xef, 0xed, 0x0b, 0x37, 0x2d, 0x47, 0xf6, 0xd3,
	0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xf1, 0x10, 0x18, 0x8a, 0x05, 0x00, 0x00,
}

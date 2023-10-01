// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candlestick/v1/candlestick.proto

package candlestick

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

type Candlestick struct {
	OpenTime                 int64    `protobuf:"varint,1,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	Open                     float64  `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	High                     float64  `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low                      float64  `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Close                    float64  `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume                   float64  `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	CloseTime                int64    `protobuf:"varint,7,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	QuoteAssetVolume         float64  `protobuf:"fixed64,8,opt,name=quote_asset_volume,json=quoteAssetVolume,proto3" json:"quote_asset_volume,omitempty"`
	TradeNum                 int64    `protobuf:"varint,9,opt,name=trade_num,json=tradeNum,proto3" json:"trade_num,omitempty"`
	TakerBuyBaseAssetVolume  float64  `protobuf:"fixed64,10,opt,name=taker_buy_base_asset_volume,json=takerBuyBaseAssetVolume,proto3" json:"taker_buy_base_asset_volume,omitempty"`
	TakerBuyQuoteAssetVolume float64  `protobuf:"fixed64,11,opt,name=taker_buy_quote_asset_volume,json=takerBuyQuoteAssetVolume,proto3" json:"taker_buy_quote_asset_volume,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Candlestick) Reset()         { *m = Candlestick{} }
func (m *Candlestick) String() string { return proto.CompactTextString(m) }
func (*Candlestick) ProtoMessage()    {}
func (*Candlestick) Descriptor() ([]byte, []int) {
	return fileDescriptor_186cda349b51dcdd, []int{0}
}

func (m *Candlestick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candlestick.Unmarshal(m, b)
}
func (m *Candlestick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candlestick.Marshal(b, m, deterministic)
}
func (m *Candlestick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candlestick.Merge(m, src)
}
func (m *Candlestick) XXX_Size() int {
	return xxx_messageInfo_Candlestick.Size(m)
}
func (m *Candlestick) XXX_DiscardUnknown() {
	xxx_messageInfo_Candlestick.DiscardUnknown(m)
}

var xxx_messageInfo_Candlestick proto.InternalMessageInfo

func (m *Candlestick) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *Candlestick) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candlestick) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candlestick) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candlestick) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candlestick) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candlestick) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *Candlestick) GetQuoteAssetVolume() float64 {
	if m != nil {
		return m.QuoteAssetVolume
	}
	return 0
}

func (m *Candlestick) GetTradeNum() int64 {
	if m != nil {
		return m.TradeNum
	}
	return 0
}

func (m *Candlestick) GetTakerBuyBaseAssetVolume() float64 {
	if m != nil {
		return m.TakerBuyBaseAssetVolume
	}
	return 0
}

func (m *Candlestick) GetTakerBuyQuoteAssetVolume() float64 {
	if m != nil {
		return m.TakerBuyQuoteAssetVolume
	}
	return 0
}

func init() {
	proto.RegisterType((*Candlestick)(nil), "candlestick.Candlestick")
}

func init() {
	proto.RegisterFile("candlestick/v1/candlestick.proto", fileDescriptor_186cda349b51dcdd)
}

var fileDescriptor_186cda349b51dcdd = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x7e, 0xb9, 0xbe, 0x5d, 0x46, 0x10, 0x0d, 0x4c, 0x61, 0x78, 0xda, 0x41, 0x56,
	0xc4, 0xa3, 0x32, 0xb4, 0xde, 0x05, 0x87, 0x78, 0xf0, 0x12, 0xd2, 0x34, 0xb4, 0x65, 0x4d, 0x53,
	0x9b, 0x64, 0x63, 0xff, 0xba, 0x27, 0xc9, 0xeb, 0x64, 0x75, 0x78, 0xfb, 0x7e, 0x3f, 0x79, 0xef,
	0xcb, 0x37, 0x3c, 0x98, 0x0b, 0x5e, 0xa5, 0xa5, 0x34, 0xb6, 0x10, 0x9b, 0x68, 0x7b, 0x17, 0x75,
	0xec, 0xb2, 0x6e, 0xb4, 0xd5, 0x64, 0xd2, 0x41, 0x37, 0xdf, 0x3d, 0x98, 0xbc, 0x1c, 0x3d, 0x99,
	0x41, 0xa8, 0x6b, 0x59, 0x31, 0x5b, 0x28, 0x49, 0x83, 0x79, 0xb0, 0xe8, 0xaf, 0xc7, 0x1e, 0xbc,
	0x17, 0x4a, 0x12, 0x02, 0x03, 0xaf, 0x69, 0x6f, 0x1e, 0x2c, 0x82, 0x35, 0x6a, 0xcf, 0xf2, 0x22,
	0xcb, 0x69, 0xbf, 0x65, 0x5e, 0x93, 0x29, 0xf4, 0x4b, 0xbd, 0xa3, 0x03, 0x44, 0x5e, 0x92, 0x73,
	0x18, 0x8a, 0x52, 0x1b, 0x49, 0x87, 0xc8, 0x5a, 0x43, 0x2e, 0x60, 0xb4, 0xd5, 0xa5, 0x53, 0x92,
	0x8e, 0x10, 0x1f, 0x1c, 0xb9, 0x06, 0xc0, 0x81, 0xb6, 0xc5, 0x19, 0xb6, 0x08, 0x91, 0x60, 0x8d,
	0x5b, 0x20, 0x5f, 0x4e, 0x5b, 0xc9, 0xb8, 0x31, 0xd2, 0xb2, 0x43, 0xc4, 0x18, 0x23, 0xa6, 0xf8,
	0xf2, 0xec, 0x1f, 0x3e, 0xda, 0xb0, 0x19, 0x84, 0xb6, 0xe1, 0xa9, 0x64, 0x95, 0x53, 0x34, 0x6c,
	0x7f, 0x84, 0xe0, 0xd5, 0x29, 0xf2, 0x08, 0x33, 0xcb, 0x37, 0xb2, 0x61, 0x89, 0xdb, 0xb3, 0x84,
	0x9b, 0x93, 0x4c, 0xc0, 0xcc, 0x4b, 0x1c, 0x89, 0xdd, 0x3e, 0xe6, 0xe6, 0x4f, 0xf4, 0x0a, 0xae,
	0x8e, 0xdb, 0xff, 0x54, 0x9a, 0xe0, 0x3a, 0xfd, 0x5d, 0x7f, 0x3b, 0xa9, 0x16, 0x3f, 0x7d, 0xae,
	0xb2, 0xc2, 0xe6, 0x2e, 0x59, 0x0a, 0xad, 0xa2, 0x1d, 0x6f, 0xea, 0xac, 0x89, 0x6c, 0x9a, 0x70,
	0x26, 0xb4, 0x52, 0xba, 0x8a, 0xf0, 0x60, 0x42, 0x97, 0xdd, 0x2b, 0x3e, 0x74, 0x74, 0x32, 0xc2,
	0x89, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xfd, 0x75, 0x59, 0xf6, 0x01, 0x00, 0x00,
}

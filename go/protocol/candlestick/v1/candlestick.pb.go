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
	proto.RegisterType((*Candlestick)(nil), "candlestick.v1.Candlestick")
}

func init() {
	proto.RegisterFile("candlestick/v1/candlestick.proto", fileDescriptor_186cda349b51dcdd)
}

var fileDescriptor_186cda349b51dcdd = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0xfe, 0xfa, 0x76, 0xa7, 0xf0, 0xa5, 0x04, 0xd1, 0x40, 0x15, 0x8a, 0xa7, 0x1e,
	0xa4, 0xa1, 0x78, 0x54, 0x04, 0xab, 0x78, 0x14, 0x2c, 0xe2, 0xc1, 0x4b, 0xc8, 0x6e, 0xc3, 0x76,
	0xe9, 0x66, 0xa7, 0x6e, 0x92, 0x96, 0xfe, 0xeb, 0x9e, 0x24, 0xb3, 0x95, 0xfe, 0xc0, 0xdb, 0x7b,
	0x9f, 0xcc, 0x3c, 0x5e, 0x18, 0x18, 0xa6, 0xaa, 0x9c, 0x17, 0xda, 0xba, 0x3c, 0x5d, 0x8a, 0xf5,
	0x44, 0x1c, 0xd8, 0xf1, 0xaa, 0x42, 0x87, 0xec, 0xff, 0x21, 0x5a, 0x4f, 0xae, 0xbf, 0x1b, 0xd0,
	0x7b, 0xda, 0x23, 0x36, 0x80, 0x18, 0x57, 0xba, 0x94, 0x2e, 0x37, 0x9a, 0x47, 0xc3, 0x68, 0xd4,
	0x9c, 0x75, 0x03, 0x78, 0xcf, 0x8d, 0x66, 0x0c, 0x5a, 0x41, 0xf3, 0xc6, 0x30, 0x1a, 0x45, 0x33,
	0xd2, 0x81, 0x2d, 0xf2, 0x6c, 0xc1, 0x9b, 0x35, 0x0b, 0x9a, 0xf5, 0xa1, 0x59, 0xe0, 0x86, 0xb7,
	0x08, 0x05, 0xc9, 0xce, 0xa0, 0x9d, 0x16, 0x68, 0x35, 0x6f, 0x13, 0xab, 0x0d, 0x3b, 0x87, 0xce,
	0x1a, 0x0b, 0x6f, 0x34, 0xef, 0x10, 0xde, 0x39, 0x76, 0x05, 0x40, 0x03, 0x75, 0x8b, 0x7f, 0xd4,
	0x22, 0x26, 0x42, 0x35, 0x6e, 0x80, 0x7d, 0x79, 0x74, 0x5a, 0x2a, 0x6b, 0xb5, 0x93, 0xbb, 0x88,
	0x2e, 0x45, 0xf4, 0xe9, 0xe5, 0x31, 0x3c, 0x7c, 0xd4, 0x61, 0x03, 0x88, 0x5d, 0xa5, 0xe6, 0x5a,
	0x96, 0xde, 0xf0, 0xb8, 0xfe, 0x11, 0x81, 0x57, 0x6f, 0xd8, 0x3d, 0x0c, 0x9c, 0x5a, 0xea, 0x4a,
	0x26, 0x7e, 0x2b, 0x13, 0x65, 0x4f, 0x32, 0x81, 0x32, 0x2f, 0x68, 0x64, 0xea, 0xb7, 0x53, 0x65,
	0x8f, 0xa2, 0x1f, 0xe0, 0x72, 0xbf, 0xfd, 0x47, 0xa5, 0x1e, 0xad, 0xf3, 0xdf, 0xf5, 0xb7, 0x93,
	0x6a, 0xd3, 0x97, 0xcf, 0xe7, 0x2c, 0x77, 0x0b, 0x9f, 0x8c, 0x53, 0x34, 0x62, 0xa3, 0xaa, 0x55,
	0x56, 0x09, 0x37, 0x4f, 0x94, 0x4c, 0xd1, 0x18, 0x2c, 0x45, 0x86, 0x82, 0xce, 0x96, 0x62, 0x21,
	0x8e, 0x4f, 0x7b, 0x77, 0x60, 0x93, 0x0e, 0x0d, 0xdd, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0xbf, 0x82, 0x3b, 0xff, 0x01, 0x00, 0x00,
}

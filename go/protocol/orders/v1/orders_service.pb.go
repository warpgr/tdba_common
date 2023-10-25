// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/v1/orders_service.proto

package orders

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

type PlaceOrderRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UsderId              string       `protobuf:"bytes,2,opt,name=usder_id,json=usderId,proto3" json:"usder_id,omitempty"`
	Order                *CreateOrder `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PlaceOrderRequest) Reset()         { *m = PlaceOrderRequest{} }
func (m *PlaceOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderRequest) ProtoMessage()    {}
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb21f6ae3f6980d6, []int{0}
}

func (m *PlaceOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceOrderRequest.Unmarshal(m, b)
}
func (m *PlaceOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceOrderRequest.Marshal(b, m, deterministic)
}
func (m *PlaceOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderRequest.Merge(m, src)
}
func (m *PlaceOrderRequest) XXX_Size() int {
	return xxx_messageInfo_PlaceOrderRequest.Size(m)
}
func (m *PlaceOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderRequest proto.InternalMessageInfo

func (m *PlaceOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlaceOrderRequest) GetUsderId() string {
	if m != nil {
		return m.UsderId
	}
	return ""
}

func (m *PlaceOrderRequest) GetOrder() *CreateOrder {
	if m != nil {
		return m.Order
	}
	return nil
}

type OrderStatusRequest struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderDescriptor      *OrderDescriptorType `protobuf:"bytes,3,opt,name=order_descriptor,json=orderDescriptor,proto3" json:"order_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderStatusRequest) Reset()         { *m = OrderStatusRequest{} }
func (m *OrderStatusRequest) String() string { return proto.CompactTextString(m) }
func (*OrderStatusRequest) ProtoMessage()    {}
func (*OrderStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb21f6ae3f6980d6, []int{1}
}

func (m *OrderStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatusRequest.Unmarshal(m, b)
}
func (m *OrderStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatusRequest.Marshal(b, m, deterministic)
}
func (m *OrderStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatusRequest.Merge(m, src)
}
func (m *OrderStatusRequest) XXX_Size() int {
	return xxx_messageInfo_OrderStatusRequest.Size(m)
}
func (m *OrderStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatusRequest proto.InternalMessageInfo

func (m *OrderStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderStatusRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OrderStatusRequest) GetOrderDescriptor() *OrderDescriptorType {
	if m != nil {
		return m.OrderDescriptor
	}
	return nil
}

type OrderStatusResponse struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UsderId              string                `protobuf:"bytes,2,opt,name=usder_id,json=usderId,proto3" json:"usder_id,omitempty"`
	Order                *OrderAndExchangeType `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OrderStatusResponse) Reset()         { *m = OrderStatusResponse{} }
func (m *OrderStatusResponse) String() string { return proto.CompactTextString(m) }
func (*OrderStatusResponse) ProtoMessage()    {}
func (*OrderStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb21f6ae3f6980d6, []int{2}
}

func (m *OrderStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatusResponse.Unmarshal(m, b)
}
func (m *OrderStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatusResponse.Marshal(b, m, deterministic)
}
func (m *OrderStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatusResponse.Merge(m, src)
}
func (m *OrderStatusResponse) XXX_Size() int {
	return xxx_messageInfo_OrderStatusResponse.Size(m)
}
func (m *OrderStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatusResponse proto.InternalMessageInfo

func (m *OrderStatusResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrderStatusResponse) GetUsderId() string {
	if m != nil {
		return m.UsderId
	}
	return ""
}

func (m *OrderStatusResponse) GetOrder() *OrderAndExchangeType {
	if m != nil {
		return m.Order
	}
	return nil
}

type OrdersStatusesRequest struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrdersDescriptors    []*OrderDescriptorType `protobuf:"bytes,3,rep,name=orders_descriptors,json=ordersDescriptors,proto3" json:"orders_descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OrdersStatusesRequest) Reset()         { *m = OrdersStatusesRequest{} }
func (m *OrdersStatusesRequest) String() string { return proto.CompactTextString(m) }
func (*OrdersStatusesRequest) ProtoMessage()    {}
func (*OrdersStatusesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb21f6ae3f6980d6, []int{3}
}

func (m *OrdersStatusesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersStatusesRequest.Unmarshal(m, b)
}
func (m *OrdersStatusesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersStatusesRequest.Marshal(b, m, deterministic)
}
func (m *OrdersStatusesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersStatusesRequest.Merge(m, src)
}
func (m *OrdersStatusesRequest) XXX_Size() int {
	return xxx_messageInfo_OrdersStatusesRequest.Size(m)
}
func (m *OrdersStatusesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersStatusesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersStatusesRequest proto.InternalMessageInfo

func (m *OrdersStatusesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrdersStatusesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OrdersStatusesRequest) GetOrdersDescriptors() []*OrderDescriptorType {
	if m != nil {
		return m.OrdersDescriptors
	}
	return nil
}

type OrdersStatusesResponse struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UsderId              string                  `protobuf:"bytes,2,opt,name=usder_id,json=usderId,proto3" json:"usder_id,omitempty"`
	Order                []*OrderAndExchangeType `protobuf:"bytes,4,rep,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OrdersStatusesResponse) Reset()         { *m = OrdersStatusesResponse{} }
func (m *OrdersStatusesResponse) String() string { return proto.CompactTextString(m) }
func (*OrdersStatusesResponse) ProtoMessage()    {}
func (*OrdersStatusesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb21f6ae3f6980d6, []int{4}
}

func (m *OrdersStatusesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersStatusesResponse.Unmarshal(m, b)
}
func (m *OrdersStatusesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersStatusesResponse.Marshal(b, m, deterministic)
}
func (m *OrdersStatusesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersStatusesResponse.Merge(m, src)
}
func (m *OrdersStatusesResponse) XXX_Size() int {
	return xxx_messageInfo_OrdersStatusesResponse.Size(m)
}
func (m *OrdersStatusesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersStatusesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersStatusesResponse proto.InternalMessageInfo

func (m *OrdersStatusesResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrdersStatusesResponse) GetUsderId() string {
	if m != nil {
		return m.UsderId
	}
	return ""
}

func (m *OrdersStatusesResponse) GetOrder() []*OrderAndExchangeType {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterType((*PlaceOrderRequest)(nil), "orders.v1.PlaceOrderRequest")
	proto.RegisterType((*OrderStatusRequest)(nil), "orders.v1.OrderStatusRequest")
	proto.RegisterType((*OrderStatusResponse)(nil), "orders.v1.OrderStatusResponse")
	proto.RegisterType((*OrdersStatusesRequest)(nil), "orders.v1.OrdersStatusesRequest")
	proto.RegisterType((*OrdersStatusesResponse)(nil), "orders.v1.OrdersStatusesResponse")
}

func init() {
	proto.RegisterFile("orders/v1/orders_service.proto", fileDescriptor_cb21f6ae3f6980d6)
}

var fileDescriptor_cb21f6ae3f6980d6 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0xaf, 0xd2, 0x30,
	0x18, 0x75, 0x4c, 0xef, 0xbd, 0xfb, 0x6e, 0xfc, 0x71, 0x6b, 0xc4, 0x49, 0xf4, 0x3a, 0xf7, 0xc4,
	0x83, 0xd9, 0x00, 0x63, 0x8c, 0xf1, 0x49, 0xd1, 0x18, 0xa2, 0x22, 0x99, 0xfa, 0xe2, 0xcb, 0x32,
	0xda, 0x06, 0x96, 0x8c, 0x75, 0xb6, 0x1d, 0xa8, 0x7f, 0x81, 0x6f, 0xfe, 0xb5, 0xbe, 0x1b, 0xda,
	0x31, 0xd8, 0x08, 0x06, 0x22, 0x6f, 0xed, 0xe9, 0xd7, 0x73, 0xce, 0x77, 0xfa, 0xa5, 0x70, 0xc9,
	0x38, 0xa1, 0x5c, 0xf8, 0xf3, 0xae, 0xaf, 0x57, 0xa1, 0xa0, 0x7c, 0x1e, 0x63, 0xea, 0x65, 0x9c,
	0x49, 0x86, 0x2c, 0x8d, 0x7a, 0xf3, 0x6e, 0xab, 0x59, 0x2f, 0xd5, 0x25, 0x6e, 0x02, 0x17, 0xa3,
	0x24, 0xc2, 0xf4, 0xe3, 0x12, 0x0c, 0xe8, 0xb7, 0x9c, 0x0a, 0x89, 0x6e, 0x40, 0x23, 0x26, 0xb6,
	0xe1, 0x18, 0x6d, 0x2b, 0x68, 0xc4, 0x04, 0xdd, 0x83, 0xb3, 0x5c, 0x10, 0xca, 0xc3, 0x98, 0xd8,
	0x0d, 0x85, 0x9e, 0xaa, 0xfd, 0x80, 0xa0, 0xc7, 0x70, 0x4d, 0xf1, 0xd9, 0xa6, 0x63, 0xb4, 0xcf,
	0x7b, 0x4d, 0xaf, 0x94, 0xf4, 0xfa, 0x9c, 0x46, 0xb2, 0x20, 0xd6, 0x45, 0xee, 0x2f, 0x03, 0x90,
	0x02, 0x3e, 0xc9, 0x48, 0xe6, 0x62, 0x97, 0xde, 0x5d, 0x38, 0xcd, 0xc5, 0xa6, 0xdc, 0xc9, 0x72,
	0x3b, 0x20, 0x68, 0x00, 0xb7, 0x14, 0x51, 0x48, 0xa8, 0xc0, 0x3c, 0xce, 0x24, 0x5b, 0x09, 0x5f,
	0x6e, 0x08, 0x2b, 0x85, 0xd7, 0x65, 0xc5, 0xe7, 0x1f, 0x19, 0x0d, 0x6e, 0xb2, 0x2a, 0xe8, 0x2e,
	0xe0, 0x76, 0xc5, 0x89, 0xc8, 0x58, 0x2a, 0xe8, 0x21, 0xad, 0x3f, 0xad, 0xb6, 0xfe, 0xb0, 0xee,
	0xe0, 0x65, 0x4a, 0xde, 0x7c, 0xc7, 0xd3, 0x28, 0x9d, 0x50, 0x65, 0xa1, 0xc8, 0xe0, 0xb7, 0x01,
	0x77, 0xd4, 0xb9, 0xd0, 0xd2, 0xf4, 0xf0, 0x18, 0x3e, 0x00, 0x2a, 0xde, 0x7b, 0x9d, 0x83, 0xb0,
	0x4d, 0xc7, 0xdc, 0x23, 0x88, 0x0b, 0x7d, 0xbc, 0x46, 0x85, 0xfb, 0x13, 0x9a, 0x75, 0x43, 0xff,
	0x91, 0xc6, 0x55, 0x65, 0x63, 0xcf, 0x34, 0x7a, 0x7f, 0x4c, 0xb8, 0x5e, 0x88, 0xeb, 0xd1, 0x45,
	0xef, 0x01, 0xd6, 0x13, 0x89, 0xee, 0x6f, 0xf0, 0x6c, 0x0d, 0x6a, 0x6b, 0xab, 0xd9, 0xea, 0x6b,
	0xba, 0x57, 0xd0, 0x10, 0xce, 0xfb, 0x51, 0x8a, 0x69, 0xa2, 0xe9, 0x1e, 0xec, 0xba, 0xb0, 0x2f,
	0xdf, 0x3b, 0x38, 0x7b, 0x4b, 0xe5, 0x91, 0xc8, 0x86, 0x60, 0x8d, 0x58, 0x72, 0x24, 0x6b, 0x1d,
	0x03, 0x05, 0x60, 0xad, 0xcc, 0x09, 0xe4, 0xd4, 0x2f, 0xd4, 0xe7, 0xad, 0xf5, 0xe8, 0x1f, 0x15,
	0xa5, 0xc7, 0x2f, 0x00, 0xa5, 0xc7, 0x63, 0x91, 0x76, 0x8c, 0x57, 0xcf, 0xbf, 0x3e, 0x9b, 0xc4,
	0x72, 0x9a, 0x8f, 0x3d, 0xcc, 0x66, 0xfe, 0x22, 0xe2, 0xd9, 0x84, 0xfb, 0x92, 0x8c, 0xa3, 0x10,
	0xb3, 0xd9, 0x8c, 0xa5, 0xbe, 0xfa, 0x9e, 0x30, 0x4b, 0xfc, 0xf2, 0xe3, 0x7a, 0xa1, 0x57, 0xe3,
	0x13, 0x75, 0xf4, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x7a, 0xb3, 0x5c, 0xfe, 0x04,
	0x00, 0x00,
}
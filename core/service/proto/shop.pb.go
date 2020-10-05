// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/shop.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 店铺
type SShop struct {
	// * 店铺编号
	Id int64 `protobuf:"zigzag64,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 商户编号
	MerchantId int64 `protobuf:"zigzag64,2,opt,name=MerchantId,proto3" json:"MerchantId,omitempty"`
	// * 店铺名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// 店铺标题
	ShopTitle string `protobuf:"bytes,4,opt,name=ShopTitle,proto3" json:"ShopTitle,omitempty"`
	// 店铺公告
	ShopNotice string `protobuf:"bytes,5,opt,name=ShopNotice,proto3" json:"ShopNotice,omitempty"`
	// 标志
	Flag int32 `protobuf:"zigzag32,6,opt,name=Flag,proto3" json:"Flag,omitempty"`
	// 店铺设置
	Config *SShopConfig `protobuf:"bytes,7,opt,name=Config,proto3" json:"Config,omitempty"`
	// * 状态
	State                int32    `protobuf:"zigzag32,8,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShop) Reset()         { *m = SShop{} }
func (m *SShop) String() string { return proto.CompactTextString(m) }
func (*SShop) ProtoMessage()    {}
func (*SShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{0}
}
func (m *SShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShop.Unmarshal(m, b)
}
func (m *SShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShop.Marshal(b, m, deterministic)
}
func (dst *SShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShop.Merge(dst, src)
}
func (m *SShop) XXX_Size() int {
	return xxx_messageInfo_SShop.Size(m)
}
func (m *SShop) XXX_DiscardUnknown() {
	xxx_messageInfo_SShop.DiscardUnknown(m)
}

var xxx_messageInfo_SShop proto.InternalMessageInfo

func (m *SShop) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShop) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SShop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SShop) GetShopTitle() string {
	if m != nil {
		return m.ShopTitle
	}
	return ""
}

func (m *SShop) GetShopNotice() string {
	if m != nil {
		return m.ShopNotice
	}
	return ""
}

func (m *SShop) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SShop) GetConfig() *SShopConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *SShop) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

// 店铺设置
type SShopConfig struct {
	// * 店铺标志
	Logo string `protobuf:"bytes,4,opt,name=Logo,proto3" json:"Logo,omitempty"`
	// * 自定义 域名
	Host string `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host,omitempty"`
	// * 个性化域名
	Alias string `protobuf:"bytes,6,opt,name=Alias,proto3" json:"Alias,omitempty"`
	// * 电话
	Tel                  string   `protobuf:"bytes,7,opt,name=Tel,proto3" json:"Tel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShopConfig) Reset()         { *m = SShopConfig{} }
func (m *SShopConfig) String() string { return proto.CompactTextString(m) }
func (*SShopConfig) ProtoMessage()    {}
func (*SShopConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{1}
}
func (m *SShopConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShopConfig.Unmarshal(m, b)
}
func (m *SShopConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShopConfig.Marshal(b, m, deterministic)
}
func (dst *SShopConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShopConfig.Merge(dst, src)
}
func (m *SShopConfig) XXX_Size() int {
	return xxx_messageInfo_SShopConfig.Size(m)
}
func (m *SShopConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SShopConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SShopConfig proto.InternalMessageInfo

func (m *SShopConfig) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShopConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShopConfig) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShopConfig) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

// 店铺编号
type ShopId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopId) Reset()         { *m = ShopId{} }
func (m *ShopId) String() string { return proto.CompactTextString(m) }
func (*ShopId) ProtoMessage()    {}
func (*ShopId) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{2}
}
func (m *ShopId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopId.Unmarshal(m, b)
}
func (m *ShopId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopId.Marshal(b, m, deterministic)
}
func (dst *ShopId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopId.Merge(dst, src)
}
func (m *ShopId) XXX_Size() int {
	return xxx_messageInfo_ShopId.Size(m)
}
func (m *ShopId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopId.DiscardUnknown(m)
}

var xxx_messageInfo_ShopId proto.InternalMessageInfo

func (m *ShopId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 门店编号
type StoreId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreId) Reset()         { *m = StoreId{} }
func (m *StoreId) String() string { return proto.CompactTextString(m) }
func (*StoreId) ProtoMessage()    {}
func (*StoreId) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{3}
}
func (m *StoreId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreId.Unmarshal(m, b)
}
func (m *StoreId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreId.Marshal(b, m, deterministic)
}
func (dst *StoreId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreId.Merge(dst, src)
}
func (m *StoreId) XXX_Size() int {
	return xxx_messageInfo_StoreId.Size(m)
}
func (m *StoreId) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreId.DiscardUnknown(m)
}

var xxx_messageInfo_StoreId proto.InternalMessageInfo

func (m *StoreId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 店铺
type SStore struct {
	Id           int64  `protobuf:"zigzag64,1,opt,name=Id,proto3" json:"Id,omitempty"`
	MerchantId   int64  `protobuf:"zigzag64,2,opt,name=MerchantId,proto3" json:"MerchantId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Alias        string `protobuf:"bytes,4,opt,name=Alias,proto3" json:"Alias,omitempty"`
	Host         string `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host,omitempty"`
	Logo         string `protobuf:"bytes,6,opt,name=Logo,proto3" json:"Logo,omitempty"`
	State        int32  `protobuf:"zigzag32,7,opt,name=State,proto3" json:"State,omitempty"`
	OpeningState int32  `protobuf:"zigzag32,8,opt,name=OpeningState,proto3" json:"OpeningState,omitempty"`
	StorePhone   string `protobuf:"bytes,9,opt,name=StorePhone,proto3" json:"StorePhone,omitempty"`
	StoreTitle   string `protobuf:"bytes,10,opt,name=StoreTitle,proto3" json:"StoreTitle,omitempty"`
	StoreNotice  string `protobuf:"bytes,11,opt,name=StoreNotice,proto3" json:"StoreNotice,omitempty"`
	Province     int32  `protobuf:"varint,12,opt,name=Province,proto3" json:"Province,omitempty"`
	City         int32  `protobuf:"varint,13,opt,name=City,proto3" json:"City,omitempty"`
	District     int32  `protobuf:"varint,14,opt,name=District,proto3" json:"District,omitempty"`
	// 地区名称
	Address string `protobuf:"bytes,15,opt,name=Address,proto3" json:"Address,omitempty"`
	// 详细地址
	DetailAddress string `protobuf:"bytes,16,opt,name=DetailAddress,proto3" json:"DetailAddress,omitempty"`
	// 纬度
	Lat float64 `protobuf:"fixed64,17,opt,name=Lat,proto3" json:"Lat,omitempty"`
	// 经度
	Lng float64 `protobuf:"fixed64,18,opt,name=Lng,proto3" json:"Lng,omitempty"`
	// 覆盖范围(公里)
	CoverRadius int32 `protobuf:"varint,19,opt,name=CoverRadius,proto3" json:"CoverRadius,omitempty"`
	// 序号
	SortNum              int32    `protobuf:"varint,20,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SStore) Reset()         { *m = SStore{} }
func (m *SStore) String() string { return proto.CompactTextString(m) }
func (*SStore) ProtoMessage()    {}
func (*SStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{4}
}
func (m *SStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SStore.Unmarshal(m, b)
}
func (m *SStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SStore.Marshal(b, m, deterministic)
}
func (dst *SStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SStore.Merge(dst, src)
}
func (m *SStore) XXX_Size() int {
	return xxx_messageInfo_SStore.Size(m)
}
func (m *SStore) XXX_DiscardUnknown() {
	xxx_messageInfo_SStore.DiscardUnknown(m)
}

var xxx_messageInfo_SStore proto.InternalMessageInfo

func (m *SStore) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SStore) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SStore) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SStore) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SStore) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SStore) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SStore) GetOpeningState() int32 {
	if m != nil {
		return m.OpeningState
	}
	return 0
}

func (m *SStore) GetStorePhone() string {
	if m != nil {
		return m.StorePhone
	}
	return ""
}

func (m *SStore) GetStoreTitle() string {
	if m != nil {
		return m.StoreTitle
	}
	return ""
}

func (m *SStore) GetStoreNotice() string {
	if m != nil {
		return m.StoreNotice
	}
	return ""
}

func (m *SStore) GetProvince() int32 {
	if m != nil {
		return m.Province
	}
	return 0
}

func (m *SStore) GetCity() int32 {
	if m != nil {
		return m.City
	}
	return 0
}

func (m *SStore) GetDistrict() int32 {
	if m != nil {
		return m.District
	}
	return 0
}

func (m *SStore) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SStore) GetDetailAddress() string {
	if m != nil {
		return m.DetailAddress
	}
	return ""
}

func (m *SStore) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *SStore) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *SStore) GetCoverRadius() int32 {
	if m != nil {
		return m.CoverRadius
	}
	return 0
}

func (m *SStore) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

// 检查店铺结果
type CheckShopResponse struct {
	ShopId int64 `protobuf:"varint,1,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	// 店铺开通状态,0:未开通 1:已开通 2:待审核 3:审核不通过
	Status               int32    `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckShopResponse) Reset()         { *m = CheckShopResponse{} }
func (m *CheckShopResponse) String() string { return proto.CompactTextString(m) }
func (*CheckShopResponse) ProtoMessage()    {}
func (*CheckShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_bc80354b34e1cf9f, []int{5}
}
func (m *CheckShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckShopResponse.Unmarshal(m, b)
}
func (m *CheckShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckShopResponse.Marshal(b, m, deterministic)
}
func (dst *CheckShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckShopResponse.Merge(dst, src)
}
func (m *CheckShopResponse) XXX_Size() int {
	return xxx_messageInfo_CheckShopResponse.Size(m)
}
func (m *CheckShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckShopResponse proto.InternalMessageInfo

func (m *CheckShopResponse) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *CheckShopResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CheckShopResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*SShop)(nil), "SShop")
	proto.RegisterType((*SShopConfig)(nil), "SShopConfig")
	proto.RegisterType((*ShopId)(nil), "ShopId")
	proto.RegisterType((*StoreId)(nil), "StoreId")
	proto.RegisterType((*SStore)(nil), "SStore")
	proto.RegisterType((*CheckShopResponse)(nil), "CheckShopResponse")
}

func init() { proto.RegisterFile("message/shop.proto", fileDescriptor_shop_bc80354b34e1cf9f) }

var fileDescriptor_shop_bc80354b34e1cf9f = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xf3, 0xe1, 0xd4, 0x93, 0xb4, 0x34, 0x4b, 0x85, 0x56, 0x08, 0x15, 0xcb, 0xea, 0xc1,
	0xa7, 0x20, 0xc1, 0x91, 0x53, 0x49, 0x85, 0x88, 0x54, 0x42, 0xb5, 0xae, 0x38, 0x80, 0x38, 0x2c,
	0xf1, 0xe2, 0xac, 0xea, 0x78, 0x23, 0xef, 0xa6, 0x12, 0x7f, 0x83, 0x7f, 0xc7, 0xbf, 0x41, 0x33,
	0xfe, 0x60, 0x2b, 0xc1, 0xad, 0xa7, 0xcc, 0x7b, 0x6f, 0x3c, 0x79, 0x33, 0x7e, 0x06, 0xb6, 0x53,
	0xd6, 0xca, 0x42, 0xbd, 0xb2, 0x5b, 0xb3, 0x5f, 0xec, 0x6b, 0xe3, 0x4c, 0xf2, 0x3b, 0x80, 0x71,
	0x96, 0x6d, 0xcd, 0x9e, 0x9d, 0xc0, 0x60, 0x95, 0xf3, 0x20, 0x0e, 0x52, 0x26, 0x06, 0xab, 0x9c,
	0x9d, 0x03, 0x7c, 0x54, 0xf5, 0x66, 0x2b, 0x2b, 0xb7, 0xca, 0xf9, 0x80, 0x78, 0x8f, 0x61, 0x0c,
	0x46, 0x6b, 0xb9, 0x53, 0x7c, 0x18, 0x07, 0x69, 0x24, 0xa8, 0x66, 0x2f, 0x20, 0xc2, 0x59, 0xb7,
	0xda, 0x95, 0x8a, 0x8f, 0x48, 0xf8, 0x4b, 0xe0, 0x44, 0x04, 0x6b, 0xe3, 0xf4, 0x46, 0xf1, 0x31,
	0xc9, 0x1e, 0x83, 0x13, 0xdf, 0x97, 0xb2, 0xe0, 0x61, 0x1c, 0xa4, 0x73, 0x41, 0x35, 0xbb, 0x80,
	0x70, 0x69, 0xaa, 0x1f, 0xba, 0xe0, 0x93, 0x38, 0x48, 0xa7, 0xaf, 0x67, 0x0b, 0x72, 0xdb, 0x70,
	0xa2, 0xd5, 0xd8, 0x19, 0x8c, 0x33, 0x27, 0x9d, 0xe2, 0x47, 0xf4, 0x68, 0x03, 0x92, 0x6f, 0x30,
	0xf5, 0x9a, 0x71, 0xfc, 0xb5, 0x29, 0x4c, 0xeb, 0x8b, 0x6a, 0xe4, 0x3e, 0x18, 0xeb, 0x5a, 0x33,
	0x54, 0xe3, 0xb0, 0xcb, 0x52, 0x4b, 0x4b, 0x3e, 0x22, 0xd1, 0x00, 0x76, 0x0a, 0xc3, 0x5b, 0x55,
	0x92, 0x8b, 0x48, 0x60, 0x99, 0x9c, 0x43, 0x88, 0xd3, 0x57, 0x39, 0x3e, 0xf1, 0x59, 0x96, 0x07,
	0x45, 0xd7, 0x1b, 0x8a, 0x06, 0x24, 0x2f, 0x61, 0x92, 0x39, 0x53, 0xab, 0xff, 0x36, 0xfc, 0x1a,
	0x41, 0x98, 0x51, 0xcb, 0xa3, 0x1c, 0xbf, 0xf7, 0x3d, 0xf2, 0x7d, 0xff, 0x6b, 0xc3, 0xee, 0x12,
	0xa1, 0x77, 0x89, 0xfe, 0x84, 0x13, 0xef, 0x84, 0x2c, 0x81, 0xd9, 0xa7, 0xbd, 0xaa, 0x74, 0x55,
	0xf8, 0xf7, 0x7d, 0xc0, 0xd1, 0x6b, 0xc5, 0x25, 0x6e, 0xb6, 0xa6, 0x52, 0x3c, 0x6a, 0x5f, 0x6b,
	0xcf, 0xf4, 0x7a, 0x93, 0x0a, 0xf0, 0xf4, 0x26, 0x16, 0x31, 0x4c, 0x09, 0xb5, 0xb9, 0x98, 0x52,
	0x83, 0x4f, 0xb1, 0xe7, 0x70, 0x74, 0x53, 0x9b, 0x7b, 0x5d, 0x6d, 0x14, 0x9f, 0xc5, 0x41, 0x3a,
	0x16, 0x3d, 0xc6, 0x5d, 0x96, 0xda, 0xfd, 0xe4, 0xc7, 0xc4, 0x53, 0x8d, 0xfd, 0x57, 0xda, 0xba,
	0x5a, 0x6f, 0x1c, 0x3f, 0x69, 0xfa, 0x3b, 0xcc, 0x38, 0x4c, 0x2e, 0xf3, 0xbc, 0x56, 0xd6, 0xf2,
	0x27, 0xf4, 0x4f, 0x1d, 0x64, 0x17, 0x70, 0x7c, 0xa5, 0x9c, 0xd4, 0x65, 0xa7, 0x9f, 0x92, 0xfe,
	0x90, 0xc4, 0x1c, 0x5c, 0x4b, 0xc7, 0xe7, 0x71, 0x90, 0x06, 0x02, 0x4b, 0x62, 0xaa, 0x82, 0xb3,
	0x96, 0xa9, 0x0a, 0xdc, 0x68, 0x69, 0xee, 0x55, 0x2d, 0x64, 0xae, 0x0f, 0x96, 0x3f, 0x25, 0x0b,
	0x3e, 0x85, 0x2e, 0x32, 0x53, 0xbb, 0xf5, 0x61, 0xc7, 0xcf, 0x48, 0xed, 0x60, 0xf2, 0x15, 0xe6,
	0xcb, 0xad, 0xda, 0xdc, 0x61, 0xb4, 0x84, 0xb2, 0x7b, 0x53, 0x59, 0xc5, 0x9e, 0x75, 0x51, 0x6b,
	0x03, 0xd4, 0x05, 0x0f, 0x79, 0x27, 0xdd, 0xc1, 0x52, 0x44, 0xc6, 0xa2, 0x45, 0xc8, 0x0b, 0xb5,
	0x93, 0xf5, 0x5d, 0x1b, 0x90, 0x16, 0xbd, 0x8b, 0xbe, 0x4c, 0x16, 0x6f, 0xe9, 0xc3, 0xff, 0x1e,
	0xd2, 0xcf, 0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xa5, 0x21, 0x87, 0x15, 0x04, 0x00,
	0x00,
}
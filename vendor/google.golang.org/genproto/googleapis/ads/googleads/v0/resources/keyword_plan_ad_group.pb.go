// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/keyword_plan_ad_group.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Keyword Planner ad group.
type KeywordPlanAdGroup struct {
	// The resource name of the Keyword Planner ad group.
	// KeywordPlanAdGroup resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanAdGroups/{kp_ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan campaign to which this ad group belongs.
	KeywordPlanCampaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3" json:"keyword_plan_campaign,omitempty"`
	// The ID of the keyword plan ad group.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the keyword plan ad group.
	//
	// This field is required and should not be empty when creating keyword plan
	// ad group.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A default ad group max cpc bid in micros in account currency for all
	// biddable keywords under the keyword plan ad group.
	// If not set, will inherit from parent campaign.
	CpcBidMicros         *wrappers.Int64Value `protobuf:"bytes,5,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeywordPlanAdGroup) Reset()         { *m = KeywordPlanAdGroup{} }
func (m *KeywordPlanAdGroup) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanAdGroup) ProtoMessage()    {}
func (*KeywordPlanAdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_ad_group_f170de358444a70a, []int{0}
}
func (m *KeywordPlanAdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanAdGroup.Unmarshal(m, b)
}
func (m *KeywordPlanAdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanAdGroup.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanAdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanAdGroup.Merge(dst, src)
}
func (m *KeywordPlanAdGroup) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanAdGroup.Size(m)
}
func (m *KeywordPlanAdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanAdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanAdGroup proto.InternalMessageInfo

func (m *KeywordPlanAdGroup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanAdGroup) GetKeywordPlanCampaign() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlanCampaign
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanAdGroup)(nil), "google.ads.googleads.v0.resources.KeywordPlanAdGroup")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/keyword_plan_ad_group.proto", fileDescriptor_keyword_plan_ad_group_f170de358444a70a)
}

var fileDescriptor_keyword_plan_ad_group_f170de358444a70a = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xfb, 0x30,
	0x1c, 0xc5, 0x69, 0xb7, 0xdf, 0x0f, 0x8c, 0xd3, 0x8b, 0x8a, 0x58, 0x54, 0x64, 0x53, 0x84, 0x81,
	0x90, 0x16, 0x15, 0x6f, 0xc4, 0x8b, 0xce, 0x8b, 0xa1, 0xa2, 0x94, 0x09, 0xbb, 0x90, 0x42, 0xc9,
	0x92, 0x18, 0xc2, 0xda, 0x24, 0x24, 0xeb, 0x86, 0xaf, 0xe3, 0xa5, 0xe0, 0x8b, 0xf8, 0x0c, 0x3e,
	0x8c, 0xf4, 0xaf, 0xc8, 0xc0, 0x79, 0x77, 0x08, 0xe7, 0x73, 0xce, 0x21, 0x7c, 0xc1, 0x15, 0x93,
	0x92, 0x25, 0xd4, 0x43, 0xc4, 0x78, 0xa5, 0xcc, 0xd5, 0xdc, 0xf7, 0x34, 0x35, 0x32, 0xd3, 0x98,
	0x1a, 0x6f, 0x4a, 0x5f, 0x16, 0x52, 0x93, 0x58, 0x25, 0x48, 0xc4, 0x88, 0xc4, 0x4c, 0xcb, 0x4c,
	0x41, 0xa5, 0xe5, 0x4c, 0x3a, 0xbd, 0x92, 0x81, 0x88, 0x18, 0xd8, 0xe0, 0x70, 0xee, 0xc3, 0x06,
	0xdf, 0x3d, 0xa8, 0x1a, 0x0a, 0x60, 0x92, 0x3d, 0x7b, 0x0b, 0x8d, 0x94, 0xa2, 0xda, 0x94, 0x11,
	0x87, 0xef, 0x36, 0x70, 0xee, 0xca, 0x8a, 0x30, 0x41, 0x22, 0x20, 0xc3, 0x3c, 0xdf, 0x39, 0x02,
	0x1b, 0x75, 0x46, 0x2c, 0x50, 0x4a, 0x5d, 0xab, 0x6b, 0xf5, 0xd7, 0x46, 0x9d, 0xfa, 0xf1, 0x01,
	0xa5, 0xd4, 0x09, 0xc1, 0xf6, 0x8f, 0x75, 0x18, 0xa5, 0x0a, 0x71, 0x26, 0x5c, 0xbb, 0x6b, 0xf5,
	0xd7, 0x4f, 0xf7, 0xab, 0x4d, 0xb0, 0xee, 0x86, 0x8f, 0x33, 0xcd, 0x05, 0x1b, 0xa3, 0x24, 0xa3,
	0xa3, 0xad, 0xe9, 0x77, 0xeb, 0x75, 0x05, 0x3a, 0x27, 0xc0, 0xe6, 0xc4, 0x6d, 0x15, 0xf8, 0xde,
	0x12, 0x7e, 0x23, 0x66, 0x17, 0xe7, 0x25, 0x6d, 0x73, 0xe2, 0xf8, 0xa0, 0x5d, 0x4c, 0x6b, 0xff,
	0xa1, 0xad, 0x70, 0x3a, 0x01, 0xd8, 0xc4, 0x0a, 0xc7, 0x13, 0x4e, 0xe2, 0x94, 0x63, 0x2d, 0x8d,
	0xfb, 0x6f, 0x75, 0x55, 0x07, 0x2b, 0x3c, 0xe0, 0xe4, 0xbe, 0x00, 0x06, 0x9f, 0x16, 0x38, 0xc6,
	0x32, 0x85, 0x2b, 0x7f, 0x7e, 0xb0, 0xb3, 0xfc, 0xad, 0x61, 0x1e, 0x1f, 0x5a, 0x4f, 0xb7, 0x15,
	0xcd, 0x64, 0x82, 0x04, 0x83, 0x52, 0x33, 0x8f, 0x51, 0x51, 0x94, 0xd7, 0x67, 0xa0, 0xb8, 0xf9,
	0xe5, 0x2a, 0x2e, 0x1b, 0xf5, 0x6a, 0xb7, 0x86, 0x41, 0xf0, 0x66, 0xf7, 0x86, 0x65, 0x64, 0x40,
	0x0c, 0x2c, 0x65, 0xae, 0xc6, 0x3e, 0x1c, 0xd5, 0xce, 0x8f, 0xda, 0x13, 0x05, 0xc4, 0x44, 0x8d,
	0x27, 0x1a, 0xfb, 0x51, 0xe3, 0x99, 0xfc, 0x2f, 0x46, 0x9c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x77, 0xaa, 0xcb, 0x45, 0x99, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/recommendation_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum_RecommendationError int32

const (
	// Enum unspecified.
	RecommendationErrorEnum_UNSPECIFIED RecommendationErrorEnum_RecommendationError = 0
	// The received error code is not known in this version.
	RecommendationErrorEnum_UNKNOWN RecommendationErrorEnum_RecommendationError = 1
	// The specified budget amount is too low e.g. lower than minimum currency
	// unit or lower than ad group minimum cost-per-click.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_SMALL RecommendationErrorEnum_RecommendationError = 2
	// The specified budget amount is too large.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_LARGE RecommendationErrorEnum_RecommendationError = 3
	// The specified budget amount is not a valid amount. e.g. not a multiple
	// of minimum currency unit.
	RecommendationErrorEnum_INVALID_BUDGET_AMOUNT RecommendationErrorEnum_RecommendationError = 4
	// The specified keyword or ad violates ad policy.
	RecommendationErrorEnum_POLICY_ERROR RecommendationErrorEnum_RecommendationError = 5
	// The specified bid amount is not valid. e.g. too many fractional digits,
	// or negative amount.
	RecommendationErrorEnum_INVALID_BID_AMOUNT RecommendationErrorEnum_RecommendationError = 6
	// The number of keywords in ad group have reached the maximum allowed.
	RecommendationErrorEnum_ADGROUP_KEYWORD_LIMIT RecommendationErrorEnum_RecommendationError = 7
	// The recommendation requested to apply has already been applied.
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_APPLIED RecommendationErrorEnum_RecommendationError = 8
	// The recommendation requested to apply has been invalidated.
	RecommendationErrorEnum_RECOMMENDATION_INVALIDATED RecommendationErrorEnum_RecommendationError = 9
	// The number of operations in a single request exceeds the maximum allowed.
	RecommendationErrorEnum_TOO_MANY_OPERATIONS RecommendationErrorEnum_RecommendationError = 10
	// There are no operations in the request.
	RecommendationErrorEnum_NO_OPERATIONS RecommendationErrorEnum_RecommendationError = 11
	// Operations with multiple recommendation types are not supported when
	// partial failure mode is not enabled.
	RecommendationErrorEnum_DIFFERENT_TYPES_NOT_SUPPORTED RecommendationErrorEnum_RecommendationError = 12
	// Request contains multiple operations with the same resource_name.
	RecommendationErrorEnum_DUPLICATE_RESOURCE_NAME RecommendationErrorEnum_RecommendationError = 13
	// The recommendation requested to dismiss has already been dismissed.
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_DISMISSED RecommendationErrorEnum_RecommendationError = 14
)

var RecommendationErrorEnum_RecommendationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BUDGET_AMOUNT_TOO_SMALL",
	3:  "BUDGET_AMOUNT_TOO_LARGE",
	4:  "INVALID_BUDGET_AMOUNT",
	5:  "POLICY_ERROR",
	6:  "INVALID_BID_AMOUNT",
	7:  "ADGROUP_KEYWORD_LIMIT",
	8:  "RECOMMENDATION_ALREADY_APPLIED",
	9:  "RECOMMENDATION_INVALIDATED",
	10: "TOO_MANY_OPERATIONS",
	11: "NO_OPERATIONS",
	12: "DIFFERENT_TYPES_NOT_SUPPORTED",
	13: "DUPLICATE_RESOURCE_NAME",
	14: "RECOMMENDATION_ALREADY_DISMISSED",
}
var RecommendationErrorEnum_RecommendationError_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"BUDGET_AMOUNT_TOO_SMALL":          2,
	"BUDGET_AMOUNT_TOO_LARGE":          3,
	"INVALID_BUDGET_AMOUNT":            4,
	"POLICY_ERROR":                     5,
	"INVALID_BID_AMOUNT":               6,
	"ADGROUP_KEYWORD_LIMIT":            7,
	"RECOMMENDATION_ALREADY_APPLIED":   8,
	"RECOMMENDATION_INVALIDATED":       9,
	"TOO_MANY_OPERATIONS":              10,
	"NO_OPERATIONS":                    11,
	"DIFFERENT_TYPES_NOT_SUPPORTED":    12,
	"DUPLICATE_RESOURCE_NAME":          13,
	"RECOMMENDATION_ALREADY_DISMISSED": 14,
}

func (x RecommendationErrorEnum_RecommendationError) String() string {
	return proto.EnumName(RecommendationErrorEnum_RecommendationError_name, int32(x))
}
func (RecommendationErrorEnum_RecommendationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_error_bf8c3c97f0057f43, []int{0, 0}
}

// Container for enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendationErrorEnum) Reset()         { *m = RecommendationErrorEnum{} }
func (m *RecommendationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RecommendationErrorEnum) ProtoMessage()    {}
func (*RecommendationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_error_bf8c3c97f0057f43, []int{0}
}
func (m *RecommendationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationErrorEnum.Unmarshal(m, b)
}
func (m *RecommendationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *RecommendationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationErrorEnum.Merge(dst, src)
}
func (m *RecommendationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RecommendationErrorEnum.Size(m)
}
func (m *RecommendationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RecommendationErrorEnum)(nil), "google.ads.googleads.v1.errors.RecommendationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.RecommendationErrorEnum_RecommendationError", RecommendationErrorEnum_RecommendationError_name, RecommendationErrorEnum_RecommendationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/recommendation_error.proto", fileDescriptor_recommendation_error_bf8c3c97f0057f43)
}

var fileDescriptor_recommendation_error_bf8c3c97f0057f43 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0x0b, 0x1b, 0xb8, 0x1b, 0x18, 0x4f, 0x50, 0x18, 0x50, 0x41, 0xc5, 0x75, 0xa2,
	0x8a, 0x2b, 0xc2, 0x95, 0x5b, 0x9f, 0x56, 0xd6, 0x12, 0xdb, 0x72, 0x92, 0x4e, 0x45, 0x95, 0xac,
	0xb2, 0x54, 0x51, 0xa5, 0x35, 0xae, 0x92, 0xb2, 0x07, 0x82, 0x3b, 0x5e, 0x82, 0x7b, 0x1e, 0x85,
	0x5b, 0x5e, 0x00, 0x25, 0x59, 0x2b, 0x86, 0xba, 0x5d, 0xe5, 0xe8, 0xf8, 0xfb, 0x7f, 0x9f, 0x13,
	0xff, 0xe8, 0x63, 0x6a, 0x6d, 0x7a, 0xb9, 0x70, 0xe7, 0x49, 0xe1, 0xd6, 0x65, 0x59, 0x5d, 0xf5,
	0xdd, 0x45, 0x9e, 0xdb, 0xbc, 0x70, 0xf3, 0xc5, 0x85, 0x5d, 0xad, 0x16, 0x59, 0x32, 0xdf, 0x2c,
	0x6d, 0x66, 0xaa, 0xae, 0xb3, 0xce, 0xed, 0xc6, 0x92, 0x6e, 0xcd, 0x3b, 0xf3, 0xa4, 0x70, 0x76,
	0x52, 0xe7, 0xaa, 0xef, 0xd4, 0xd2, 0xd3, 0xd7, 0x5b, 0xeb, 0xf5, 0xd2, 0x9d, 0x67, 0x99, 0xdd,
	0x54, 0x16, 0x45, 0xad, 0xee, 0xfd, 0x6c, 0xa1, 0x8e, 0xbe, 0x61, 0x0e, 0xa5, 0x0c, 0xb2, 0xaf,
	0xab, 0xde, 0xf7, 0x16, 0x3a, 0xd9, 0x73, 0x46, 0x9e, 0xa0, 0x76, 0x2c, 0x42, 0x05, 0x43, 0x3e,
	0xe2, 0xc0, 0xf0, 0x3d, 0xd2, 0x46, 0x87, 0xb1, 0x38, 0x13, 0xf2, 0x5c, 0xe0, 0x06, 0x79, 0x85,
	0x3a, 0x83, 0x98, 0x8d, 0x21, 0x32, 0x34, 0x90, 0xb1, 0x88, 0x4c, 0x24, 0xa5, 0x09, 0x03, 0xea,
	0xfb, 0xb8, 0xb9, 0xff, 0xd0, 0xa7, 0x7a, 0x0c, 0xb8, 0x45, 0x5e, 0xa2, 0x67, 0x5c, 0x4c, 0xa8,
	0xcf, 0x99, 0xb9, 0x01, 0xe1, 0xfb, 0x04, 0xa3, 0x23, 0x25, 0x7d, 0x3e, 0x9c, 0x1a, 0xd0, 0x5a,
	0x6a, 0xfc, 0x80, 0x3c, 0x47, 0x64, 0x07, 0x73, 0xb6, 0x25, 0x0f, 0x4a, 0x13, 0xca, 0xc6, 0x5a,
	0xc6, 0xca, 0x9c, 0xc1, 0xf4, 0x5c, 0x6a, 0x66, 0x7c, 0x1e, 0xf0, 0x08, 0x1f, 0x92, 0x1e, 0xea,
	0x6a, 0x18, 0xca, 0x20, 0x00, 0xc1, 0x68, 0xc4, 0xa5, 0x30, 0xd4, 0xd7, 0x40, 0xd9, 0xd4, 0x50,
	0xa5, 0xfc, 0x72, 0x95, 0x87, 0xa4, 0x8b, 0x4e, 0xff, 0x63, 0xae, 0x6f, 0xa1, 0x11, 0x30, 0xfc,
	0x88, 0x74, 0xd0, 0x49, 0x39, 0x72, 0x40, 0xc5, 0xd4, 0x48, 0x05, 0xba, 0x62, 0x42, 0x8c, 0xc8,
	0x53, 0x74, 0x2c, 0xe4, 0xbf, 0xad, 0x36, 0x79, 0x87, 0xde, 0x30, 0x3e, 0x1a, 0x81, 0x86, 0x72,
	0xd1, 0xa9, 0x82, 0xd0, 0x08, 0x19, 0x99, 0x30, 0x56, 0x4a, 0xea, 0xd2, 0xee, 0xa8, 0xfc, 0x1f,
	0x2c, 0x56, 0x3e, 0x1f, 0xd2, 0x08, 0x8c, 0x86, 0x50, 0xc6, 0x7a, 0x08, 0x46, 0xd0, 0x00, 0xf0,
	0x31, 0x79, 0x8f, 0xde, 0xde, 0x32, 0x2f, 0xe3, 0x61, 0xc0, 0xc3, 0x10, 0x18, 0x7e, 0x3c, 0xf8,
	0xd3, 0x40, 0xbd, 0x0b, 0xbb, 0x72, 0xee, 0x8e, 0xc1, 0xe0, 0xc5, 0x9e, 0x97, 0x54, 0x65, 0x04,
	0x54, 0xe3, 0x33, 0xbb, 0xd6, 0xa6, 0xf6, 0x72, 0x9e, 0xa5, 0x8e, 0xcd, 0x53, 0x37, 0x5d, 0x64,
	0x55, 0x40, 0xb6, 0x69, 0x5c, 0x2f, 0x8b, 0xdb, 0xc2, 0xf9, 0xa9, 0xfe, 0x7c, 0x6b, 0xb6, 0xc6,
	0x94, 0xfe, 0x68, 0x76, 0xc7, 0xb5, 0x19, 0x4d, 0x0a, 0xa7, 0x2e, 0xcb, 0x6a, 0xd2, 0x77, 0xaa,
	0x2b, 0x8b, 0x5f, 0x5b, 0x60, 0x46, 0x93, 0x62, 0xb6, 0x03, 0x66, 0x93, 0xfe, 0xac, 0x06, 0x7e,
	0x37, 0x7b, 0x75, 0xd7, 0xf3, 0x68, 0x52, 0x78, 0xde, 0x0e, 0xf1, 0xbc, 0x49, 0xdf, 0xf3, 0x6a,
	0xe8, 0xcb, 0x41, 0x35, 0xdd, 0x87, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x0f, 0x94, 0x12,
	0x39, 0x03, 0x00, 0x00,
}
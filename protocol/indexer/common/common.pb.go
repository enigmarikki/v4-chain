// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/indexer/common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OrderRemovalReason is an enum of all the reasons an order was removed.
type OrderRemovalReason int32

const (
	// Default value, this is invalid and unused.
	OrderRemovalReason_ORDER_REMOVAL_REASON_UNSPECIFIED OrderRemovalReason = 0
	// The order was removed due to being expired.
	OrderRemovalReason_ORDER_REMOVAL_REASON_EXPIRED OrderRemovalReason = 1
	// The order was removed due to being canceled by a user.
	OrderRemovalReason_ORDER_REMOVAL_REASON_USER_CANCELED OrderRemovalReason = 2
	// The order was removed due to being undercollateralized.
	OrderRemovalReason_ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED OrderRemovalReason = 3
	// The order caused an internal error during order placement and was
	// removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_INTERNAL_ERROR OrderRemovalReason = 4
	// The order would have matched against another order placed by the same
	// subaccount and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_SELF_TRADE_ERROR OrderRemovalReason = 5
	// The order would have matched against maker orders on the orderbook
	// despite being a post-only order and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER OrderRemovalReason = 6
	// The order was an ICO order and would have been placed on the orderbook as
	// resting liquidity and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK OrderRemovalReason = 7
	// The order was a fill-or-kill order that could not be fully filled and was
	// removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED OrderRemovalReason = 8
	// The order was a reduce-only order that was removed due to either:
	// - being a taker order and fully-filling the order would flip the side of
	//    the subaccount's position, in this case the remaining size of the
	//    order is removed
	// - being a maker order resting on the book and being removed when either
	//    the subaccount's position is closed or flipped sides
	OrderRemovalReason_ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE OrderRemovalReason = 9
	// The order should be expired, according to the Indexer's cached data, but
	// the Indexer has yet to receive a message to remove the order. In order to
	// keep the data cached by the Indexer up-to-date and accurate, clear out
	// the data if it's expired by sending an order removal with this reason.
	// Protocol should never send this reason to Indexer.
	OrderRemovalReason_ORDER_REMOVAL_REASON_INDEXER_EXPIRED OrderRemovalReason = 10
	// The order has been replaced.
	OrderRemovalReason_ORDER_REMOVAL_REASON_REPLACED OrderRemovalReason = 11
)

var OrderRemovalReason_name = map[int32]string{
	0:  "ORDER_REMOVAL_REASON_UNSPECIFIED",
	1:  "ORDER_REMOVAL_REASON_EXPIRED",
	2:  "ORDER_REMOVAL_REASON_USER_CANCELED",
	3:  "ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED",
	4:  "ORDER_REMOVAL_REASON_INTERNAL_ERROR",
	5:  "ORDER_REMOVAL_REASON_SELF_TRADE_ERROR",
	6:  "ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER",
	7:  "ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK",
	8:  "ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED",
	9:  "ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE",
	10: "ORDER_REMOVAL_REASON_INDEXER_EXPIRED",
	11: "ORDER_REMOVAL_REASON_REPLACED",
}

var OrderRemovalReason_value = map[string]int32{
	"ORDER_REMOVAL_REASON_UNSPECIFIED":                            0,
	"ORDER_REMOVAL_REASON_EXPIRED":                                1,
	"ORDER_REMOVAL_REASON_USER_CANCELED":                          2,
	"ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED":                    3,
	"ORDER_REMOVAL_REASON_INTERNAL_ERROR":                         4,
	"ORDER_REMOVAL_REASON_SELF_TRADE_ERROR":                       5,
	"ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER":      6,
	"ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK": 7,
	"ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED":    8,
	"ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE":                     9,
	"ORDER_REMOVAL_REASON_INDEXER_EXPIRED":                        10,
	"ORDER_REMOVAL_REASON_REPLACED":                               11,
}

func (x OrderRemovalReason) String() string {
	return proto.EnumName(OrderRemovalReason_name, int32(x))
}

func (OrderRemovalReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc7f22a39c006b65, []int{0}
}

func init() {
	proto.RegisterEnum("dydxprotocol.indexer.common.OrderRemovalReason", OrderRemovalReason_name, OrderRemovalReason_value)
}

func init() {
	proto.RegisterFile("dydxprotocol/indexer/common/common.proto", fileDescriptor_cc7f22a39c006b65)
}

var fileDescriptor_cc7f22a39c006b65 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x13, 0x28, 0x05, 0xcc, 0xc5, 0xf2, 0x15, 0x88, 0x0a, 0x14, 0xba, 0x05, 0xb4, 0x7b,
	0x00, 0x21, 0x04, 0x48, 0xc8, 0x6b, 0x4f, 0xa4, 0x68, 0x9d, 0x38, 0x1a, 0x27, 0xd0, 0xee, 0x65,
	0xd4, 0x36, 0x11, 0x54, 0xea, 0x6e, 0x50, 0x28, 0x55, 0x79, 0x0b, 0xde, 0x87, 0x17, 0xe0, 0xd8,
	0x23, 0x47, 0xb4, 0xfb, 0x22, 0x88, 0x6c, 0x40, 0x80, 0xbc, 0x17, 0xfb, 0x30, 0xdf, 0xff, 0xcf,
	0xd8, 0xf3, 0xb3, 0x41, 0xf5, 0xb9, 0x3a, 0xff, 0xd0, 0x36, 0xa7, 0xcd, 0x51, 0x73, 0x32, 0x3a,
	0x9e, 0x57, 0xf5, 0x79, 0xdd, 0x8e, 0x8e, 0x9a, 0xd9, 0xac, 0x99, 0xf7, 0xd7, 0xb0, 0x2b, 0x8b,
	0x9b, 0x7f, 0x93, 0xc3, 0x9e, 0x1c, 0xae, 0x90, 0x87, 0x5f, 0x37, 0x98, 0xb0, 0x6d, 0x55, 0xb7,
	0x58, 0xcf, 0x9a, 0xb3, 0x83, 0x13, 0xac, 0x0f, 0x3e, 0x36, 0x73, 0xb1, 0xcd, 0xb6, 0x2c, 0x6a,
	0x40, 0x42, 0x48, 0xed, 0x1b, 0x69, 0x08, 0x41, 0x3a, 0x9b, 0x51, 0x99, 0xb9, 0x1c, 0x54, 0x12,
	0x27, 0xa0, 0x79, 0x20, 0xb6, 0xd8, 0x2d, 0x2f, 0x05, 0x7b, 0x79, 0x82, 0xa0, 0x79, 0x28, 0x1e,
	0xb0, 0xbb, 0x7e, 0x1f, 0x07, 0x48, 0x4a, 0x66, 0x0a, 0x0c, 0x68, 0x7e, 0x49, 0x3c, 0x66, 0x83,
	0x35, 0xfd, 0x34, 0xa0, 0xb2, 0xc6, 0xc8, 0x02, 0x50, 0x9a, 0x64, 0x0a, 0x9a, 0x5f, 0x16, 0x3b,
	0xec, 0x9e, 0x97, 0x4e, 0xb2, 0x02, 0x30, 0x93, 0x86, 0x00, 0xd1, 0x22, 0xdf, 0x10, 0xbb, 0xec,
	0xbe, 0x17, 0x74, 0x60, 0x62, 0x2a, 0x50, 0x6a, 0xe8, 0xd1, 0x2b, 0xe2, 0x05, 0x7b, 0xe6, 0x45,
	0x73, 0xeb, 0x0a, 0xb2, 0x99, 0xd9, 0xa7, 0xb7, 0xb6, 0x34, 0x9a, 0x14, 0x5a, 0xe7, 0x28, 0x95,
	0x13, 0x40, 0xea, 0x04, 0x7c, 0x53, 0xbc, 0x66, 0x2f, 0xfd, 0xf3, 0xa4, 0x29, 0xe8, 0x44, 0x16,
	0x40, 0xf6, 0xf7, 0x6b, 0x7b, 0x17, 0x84, 0xce, 0x95, 0xc6, 0xd6, 0x4e, 0xf8, 0x55, 0xf1, 0x8a,
	0x3d, 0xf7, 0x1a, 0xc4, 0x76, 0xb2, 0x6a, 0x42, 0xaa, 0x93, 0x65, 0xb6, 0xa0, 0x31, 0x50, 0x5c,
	0x1a, 0xb3, 0xdf, 0x9d, 0xa0, 0xf9, 0x35, 0xf1, 0x88, 0xed, 0x78, 0xd5, 0x08, 0xba, 0x54, 0xb0,
	0x1a, 0x1e, 0xc1, 0x25, 0x53, 0xe0, 0xd7, 0xc5, 0x80, 0x6d, 0xaf, 0xf9, 0x3b, 0x0d, 0x7b, 0x80,
	0x7f, 0x76, 0xc7, 0xc4, 0x1d, 0x76, 0x7b, 0x8d, 0x6d, 0x6e, 0xa4, 0x02, 0xcd, 0x6f, 0x8c, 0xd5,
	0xb7, 0x45, 0x14, 0x5e, 0x2c, 0xa2, 0xf0, 0xc7, 0x22, 0x0a, 0xbf, 0x2c, 0xa3, 0xe0, 0x62, 0x19,
	0x05, 0xdf, 0x97, 0x51, 0x30, 0xdd, 0x7d, 0x77, 0x7c, 0xfa, 0xfe, 0xd3, 0xe1, 0xaf, 0x98, 0x8d,
	0xfe, 0x49, 0xea, 0xd9, 0xd3, 0xff, 0xc2, 0x7a, 0xb8, 0xd9, 0xd5, 0x9e, 0xfc, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xef, 0x7d, 0x7e, 0xd2, 0x02, 0x00, 0x00,
}
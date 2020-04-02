// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter/enum.proto

package filter

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type HistoryEventFilterType int32

const (
	HistoryEventFilterTypeAllEvent   HistoryEventFilterType = 0
	HistoryEventFilterTypeCloseEvent HistoryEventFilterType = 1
)

var HistoryEventFilterType_name = map[int32]string{
	0: "HistoryEventFilterTypeAllEvent",
	1: "HistoryEventFilterTypeCloseEvent",
}

var HistoryEventFilterType_value = map[string]int32{
	"HistoryEventFilterTypeAllEvent":   0,
	"HistoryEventFilterTypeCloseEvent": 1,
}

func (HistoryEventFilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a536f7c83c29d2f9, []int{0}
}

func init() {
	proto.RegisterEnum("filter.HistoryEventFilterType", HistoryEventFilterType_name, HistoryEventFilterType_value)
}

func init() { proto.RegisterFile("filter/enum.proto", fileDescriptor_a536f7c83c29d2f9) }

var fileDescriptor_a536f7c83c29d2f9 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0xd2, 0x4f, 0xcd, 0x2b, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0x08, 0x69, 0x25, 0x71, 0x89, 0x79, 0x64, 0x16, 0x97, 0xe4, 0x17, 0x55, 0xba, 0x96, 0xa5, 0xe6,
	0x95, 0xb8, 0x81, 0x45, 0x43, 0x2a, 0x0b, 0x52, 0x85, 0x94, 0xb8, 0xe4, 0xb0, 0xcb, 0x38, 0xe6,
	0xe4, 0x80, 0x45, 0x04, 0x18, 0x84, 0x54, 0xb8, 0x14, 0xb0, 0xab, 0x71, 0xce, 0xc9, 0x2f, 0x4e,
	0x85, 0xa8, 0x62, 0x74, 0x2a, 0xbb, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f,
	0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0x25, 0x91, 0x99,
	0xaf, 0x57, 0x92, 0x9a, 0x5b, 0x90, 0x5f, 0x94, 0x98, 0x03, 0x71, 0xb4, 0x1e, 0xc4, 0xcd, 0x01,
	0x8c, 0x51, 0x2a, 0xe9, 0x48, 0x72, 0x99, 0xf9, 0xfa, 0x30, 0xb6, 0x2e, 0x58, 0x9d, 0x3e, 0x44,
	0x5d, 0x12, 0x1b, 0x98, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x47, 0x61, 0x74, 0xb6, 0xff,
	0x00, 0x00, 0x00,
}

func (x HistoryEventFilterType) String() string {
	s, ok := HistoryEventFilterType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

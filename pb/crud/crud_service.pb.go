// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/crud_service.proto

/*
Package crud is a generated protocol buffer package.

It is generated from these files:
	pb/crud_service.proto

It has these top-level messages:
	Key
	Record
	UpsertResponse
	DeleteResponse
*/
package crud

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

type Key struct {
	MovieId int32 `protobuf:"varint,1,opt,name=movieId" json:"movieId,omitempty"`
	UserId  int32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Key) GetMovieId() int32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *Key) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Record struct {
	MovieId int32   `protobuf:"varint,1,opt,name=movieId" json:"movieId,omitempty"`
	UserId  int32   `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rating  float32 `protobuf:"fixed32,3,opt,name=rating" json:"rating,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetMovieId() int32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *Record) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Record) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type UpsertResponse struct {
}

func (m *UpsertResponse) Reset()                    { *m = UpsertResponse{} }
func (m *UpsertResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertResponse) ProtoMessage()               {}
func (*UpsertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Key)(nil), "crud.Key")
	proto.RegisterType((*Record)(nil), "crud.Record")
	proto.RegisterType((*UpsertResponse)(nil), "crud.UpsertResponse")
	proto.RegisterType((*DeleteResponse)(nil), "crud.DeleteResponse")
}

func init() { proto.RegisterFile("pb/crud_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xdb, 0x6d, 0x56, 0x7c, 0x8a, 0x48, 0x50, 0x29, 0x3b, 0xcd, 0x9c, 0x06, 0xb2, 0x04,
	0xf4, 0xe0, 0x5d, 0x77, 0x19, 0xbb, 0x45, 0x76, 0xf1, 0x22, 0x69, 0xf2, 0xac, 0x05, 0x97, 0x94,
	0x24, 0xad, 0xf4, 0x43, 0xf8, 0x9d, 0xa5, 0x4b, 0x07, 0xce, 0xdb, 0x8e, 0xbf, 0xdf, 0xcb, 0xff,
	0xe5, 0xbd, 0x04, 0x6e, 0xea, 0x82, 0x2b, 0xd7, 0xe8, 0x77, 0x8f, 0xae, 0xad, 0x14, 0xb2, 0xda,
	0xd9, 0x60, 0xc9, 0xa4, 0x77, 0xf4, 0x09, 0xc6, 0x6b, 0xec, 0x48, 0x0e, 0xa7, 0x5b, 0xdb, 0x56,
	0xb8, 0xd2, 0x79, 0x3a, 0x4b, 0xe7, 0x27, 0x62, 0x8f, 0xe4, 0x16, 0xb2, 0xc6, 0xa3, 0x5b, 0xe9,
	0x7c, 0xb4, 0x2b, 0x0c, 0x44, 0x05, 0x64, 0x02, 0x95, 0x75, 0xfa, 0xf8, 0x6c, 0xef, 0x9d, 0x0c,
	0x95, 0x29, 0xf3, 0xf1, 0x2c, 0x9d, 0x8f, 0xc4, 0x40, 0xf4, 0x0a, 0x2e, 0x37, 0xb5, 0x47, 0x17,
	0x04, 0xfa, 0xda, 0x1a, 0x8f, 0xbd, 0x59, 0xe2, 0x17, 0x06, 0xdc, 0x9b, 0x87, 0x9f, 0x14, 0xce,
	0x5f, 0xc4, 0x66, 0xf9, 0x1a, 0x97, 0x21, 0x0c, 0xb2, 0x98, 0x21, 0x17, 0xac, 0xdf, 0x88, 0xc5,
	0xa9, 0xa6, 0xd7, 0x91, 0xfe, 0xf5, 0x4b, 0xc8, 0x1d, 0x4c, 0x04, 0x4a, 0x4d, 0xce, 0x62, 0x7d,
	0x8d, 0xdd, 0xf4, 0x20, 0x48, 0x13, 0x72, 0x0f, 0x59, 0xbc, 0xf4, 0xef, 0xa1, 0xa1, 0xdf, 0xe1,
	0x34, 0x34, 0x79, 0xe6, 0x6f, 0x8b, 0xb2, 0x0a, 0x9f, 0x4d, 0xc1, 0x94, 0xdd, 0xf2, 0x0f, 0x87,
	0x5a, 0x77, 0x65, 0xcb, 0x95, 0xf4, 0x5e, 0x1a, 0xed, 0xe4, 0xe2, 0x5b, 0x1a, 0x23, 0x0b, 0xe4,
	0xc3, 0x2f, 0x14, 0xd9, 0xee, 0xf9, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x07, 0xbb,
	0xad, 0x97, 0x01, 0x00, 0x00,
}
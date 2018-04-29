// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

/*
Package practical_grpc_v1 is a generated protocol buffer package.

It is generated from these files:
	database.proto
	archiver.proto
	tokenizer.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
	ZipRequest
	ZipResponse
	TokenizeRequest
	TokenizeResponse
*/
package practical_grpc_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SearchRequest struct {
	Term       string `protobuf:"bytes,1,opt,name=term" json:"term,omitempty"`
	MaxResults int64  `protobuf:"varint,2,opt,name=max_results,json=maxResults" json:"max_results,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *SearchRequest) GetMaxResults() int64 {
	if m != nil {
		return m.MaxResults
	}
	return 0
}

type SearchResponse struct {
	MatchedTerm string `protobuf:"bytes,1,opt,name=matched_term,json=matchedTerm" json:"matched_term,omitempty"`
	Content     string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Rank        int32  `protobuf:"varint,3,opt,name=rank" json:"rank,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetMatchedTerm() string {
	if m != nil {
		return m.MatchedTerm
	}
	return ""
}

func (m *SearchResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SearchResponse) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "practical.grpc.v1.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "practical.grpc.v1.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Database service

type DatabaseClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (Database_SearchClient, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (Database_SearchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Database_serviceDesc.Streams[0], c.cc, "/practical.grpc.v1.Database/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Database_SearchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type databaseSearchClient struct {
	grpc.ClientStream
}

func (x *databaseSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Database service

type DatabaseServer interface {
	Search(*SearchRequest, Database_SearchServer) error
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseServer).Search(m, &databaseSearchServer{stream})
}

type Database_SearchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type databaseSearchServer struct {
	grpc.ServerStream
}

func (x *databaseSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "practical.grpc.v1.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _Database_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "database.proto",
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x89, 0xd5, 0x6a, 0xa7, 0x5a, 0x30, 0xa7, 0xe0, 0xc5, 0xb4, 0xa7, 0x3d, 0x05, 0xff,
	0x7c, 0x85, 0xde, 0x85, 0xe8, 0xcd, 0x43, 0x99, 0xa6, 0x83, 0x15, 0x9b, 0x3f, 0x26, 0x53, 0xd9,
	0x8f, 0x2f, 0x9b, 0xdd, 0x05, 0x45, 0xf0, 0xf6, 0xe6, 0x31, 0xfc, 0xe6, 0xbd, 0x81, 0xc5, 0x0e,
	0x19, 0xb7, 0x58, 0xc8, 0xa4, 0x1c, 0x39, 0xca, 0xeb, 0x94, 0xd1, 0xf1, 0xbb, 0xc3, 0x83, 0x79,
	0xcb, 0xc9, 0x99, 0xaf, 0xfb, 0xd5, 0x1a, 0xae, 0x9e, 0x09, 0xb3, 0xdb, 0x5b, 0xfa, 0x3c, 0x52,
	0x61, 0x29, 0xe1, 0x94, 0x29, 0x7b, 0x25, 0xb4, 0x68, 0x66, 0xb6, 0x6a, 0x79, 0x0b, 0x73, 0x8f,
	0xed, 0x26, 0x53, 0x39, 0x1e, 0xb8, 0xa8, 0x13, 0x2d, 0x9a, 0x89, 0x05, 0x8f, 0xad, 0xed, 0x9d,
	0x15, 0xc2, 0x62, 0xa4, 0x94, 0x14, 0x43, 0x21, 0xb9, 0x84, 0x4b, 0x8f, 0xec, 0xf6, 0xb4, 0xdb,
	0xfc, 0xc0, 0xcd, 0x07, 0xef, 0xa5, 0xa3, 0x2a, 0x38, 0x77, 0x31, 0x30, 0x05, 0xae, 0xc4, 0x99,
	0x1d, 0xc7, 0x2e, 0x43, 0xc6, 0xf0, 0xa1, 0x26, 0x5a, 0x34, 0x67, 0xb6, 0xea, 0x87, 0x57, 0xb8,
	0x58, 0x0f, 0x6d, 0xe4, 0x13, 0x4c, 0xfb, 0x73, 0x52, 0x9b, 0x3f, 0x95, 0xcc, 0xaf, 0x3e, 0x37,
	0xcb, 0x7f, 0x36, 0xfa, 0xac, 0x77, 0x62, 0x3b, 0xad, 0xff, 0x79, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0x74, 0x31, 0x96, 0x31, 0x01, 0x00, 0x00,
}
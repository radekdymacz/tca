// Code generated by protoc-gen-go.
// source: name.proto
// DO NOT EDIT!

/*
Package name is a generated protocol buffer package.

It is generated from these files:
	name.proto

It has these top-level messages:
	Request
	Response
*/
package name

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Response example {"name":"Δαμέας","surname":"Γιάνναρης","gender":"male","region":"Greece"}
type Response struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Surname string `protobuf:"bytes,2,opt,name=surname" json:"surname,omitempty"`
	Gender  string `protobuf:"bytes,3,opt,name=gender" json:"gender,omitempty"`
	Region  string `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Response) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Response) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "name.Request")
	proto.RegisterType((*Response)(nil), "name.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Publisher API

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

type publisher struct {
	c     client.Client
	topic string
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	return p.c.Publish(ctx, p.c.NewPublication(p.topic, msg), opts...)
}

func NewPublisher(topic string, c client.Client) Publisher {
	if c == nil {
		c = client.NewClient()
	}
	return &publisher{c, topic}
}

// Subscriber API

func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}

// Client API for Name service

type NameClient interface {
	Random(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type nameClient struct {
	c           client.Client
	serviceName string
}

func NewNameClient(serviceName string, c client.Client) NameClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "name"
	}
	return &nameClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *nameClient) Random(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Name.Random", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Name service

type NameHandler interface {
	Random(context.Context, *Request, *Response) error
}

func RegisterNameHandler(s server.Server, hdlr NameHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Name{hdlr}, opts...))
}

type Name struct {
	NameHandler
}

func (h *Name) Random(ctx context.Context, in *Request, out *Response) error {
	return h.NameHandler.Random(ctx, in, out)
}

func init() { proto.RegisterFile("name.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0x8e, 0x31, 0x0e, 0xc2, 0x30,
	0x0c, 0x45, 0x29, 0x54, 0x29, 0xb5, 0x04, 0x83, 0x07, 0x14, 0x31, 0xa1, 0x4c, 0xb0, 0x54, 0x02,
	0xee, 0xc1, 0x90, 0x1b, 0x14, 0xd5, 0x2a, 0x0c, 0xb5, 0x4b, 0xd2, 0xde, 0x1f, 0xc5, 0x69, 0x37,
	0xbf, 0xf7, 0x65, 0xe9, 0x01, 0x70, 0x3b, 0x50, 0x33, 0x06, 0x99, 0x04, 0xcb, 0x74, 0xbb, 0x1a,
	0x2a, 0x4f, 0xbf, 0x99, 0xe2, 0xe4, 0x3e, 0xb0, 0xf7, 0x14, 0x47, 0xe1, 0x48, 0x88, 0xa0, 0xb3,
	0x2d, 0x2e, 0xc5, 0xb5, 0xf6, 0x7a, 0xa3, 0x85, 0x2a, 0xce, 0x41, 0xf5, 0x56, 0xf5, 0x8a, 0x78,
	0x02, 0xd3, 0x13, 0x77, 0x14, 0xec, 0x4e, 0x87, 0x85, 0x92, 0x0f, 0xd4, 0x7f, 0x85, 0x6d, 0x99,
	0x7d, 0xa6, 0xc7, 0x1d, 0xca, 0x57, 0xfa, 0xbb, 0x81, 0xf1, 0x2d, 0x77, 0x32, 0xe0, 0xa1, 0xd1,
	0xb2, 0x25, 0xe5, 0x7c, 0x5c, 0x31, 0xe7, 0xb8, 0xcd, 0xdb, 0x68, 0xf4, 0xf3, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xf5, 0x96, 0xc5, 0xc2, 0x00, 0x00, 0x00,
}
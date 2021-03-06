// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	event.proto

It has these top-level messages:
	Event
	EventsList
	Filter
	EventId
	LoadEventsRequest
	LoadEventsResponse
	StreamEventsRequest
	StreamEventsResponse
	StreamId
	Void
*/
package server

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Severity int32

const (
	Severity_DEBUG    Severity = 0
	Severity_INFO     Severity = 1
	Severity_WARN     Severity = 2
	Severity_ERROR    Severity = 3
	Severity_CRITICAL Severity = 4
)

var Severity_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARN",
	3: "ERROR",
	4: "CRITICAL",
}
var Severity_value = map[string]int32{
	"DEBUG":    0,
	"INFO":     1,
	"WARN":     2,
	"ERROR":    3,
	"CRITICAL": 4,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}
func (Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptorEvent, []int{0} }

type Event struct {
	Id        int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg       string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	System    string   `protobuf:"bytes,4,opt,name=system,proto3" json:"system,omitempty"`
	Severity  Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=server.Severity" json:"severity,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{0} }

func (m *Event) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Event) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_DEBUG
}

type EventsList struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *EventsList) Reset()                    { *m = EventsList{} }
func (m *EventsList) String() string            { return proto.CompactTextString(m) }
func (*EventsList) ProtoMessage()               {}
func (*EventsList) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{1} }

func (m *EventsList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Filter struct {
	System     string `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	SeverityId int32  `protobuf:"varint,2,opt,name=severityId,proto3" json:"severityId,omitempty"`
	Msg        string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	BaseDate   int64  `protobuf:"varint,4,opt,name=baseDate,proto3" json:"baseDate,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{2} }

func (m *Filter) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Filter) GetSeverityId() int32 {
	if m != nil {
		return m.SeverityId
	}
	return 0
}

func (m *Filter) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Filter) GetBaseDate() int64 {
	if m != nil {
		return m.BaseDate
	}
	return 0
}

type EventId struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *EventId) Reset()                    { *m = EventId{} }
func (m *EventId) String() string            { return proto.CompactTextString(m) }
func (*EventId) ProtoMessage()               {}
func (*EventId) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{3} }

func (m *EventId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LoadEventsRequest struct {
	Filter *Filter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
}

func (m *LoadEventsRequest) Reset()                    { *m = LoadEventsRequest{} }
func (m *LoadEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadEventsRequest) ProtoMessage()               {}
func (*LoadEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{4} }

func (m *LoadEventsRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type LoadEventsResponse struct {
	Events *EventsList `protobuf:"bytes,1,opt,name=events" json:"events,omitempty"`
}

func (m *LoadEventsResponse) Reset()                    { *m = LoadEventsResponse{} }
func (m *LoadEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadEventsResponse) ProtoMessage()               {}
func (*LoadEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{5} }

func (m *LoadEventsResponse) GetEvents() *EventsList {
	if m != nil {
		return m.Events
	}
	return nil
}

type StreamEventsRequest struct {
	Filter *Filter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
}

func (m *StreamEventsRequest) Reset()                    { *m = StreamEventsRequest{} }
func (m *StreamEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamEventsRequest) ProtoMessage()               {}
func (*StreamEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{6} }

func (m *StreamEventsRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type StreamEventsResponse struct {
	Event    *Event    `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	StreamId *StreamId `protobuf:"bytes,2,opt,name=streamId" json:"streamId,omitempty"`
	Filter   *Filter   `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (m *StreamEventsResponse) Reset()                    { *m = StreamEventsResponse{} }
func (m *StreamEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamEventsResponse) ProtoMessage()               {}
func (*StreamEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{7} }

func (m *StreamEventsResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamEventsResponse) GetStreamId() *StreamId {
	if m != nil {
		return m.StreamId
	}
	return nil
}

func (m *StreamEventsResponse) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type StreamId struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *StreamId) Reset()                    { *m = StreamId{} }
func (m *StreamId) String() string            { return proto.CompactTextString(m) }
func (*StreamId) ProtoMessage()               {}
func (*StreamId) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{8} }

func (m *StreamId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{9} }

func init() {
	proto.RegisterType((*Event)(nil), "server.Event")
	proto.RegisterType((*EventsList)(nil), "server.EventsList")
	proto.RegisterType((*Filter)(nil), "server.Filter")
	proto.RegisterType((*EventId)(nil), "server.EventId")
	proto.RegisterType((*LoadEventsRequest)(nil), "server.LoadEventsRequest")
	proto.RegisterType((*LoadEventsResponse)(nil), "server.LoadEventsResponse")
	proto.RegisterType((*StreamEventsRequest)(nil), "server.StreamEventsRequest")
	proto.RegisterType((*StreamEventsResponse)(nil), "server.StreamEventsResponse")
	proto.RegisterType((*StreamId)(nil), "server.StreamId")
	proto.RegisterType((*Void)(nil), "server.Void")
	proto.RegisterEnum("server.Severity", Severity_name, Severity_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventLogger service

type EventLoggerClient interface {
	LoadEvents(ctx context.Context, in *LoadEventsRequest, opts ...grpc.CallOption) (*LoadEventsResponse, error)
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (EventLogger_StreamEventsClient, error)
	StopStreaming(ctx context.Context, in *StreamId, opts ...grpc.CallOption) (*Void, error)
}

type eventLoggerClient struct {
	cc *grpc.ClientConn
}

func NewEventLoggerClient(cc *grpc.ClientConn) EventLoggerClient {
	return &eventLoggerClient{cc}
}

func (c *eventLoggerClient) LoadEvents(ctx context.Context, in *LoadEventsRequest, opts ...grpc.CallOption) (*LoadEventsResponse, error) {
	out := new(LoadEventsResponse)
	err := grpc.Invoke(ctx, "/server.EventLogger/LoadEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventLoggerClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (EventLogger_StreamEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventLogger_serviceDesc.Streams[0], c.cc, "/server.EventLogger/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventLoggerStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventLogger_StreamEventsClient interface {
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type eventLoggerStreamEventsClient struct {
	grpc.ClientStream
}

func (x *eventLoggerStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventLoggerClient) StopStreaming(ctx context.Context, in *StreamId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/server.EventLogger/StopStreaming", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventLogger service

type EventLoggerServer interface {
	LoadEvents(context.Context, *LoadEventsRequest) (*LoadEventsResponse, error)
	StreamEvents(*StreamEventsRequest, EventLogger_StreamEventsServer) error
	StopStreaming(context.Context, *StreamId) (*Void, error)
}

func RegisterEventLoggerServer(s *grpc.Server, srv EventLoggerServer) {
	s.RegisterService(&_EventLogger_serviceDesc, srv)
}

func _EventLogger_LoadEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventLoggerServer).LoadEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.EventLogger/LoadEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventLoggerServer).LoadEvents(ctx, req.(*LoadEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventLogger_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventLoggerServer).StreamEvents(m, &eventLoggerStreamEventsServer{stream})
}

type EventLogger_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	grpc.ServerStream
}

type eventLoggerStreamEventsServer struct {
	grpc.ServerStream
}

func (x *eventLoggerStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventLogger_StopStreaming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventLoggerServer).StopStreaming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.EventLogger/StopStreaming",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventLoggerServer).StopStreaming(ctx, req.(*StreamId))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventLogger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.EventLogger",
	HandlerType: (*EventLoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadEvents",
			Handler:    _EventLogger_LoadEvents_Handler,
		},
		{
			MethodName: "StopStreaming",
			Handler:    _EventLogger_StopStreaming_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _EventLogger_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptorEvent) }

var fileDescriptorEvent = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xed, 0x34, 0x4d, 0x4c, 0x6f, 0xba, 0x25, 0x5e, 0x45, 0xb2, 0x71, 0x91, 0x12, 0x51, 0xca,
	0x22, 0x45, 0xbb, 0x8f, 0x22, 0xb8, 0x1f, 0x5d, 0x09, 0xd4, 0x5d, 0x98, 0xfa, 0xf1, 0x9c, 0x25,
	0x63, 0x09, 0x98, 0xa6, 0x66, 0xc6, 0xc2, 0xfe, 0x08, 0x1f, 0xfc, 0x73, 0xfe, 0x1e, 0xc9, 0xcd,
	0x4c, 0x9b, 0x6c, 0x17, 0x1f, 0x7c, 0x9b, 0xb9, 0xe7, 0x9e, 0xb9, 0xe7, 0x9c, 0x9b, 0x80, 0x27,
	0x36, 0x62, 0xa5, 0x26, 0xeb, 0xb2, 0x50, 0x05, 0x3a, 0x52, 0x94, 0x1b, 0x51, 0x46, 0xbf, 0x18,
	0xd8, 0xb3, 0xaa, 0x8e, 0x43, 0xe8, 0x66, 0x69, 0xc0, 0x46, 0x6c, 0x6c, 0xf1, 0x6e, 0x96, 0xa2,
	0x0f, 0x56, 0x2e, 0x97, 0x41, 0x77, 0xc4, 0xc6, 0x7d, 0x5e, 0x1d, 0xf1, 0x08, 0xfa, 0x2a, 0xcb,
	0x85, 0x54, 0x49, 0xbe, 0x0e, 0x2c, 0x6a, 0xdc, 0x15, 0xf0, 0x09, 0x38, 0xf2, 0x56, 0x2a, 0x91,
	0x07, 0x3d, 0xa2, 0xe8, 0x1b, 0xbe, 0x02, 0x57, 0x8a, 0x8d, 0x28, 0x33, 0x75, 0x1b, 0xd8, 0x23,
	0x36, 0x1e, 0x4e, 0xfd, 0x49, 0x3d, 0x7c, 0xb2, 0xd0, 0x75, 0xbe, 0xed, 0x88, 0x4e, 0x00, 0x48,
	0x8e, 0x9c, 0x67, 0x52, 0xe1, 0x0b, 0x70, 0x48, 0xb4, 0x0c, 0xd8, 0xc8, 0x1a, 0x7b, 0xd3, 0x03,
	0xc3, 0xa4, 0x1e, 0xae, 0xc1, 0x68, 0x05, 0xce, 0x65, 0xf6, 0x5d, 0x89, 0xb2, 0x21, 0x82, 0xb5,
	0x44, 0x3c, 0x03, 0x30, 0x23, 0xe2, 0x94, 0x3c, 0xd9, 0xbc, 0x51, 0x31, 0x66, 0xad, 0x9d, 0xd9,
	0x10, 0xdc, 0x9b, 0x44, 0x8a, 0x8b, 0x44, 0x09, 0x32, 0x64, 0xf1, 0xed, 0x3d, 0x3a, 0x84, 0x07,
	0x24, 0x20, 0x4e, 0xef, 0xa6, 0x16, 0xbd, 0x85, 0x87, 0xf3, 0x22, 0x49, 0x6b, 0x0f, 0x5c, 0xfc,
	0xf8, 0x29, 0xa4, 0xc2, 0x97, 0xe0, 0x7c, 0x23, 0x7d, 0xd4, 0xe8, 0x4d, 0x87, 0xc6, 0x46, 0xad,
	0x9a, 0x6b, 0x34, 0x7a, 0x0f, 0xd8, 0x24, 0xcb, 0x75, 0xb1, 0x92, 0x02, 0x8f, 0x1b, 0x21, 0x54,
	0x6c, 0x6c, 0x85, 0x40, 0x41, 0x6d, 0x93, 0x78, 0x07, 0x8f, 0x16, 0xaa, 0x14, 0x49, 0xfe, 0x7f,
	0x02, 0x7e, 0x33, 0x78, 0xdc, 0xe6, 0x6b, 0x0d, 0xcf, 0xc1, 0xa6, 0x09, 0x9a, 0x7f, 0x67, 0x0f,
	0x35, 0x46, 0x9b, 0x26, 0xb2, 0x8e, 0xd8, 0x6b, 0x6c, 0x5a, 0xd7, 0xf9, 0xb6, 0xa3, 0xa1, 0xc9,
	0xfa, 0xa7, 0xa6, 0x10, 0x5c, 0xc3, 0xde, 0x4b, 0xdb, 0x81, 0xde, 0x97, 0x22, 0x4b, 0x8f, 0xcf,
	0xc0, 0x35, 0xdf, 0x12, 0xf6, 0xc1, 0xbe, 0x98, 0x9d, 0x7d, 0xfe, 0xe0, 0x77, 0xd0, 0x85, 0x5e,
	0x7c, 0x75, 0x79, 0xed, 0xb3, 0xea, 0xf4, 0xf5, 0x94, 0x5f, 0xf9, 0xdd, 0x0a, 0x9e, 0x71, 0x7e,
	0xcd, 0x7d, 0x0b, 0x07, 0xe0, 0x9e, 0xf3, 0xf8, 0x53, 0x7c, 0x7e, 0x3a, 0xf7, 0x7b, 0xd3, 0x3f,
	0x0c, 0x3c, 0xb2, 0x33, 0x2f, 0x96, 0x4b, 0x51, 0xe2, 0x0c, 0x60, 0xb7, 0x0c, 0x3c, 0x34, 0xea,
	0xf6, 0xb6, 0x1b, 0x86, 0xf7, 0x41, 0x75, 0x6e, 0x51, 0x07, 0x3f, 0xc2, 0xa0, 0x99, 0x28, 0x3e,
	0x6d, 0x47, 0xd2, 0x7e, 0xea, 0xe8, 0x7e, 0xd0, 0x3c, 0xf6, 0x9a, 0xe1, 0x1b, 0x38, 0x58, 0xa8,
	0x62, 0x5d, 0xe3, 0xd9, 0x6a, 0x89, 0x7b, 0x11, 0x87, 0x03, 0x53, 0xa9, 0xa2, 0x89, 0x3a, 0x37,
	0x0e, 0xfd, 0xf1, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xa6, 0x85, 0x76, 0x00, 0x04,
	0x00, 0x00,
}

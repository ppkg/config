// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/appconfig.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ConfigRequest struct {
	AppName     string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	EnvCode     string `protobuf:"bytes,2,opt,name=env_code,json=envCode,proto3" json:"env_code,omitempty"`
	VersionName string `protobuf:"bytes,3,opt,name=version_name,json=versionName,proto3" json:"version_name,omitempty"`
	//app、环境、版本 维度下查询，如果key不为空，则只查这个key的value，否则查所有配置
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c9ad451962778ea, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ConfigRequest) GetEnvCode() string {
	if m != nil {
		return m.EnvCode
	}
	return ""
}

func (m *ConfigRequest) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *ConfigRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type BaseResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c9ad451962778ea, []int{1}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "proto.ConfigRequest")
	proto.RegisterType((*BaseResponse)(nil), "proto.BaseResponse")
}

func init() { proto.RegisterFile("proto/appconfig.proto", fileDescriptor_7c9ad451962778ea) }

var fileDescriptor_7c9ad451962778ea = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4f, 0x02, 0x41,
	0x10, 0x85, 0x83, 0x70, 0x22, 0x23, 0x1a, 0x33, 0x62, 0x72, 0x5a, 0x29, 0x95, 0x15, 0x26, 0x5a,
	0x5a, 0x29, 0x85, 0x56, 0x16, 0x5b, 0xd8, 0x92, 0x15, 0x9e, 0x84, 0xe8, 0xed, 0x8c, 0xbb, 0xc7,
	0x25, 0xfa, 0xeb, 0xcd, 0xcd, 0x5e, 0x8c, 0x54, 0x3b, 0xef, 0x7d, 0x99, 0x7d, 0xf3, 0xe8, 0x4c,
	0xa3, 0xd4, 0x72, 0xe3, 0x55, 0x97, 0x12, 0xde, 0x37, 0xeb, 0x99, 0x69, 0x2e, 0xec, 0x99, 0xfe,
	0xd0, 0xd1, 0xdc, 0x6c, 0x87, 0xaf, 0x2d, 0x52, 0xcd, 0xe7, 0x74, 0xe0, 0x55, 0x17, 0xc1, 0x57,
	0x28, 0x7b, 0x97, 0xbd, 0xeb, 0x91, 0x1b, 0x7a, 0xd5, 0x17, 0x5f, 0xa1, 0x45, 0x08, 0xcd, 0x62,
	0x29, 0x2b, 0x94, 0x7b, 0x19, 0x21, 0x34, 0x73, 0x59, 0x81, 0xaf, 0x68, 0xdc, 0x20, 0xa6, 0x8d,
	0x84, 0xbc, 0xd9, 0x37, 0x7c, 0xd8, 0x79, 0xb6, 0x7d, 0x42, 0xfd, 0x0f, 0x7c, 0x97, 0x03, 0x23,
	0xed, 0x38, 0x75, 0x34, 0x7e, 0xf4, 0x09, 0x0e, 0x49, 0x25, 0x24, 0x30, 0xd3, 0xc0, 0xfe, 0x6e,
	0x63, 0x0b, 0x67, 0x33, 0x97, 0x34, 0xac, 0x90, 0x92, 0x5f, 0xff, 0x45, 0x76, 0x92, 0x27, 0x54,
	0x20, 0x46, 0x89, 0x5d, 0x56, 0x16, 0xb7, 0xcf, 0x34, 0x7a, 0x50, 0xcd, 0x95, 0xf8, 0x9e, 0x8e,
	0x9f, 0x50, 0x67, 0xf1, 0xea, 0x3f, 0xb7, 0xe0, 0x49, 0x6e, 0x3f, 0xdb, 0xe9, 0x7c, 0x71, 0xda,
	0xb9, 0xff, 0xaf, 0x79, 0xdb, 0x37, 0xef, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x44, 0xe2, 0x00,
	0x87, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppConfigClient is the client API for AppConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppConfigClient interface {
	GetConfigValue(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type appConfigClient struct {
	cc *grpc.ClientConn
}

func NewAppConfigClient(cc *grpc.ClientConn) AppConfigClient {
	return &appConfigClient{cc}
}

func (c *appConfigClient) GetConfigValue(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/proto.AppConfig/GetConfigValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppConfigServer is the server API for AppConfig service.
type AppConfigServer interface {
	GetConfigValue(context.Context, *ConfigRequest) (*BaseResponse, error)
}

// UnimplementedAppConfigServer can be embedded to have forward compatible implementations.
type UnimplementedAppConfigServer struct {
}

func (*UnimplementedAppConfigServer) GetConfigValue(ctx context.Context, req *ConfigRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigValue not implemented")
}

func RegisterAppConfigServer(s *grpc.Server, srv AppConfigServer) {
	s.RegisterService(&_AppConfig_serviceDesc, srv)
}

func _AppConfig_GetConfigValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigServer).GetConfigValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AppConfig/GetConfigValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigServer).GetConfigValue(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AppConfig",
	HandlerType: (*AppConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigValue",
			Handler:    _AppConfig_GetConfigValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/appconfig.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	app.proto
	job.proto
	pilot.proto
	runtime_env.proto
	task.proto

It has these top-level messages:
	CreateAppRequest
	CreateAppResponse
	ModifyAppRequest
	ModifyAppResponse
	DeleteAppRequest
	DeleteAppResponse
	App
	DescribeAppsRequest
	DescribeAppsResponse
	CreateJobRequest
	CreateJobResponse
	Job
	DescribeJobsRequest
	DescribeJobsResponse
	SendTaskRequest
	SendTaskResponse
	GetTaskStatusRequest
	TaskStatus
	GetTaskStatusResponse
	RuntimeEnv
	CreateRuntimeEnvRequest
	CreateRuntimeEnvResponse
	DescribeRuntimeEnvsRequest
	DescribeRuntimeEnvsResponse
	ModifyRuntimeEnvRequest
	ModifyRuntimeEnvResponse
	DeleteRuntimeEnvRequest
	DeleteRuntimeEnvResponse
	RuntimeEnvCredential
	CreateRuntimeEnvCredentialRequset
	CreateRuntimeEnvCredentialResponse
	DescribeRuntimeEnvCredentialsRequset
	DescribeRuntimeEnvCredentialsResponse
	ModifyRuntimeEnvCredentialRequest
	ModifyRuntimeEnvCredentialResponse
	DeleteRuntimeEnvCredentialRequset
	DeleteRuntimeEnvCredentialResponse
	AttachCredentialToRuntimeEnvRequset
	AttachCredentialToRuntimeEnvResponse
	DetachCredentialFromRuntimeEnvRequset
	DetachCredentialFromRuntimeEnvResponse
	CreateTaskRequest
	CreateTaskResponse
	Task
	DescribeTasksRequest
	DescribeTasksResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

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

type CreateAppRequest struct {
	X           string `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      string `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner       string `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	ChartName   string `protobuf:"bytes,5,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Home        string `protobuf:"bytes,8,opt,name=home" json:"home,omitempty"`
	Icon        string `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	Screenshots string `protobuf:"bytes,10,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers string `protobuf:"bytes,11,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     string `protobuf:"bytes,12,opt,name=sources" json:"sources,omitempty"`
	Readme      string `protobuf:"bytes,13,opt,name=readme" json:"readme,omitempty"`
}

func (m *CreateAppRequest) Reset()                    { *m = CreateAppRequest{} }
func (m *CreateAppRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateAppRequest) ProtoMessage()               {}
func (*CreateAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateAppRequest) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *CreateAppRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAppRequest) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *CreateAppRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CreateAppRequest) GetChartName() string {
	if m != nil {
		return m.ChartName
	}
	return ""
}

func (m *CreateAppRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateAppRequest) GetHome() string {
	if m != nil {
		return m.Home
	}
	return ""
}

func (m *CreateAppRequest) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *CreateAppRequest) GetScreenshots() string {
	if m != nil {
		return m.Screenshots
	}
	return ""
}

func (m *CreateAppRequest) GetMaintainers() string {
	if m != nil {
		return m.Maintainers
	}
	return ""
}

func (m *CreateAppRequest) GetSources() string {
	if m != nil {
		return m.Sources
	}
	return ""
}

func (m *CreateAppRequest) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

type CreateAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *CreateAppResponse) Reset()                    { *m = CreateAppResponse{} }
func (m *CreateAppResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateAppResponse) ProtoMessage()               {}
func (*CreateAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type ModifyAppRequest struct {
	AppId       string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      string `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner       string `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	ChartName   string `protobuf:"bytes,5,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Home        string `protobuf:"bytes,8,opt,name=home" json:"home,omitempty"`
	Icon        string `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	Screenshots string `protobuf:"bytes,10,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers string `protobuf:"bytes,11,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     string `protobuf:"bytes,12,opt,name=sources" json:"sources,omitempty"`
	Readme      string `protobuf:"bytes,13,opt,name=readme" json:"readme,omitempty"`
}

func (m *ModifyAppRequest) Reset()                    { *m = ModifyAppRequest{} }
func (m *ModifyAppRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyAppRequest) ProtoMessage()               {}
func (*ModifyAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModifyAppRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ModifyAppRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModifyAppRequest) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *ModifyAppRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ModifyAppRequest) GetChartName() string {
	if m != nil {
		return m.ChartName
	}
	return ""
}

func (m *ModifyAppRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ModifyAppRequest) GetHome() string {
	if m != nil {
		return m.Home
	}
	return ""
}

func (m *ModifyAppRequest) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ModifyAppRequest) GetScreenshots() string {
	if m != nil {
		return m.Screenshots
	}
	return ""
}

func (m *ModifyAppRequest) GetMaintainers() string {
	if m != nil {
		return m.Maintainers
	}
	return ""
}

func (m *ModifyAppRequest) GetSources() string {
	if m != nil {
		return m.Sources
	}
	return ""
}

func (m *ModifyAppRequest) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

type ModifyAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *ModifyAppResponse) Reset()                    { *m = ModifyAppResponse{} }
func (m *ModifyAppResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyAppResponse) ProtoMessage()               {}
func (*ModifyAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ModifyAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type DeleteAppRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *DeleteAppRequest) Reset()                    { *m = DeleteAppRequest{} }
func (m *DeleteAppRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteAppRequest) ProtoMessage()               {}
func (*DeleteAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteAppRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type DeleteAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *DeleteAppResponse) Reset()                    { *m = DeleteAppResponse{} }
func (m *DeleteAppResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteAppResponse) ProtoMessage()               {}
func (*DeleteAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type App struct {
	AppId       *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Status      *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	Home        *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=home" json:"home,omitempty"`
	Icon        *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=icon" json:"icon,omitempty"`
	Screenshots *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=sources" json:"sources,omitempty"`
	Readme      *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=readme" json:"readme,omitempty"`
	ChartName   *google_protobuf.StringValue `protobuf:"bytes,12,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Owner       *google_protobuf.StringValue `protobuf:"bytes,13,opt,name=owner" json:"owner,omitempty"`
	CreateTime  *google_protobuf1.Timestamp  `protobuf:"bytes,14,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime  *google_protobuf1.Timestamp  `protobuf:"bytes,15,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *App) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *App) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *App) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *App) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *App) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *App) GetHome() *google_protobuf.StringValue {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *App) GetIcon() *google_protobuf.StringValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *App) GetScreenshots() *google_protobuf.StringValue {
	if m != nil {
		return m.Screenshots
	}
	return nil
}

func (m *App) GetMaintainers() *google_protobuf.StringValue {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func (m *App) GetSources() *google_protobuf.StringValue {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *App) GetReadme() *google_protobuf.StringValue {
	if m != nil {
		return m.Readme
	}
	return nil
}

func (m *App) GetChartName() *google_protobuf.StringValue {
	if m != nil {
		return m.ChartName
	}
	return nil
}

func (m *App) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *App) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *App) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeAppsRequest struct {
	AppId  string   `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId string   `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Status []string `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	Limit  uint32   `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	Offset uint32   `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeAppsRequest) Reset()                    { *m = DescribeAppsRequest{} }
func (m *DescribeAppsRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeAppsRequest) ProtoMessage()               {}
func (*DescribeAppsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DescribeAppsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DescribeAppsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribeAppsRequest) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *DescribeAppsRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeAppsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeAppsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeAppsResponse struct {
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	AppSet     []*App `protobuf:"bytes,2,rep,name=app_set,json=appSet" json:"app_set,omitempty"`
}

func (m *DescribeAppsResponse) Reset()                    { *m = DescribeAppsResponse{} }
func (m *DescribeAppsResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeAppsResponse) ProtoMessage()               {}
func (*DescribeAppsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DescribeAppsResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeAppsResponse) GetAppSet() []*App {
	if m != nil {
		return m.AppSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAppRequest)(nil), "openpitrix.CreateAppRequest")
	proto.RegisterType((*CreateAppResponse)(nil), "openpitrix.CreateAppResponse")
	proto.RegisterType((*ModifyAppRequest)(nil), "openpitrix.ModifyAppRequest")
	proto.RegisterType((*ModifyAppResponse)(nil), "openpitrix.ModifyAppResponse")
	proto.RegisterType((*DeleteAppRequest)(nil), "openpitrix.DeleteAppRequest")
	proto.RegisterType((*DeleteAppResponse)(nil), "openpitrix.DeleteAppResponse")
	proto.RegisterType((*App)(nil), "openpitrix.App")
	proto.RegisterType((*DescribeAppsRequest)(nil), "openpitrix.DescribeAppsRequest")
	proto.RegisterType((*DescribeAppsResponse)(nil), "openpitrix.DescribeAppsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AppManager service

type AppManagerClient interface {
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	DescribeApps(ctx context.Context, in *DescribeAppsRequest, opts ...grpc.CallOption) (*DescribeAppsResponse, error)
	ModifyApp(ctx context.Context, in *ModifyAppRequest, opts ...grpc.CallOption) (*ModifyAppResponse, error)
	DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error)
}

type appManagerClient struct {
	cc *grpc.ClientConn
}

func NewAppManagerClient(cc *grpc.ClientConn) AppManagerClient {
	return &appManagerClient{cc}
}

func (c *appManagerClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/CreateApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) DescribeApps(ctx context.Context, in *DescribeAppsRequest, opts ...grpc.CallOption) (*DescribeAppsResponse, error) {
	out := new(DescribeAppsResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/DescribeApps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) ModifyApp(ctx context.Context, in *ModifyAppRequest, opts ...grpc.CallOption) (*ModifyAppResponse, error) {
	out := new(ModifyAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/ModifyApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error) {
	out := new(DeleteAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/DeleteApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppManager service

type AppManagerServer interface {
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	DescribeApps(context.Context, *DescribeAppsRequest) (*DescribeAppsResponse, error)
	ModifyApp(context.Context, *ModifyAppRequest) (*ModifyAppResponse, error)
	DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error)
}

func RegisterAppManagerServer(s *grpc.Server, srv AppManagerServer) {
	s.RegisterService(&_AppManager_serviceDesc, srv)
}

func _AppManager_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_DescribeApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).DescribeApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/DescribeApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).DescribeApps(ctx, req.(*DescribeAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_ModifyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).ModifyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/ModifyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).ModifyApp(ctx, req.(*ModifyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).DeleteApp(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.AppManager",
	HandlerType: (*AppManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _AppManager_CreateApp_Handler,
		},
		{
			MethodName: "DescribeApps",
			Handler:    _AppManager_DescribeApps_Handler,
		},
		{
			MethodName: "ModifyApp",
			Handler:    _AppManager_ModifyApp_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _AppManager_DeleteApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4f, 0x73, 0x1b, 0x35,
	0x14, 0xc7, 0x76, 0xec, 0xc4, 0xcf, 0x31, 0x0d, 0x22, 0x2d, 0x8b, 0x49, 0xc9, 0xb2, 0xa7, 0x12,
	0xe8, 0x6e, 0xea, 0x96, 0x1e, 0xd2, 0x99, 0xcc, 0x98, 0xf6, 0x40, 0x0f, 0x85, 0x8e, 0xcb, 0x70,
	0xe0, 0xe2, 0x51, 0xd6, 0xb2, 0xad, 0xce, 0x5a, 0x7a, 0x95, 0x64, 0x12, 0x8e, 0xf0, 0x11, 0xc2,
	0x07, 0xe0, 0xb3, 0x70, 0xe6, 0xc8, 0x0c, 0x07, 0xce, 0x7c, 0x10, 0x46, 0xd2, 0xda, 0xd9, 0x8d,
	0xd3, 0xf1, 0x1e, 0x38, 0xf6, 0xe4, 0xd5, 0xef, 0xfd, 0xd7, 0xfb, 0xe9, 0x3d, 0x43, 0x9b, 0x22,
	0xc6, 0xa8, 0xa4, 0x91, 0x04, 0x24, 0x32, 0x81, 0xdc, 0x28, 0x7e, 0xd1, 0xfb, 0x74, 0x2a, 0xe5,
	0x34, 0x63, 0x89, 0x93, 0x9c, 0x2d, 0x26, 0xc9, 0xb9, 0xa2, 0x88, 0x4c, 0x69, 0xaf, 0xdb, 0x3b,
	0xbc, 0x2e, 0x37, 0x7c, 0xce, 0xb4, 0xa1, 0xf3, 0xdc, 0x59, 0xef, 0x20, 0x57, 0xa0, 0xc8, 0x13,
	0x2a, 0x84, 0x34, 0xd4, 0x70, 0x29, 0x96, 0xe6, 0x5f, 0xba, 0x9f, 0xf4, 0xfe, 0x94, 0x89, 0xfb,
	0xfa, 0x9c, 0x4e, 0xa7, 0x4c, 0x25, 0x12, 0x9d, 0xc6, 0xba, 0x76, 0xf4, 0x47, 0x1d, 0xf6, 0x9e,
	0x2a, 0x46, 0x0d, 0x1b, 0x20, 0x0e, 0xd9, 0x9b, 0x05, 0xd3, 0x86, 0x74, 0xa0, 0x36, 0x0a, 0x6a,
	0x61, 0xed, 0x5e, 0x7b, 0xf8, 0x1e, 0x21, 0xb0, 0x25, 0xe8, 0x9c, 0x05, 0x75, 0x77, 0x76, 0xdf,
	0xe4, 0x23, 0xd8, 0x56, 0x0c, 0xe5, 0x88, 0x8f, 0x83, 0x86, 0x83, 0x5b, 0xf6, 0xf8, 0x7c, 0x4c,
	0xf6, 0xa1, 0x29, 0xcf, 0x05, 0x53, 0xc1, 0x96, 0x83, 0xfd, 0x81, 0xdc, 0x05, 0x48, 0x67, 0x54,
	0x99, 0x91, 0x73, 0xd4, 0x74, 0xa2, 0xb6, 0x43, 0xbe, 0xb5, 0xde, 0x42, 0xe8, 0x8c, 0x99, 0x4e,
	0x15, 0x77, 0x59, 0x06, 0x2d, 0x27, 0x2f, 0x42, 0x36, 0x87, 0x99, 0x9c, 0xb3, 0x60, 0xc7, 0xe7,
	0x60, 0xbf, 0x2d, 0xc6, 0x53, 0x29, 0x82, 0xb6, 0xc7, 0xec, 0xb7, 0xf5, 0xa4, 0x53, 0xc5, 0x98,
	0xd0, 0x33, 0x69, 0x74, 0x00, 0xde, 0x53, 0x01, 0xb2, 0x1a, 0x73, 0xca, 0x85, 0xa1, 0x5c, 0x30,
	0xa5, 0x83, 0x8e, 0xd7, 0x28, 0x40, 0x24, 0x80, 0x6d, 0x2d, 0x17, 0x2a, 0x65, 0x3a, 0xd8, 0x75,
	0xd2, 0xe5, 0x91, 0xdc, 0x81, 0x96, 0x62, 0x74, 0x3c, 0x67, 0x41, 0x77, 0x59, 0xb4, 0x3d, 0x45,
	0x8f, 0xe1, 0x83, 0xc2, 0x15, 0x6a, 0x94, 0x42, 0x33, 0xf2, 0x19, 0x34, 0x28, 0xa2, 0xbb, 0xc5,
	0x4e, 0xff, 0x56, 0x7c, 0xd5, 0xff, 0xd8, 0x6a, 0x59, 0x59, 0xf4, 0x67, 0x1d, 0xf6, 0x5e, 0xc8,
	0x31, 0x9f, 0xfc, 0x5c, 0xb8, 0xfb, 0xdb, 0xd0, 0xa2, 0x88, 0xf6, 0x66, 0x7d, 0x03, 0x9a, 0x14,
	0xf1, 0xf9, 0xf8, 0x5d, 0x17, 0xaa, 0x74, 0xa1, 0x70, 0x99, 0xd5, 0xbb, 0xf0, 0x39, 0xec, 0x3d,
	0x63, 0x19, 0x2b, 0x3d, 0x80, 0x9b, 0x9b, 0x60, 0x43, 0x14, 0x54, 0xab, 0x87, 0xf8, 0xbb, 0x05,
	0x8d, 0x01, 0x22, 0x79, 0x58, 0x72, 0xdb, 0xe9, 0x1f, 0xc4, 0xfe, 0x25, 0xc7, 0xcb, 0xa7, 0x1e,
	0xbf, 0x32, 0x8a, 0x8b, 0xe9, 0x0f, 0x34, 0x5b, 0xb0, 0x65, 0xe7, 0x8f, 0x0b, 0x9d, 0xdf, 0x64,
	0xe2, 0x79, 0xf1, 0x55, 0x99, 0x17, 0x9b, 0x8c, 0x96, 0xac, 0x39, 0x2d, 0x13, 0x60, 0xab, 0x82,
	0x69, 0x89, 0x1e, 0x8f, 0xa0, 0xa5, 0x0d, 0x35, 0x0b, 0xed, 0xb8, 0xb5, 0x31, 0xaa, 0xd7, 0xb5,
	0xe5, 0x39, 0x52, 0xb5, 0xaa, 0x94, 0xe7, 0x28, 0x77, 0x9c, 0x53, 0x6e, 0xbb, 0x8a, 0x85, 0x23,
	0xe4, 0x69, 0x99, 0x90, 0x3b, 0x55, 0x2a, 0x2b, 0xd2, 0xf5, 0xb4, 0x4c, 0xd7, 0x76, 0x15, 0xfb,
	0x22, 0x99, 0x1f, 0x5f, 0x91, 0x19, 0x2a, 0xd8, 0xae, 0xa8, 0xfe, 0x68, 0x45, 0xf5, 0x4e, 0xb5,
	0x3e, 0x5a, 0x5d, 0xf2, 0xa4, 0xf4, 0xce, 0x77, 0x2b, 0x58, 0x16, 0xa6, 0x40, 0x7f, 0x39, 0x3a,
	0xba, 0x55, 0x18, 0xea, 0x07, 0xcb, 0x13, 0xe8, 0xa4, 0x6e, 0xfe, 0x8d, 0xec, 0xa6, 0x0a, 0xde,
	0x77, 0x96, 0xbd, 0x35, 0xcb, 0xef, 0x97, 0x6b, 0x6c, 0x08, 0x5e, 0xdd, 0x02, 0xd6, 0xd8, 0x33,
	0xc1, 0x1b, 0xdf, 0xda, 0x6c, 0xec, 0xd5, 0x2d, 0x10, 0xfd, 0x5e, 0x83, 0x0f, 0x9f, 0x39, 0x0a,
	0x9e, 0xd9, 0x37, 0xa9, 0xff, 0xcf, 0x21, 0x7a, 0x67, 0x45, 0xe7, 0xad, 0xb0, 0x61, 0xf1, 0x9c,
	0xb0, 0xfb, 0xd0, 0xcc, 0xf8, 0x9c, 0x1b, 0xc7, 0xf2, 0xee, 0xd0, 0x1f, 0xac, 0xb6, 0x9c, 0x4c,
	0x34, 0x33, 0x8e, 0xc8, 0xdd, 0x61, 0x7e, 0x8a, 0x28, 0xec, 0x97, 0x13, 0xcc, 0xa7, 0xc6, 0x21,
	0x74, 0x8c, 0x34, 0x34, 0x1b, 0xa5, 0x72, 0x21, 0x8c, 0x4b, 0xb3, 0x3b, 0x04, 0x07, 0x3d, 0xb5,
	0x08, 0xb9, 0x07, 0xdb, 0xb6, 0x04, 0xeb, 0xb1, 0x1e, 0x36, 0x6e, 0x1a, 0x2d, 0xb6, 0xc4, 0x57,
	0xcc, 0xf4, 0xff, 0x69, 0x00, 0x0c, 0x10, 0x5f, 0x50, 0x41, 0xa7, 0x4c, 0x91, 0x0c, 0xda, 0xab,
	0x6d, 0x44, 0x0e, 0x8a, 0x46, 0xd7, 0xf7, 0x7c, 0xef, 0xee, 0x5b, 0xa4, 0x3e, 0xc7, 0x28, 0xba,
	0x1c, 0xec, 0x92, 0xbc, 0x57, 0x21, 0x45, 0xfc, 0xf5, 0xaf, 0x7f, 0x7f, 0xab, 0x77, 0xa3, 0x9d,
	0xe4, 0xa7, 0x07, 0x09, 0x45, 0xd4, 0x27, 0xb5, 0x23, 0xf2, 0x4b, 0x0d, 0x76, 0x8b, 0x05, 0x92,
	0xc3, 0xa2, 0xcf, 0x1b, 0x7a, 0xd3, 0x0b, 0xdf, 0xae, 0x90, 0xc7, 0x8d, 0x2f, 0x07, 0x9f, 0x90,
	0x8f, 0xc7, 0xb9, 0xc8, 0x46, 0xd6, 0xe1, 0x39, 0x37, 0xb3, 0x70, 0xc2, 0x33, 0xc3, 0x94, 0x4b,
	0x03, 0xc8, 0x2a, 0x0d, 0x5b, 0xf1, 0x6a, 0xf2, 0x97, 0x2b, 0xbe, 0xbe, 0x5d, 0xcb, 0x15, 0xaf,
	0xad, 0x8b, 0xbc, 0xe2, 0xb9, 0xc3, 0xaf, 0x2a, 0xee, 0x97, 0x2a, 0xce, 0xa0, 0xbd, 0x5a, 0x02,
	0xe5, 0x68, 0xd7, 0xd7, 0x48, 0x39, 0xda, 0xda, 0xe6, 0xc8, 0xa3, 0x8d, 0x1d, 0x7e, 0x15, 0xed,
	0xa8, 0x18, 0xed, 0xeb, 0x8b, 0xcb, 0xc1, 0x1b, 0xf2, 0x0d, 0x90, 0xef, 0x90, 0x89, 0x97, 0x5c,
	0x31, 0x7e, 0x11, 0xbe, 0x54, 0xf2, 0x35, 0x4b, 0x4d, 0xf4, 0xc5, 0x4d, 0x28, 0xb9, 0x3d, 0x33,
	0x06, 0xf5, 0x49, 0x92, 0x14, 0x42, 0x73, 0xd9, 0x6f, 0x1e, 0xc7, 0xc7, 0xf1, 0x83, 0xa3, 0x5a,
	0xad, 0xbf, 0x47, 0x11, 0x33, 0x9e, 0xba, 0x7f, 0x81, 0xc9, 0x6b, 0x2d, 0xc5, 0xc9, 0x1a, 0xf2,
	0x63, 0x1d, 0xcf, 0xce, 0x5a, 0xee, 0xed, 0x3d, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x6c,
	0x76, 0xa2, 0xc6, 0x0a, 0x00, 0x00,
}

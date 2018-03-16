// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

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

type CreateJobRequest struct {
	X         *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	ClusterId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	AppId     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	VersionId *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId" json:"version_id,omitempty"`
	JobAction *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=job_action,json=jobAction" json:"job_action,omitempty"`
	Runtime   *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=runtime" json:"runtime,omitempty"`
	Directive *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=directive" json:"directive,omitempty"`
}

func (m *CreateJobRequest) Reset()                    { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()               {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CreateJobRequest) GetX() *google_protobuf.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateJobRequest) GetClusterId() *google_protobuf.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobRequest) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobRequest) GetVersionId() *google_protobuf.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateJobRequest) GetJobAction() *google_protobuf.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *CreateJobRequest) GetRuntime() *google_protobuf.StringValue {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *CreateJobRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type CreateJobResponse struct {
	JobId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	ClusterId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	AppId     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	VersionId *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId" json:"version_id,omitempty"`
}

func (m *CreateJobResponse) Reset()                    { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()               {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateJobResponse) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateJobResponse) GetClusterId() *google_protobuf.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *CreateJobResponse) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateJobResponse) GetVersionId() *google_protobuf.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

type Job struct {
	JobId      *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	ClusterId  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	AppId      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	VersionId  *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId" json:"version_id,omitempty"`
	JobAction  *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=job_action,json=jobAction" json:"job_action,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	ErrorCode  *google_protobuf.UInt32Value `protobuf:"bytes,7,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=directive" json:"directive,omitempty"`
	Executor   *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=executor" json:"executor,omitempty"`
	TaskCount  *google_protobuf.UInt32Value `protobuf:"bytes,10,opt,name=task_count,json=taskCount" json:"task_count,omitempty"`
	Owner      *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=owner" json:"owner,omitempty"`
	Runtime    *google_protobuf.StringValue `protobuf:"bytes,12,opt,name=runtime" json:"runtime,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,13,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,14,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Job) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Job) GetClusterId() *google_protobuf.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *Job) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *Job) GetVersionId() *google_protobuf.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *Job) GetJobAction() *google_protobuf.StringValue {
	if m != nil {
		return m.JobAction
	}
	return nil
}

func (m *Job) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Job) GetErrorCode() *google_protobuf.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Job) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Job) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Job) GetTaskCount() *google_protobuf.UInt32Value {
	if m != nil {
		return m.TaskCount
	}
	return nil
}

func (m *Job) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Job) GetRuntime() *google_protobuf.StringValue {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Job) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Job) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeJobsRequest struct {
	JobId     []string                     `protobuf:"bytes,1,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	ClusterId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	AppId     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	VersionId *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=version_id,json=versionId" json:"version_id,omitempty"`
	Executor  *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=executor" json:"executor,omitempty"`
	Runtime   *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=runtime" json:"runtime,omitempty"`
	Status    []string                     `protobuf:"bytes,7,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset uint32 `protobuf:"varint,9,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeJobsRequest) Reset()                    { *m = DescribeJobsRequest{} }
func (m *DescribeJobsRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeJobsRequest) ProtoMessage()               {}
func (*DescribeJobsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DescribeJobsRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeJobsRequest) GetClusterId() *google_protobuf.StringValue {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *DescribeJobsRequest) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *DescribeJobsRequest) GetVersionId() *google_protobuf.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *DescribeJobsRequest) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeJobsRequest) GetRuntime() *google_protobuf.StringValue {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *DescribeJobsRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeJobsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeJobsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeJobsResponse struct {
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	JobSet     []*Job `protobuf:"bytes,2,rep,name=job_set,json=jobSet" json:"job_set,omitempty"`
}

func (m *DescribeJobsResponse) Reset()                    { *m = DescribeJobsResponse{} }
func (m *DescribeJobsResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeJobsResponse) ProtoMessage()               {}
func (*DescribeJobsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DescribeJobsResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeJobsResponse) GetJobSet() []*Job {
	if m != nil {
		return m.JobSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateJobRequest)(nil), "openpitrix.CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "openpitrix.CreateJobResponse")
	proto.RegisterType((*Job)(nil), "openpitrix.Job")
	proto.RegisterType((*DescribeJobsRequest)(nil), "openpitrix.DescribeJobsRequest")
	proto.RegisterType((*DescribeJobsResponse)(nil), "openpitrix.DescribeJobsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for JobManager service

type JobManagerClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error)
}

type jobManagerClient struct {
	cc *grpc.ClientConn
}

func NewJobManagerClient(cc *grpc.ClientConn) JobManagerClient {
	return &jobManagerClient{cc}
}

func (c *jobManagerClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := grpc.Invoke(ctx, "/openpitrix.JobManager/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobManagerClient) DescribeJobs(ctx context.Context, in *DescribeJobsRequest, opts ...grpc.CallOption) (*DescribeJobsResponse, error) {
	out := new(DescribeJobsResponse)
	err := grpc.Invoke(ctx, "/openpitrix.JobManager/DescribeJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobManager service

type JobManagerServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	DescribeJobs(context.Context, *DescribeJobsRequest) (*DescribeJobsResponse, error)
}

func RegisterJobManagerServer(s *grpc.Server, srv JobManagerServer) {
	s.RegisterService(&_JobManager_serviceDesc, srv)
}

func _JobManager_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobManager_DescribeJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).DescribeJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.JobManager/DescribeJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).DescribeJobs(ctx, req.(*DescribeJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.JobManager",
	HandlerType: (*JobManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobManager_CreateJob_Handler,
		},
		{
			MethodName: "DescribeJobs",
			Handler:    _JobManager_DescribeJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x96, 0x4f, 0x73, 0xd3, 0x38,
	0x18, 0xc6, 0xd7, 0xc9, 0x26, 0x69, 0x94, 0x66, 0xb7, 0xab, 0x6d, 0x77, 0xbc, 0xd9, 0x2e, 0xcd,
	0xf8, 0x14, 0x0a, 0xb5, 0xdb, 0x94, 0x61, 0x98, 0xf4, 0x14, 0xca, 0x81, 0x76, 0x86, 0xa1, 0x93,
	0x16, 0x0e, 0x5c, 0x32, 0xb2, 0xad, 0xa4, 0x32, 0xae, 0xa4, 0x4a, 0x72, 0xd3, 0x2b, 0x7c, 0x84,
	0xf6, 0xcc, 0x97, 0xe0, 0xab, 0xf0, 0x15, 0xf8, 0x0e, 0xdc, 0x18, 0x46, 0xb2, 0x13, 0x4c, 0xff,
	0x0c, 0x2e, 0x9c, 0x3a, 0x9c, 0x32, 0x91, 0x7e, 0x8f, 0xde, 0x57, 0x7a, 0x9f, 0x57, 0x32, 0xa8,
	0x47, 0xcc, 0x77, 0xb9, 0x60, 0x8a, 0x41, 0xc0, 0x38, 0xa6, 0x9c, 0x28, 0x41, 0x4e, 0x5b, 0x77,
	0xc6, 0x8c, 0x8d, 0x63, 0xec, 0x99, 0x19, 0x3f, 0x19, 0x79, 0x13, 0x81, 0x38, 0xc7, 0x42, 0xa6,
	0x6c, 0x6b, 0xe5, 0xe2, 0xbc, 0x22, 0x47, 0x58, 0x2a, 0x74, 0xc4, 0x33, 0x60, 0x39, 0x03, 0x10,
	0x27, 0x1e, 0xa2, 0x94, 0x29, 0xa4, 0x08, 0xa3, 0x53, 0xf9, 0x7d, 0xf3, 0x13, 0xac, 0x8d, 0x31,
	0x5d, 0x93, 0x13, 0x34, 0x1e, 0x63, 0xe1, 0x31, 0x6e, 0x88, 0xcb, 0xb4, 0xf3, 0xae, 0x0c, 0x16,
	0xb6, 0x05, 0x46, 0x0a, 0xef, 0x32, 0x7f, 0x80, 0x8f, 0x13, 0x2c, 0x15, 0xbc, 0x0b, 0xac, 0xa1,
	0x6d, 0xb5, 0xad, 0x4e, 0xa3, 0xbb, 0xec, 0xa6, 0xc1, 0xdc, 0x69, 0x36, 0xee, 0xbe, 0x12, 0x84,
	0x8e, 0x5f, 0xa2, 0x38, 0xc1, 0x83, 0xdf, 0xe0, 0x16, 0x00, 0x41, 0x9c, 0x48, 0x85, 0xc5, 0x90,
	0x84, 0x76, 0xa9, 0x80, 0xa6, 0x9e, 0xf1, 0x3b, 0x21, 0xdc, 0x04, 0x55, 0xc4, 0xb9, 0x16, 0x96,
	0x0b, 0x08, 0x2b, 0x88, 0xf3, 0x9d, 0x50, 0x47, 0x3c, 0xc1, 0x42, 0x12, 0x46, 0xb5, 0xf0, 0xf7,
	0x22, 0x11, 0x33, 0x3e, 0x15, 0x47, 0xcc, 0x1f, 0xa2, 0x40, 0x9f, 0x81, 0x5d, 0x29, 0x22, 0x8e,
	0x98, 0xdf, 0x37, 0x38, 0x7c, 0x08, 0x6a, 0x22, 0xa1, 0xba, 0x1a, 0x76, 0xb5, 0x80, 0x72, 0x0a,
	0xc3, 0x1e, 0xa8, 0x87, 0x44, 0xe0, 0x40, 0x91, 0x13, 0x6c, 0xd7, 0x8a, 0xc4, 0x9c, 0xe1, 0xce,
	0x27, 0x0b, 0xfc, 0x95, 0xab, 0x8f, 0xe4, 0x8c, 0x4a, 0xac, 0x0f, 0x4e, 0x6f, 0x83, 0x84, 0x85,
	0xaa, 0x54, 0x89, 0x98, 0x9f, 0xee, 0xfd, 0x16, 0x95, 0xca, 0x79, 0x5f, 0x05, 0xe5, 0x5d, 0xe6,
	0xff, 0x0a, 0x7b, 0xfd, 0x39, 0x5b, 0x3e, 0x00, 0x55, 0xa9, 0x90, 0x4a, 0x64, 0x21, 0x57, 0x66,
	0xac, 0x0e, 0x89, 0x85, 0x60, 0x62, 0x18, 0xb0, 0xf0, 0x7a, 0x57, 0xbe, 0xd8, 0xa1, 0x6a, 0xb3,
	0x9b, 0x85, 0x34, 0xfc, 0x36, 0x0b, 0x2f, 0x38, 0x7a, 0xee, 0x46, 0x8e, 0x86, 0x8f, 0xc0, 0x1c,
	0x3e, 0xc5, 0x41, 0xa2, 0x98, 0xb0, 0xeb, 0x05, 0xa4, 0x33, 0x5a, 0xa7, 0xac, 0x90, 0x7c, 0x3d,
	0x0c, 0x58, 0x42, 0x95, 0x0d, 0x8a, 0xa4, 0xac, 0xf9, 0x6d, 0x8d, 0xc3, 0x2e, 0xa8, 0xb0, 0x09,
	0xc5, 0xc2, 0x6e, 0x14, 0xa9, 0xa9, 0x41, 0xf3, 0x0d, 0x3f, 0x7f, 0x93, 0x86, 0xdf, 0x02, 0x8d,
	0xc0, 0xf4, 0xec, 0xd0, 0x68, 0x9b, 0x46, 0xdb, 0xba, 0xa4, 0x3d, 0x98, 0xde, 0xeb, 0x03, 0x90,
	0xe2, 0x07, 0x99, 0x38, 0x2d, 0x51, 0x2a, 0xfe, 0xe3, 0xfb, 0xe2, 0x14, 0xd7, 0x03, 0xce, 0x79,
	0x19, 0xfc, 0xfd, 0x04, 0xcb, 0x40, 0x10, 0x5f, 0x5f, 0x18, 0x72, 0x7a, 0xa3, 0x2f, 0xe5, 0x9a,
	0xa8, 0xdc, 0xa9, 0xdf, 0xd2, 0x36, 0xc9, 0x5b, 0xa7, 0x72, 0x23, 0xeb, 0xfc, 0xe8, 0xd5, 0xfd,
	0xcf, 0xac, 0xb7, 0x6a, 0xe6, 0xdc, 0xa6, 0xdd, 0xb3, 0x08, 0x2a, 0x31, 0x39, 0x22, 0xca, 0x98,
	0xbf, 0x39, 0x48, 0xff, 0x68, 0x9a, 0x8d, 0x46, 0x12, 0x2b, 0x63, 0xec, 0xe6, 0x20, 0xfb, 0xe7,
	0x20, 0xb0, 0xf8, 0x6d, 0x51, 0xb2, 0x6b, 0x7c, 0x05, 0x34, 0x14, 0x53, 0x28, 0xce, 0x1c, 0x6d,
	0x19, 0x11, 0x30, 0x43, 0xa9, 0x69, 0x3b, 0xa0, 0xa6, 0xcb, 0xa6, 0x57, 0x2c, 0xb5, 0xcb, 0x9d,
	0x46, 0xf7, 0x4f, 0xf7, 0xeb, 0x87, 0x84, 0xab, 0x5f, 0x04, 0x5d, 0xd6, 0x7d, 0xac, 0xba, 0x9f,
	0x2d, 0x00, 0x76, 0x99, 0xff, 0x0c, 0x51, 0x34, 0xc6, 0x02, 0xc6, 0xa0, 0x3e, 0x7b, 0x35, 0xe0,
	0x72, 0x5e, 0x74, 0xf1, 0xb1, 0x6f, 0xfd, 0x7f, 0xcd, 0x6c, 0x9a, 0xa3, 0xe3, 0x9c, 0xf5, 0xe7,
	0x61, 0xe6, 0xcf, 0x76, 0xc4, 0xfc, 0xb7, 0x1f, 0x3e, 0x9e, 0x97, 0x9a, 0xce, 0x9c, 0x77, 0xb2,
	0xe1, 0x45, 0xcc, 0x97, 0x3d, 0x6b, 0x15, 0xbe, 0xb1, 0xc0, 0x7c, 0x7e, 0x83, 0x70, 0x25, 0xbf,
	0xe6, 0x15, 0x7e, 0x6c, 0xb5, 0xaf, 0x07, 0xb2, 0xb8, 0xee, 0x59, 0xff, 0x3f, 0xf8, 0x6f, 0x98,
	0x4d, 0xe9, 0xc8, 0xb2, 0x3d, 0x21, 0xea, 0xb0, 0x3d, 0x22, 0xb1, 0xc2, 0xc2, 0xa4, 0x01, 0xe0,
	0x2c, 0x8d, 0xc7, 0xa7, 0x67, 0xfd, 0x63, 0xf8, 0x14, 0xc0, 0xe7, 0x1c, 0xd3, 0x3d, 0xb3, 0x6e,
	0x7b, 0x4f, 0xb0, 0x08, 0x07, 0xca, 0xb9, 0x77, 0xd5, 0x28, 0x5c, 0x3a, 0x54, 0x8a, 0xcb, 0x9e,
	0xe7, 0xe5, 0x32, 0x21, 0xac, 0x5b, 0x59, 0x77, 0xd7, 0xdd, 0x8d, 0x55, 0xcb, 0xea, 0x2e, 0x20,
	0xce, 0x63, 0x12, 0x98, 0xcf, 0x25, 0x2f, 0x92, 0x8c, 0xf6, 0x2e, 0x8d, 0xbc, 0x2a, 0x71, 0xdf,
	0xaf, 0x1a, 0x0b, 0x6d, 0x7e, 0x09, 0x00, 0x00, 0xff, 0xff, 0x17, 0xa3, 0xd3, 0x64, 0xef, 0x09,
	0x00, 0x00,
}

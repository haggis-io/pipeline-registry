// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/cloudprofiler/v2/profiler.proto

/*
Package cloudprofiler is a generated protocol buffer package.

It is generated from these files:
	google/devtools/cloudprofiler/v2/profiler.proto

It has these top-level messages:
	CreateProfileRequest
	UpdateProfileRequest
	Profile
	Deployment
*/
package cloudprofiler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

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

// ProfileType is type of profiling data.
// NOTE: the enumeration member names are used (in lowercase) as unique string
// identifiers of profile types, so they must not be renamed.
type ProfileType int32

const (
	// Unspecified profile type.
	ProfileType_PROFILE_TYPE_UNSPECIFIED ProfileType = 0
	// Thread CPU time sampling.
	ProfileType_CPU ProfileType = 1
	// Wallclock time sampling. More expensive as stops all threads.
	ProfileType_WALL ProfileType = 2
	// Heap allocation sampling.
	ProfileType_HEAP ProfileType = 3
	// Single-shot collection of all thread stacks.
	ProfileType_THREADS ProfileType = 4
	// Synchronization contention profile.
	ProfileType_CONTENTION ProfileType = 5
)

var ProfileType_name = map[int32]string{
	0: "PROFILE_TYPE_UNSPECIFIED",
	1: "CPU",
	2: "WALL",
	3: "HEAP",
	4: "THREADS",
	5: "CONTENTION",
}
var ProfileType_value = map[string]int32{
	"PROFILE_TYPE_UNSPECIFIED": 0,
	"CPU":                      1,
	"WALL":                     2,
	"HEAP":                     3,
	"THREADS":                  4,
	"CONTENTION":               5,
}

func (x ProfileType) String() string {
	return proto.EnumName(ProfileType_name, int32(x))
}
func (ProfileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CreateProfileRequest describes a profile resource creation request.
// Deployment field must be populated for both online and offline modes.
// For the online mode, profile field is not set and the profile_type specifies
// the list of profile types supported by the agent. The creation call will hang
// until a profile of one of these types needs to be collected. For offline
// mode, profile field must be set, profile_type must be empty, and deployment
// field must be identical to the deployment in the profile.
type CreateProfileRequest struct {
	// Deployment details.
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment" json:"deployment,omitempty"`
	// Online mode: One or more profile types that the agent is capable of
	// providing.
	ProfileType []ProfileType `protobuf:"varint,2,rep,packed,name=profile_type,json=profileType,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
	// Offline mode: Contents of the profile to create.
	Profile *Profile `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
}

func (m *CreateProfileRequest) Reset()                    { *m = CreateProfileRequest{} }
func (m *CreateProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateProfileRequest) ProtoMessage()               {}
func (*CreateProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateProfileRequest) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *CreateProfileRequest) GetProfileType() []ProfileType {
	if m != nil {
		return m.ProfileType
	}
	return nil
}

func (m *CreateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

// UpdateProfileRequest contains the profile to update.
type UpdateProfileRequest struct {
	// Profile to update
	Profile *Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *UpdateProfileRequest) Reset()                    { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()               {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

// Profile resource.
type Profile struct {
	// Opaque, server-assigned, unique ID for this profile.
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Type of profile.
	// Input (for the offline mode) or output (for the online mode).
	ProfileType ProfileType `protobuf:"varint,2,opt,name=profile_type,json=profileType,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
	// Deployment this profile corresponds to.
	Deployment *Deployment `protobuf:"bytes,3,opt,name=deployment" json:"deployment,omitempty"`
	// Duration of the profiling session.
	// Input (for the offline mode) or output (for the online mode).
	// The field represents requested profiling duration. It may slightly differ
	// from the effective profiling duration, which is recorded in the profile
	// data, in case the profiling can't be stopped immediately (e.g. in case
	// stopping the profiling is handled asynchronously).
	Duration *google_protobuf1.Duration `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
	// Profile bytes, as a gzip compressed serialized proto, the format is
	// https://github.com/google/pprof/blob/master/proto/profile.proto.
	ProfileBytes []byte `protobuf:"bytes,5,opt,name=profile_bytes,json=profileBytes,proto3" json:"profile_bytes,omitempty"`
	// Labels associated to this specific profile. These labels will get merged
	// with the deployment labels for the final data set.
	// See documentation on deployment labels for validation rules and limits.
	// Input only, will not be populated on responses.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetProfileType() ProfileType {
	if m != nil {
		return m.ProfileType
	}
	return ProfileType_PROFILE_TYPE_UNSPECIFIED
}

func (m *Profile) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *Profile) GetDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Profile) GetProfileBytes() []byte {
	if m != nil {
		return m.ProfileBytes
	}
	return nil
}

func (m *Profile) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Deployment contains the deployment identification information.
type Deployment struct {
	// Project ID is the ID of a cloud project.
	// Validation regex: `^[a-z][-a-z0-9:.]{4,61}[a-z0-9]$`.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	// Target is the service name used to group related deployments:
	// * Service name for GAE Flex / Standard.
	// * Cluster and container name for GKE.
	// * User-specified string for direct GCE profiling (e.g. Java).
	// * Job name for Dataflow.
	// Validation regex: `^[a-z]([-a-z0-9_.]{0,253}[a-z0-9])?$`.
	Target string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	// Labels identify the deployment within the user universe and same target.
	// Validation regex for label names: `^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// Value for an individual label must be <= 512 bytes, the total
	// size of all label names and values must be <= 1024 bytes.
	//
	// Label named "language" can be used to record the programming language of
	// the profiled deployment. The standard choices for the value include "java",
	// "go", "python", "ruby", "nodejs", "php", "dotnet".
	//
	// For deployments running on Google Cloud Platform, "zone" or "region" label
	// should be present describing the deployment location. An example of a zone
	// is "us-central1-a", an example of a region is "us-central1" or
	// "us-central".
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Deployment) Reset()                    { *m = Deployment{} }
func (m *Deployment) String() string            { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()               {}
func (*Deployment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Deployment) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Deployment) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Deployment) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateProfileRequest)(nil), "google.devtools.cloudprofiler.v2.CreateProfileRequest")
	proto.RegisterType((*UpdateProfileRequest)(nil), "google.devtools.cloudprofiler.v2.UpdateProfileRequest")
	proto.RegisterType((*Profile)(nil), "google.devtools.cloudprofiler.v2.Profile")
	proto.RegisterType((*Deployment)(nil), "google.devtools.cloudprofiler.v2.Deployment")
	proto.RegisterEnum("google.devtools.cloudprofiler.v2.ProfileType", ProfileType_name, ProfileType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProfilerService service

type ProfilerServiceClient interface {
	// CreateProfile creates a new profile resource.
	//
	// In the online creation mode:
	// * The server ensures that the new profiles are created at a constant rate
	//   per deployment, so the creation request may hang for some time until the
	//   next profile session is available.
	// * The request may fail with ABORTED error if the creation is not
	//   available within ~1m, the response will indicate the duration of the
	//   backoff the client should take before attempting creating a profile
	//   again. The backoff duration is returned in google.rpc.RetryInfo extension
	//   on the response status. To a gRPC client, the extension will be return as
	//   a binary-serialized proto in the trailing metadata item named
	//   "google.rpc.retryinfo-bin".
	//
	// In the offline creation mode:
	// * The client provides the profile to create along with the profile bytes,
	//   the server records it.
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	// UpdateProfile updates the profile bytes and labels on the profile resource
	// created in the online mode.
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Profile, error)
}

type profilerServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfilerServiceClient(cc *grpc.ClientConn) ProfilerServiceClient {
	return &profilerServiceClient{cc}
}

func (c *profilerServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/google.devtools.cloudprofiler.v2.ProfilerService/UpdateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfilerService service

type ProfilerServiceServer interface {
	// CreateProfile creates a new profile resource.
	//
	// In the online creation mode:
	// * The server ensures that the new profiles are created at a constant rate
	//   per deployment, so the creation request may hang for some time until the
	//   next profile session is available.
	// * The request may fail with ABORTED error if the creation is not
	//   available within ~1m, the response will indicate the duration of the
	//   backoff the client should take before attempting creating a profile
	//   again. The backoff duration is returned in google.rpc.RetryInfo extension
	//   on the response status. To a gRPC client, the extension will be return as
	//   a binary-serialized proto in the trailing metadata item named
	//   "google.rpc.retryinfo-bin".
	//
	// In the offline creation mode:
	// * The client provides the profile to create along with the profile bytes,
	//   the server records it.
	CreateProfile(context.Context, *CreateProfileRequest) (*Profile, error)
	// UpdateProfile updates the profile bytes and labels on the profile resource
	// created in the online mode.
	UpdateProfile(context.Context, *UpdateProfileRequest) (*Profile, error)
}

func RegisterProfilerServiceServer(s *grpc.Server, srv ProfilerServiceServer) {
	s.RegisterService(&_ProfilerService_serviceDesc, srv)
}

func _ProfilerService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudprofiler.v2.ProfilerService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfilerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.cloudprofiler.v2.ProfilerService",
	HandlerType: (*ProfilerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _ProfilerService_CreateProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfilerService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/cloudprofiler/v2/profiler.proto",
}

func init() { proto.RegisterFile("google/devtools/cloudprofiler/v2/profiler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0xe3, 0x34, 0x69, 0x27, 0x6d, 0xb1, 0x56, 0x15, 0x32, 0x51, 0x81, 0x28, 0x5c, 0x42,
	0x44, 0x6d, 0xc9, 0x55, 0x51, 0x5b, 0xc4, 0xa1, 0x4d, 0x5c, 0x35, 0x52, 0x9a, 0x58, 0x6e, 0x2a,
	0x04, 0x1c, 0x22, 0xa7, 0xde, 0x5a, 0x06, 0xc7, 0x6b, 0xec, 0x4d, 0xa4, 0xa8, 0xea, 0x85, 0x57,
	0xe0, 0x11, 0xb8, 0xc2, 0xbb, 0x20, 0xf1, 0x0a, 0x3c, 0x00, 0x77, 0x2e, 0xc8, 0xf6, 0x3a, 0x3f,
	0x50, 0x94, 0x94, 0x72, 0xdb, 0x99, 0x9d, 0xef, 0xdb, 0x6f, 0x66, 0xd7, 0x9f, 0x41, 0xb1, 0x29,
	0xb5, 0x5d, 0xa2, 0x58, 0x64, 0xc8, 0x28, 0x75, 0x43, 0xe5, 0xdc, 0xa5, 0x03, 0xcb, 0x0f, 0xe8,
	0x85, 0xe3, 0x92, 0x40, 0x19, 0xaa, 0x4a, 0xba, 0x96, 0xfd, 0x80, 0x32, 0x8a, 0x4b, 0x09, 0x40,
	0x4e, 0x01, 0xf2, 0x0c, 0x40, 0x1e, 0xaa, 0xc5, 0x4d, 0x4e, 0x69, 0xfa, 0x8e, 0x62, 0x7a, 0x1e,
	0x65, 0x26, 0x73, 0xa8, 0x17, 0x26, 0xf8, 0xe2, 0x43, 0xbe, 0x1b, 0x47, 0xbd, 0xc1, 0x85, 0x62,
	0x0d, 0x82, 0xb8, 0x80, 0xef, 0x3f, 0xfa, 0x7d, 0x9f, 0x39, 0x7d, 0x12, 0x32, 0xb3, 0xef, 0x27,
	0x05, 0xe5, 0x9f, 0x08, 0x36, 0x6a, 0x01, 0x31, 0x19, 0xd1, 0x93, 0x43, 0x0d, 0xf2, 0x7e, 0x40,
	0x42, 0x86, 0x9b, 0x00, 0x16, 0xf1, 0x5d, 0x3a, 0xea, 0x13, 0x8f, 0x49, 0xa8, 0x84, 0x2a, 0x05,
	0xf5, 0xa9, 0x3c, 0x4f, 0xae, 0x5c, 0x1f, 0x63, 0x8c, 0x29, 0x3c, 0xd6, 0x61, 0x95, 0x57, 0x75,
	0xd9, 0xc8, 0x27, 0x52, 0xa6, 0x24, 0x54, 0xd6, 0xd5, 0xad, 0xf9, 0x7c, 0x5c, 0x55, 0x67, 0xe4,
	0x13, 0xa3, 0xe0, 0x4f, 0x02, 0x5c, 0x83, 0x3c, 0x0f, 0x25, 0x21, 0x16, 0xf7, 0x64, 0x61, 0x32,
	0x23, 0x45, 0x96, 0xdf, 0xc0, 0xc6, 0x99, 0x6f, 0xfd, 0xd9, 0xfc, 0x14, 0x39, 0xfa, 0x67, 0xf2,
	0x4f, 0x02, 0xe4, 0x79, 0x12, 0x63, 0xc8, 0x7a, 0x66, 0x3f, 0x61, 0x5b, 0x31, 0xe2, 0xf5, 0x35,
	0x33, 0x41, 0xb7, 0x9c, 0xc9, 0xec, 0x9d, 0x09, 0xb7, 0xbc, 0xb3, 0x1d, 0x58, 0x4e, 0x5f, 0x93,
	0x94, 0x8d, 0xb9, 0xee, 0xa7, 0x5c, 0xe9, 0x73, 0x92, 0xeb, 0xbc, 0xc0, 0x18, 0x97, 0xe2, 0xc7,
	0xb0, 0x96, 0xb6, 0xd5, 0x1b, 0x31, 0x12, 0x4a, 0x4b, 0x25, 0x54, 0x59, 0x35, 0xd2, 0x5e, 0x0f,
	0xa3, 0x1c, 0x3e, 0x81, 0x9c, 0x6b, 0xf6, 0x88, 0x1b, 0x4a, 0xb9, 0x92, 0x50, 0x29, 0xa8, 0x3b,
	0x0b, 0x77, 0x2d, 0x37, 0x63, 0x9c, 0xe6, 0xb1, 0x60, 0x64, 0x70, 0x92, 0xe2, 0x1e, 0x14, 0xa6,
	0xd2, 0x58, 0x04, 0xe1, 0x1d, 0x19, 0xf1, 0x61, 0x47, 0x4b, 0xbc, 0x01, 0x4b, 0x43, 0xd3, 0x1d,
	0x24, 0x43, 0x5e, 0x31, 0x92, 0x60, 0x3f, 0xb3, 0x8b, 0xca, 0x5f, 0x11, 0xc0, 0x64, 0x00, 0xf8,
	0x01, 0x80, 0x1f, 0xd0, 0xb7, 0xe4, 0x9c, 0x75, 0x1d, 0x8b, 0x33, 0xac, 0xf0, 0x4c, 0xc3, 0xc2,
	0xf7, 0x20, 0xc7, 0xcc, 0xc0, 0x26, 0x8c, 0x13, 0xf1, 0x08, 0xeb, 0xe3, 0x7e, 0x84, 0xb8, 0x9f,
	0xdd, 0x9b, 0x4c, 0xfd, 0x3f, 0xb7, 0x54, 0x25, 0x50, 0x98, 0x7a, 0x22, 0x78, 0x13, 0x24, 0xdd,
	0x68, 0x1f, 0x35, 0x9a, 0x5a, 0xb7, 0xf3, 0x4a, 0xd7, 0xba, 0x67, 0xad, 0x53, 0x5d, 0xab, 0x35,
	0x8e, 0x1a, 0x5a, 0x5d, 0xbc, 0x83, 0xf3, 0x20, 0xd4, 0xf4, 0x33, 0x11, 0xe1, 0x65, 0xc8, 0xbe,
	0x3c, 0x68, 0x36, 0xc5, 0x4c, 0xb4, 0x3a, 0xd6, 0x0e, 0x74, 0x51, 0xc0, 0x05, 0xc8, 0x77, 0x8e,
	0x0d, 0xed, 0xa0, 0x7e, 0x2a, 0x66, 0xf1, 0x3a, 0x40, 0xad, 0xdd, 0xea, 0x68, 0xad, 0x4e, 0xa3,
	0xdd, 0x12, 0x97, 0xd4, 0x1f, 0x19, 0xb8, 0xcb, 0xcf, 0x09, 0x4e, 0x49, 0x30, 0x74, 0xce, 0x09,
	0xfe, 0x8c, 0x60, 0x6d, 0xc6, 0x4e, 0xf0, 0xb3, 0xf9, 0x93, 0xb8, 0xce, 0x7f, 0x8a, 0x8b, 0x7f,
	0x71, 0xe5, 0xdd, 0x0f, 0xdf, 0xbe, 0x7f, 0xcc, 0xa8, 0xe5, 0x2d, 0x6e, 0xb0, 0xd1, 0x5d, 0x85,
	0xca, 0xe5, 0xe4, 0x29, 0xcb, 0x93, 0x2b, 0xbd, 0x4a, 0x1d, 0x38, 0xdc, 0x47, 0x55, 0xfc, 0x05,
	0xc1, 0xda, 0x8c, 0x01, 0x2c, 0x22, 0xf7, 0x3a, 0xc7, 0xb8, 0x89, 0xdc, 0xbd, 0x58, 0xee, 0xb6,
	0x5a, 0x89, 0xe4, 0x5e, 0xf2, 0x0a, 0x39, 0xb2, 0x84, 0x17, 0x63, 0xf1, 0xd5, 0xb1, 0x4c, 0xa5,
	0x7a, 0xb5, 0x9f, 0x5a, 0xca, 0x61, 0xfb, 0xf5, 0x09, 0x3f, 0xc6, 0xa6, 0xae, 0xe9, 0xd9, 0x32,
	0x0d, 0x6c, 0xc5, 0x26, 0x5e, 0xfc, 0x3d, 0xf2, 0x9f, 0x8f, 0xe9, 0x3b, 0xe1, 0xdf, 0x7f, 0x40,
	0xcf, 0x67, 0x12, 0xbd, 0x5c, 0x8c, 0xdc, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x83, 0x3f, 0xa6,
	0xd9, 0xb9, 0x06, 0x00, 0x00,
}

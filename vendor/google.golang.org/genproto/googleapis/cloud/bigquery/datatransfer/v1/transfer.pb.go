// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/datatransfer/v1/transfer.proto

package datatransfer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents data transfer type.
type TransferType int32

const (
	// Invalid or Unknown transfer type placeholder.
	TransferType_TRANSFER_TYPE_UNSPECIFIED TransferType = 0
	// Batch data transfer.
	TransferType_BATCH TransferType = 1
	// Streaming data transfer. Streaming data source currently doesn't
	// support multiple transfer configs per project.
	TransferType_STREAMING TransferType = 2
)

var TransferType_name = map[int32]string{
	0: "TRANSFER_TYPE_UNSPECIFIED",
	1: "BATCH",
	2: "STREAMING",
}
var TransferType_value = map[string]int32{
	"TRANSFER_TYPE_UNSPECIFIED": 0,
	"BATCH":                     1,
	"STREAMING":                 2,
}

func (x TransferType) String() string {
	return proto.EnumName(TransferType_name, int32(x))
}
func (TransferType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Represents data transfer run state.
type TransferState int32

const (
	// State placeholder.
	TransferState_TRANSFER_STATE_UNSPECIFIED TransferState = 0
	// Data transfer is scheduled and is waiting to be picked up by
	// data transfer backend.
	TransferState_PENDING TransferState = 2
	// Data transfer is in progress.
	TransferState_RUNNING TransferState = 3
	// Data transfer completed successsfully.
	TransferState_SUCCEEDED TransferState = 4
	// Data transfer failed.
	TransferState_FAILED TransferState = 5
	// Data transfer is cancelled.
	TransferState_CANCELLED TransferState = 6
)

var TransferState_name = map[int32]string{
	0: "TRANSFER_STATE_UNSPECIFIED",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
	5: "FAILED",
	6: "CANCELLED",
}
var TransferState_value = map[string]int32{
	"TRANSFER_STATE_UNSPECIFIED": 0,
	"PENDING":                    2,
	"RUNNING":                    3,
	"SUCCEEDED":                  4,
	"FAILED":                     5,
	"CANCELLED":                  6,
}

func (x TransferState) String() string {
	return proto.EnumName(TransferState_name, int32(x))
}
func (TransferState) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Represents data transfer user facing message severity.
type TransferMessage_MessageSeverity int32

const (
	// No severity specified.
	TransferMessage_MESSAGE_SEVERITY_UNSPECIFIED TransferMessage_MessageSeverity = 0
	// Informational message.
	TransferMessage_INFO TransferMessage_MessageSeverity = 1
	// Warning message.
	TransferMessage_WARNING TransferMessage_MessageSeverity = 2
	// Error message.
	TransferMessage_ERROR TransferMessage_MessageSeverity = 3
)

var TransferMessage_MessageSeverity_name = map[int32]string{
	0: "MESSAGE_SEVERITY_UNSPECIFIED",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var TransferMessage_MessageSeverity_value = map[string]int32{
	"MESSAGE_SEVERITY_UNSPECIFIED": 0,
	"INFO":                         1,
	"WARNING":                      2,
	"ERROR":                        3,
}

func (x TransferMessage_MessageSeverity) String() string {
	return proto.EnumName(TransferMessage_MessageSeverity_name, int32(x))
}
func (TransferMessage_MessageSeverity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2, 0}
}

// Represents a data transfer configuration. A transfer configuration
// contains all metadata needed to perform a data transfer. For example,
// `destination_dataset_id` specifies where data should be stored.
// When a new transfer configuration is created, the specified
// `destination_dataset_id` is created when needed and shared with the
// appropriate data source service account.
// Next id: 20
type TransferConfig struct {
	// The resource name of the transfer config.
	// Transfer config names have the form
	// `projects/{project_id}/transferConfigs/{config_id}`.
	// Where `config_id` is usually a uuid, even though it is not
	// guaranteed or required. The name is ignored when creating a transfer
	// config.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The BigQuery target dataset id.
	DestinationDatasetId string `protobuf:"bytes,2,opt,name=destination_dataset_id,json=destinationDatasetId" json:"destination_dataset_id,omitempty"`
	// User specified display name for the data transfer.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Data source id. Cannot be changed once data transfer is created.
	DataSourceId string `protobuf:"bytes,5,opt,name=data_source_id,json=dataSourceId" json:"data_source_id,omitempty"`
	// Data transfer specific parameters.
	Params *google_protobuf1.Struct `protobuf:"bytes,9,opt,name=params" json:"params,omitempty"`
	// Data transfer schedule.
	// If the data source does not support a custom schedule, this should be
	// empty. If it is empty, the default value for the data source will be
	// used.
	// The specified times are in UTC.
	// Examples of valid format:
	// `1st,3rd monday of month 15:30`,
	// `every wed,fri of jan,jun 13:15`, and
	// `first sunday of quarter 00:00`.
	// See more explanation about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	Schedule string `protobuf:"bytes,7,opt,name=schedule" json:"schedule,omitempty"`
	// The number of days to look back to automatically refresh the data.
	// For example, if `data_refresh_window_days = 10`, then every day
	// BigQuery reingests data for [today-10, today-1], rather than ingesting data
	// for just [today-1].
	// Only valid if the data source supports the feature. Set the value to  0
	// to use the default value.
	DataRefreshWindowDays int32 `protobuf:"varint,12,opt,name=data_refresh_window_days,json=dataRefreshWindowDays" json:"data_refresh_window_days,omitempty"`
	// Is this config disabled. When set to true, no runs are scheduled
	// for a given transfer.
	Disabled bool `protobuf:"varint,13,opt,name=disabled" json:"disabled,omitempty"`
	// Output only. Data transfer modification time. Ignored by server on input.
	UpdateTime *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// Output only. Next time when data transfer will run.
	NextRunTime *google_protobuf2.Timestamp `protobuf:"bytes,8,opt,name=next_run_time,json=nextRunTime" json:"next_run_time,omitempty"`
	// Output only. State of the most recently updated transfer run.
	State TransferState `protobuf:"varint,10,opt,name=state,enum=google.cloud.bigquery.datatransfer.v1.TransferState" json:"state,omitempty"`
	// Output only. Unique ID of the user on whose behalf transfer is done.
	// Applicable only to data sources that do not support service accounts.
	// When set to 0, the data source service account credentials are used.
	// May be negative. Note, that this identifier is not stable.
	// It may change over time even for the same user.
	UserId int64 `protobuf:"varint,11,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Output only. Region in which BigQuery dataset is located.
	DatasetRegion string `protobuf:"bytes,14,opt,name=dataset_region,json=datasetRegion" json:"dataset_region,omitempty"`
}

func (m *TransferConfig) Reset()                    { *m = TransferConfig{} }
func (m *TransferConfig) String() string            { return proto.CompactTextString(m) }
func (*TransferConfig) ProtoMessage()               {}
func (*TransferConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TransferConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransferConfig) GetDestinationDatasetId() string {
	if m != nil {
		return m.DestinationDatasetId
	}
	return ""
}

func (m *TransferConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TransferConfig) GetDataSourceId() string {
	if m != nil {
		return m.DataSourceId
	}
	return ""
}

func (m *TransferConfig) GetParams() *google_protobuf1.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TransferConfig) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *TransferConfig) GetDataRefreshWindowDays() int32 {
	if m != nil {
		return m.DataRefreshWindowDays
	}
	return 0
}

func (m *TransferConfig) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *TransferConfig) GetUpdateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *TransferConfig) GetNextRunTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.NextRunTime
	}
	return nil
}

func (m *TransferConfig) GetState() TransferState {
	if m != nil {
		return m.State
	}
	return TransferState_TRANSFER_STATE_UNSPECIFIED
}

func (m *TransferConfig) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TransferConfig) GetDatasetRegion() string {
	if m != nil {
		return m.DatasetRegion
	}
	return ""
}

// Represents a data transfer run.
// Next id: 27
type TransferRun struct {
	// The resource name of the transfer run.
	// Transfer run names have the form
	// `projects/{project_id}/locations/{location}/transferConfigs/{config_id}/runs/{run_id}`.
	// The name is ignored when creating a transfer run.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Minimum time after which a transfer run can be started.
	ScheduleTime *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=schedule_time,json=scheduleTime" json:"schedule_time,omitempty"`
	// For batch transfer runs, specifies the date and time that
	// data should be ingested.
	RunTime *google_protobuf2.Timestamp `protobuf:"bytes,10,opt,name=run_time,json=runTime" json:"run_time,omitempty"`
	// Status of the transfer run.
	ErrorStatus *google_rpc.Status `protobuf:"bytes,21,opt,name=error_status,json=errorStatus" json:"error_status,omitempty"`
	// Output only. Time when transfer run was started.
	// Parameter ignored by server for input requests.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Output only. Time when transfer run ended.
	// Parameter ignored by server for input requests.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Output only. Last time the data transfer run state was updated.
	UpdateTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// Output only. Data transfer specific parameters.
	Params *google_protobuf1.Struct `protobuf:"bytes,9,opt,name=params" json:"params,omitempty"`
	// Output only. The BigQuery target dataset id.
	DestinationDatasetId string `protobuf:"bytes,2,opt,name=destination_dataset_id,json=destinationDatasetId" json:"destination_dataset_id,omitempty"`
	// Output only. Data source id.
	DataSourceId string `protobuf:"bytes,7,opt,name=data_source_id,json=dataSourceId" json:"data_source_id,omitempty"`
	// Data transfer run state. Ignored for input requests.
	State TransferState `protobuf:"varint,8,opt,name=state,enum=google.cloud.bigquery.datatransfer.v1.TransferState" json:"state,omitempty"`
	// Output only. Unique ID of the user on whose behalf transfer is done.
	// Applicable only to data sources that do not support service accounts.
	// When set to 0, the data source service account credentials are used.
	// May be negative. Note, that this identifier is not stable.
	// It may change over time even for the same user.
	UserId int64 `protobuf:"varint,11,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Output only. Describes the schedule of this transfer run if it was
	// created as part of a regular schedule. For batch transfer runs that are
	// scheduled manually, this is empty.
	// NOTE: the system might choose to delay the schedule depending on the
	// current load, so `schedule_time` doesn't always matches this.
	Schedule string `protobuf:"bytes,12,opt,name=schedule" json:"schedule,omitempty"`
}

func (m *TransferRun) Reset()                    { *m = TransferRun{} }
func (m *TransferRun) String() string            { return proto.CompactTextString(m) }
func (*TransferRun) ProtoMessage()               {}
func (*TransferRun) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TransferRun) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransferRun) GetScheduleTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *TransferRun) GetRunTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.RunTime
	}
	return nil
}

func (m *TransferRun) GetErrorStatus() *google_rpc.Status {
	if m != nil {
		return m.ErrorStatus
	}
	return nil
}

func (m *TransferRun) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TransferRun) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TransferRun) GetUpdateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *TransferRun) GetParams() *google_protobuf1.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *TransferRun) GetDestinationDatasetId() string {
	if m != nil {
		return m.DestinationDatasetId
	}
	return ""
}

func (m *TransferRun) GetDataSourceId() string {
	if m != nil {
		return m.DataSourceId
	}
	return ""
}

func (m *TransferRun) GetState() TransferState {
	if m != nil {
		return m.State
	}
	return TransferState_TRANSFER_STATE_UNSPECIFIED
}

func (m *TransferRun) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TransferRun) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

// Represents a user facing message for a particular data transfer run.
type TransferMessage struct {
	// Time when message was logged.
	MessageTime *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=message_time,json=messageTime" json:"message_time,omitempty"`
	// Message severity.
	Severity TransferMessage_MessageSeverity `protobuf:"varint,2,opt,name=severity,enum=google.cloud.bigquery.datatransfer.v1.TransferMessage_MessageSeverity" json:"severity,omitempty"`
	// Message text.
	MessageText string `protobuf:"bytes,3,opt,name=message_text,json=messageText" json:"message_text,omitempty"`
}

func (m *TransferMessage) Reset()                    { *m = TransferMessage{} }
func (m *TransferMessage) String() string            { return proto.CompactTextString(m) }
func (*TransferMessage) ProtoMessage()               {}
func (*TransferMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TransferMessage) GetMessageTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.MessageTime
	}
	return nil
}

func (m *TransferMessage) GetSeverity() TransferMessage_MessageSeverity {
	if m != nil {
		return m.Severity
	}
	return TransferMessage_MESSAGE_SEVERITY_UNSPECIFIED
}

func (m *TransferMessage) GetMessageText() string {
	if m != nil {
		return m.MessageText
	}
	return ""
}

func init() {
	proto.RegisterType((*TransferConfig)(nil), "google.cloud.bigquery.datatransfer.v1.TransferConfig")
	proto.RegisterType((*TransferRun)(nil), "google.cloud.bigquery.datatransfer.v1.TransferRun")
	proto.RegisterType((*TransferMessage)(nil), "google.cloud.bigquery.datatransfer.v1.TransferMessage")
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferType", TransferType_name, TransferType_value)
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferState", TransferState_name, TransferState_value)
	proto.RegisterEnum("google.cloud.bigquery.datatransfer.v1.TransferMessage_MessageSeverity", TransferMessage_MessageSeverity_name, TransferMessage_MessageSeverity_value)
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/datatransfer/v1/transfer.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xf9, 0xcf, 0x71, 0x92, 0x8d, 0x46, 0x2c, 0x35, 0xd5, 0x02, 0xa1, 0xa2, 0x52, 0xd8,
	0x0b, 0x5b, 0x2d, 0x5d, 0x21, 0xb4, 0x02, 0x94, 0x1f, 0x27, 0x04, 0x6d, 0xb3, 0xd9, 0xb1, 0xdb,
	0xd5, 0xa2, 0x4a, 0xd6, 0x24, 0x9e, 0x7a, 0x2d, 0x25, 0xb6, 0x99, 0x19, 0x77, 0x9b, 0x87, 0xe0,
	0x25, 0xb8, 0xe4, 0x82, 0x07, 0xe1, 0x82, 0xd7, 0xe0, 0x35, 0x90, 0xc7, 0x76, 0x94, 0xcd, 0x56,
	0x4a, 0x8b, 0xc4, 0x55, 0xe7, 0xcc, 0xf9, 0xbe, 0xaf, 0xdf, 0x9c, 0x1f, 0x2b, 0x70, 0xe6, 0x85,
	0xa1, 0xb7, 0xa4, 0xc6, 0x62, 0x19, 0xc6, 0xae, 0x31, 0xf7, 0xbd, 0x5f, 0x63, 0xca, 0xd6, 0x86,
	0x4b, 0x04, 0x11, 0x8c, 0x04, 0xfc, 0x9a, 0x32, 0xe3, 0xe6, 0xc4, 0xc8, 0xcf, 0x7a, 0xc4, 0x42,
	0x11, 0xa2, 0xe3, 0x94, 0xa5, 0x4b, 0x96, 0x9e, 0xb3, 0xf4, 0x6d, 0x96, 0x7e, 0x73, 0x72, 0xf8,
	0x24, 0x13, 0x27, 0x91, 0x6f, 0x90, 0x20, 0x08, 0x05, 0x11, 0x7e, 0x18, 0xf0, 0x54, 0x64, 0x93,
	0x95, 0xd1, 0x3c, 0xbe, 0x36, 0xb8, 0x60, 0xf1, 0x42, 0x64, 0xd9, 0x2f, 0x76, 0xb3, 0xc2, 0x5f,
	0x51, 0x2e, 0xc8, 0x2a, 0xca, 0x00, 0x07, 0x19, 0x80, 0x45, 0x0b, 0x83, 0x0b, 0x22, 0xe2, 0x4c,
	0xf7, 0xe8, 0xef, 0x12, 0xb4, 0xec, 0xcc, 0xc5, 0x20, 0x0c, 0xae, 0x7d, 0x0f, 0x21, 0x28, 0x05,
	0x64, 0x45, 0x35, 0xa5, 0xa3, 0x74, 0xeb, 0x58, 0x9e, 0xd1, 0x19, 0x7c, 0xe2, 0x52, 0x2e, 0xfc,
	0x40, 0x9a, 0x72, 0x12, 0xef, 0x9c, 0x0a, 0xc7, 0x77, 0xb5, 0x82, 0x44, 0x7d, 0xbc, 0x95, 0x1d,
	0xa6, 0xc9, 0x89, 0x8b, 0xbe, 0x84, 0x86, 0xeb, 0xf3, 0x68, 0x49, 0xd6, 0x8e, 0x54, 0x2c, 0x4a,
	0xac, 0x9a, 0xdd, 0x4d, 0x13, 0xe1, 0xaf, 0xa0, 0x95, 0x88, 0x39, 0x3c, 0x8c, 0xd9, 0x82, 0x26,
	0x82, 0x65, 0x09, 0x6a, 0x24, 0xb7, 0x96, 0xbc, 0x9c, 0xb8, 0xc8, 0x80, 0x4a, 0x44, 0x18, 0x59,
	0x71, 0xad, 0xde, 0x51, 0xba, 0xea, 0xe9, 0x81, 0x9e, 0xd5, 0x34, 0x7f, 0xb0, 0x6e, 0xc9, 0x72,
	0xe0, 0x0c, 0x86, 0x0e, 0xa1, 0xc6, 0x17, 0x6f, 0xa9, 0x1b, 0x2f, 0xa9, 0x56, 0x95, 0x82, 0x9b,
	0x18, 0x7d, 0x0b, 0x9a, 0xfc, 0x97, 0x8c, 0x5e, 0x33, 0xca, 0xdf, 0x3a, 0xef, 0xfc, 0xc0, 0x0d,
	0xdf, 0x39, 0x2e, 0x59, 0x73, 0xad, 0xd1, 0x51, 0xba, 0x65, 0xfc, 0x38, 0xc9, 0xe3, 0x34, 0xfd,
	0x5a, 0x66, 0x87, 0x64, 0x2d, 0x45, 0x5d, 0x9f, 0x93, 0xf9, 0x92, 0xba, 0x5a, 0xb3, 0xa3, 0x74,
	0x6b, 0x78, 0x13, 0xa3, 0xe7, 0xa0, 0xc6, 0x91, 0x4b, 0x04, 0x75, 0x92, 0xd2, 0x6b, 0x25, 0x69,
	0xf3, 0xf0, 0x03, 0x9b, 0x76, 0xde, 0x17, 0x0c, 0x29, 0x3c, 0xb9, 0x40, 0x3f, 0x40, 0x33, 0xa0,
	0xb7, 0xc2, 0x61, 0x71, 0x90, 0xd2, 0x6b, 0x7b, 0xe9, 0x6a, 0x42, 0xc0, 0x71, 0x20, 0xf9, 0x3f,
	0x43, 0x39, 0x69, 0x2a, 0xd5, 0xa0, 0xa3, 0x74, 0x5b, 0xa7, 0x67, 0xfa, 0xbd, 0x26, 0x4e, 0xcf,
	0xfb, 0x6e, 0x25, 0x5c, 0x9c, 0x4a, 0xa0, 0x03, 0xa8, 0xc6, 0x9c, 0xb2, 0xa4, 0x13, 0x6a, 0x47,
	0xe9, 0x16, 0x71, 0x25, 0x09, 0x27, 0x2e, 0x3a, 0x4e, 0x3b, 0x95, 0xb4, 0x9d, 0x51, 0xcf, 0x0f,
	0x03, 0xad, 0x25, 0x0b, 0xdb, 0xcc, 0x6e, 0xb1, 0xbc, 0x3c, 0xfa, 0xad, 0x0c, 0x6a, 0x2e, 0x8c,
	0xe3, 0xe0, 0xce, 0x69, 0xfa, 0x11, 0x9a, 0x79, 0x37, 0xd2, 0xf7, 0x16, 0xf7, 0xbe, 0xb7, 0x91,
	0x13, 0xe4, 0x83, 0x9f, 0x41, 0x6d, 0x53, 0x2b, 0xd8, 0xcb, 0xad, 0xb2, 0xac, 0x4e, 0xcf, 0xa0,
	0x41, 0x19, 0x0b, 0x99, 0x93, 0xae, 0x80, 0xf6, 0x58, 0x52, 0x51, 0x4e, 0x65, 0xd1, 0x42, 0xb7,
	0x64, 0x06, 0xab, 0x12, 0x97, 0x06, 0xe8, 0x3b, 0x00, 0x2e, 0x08, 0x13, 0xf7, 0x6d, 0x6d, 0x5d,
	0xa2, 0x73, 0xa3, 0x34, 0x70, 0x53, 0x62, 0x79, 0xbf, 0x51, 0x1a, 0xb8, 0x92, 0xb6, 0x33, 0x4d,
	0x95, 0x07, 0x4d, 0xd3, 0x83, 0x97, 0xe5, 0xbf, 0x2d, 0xf7, 0x87, 0x9b, 0x5b, 0xbd, 0x63, 0x73,
	0x37, 0xa3, 0x59, 0xfb, 0x1f, 0x47, 0x73, 0x7b, 0xdb, 0x1b, 0xef, 0x6f, 0xfb, 0xd1, 0x9f, 0x05,
	0x78, 0x94, 0xab, 0x9d, 0x53, 0xce, 0x89, 0x47, 0xd1, 0xf7, 0xd0, 0x58, 0xa5, 0xc7, 0xb4, 0xbe,
	0xca, 0xfe, 0x75, 0xcb, 0xf0, 0xb2, 0xc0, 0x73, 0xa8, 0x71, 0x7a, 0x43, 0x99, 0x2f, 0xd6, 0xb2,
	0x42, 0xad, 0xd3, 0xd1, 0x03, 0x9f, 0x95, 0x19, 0xd1, 0xb3, 0xbf, 0x56, 0xa6, 0x86, 0x37, 0xba,
	0xc9, 0xa7, 0x73, 0x63, 0x91, 0xde, 0x8a, 0xfc, 0xd3, 0x99, 0xdb, 0xa0, 0xb7, 0xe2, 0xe8, 0x02,
	0x1e, 0xed, 0xf0, 0x51, 0x07, 0x9e, 0x9c, 0x9b, 0x96, 0xd5, 0x1b, 0x9b, 0x8e, 0x65, 0x5e, 0x9a,
	0x78, 0x62, 0xbf, 0x71, 0x2e, 0xa6, 0xd6, 0xcc, 0x1c, 0x4c, 0x46, 0x13, 0x73, 0xd8, 0xfe, 0x08,
	0xd5, 0xa0, 0x34, 0x99, 0x8e, 0x5e, 0xb6, 0x15, 0xa4, 0x42, 0xf5, 0x75, 0x0f, 0x4f, 0x27, 0xd3,
	0x71, 0xbb, 0x80, 0xea, 0x50, 0x36, 0x31, 0x7e, 0x89, 0xdb, 0xc5, 0xa7, 0x63, 0x68, 0xe4, 0x36,
	0xed, 0x75, 0x44, 0xd1, 0x67, 0xf0, 0xa9, 0x8d, 0x7b, 0x53, 0x6b, 0x64, 0x62, 0xc7, 0x7e, 0x33,
	0x33, 0x77, 0x04, 0xeb, 0x50, 0xee, 0xf7, 0xec, 0xc1, 0x4f, 0x6d, 0x05, 0x35, 0xa1, 0x6e, 0xd9,
	0xd8, 0xec, 0x9d, 0x4b, 0xcd, 0xa7, 0x1c, 0x9a, 0xef, 0xb5, 0x11, 0x7d, 0x0e, 0x87, 0x1b, 0x25,
	0xcb, 0xee, 0xd9, 0xbb, 0x52, 0x2a, 0x54, 0x67, 0xe6, 0x74, 0x98, 0x3a, 0x52, 0xa1, 0x8a, 0x2f,
	0xa6, 0xd2, 0x5e, 0x51, 0x2a, 0x5f, 0x0c, 0x06, 0xa6, 0x39, 0x34, 0x87, 0xed, 0x12, 0x02, 0xa8,
	0x8c, 0x7a, 0x93, 0x17, 0xe6, 0xb0, 0x5d, 0x4e, 0x52, 0x83, 0xde, 0x74, 0x60, 0xbe, 0x48, 0xc2,
	0x4a, 0xff, 0x1f, 0x05, 0xbe, 0x5e, 0x84, 0xab, 0xfb, 0xf5, 0xa3, 0xbf, 0x31, 0x38, 0x4b, 0x5a,
	0x3e, 0x53, 0x7e, 0x79, 0x95, 0xf1, 0xbc, 0x70, 0x49, 0x02, 0x4f, 0x0f, 0x99, 0x67, 0x78, 0x34,
	0x90, 0x03, 0x61, 0xa4, 0x29, 0x12, 0xf9, 0x7c, 0xcf, 0x0f, 0x80, 0xe7, 0xdb, 0xf1, 0xef, 0x85,
	0xf2, 0x78, 0xd0, 0x1f, 0xda, 0x7f, 0x14, 0x8e, 0xc7, 0xa9, 0xf6, 0x40, 0x7a, 0xea, 0xfb, 0xde,
	0x2b, 0xe9, 0x29, 0xd9, 0xa8, 0xdc, 0x86, 0x7e, 0x79, 0xf2, 0x57, 0x8e, 0xbb, 0x92, 0xb8, 0xab,
	0x1c, 0x77, 0xb5, 0x8d, 0xbb, 0xba, 0x3c, 0x99, 0x57, 0xa4, 0xab, 0x6f, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x74, 0xf0, 0x31, 0x95, 0x08, 0x00, 0x00,
}

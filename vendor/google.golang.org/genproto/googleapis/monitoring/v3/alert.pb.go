// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/alert.proto

/*
Package monitoring is a generated protocol buffer package.

It is generated from these files:
	google/monitoring/v3/alert.proto
	google/monitoring/v3/alert_service.proto
	google/monitoring/v3/common.proto
	google/monitoring/v3/group.proto
	google/monitoring/v3/group_service.proto
	google/monitoring/v3/metric.proto
	google/monitoring/v3/metric_service.proto
	google/monitoring/v3/mutation_record.proto
	google/monitoring/v3/notification.proto
	google/monitoring/v3/notification_service.proto
	google/monitoring/v3/uptime.proto
	google/monitoring/v3/uptime_service.proto

It has these top-level messages:
	AlertPolicy
	CreateAlertPolicyRequest
	GetAlertPolicyRequest
	ListAlertPoliciesRequest
	ListAlertPoliciesResponse
	UpdateAlertPolicyRequest
	DeleteAlertPolicyRequest
	TypedValue
	TimeInterval
	Aggregation
	Group
	ListGroupsRequest
	ListGroupsResponse
	GetGroupRequest
	CreateGroupRequest
	UpdateGroupRequest
	DeleteGroupRequest
	ListGroupMembersRequest
	ListGroupMembersResponse
	Point
	TimeSeries
	ListMonitoredResourceDescriptorsRequest
	ListMonitoredResourceDescriptorsResponse
	GetMonitoredResourceDescriptorRequest
	ListMetricDescriptorsRequest
	ListMetricDescriptorsResponse
	GetMetricDescriptorRequest
	CreateMetricDescriptorRequest
	DeleteMetricDescriptorRequest
	ListTimeSeriesRequest
	ListTimeSeriesResponse
	CreateTimeSeriesRequest
	CreateTimeSeriesError
	MutationRecord
	NotificationChannelDescriptor
	NotificationChannel
	ListNotificationChannelDescriptorsRequest
	ListNotificationChannelDescriptorsResponse
	GetNotificationChannelDescriptorRequest
	CreateNotificationChannelRequest
	ListNotificationChannelsRequest
	ListNotificationChannelsResponse
	GetNotificationChannelRequest
	UpdateNotificationChannelRequest
	DeleteNotificationChannelRequest
	SendNotificationChannelVerificationCodeRequest
	GetNotificationChannelVerificationCodeRequest
	GetNotificationChannelVerificationCodeResponse
	VerifyNotificationChannelRequest
	UptimeCheckConfig
	UptimeCheckIp
	ListUptimeCheckConfigsRequest
	ListUptimeCheckConfigsResponse
	GetUptimeCheckConfigRequest
	CreateUptimeCheckConfigRequest
	UpdateUptimeCheckConfigRequest
	DeleteUptimeCheckConfigRequest
	ListUptimeCheckIpsRequest
	ListUptimeCheckIpsResponse
*/
package monitoring

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf4 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Operators for combining conditions.
type AlertPolicy_ConditionCombinerType int32

const (
	// An unspecified combiner.
	AlertPolicy_COMBINE_UNSPECIFIED AlertPolicy_ConditionCombinerType = 0
	// Combine conditions using the logical `AND` operator. An
	// incident is created only if all conditions are met
	// simultaneously. This combiner is satisfied if all conditions are
	// met, even if they are met on completely different resources.
	AlertPolicy_AND AlertPolicy_ConditionCombinerType = 1
	// Combine conditions using the logical `OR` operator. An incident
	// is created if any of the listed conditions is met.
	AlertPolicy_OR AlertPolicy_ConditionCombinerType = 2
	// Combine conditions using logical `AND` operator, but unlike the regular
	// `AND` option, an incident is created only if all conditions are met
	// simultaneously on at least one resource.
	AlertPolicy_AND_WITH_MATCHING_RESOURCE AlertPolicy_ConditionCombinerType = 3
)

var AlertPolicy_ConditionCombinerType_name = map[int32]string{
	0: "COMBINE_UNSPECIFIED",
	1: "AND",
	2: "OR",
	3: "AND_WITH_MATCHING_RESOURCE",
}
var AlertPolicy_ConditionCombinerType_value = map[string]int32{
	"COMBINE_UNSPECIFIED":        0,
	"AND":                        1,
	"OR":                         2,
	"AND_WITH_MATCHING_RESOURCE": 3,
}

func (x AlertPolicy_ConditionCombinerType) String() string {
	return proto.EnumName(AlertPolicy_ConditionCombinerType_name, int32(x))
}
func (AlertPolicy_ConditionCombinerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// A description of the conditions under which some aspect of your system is
// considered to be "unhealthy" and the ways to notify people or services about
// this state. For an overview of alert policies, see
// [Introduction to Alerting](/monitoring/alerts/).
type AlertPolicy struct {
	// Required if the policy exists. The resource name for this policy. The
	// syntax is:
	//
	//     projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	//
	// `[ALERT_POLICY_ID]` is assigned by Stackdriver Monitoring when the policy
	// is created.  When calling the
	// [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	// method, do not include the `name` field in the alerting policy passed as
	// part of the request.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Documentation that is included with notifications and incidents related to
	// this policy. Best practice is for the documentation to include information
	// to help responders understand, mitigate, escalate, and correct the
	// underlying problems detected by the alerting policy. Notification channels
	// that have limited capacity might not show this documentation.
	Documentation *AlertPolicy_Documentation `protobuf:"bytes,13,opt,name=documentation" json:"documentation,omitempty"`
	// User-supplied key/value data to be used for organizing and
	// identifying the `AlertPolicy` objects.
	//
	// The field can contain up to 64 entries. Each key and value is limited to
	// 63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	// values can contain only lowercase letters, numerals, underscores, and
	// dashes. Keys must begin with a letter.
	UserLabels map[string]string `protobuf:"bytes,16,rep,name=user_labels,json=userLabels" json:"user_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// A list of conditions for the policy. The conditions are combined by AND or
	// OR according to the `combiner` field. If the combined conditions evaluate
	// to true, then an incident is created. A policy can have from one to six
	// conditions.
	Conditions []*AlertPolicy_Condition `protobuf:"bytes,12,rep,name=conditions" json:"conditions,omitempty"`
	// How to combine the results of multiple conditions
	// to determine if an incident should be opened.
	Combiner AlertPolicy_ConditionCombinerType `protobuf:"varint,6,opt,name=combiner,enum=google.monitoring.v3.AlertPolicy_ConditionCombinerType" json:"combiner,omitempty"`
	// Whether or not the policy is enabled. On write, the default interpretation
	// if unset is that the policy is enabled. On read, clients should not make
	// any assumption about the state if it has not been populated. The
	// field should always be populated on List and Get operations, unless
	// a field projection has been specified that strips it out.
	Enabled *google_protobuf4.BoolValue `protobuf:"bytes,17,opt,name=enabled" json:"enabled,omitempty"`
	// Identifies the notification channels to which notifications should be sent
	// when incidents are opened or closed or when new violations occur on
	// an already opened incident. Each element of this array corresponds to
	// the `name` field in each of the
	// [`NotificationChannel`][google.monitoring.v3.NotificationChannel]
	// objects that are returned from the [`ListNotificationChannels`]
	// [google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	// method. The syntax of the entries in this field is:
	//
	//     projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	NotificationChannels []string `protobuf:"bytes,14,rep,name=notification_channels,json=notificationChannels" json:"notification_channels,omitempty"`
	// A read-only record of the creation of the alerting policy. If provided
	// in a call to create or update, this field will be ignored.
	CreationRecord *MutationRecord `protobuf:"bytes,10,opt,name=creation_record,json=creationRecord" json:"creation_record,omitempty"`
	// A read-only record of the most recent change to the alerting policy. If
	// provided in a call to create or update, this field will be ignored.
	MutationRecord *MutationRecord `protobuf:"bytes,11,opt,name=mutation_record,json=mutationRecord" json:"mutation_record,omitempty"`
}

func (m *AlertPolicy) Reset()                    { *m = AlertPolicy{} }
func (m *AlertPolicy) String() string            { return proto.CompactTextString(m) }
func (*AlertPolicy) ProtoMessage()               {}
func (*AlertPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlertPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertPolicy) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AlertPolicy) GetDocumentation() *AlertPolicy_Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *AlertPolicy) GetUserLabels() map[string]string {
	if m != nil {
		return m.UserLabels
	}
	return nil
}

func (m *AlertPolicy) GetConditions() []*AlertPolicy_Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *AlertPolicy) GetCombiner() AlertPolicy_ConditionCombinerType {
	if m != nil {
		return m.Combiner
	}
	return AlertPolicy_COMBINE_UNSPECIFIED
}

func (m *AlertPolicy) GetEnabled() *google_protobuf4.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *AlertPolicy) GetNotificationChannels() []string {
	if m != nil {
		return m.NotificationChannels
	}
	return nil
}

func (m *AlertPolicy) GetCreationRecord() *MutationRecord {
	if m != nil {
		return m.CreationRecord
	}
	return nil
}

func (m *AlertPolicy) GetMutationRecord() *MutationRecord {
	if m != nil {
		return m.MutationRecord
	}
	return nil
}

// A content string and a MIME type that describes the content string's
// format.
type AlertPolicy_Documentation struct {
	// The text of the documentation, interpreted according to `mime_type`.
	// The content may not exceed 8,192 Unicode characters and may not exceed
	// more than 10,240 bytes when encoded in UTF-8 format, whichever is
	// smaller.
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	// The format of the `content` field. Presently, only the value
	// `"text/markdown"` is supported. See
	// [Markdown](https://en.wikipedia.org/wiki/Markdown) for more information.
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType" json:"mime_type,omitempty"`
}

func (m *AlertPolicy_Documentation) Reset()                    { *m = AlertPolicy_Documentation{} }
func (m *AlertPolicy_Documentation) String() string            { return proto.CompactTextString(m) }
func (*AlertPolicy_Documentation) ProtoMessage()               {}
func (*AlertPolicy_Documentation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AlertPolicy_Documentation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *AlertPolicy_Documentation) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// A condition is a true/false test that determines when an alerting policy
// should open an incident. If a condition evaluates to true, it signifies
// that something is wrong.
type AlertPolicy_Condition struct {
	// Required if the condition exists. The unique resource name for this
	// condition. Its syntax is:
	//
	//     projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	//
	// `[CONDITION_ID]` is assigned by Stackdriver Monitoring when the
	// condition is created as part of a new or updated alerting policy.
	//
	// When calling the
	// [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	// method, do not include the `name` field in the conditions of the
	// requested alerting policy. Stackdriver Monitoring creates the
	// condition identifiers and includes them in the new policy.
	//
	// When calling the
	// [alertPolicies.update][google.monitoring.v3.AlertPolicyService.UpdateAlertPolicy]
	// method to update a policy, including a condition `name` causes the
	// existing condition to be updated. Conditions without names are added to
	// the updated policy. Existing conditions are deleted if they are not
	// updated.
	//
	// Best practice is to preserve `[CONDITION_ID]` if you make only small
	// changes, such as those to condition thresholds, durations, or trigger
	// values.  Otherwise, treat the change as a new condition and let the
	// existing condition be deleted.
	Name string `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	// A short name or phrase used to identify the condition in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same policy.
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Only one of the following condition types will be specified.
	//
	// Types that are valid to be assigned to Condition:
	//	*AlertPolicy_Condition_ConditionThreshold
	//	*AlertPolicy_Condition_ConditionAbsent
	Condition isAlertPolicy_Condition_Condition `protobuf_oneof:"condition"`
}

func (m *AlertPolicy_Condition) Reset()                    { *m = AlertPolicy_Condition{} }
func (m *AlertPolicy_Condition) String() string            { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition) ProtoMessage()               {}
func (*AlertPolicy_Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type isAlertPolicy_Condition_Condition interface {
	isAlertPolicy_Condition_Condition()
}

type AlertPolicy_Condition_ConditionThreshold struct {
	ConditionThreshold *AlertPolicy_Condition_MetricThreshold `protobuf:"bytes,1,opt,name=condition_threshold,json=conditionThreshold,oneof"`
}
type AlertPolicy_Condition_ConditionAbsent struct {
	ConditionAbsent *AlertPolicy_Condition_MetricAbsence `protobuf:"bytes,2,opt,name=condition_absent,json=conditionAbsent,oneof"`
}

func (*AlertPolicy_Condition_ConditionThreshold) isAlertPolicy_Condition_Condition() {}
func (*AlertPolicy_Condition_ConditionAbsent) isAlertPolicy_Condition_Condition()    {}

func (m *AlertPolicy_Condition) GetCondition() isAlertPolicy_Condition_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *AlertPolicy_Condition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertPolicy_Condition) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AlertPolicy_Condition) GetConditionThreshold() *AlertPolicy_Condition_MetricThreshold {
	if x, ok := m.GetCondition().(*AlertPolicy_Condition_ConditionThreshold); ok {
		return x.ConditionThreshold
	}
	return nil
}

func (m *AlertPolicy_Condition) GetConditionAbsent() *AlertPolicy_Condition_MetricAbsence {
	if x, ok := m.GetCondition().(*AlertPolicy_Condition_ConditionAbsent); ok {
		return x.ConditionAbsent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AlertPolicy_Condition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AlertPolicy_Condition_OneofMarshaler, _AlertPolicy_Condition_OneofUnmarshaler, _AlertPolicy_Condition_OneofSizer, []interface{}{
		(*AlertPolicy_Condition_ConditionThreshold)(nil),
		(*AlertPolicy_Condition_ConditionAbsent)(nil),
	}
}

func _AlertPolicy_Condition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AlertPolicy_Condition)
	// condition
	switch x := m.Condition.(type) {
	case *AlertPolicy_Condition_ConditionThreshold:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConditionThreshold); err != nil {
			return err
		}
	case *AlertPolicy_Condition_ConditionAbsent:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConditionAbsent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AlertPolicy_Condition.Condition has unexpected type %T", x)
	}
	return nil
}

func _AlertPolicy_Condition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AlertPolicy_Condition)
	switch tag {
	case 1: // condition.condition_threshold
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AlertPolicy_Condition_MetricThreshold)
		err := b.DecodeMessage(msg)
		m.Condition = &AlertPolicy_Condition_ConditionThreshold{msg}
		return true, err
	case 2: // condition.condition_absent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AlertPolicy_Condition_MetricAbsence)
		err := b.DecodeMessage(msg)
		m.Condition = &AlertPolicy_Condition_ConditionAbsent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AlertPolicy_Condition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AlertPolicy_Condition)
	// condition
	switch x := m.Condition.(type) {
	case *AlertPolicy_Condition_ConditionThreshold:
		s := proto.Size(x.ConditionThreshold)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AlertPolicy_Condition_ConditionAbsent:
		s := proto.Size(x.ConditionAbsent)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specifies how many time series must fail a predicate to trigger a
// condition. If not specified, then a `{count: 1}` trigger is used.
type AlertPolicy_Condition_Trigger struct {
	// A type of trigger.
	//
	// Types that are valid to be assigned to Type:
	//	*AlertPolicy_Condition_Trigger_Count
	//	*AlertPolicy_Condition_Trigger_Percent
	Type isAlertPolicy_Condition_Trigger_Type `protobuf_oneof:"type"`
}

func (m *AlertPolicy_Condition_Trigger) Reset()         { *m = AlertPolicy_Condition_Trigger{} }
func (m *AlertPolicy_Condition_Trigger) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_Trigger) ProtoMessage()    {}
func (*AlertPolicy_Condition_Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0}
}

type isAlertPolicy_Condition_Trigger_Type interface {
	isAlertPolicy_Condition_Trigger_Type()
}

type AlertPolicy_Condition_Trigger_Count struct {
	Count int32 `protobuf:"varint,1,opt,name=count,oneof"`
}
type AlertPolicy_Condition_Trigger_Percent struct {
	Percent float64 `protobuf:"fixed64,2,opt,name=percent,oneof"`
}

func (*AlertPolicy_Condition_Trigger_Count) isAlertPolicy_Condition_Trigger_Type()   {}
func (*AlertPolicy_Condition_Trigger_Percent) isAlertPolicy_Condition_Trigger_Type() {}

func (m *AlertPolicy_Condition_Trigger) GetType() isAlertPolicy_Condition_Trigger_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AlertPolicy_Condition_Trigger) GetCount() int32 {
	if x, ok := m.GetType().(*AlertPolicy_Condition_Trigger_Count); ok {
		return x.Count
	}
	return 0
}

func (m *AlertPolicy_Condition_Trigger) GetPercent() float64 {
	if x, ok := m.GetType().(*AlertPolicy_Condition_Trigger_Percent); ok {
		return x.Percent
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AlertPolicy_Condition_Trigger) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AlertPolicy_Condition_Trigger_OneofMarshaler, _AlertPolicy_Condition_Trigger_OneofUnmarshaler, _AlertPolicy_Condition_Trigger_OneofSizer, []interface{}{
		(*AlertPolicy_Condition_Trigger_Count)(nil),
		(*AlertPolicy_Condition_Trigger_Percent)(nil),
	}
}

func _AlertPolicy_Condition_Trigger_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AlertPolicy_Condition_Trigger)
	// type
	switch x := m.Type.(type) {
	case *AlertPolicy_Condition_Trigger_Count:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Count))
	case *AlertPolicy_Condition_Trigger_Percent:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Percent))
	case nil:
	default:
		return fmt.Errorf("AlertPolicy_Condition_Trigger.Type has unexpected type %T", x)
	}
	return nil
}

func _AlertPolicy_Condition_Trigger_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AlertPolicy_Condition_Trigger)
	switch tag {
	case 1: // type.count
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &AlertPolicy_Condition_Trigger_Count{int32(x)}
		return true, err
	case 2: // type.percent
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Type = &AlertPolicy_Condition_Trigger_Percent{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _AlertPolicy_Condition_Trigger_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AlertPolicy_Condition_Trigger)
	// type
	switch x := m.Type.(type) {
	case *AlertPolicy_Condition_Trigger_Count:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Count))
	case *AlertPolicy_Condition_Trigger_Percent:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A condition type that compares a collection of time series
// against a threshold.
type AlertPolicy_Condition_MetricThreshold struct {
	// A [filter](/monitoring/api/v3/filters) that
	// identifies which time series should be compared with the threshold.
	//
	// The filter is similar to the one that is specified in the
	// [`MetricService.ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that
	// call is useful to verify the time series that will be retrieved /
	// processed) and must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Specifies the alignment of data points in individual time series as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.
	//
	// This field is similar to the one in the
	// [`MetricService.ListTimeSeries` request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
	// It is advisable to use the `ListTimeSeries` method when debugging this field.
	Aggregations []*Aggregation `protobuf:"bytes,8,rep,name=aggregations" json:"aggregations,omitempty"`
	// A [filter](/monitoring/api/v3/filters) that identifies a time
	// series that should be used as the denominator of a ratio that will be
	// compared with the threshold. If a `denominator_filter` is specified,
	// the time series specified by the `filter` field will be used as the
	// numerator.
	//
	// The filter is similar to the one that is specified in the
	// [`MetricService.ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that
	// call is useful to verify the time series that will be retrieved /
	// processed) and must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	DenominatorFilter string `protobuf:"bytes,9,opt,name=denominator_filter,json=denominatorFilter" json:"denominator_filter,omitempty"`
	// Specifies the alignment of data points in individual time series
	// selected by `denominatorFilter` as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resources).
	//
	// When computing ratios, the `aggregations` and
	// `denominator_aggregations` fields must use the same alignment period
	// and produce time series that have the same periodicity and labels.
	//
	// This field is similar to the one in the
	// [`MetricService.ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list). It
	// is advisable to use the `ListTimeSeries` method when debugging this
	// field.
	DenominatorAggregations []*Aggregation `protobuf:"bytes,10,rep,name=denominator_aggregations,json=denominatorAggregations" json:"denominator_aggregations,omitempty"`
	// The comparison to apply between the time series (indicated by `filter`
	// and `aggregation`) and the threshold (indicated by `threshold_value`).
	// The comparison is applied on each time series, with the time series
	// on the left-hand side and the threshold on the right-hand side.
	//
	// Only `COMPARISON_LT` and `COMPARISON_GT` are supported currently.
	Comparison ComparisonType `protobuf:"varint,4,opt,name=comparison,enum=google.monitoring.v3.ComparisonType" json:"comparison,omitempty"`
	// A value against which to compare the time series.
	ThresholdValue float64 `protobuf:"fixed64,5,opt,name=threshold_value,json=thresholdValue" json:"threshold_value,omitempty"`
	// The amount of time that a time series must violate the
	// threshold to be considered failing. Currently, only values
	// that are a multiple of a minute--e.g.  60, 120, or 300
	// seconds--are supported. If an invalid value is given, an
	// error will be returned. The `Duration.nanos` field is
	// ignored. When choosing a duration, it is useful to keep in mind the
	// frequency of the underlying time series data (which may also be
	// affected by any alignments specified in the `aggregation` field);
	// a good duration is long enough so that a single outlier does not
	// generate spurious alerts, but short enough that unhealthy states
	// are detected and alerted on quickly.
	Duration *google_protobuf3.Duration `protobuf:"bytes,6,opt,name=duration" json:"duration,omitempty"`
	// The number/percent of time series for which the comparison must hold
	// in order for the condition to trigger. If unspecified, then the
	// condition will trigger if the comparison is true for any of the
	// time series that have been identified by `filter` and `aggregations`,
	// or by the ratio, if `denominator_filter` and `denominator_aggregations`
	// are specified.
	Trigger *AlertPolicy_Condition_Trigger `protobuf:"bytes,7,opt,name=trigger" json:"trigger,omitempty"`
}

func (m *AlertPolicy_Condition_MetricThreshold) Reset()         { *m = AlertPolicy_Condition_MetricThreshold{} }
func (m *AlertPolicy_Condition_MetricThreshold) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_MetricThreshold) ProtoMessage()    {}
func (*AlertPolicy_Condition_MetricThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 1}
}

func (m *AlertPolicy_Condition_MetricThreshold) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricThreshold) GetAggregations() []*Aggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDenominatorFilter() string {
	if m != nil {
		return m.DenominatorFilter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDenominatorAggregations() []*Aggregation {
	if m != nil {
		return m.DenominatorAggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetComparison() ComparisonType {
	if m != nil {
		return m.Comparison
	}
	return ComparisonType_COMPARISON_UNSPECIFIED
}

func (m *AlertPolicy_Condition_MetricThreshold) GetThresholdValue() float64 {
	if m != nil {
		return m.ThresholdValue
	}
	return 0
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDuration() *google_protobuf3.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetTrigger() *AlertPolicy_Condition_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

// A condition type that checks that monitored resources
// are reporting data. The configuration defines a metric and
// a set of monitored resources. The predicate is considered in violation
// when a time series for the specified metric of a monitored
// resource does not include any data in the specified `duration`.
type AlertPolicy_Condition_MetricAbsence struct {
	// A [filter](/monitoring/api/v3/filters) that
	// identifies which time series should be compared with the threshold.
	//
	// The filter is similar to the one that is specified in the
	// [`MetricService.ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that
	// call is useful to verify the time series that will be retrieved /
	// processed) and must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	Filter string `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	// Specifies the alignment of data points in individual time series as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.
	//
	// This field is similar to the
	// one in the [`MetricService.ListTimeSeries` request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
	// It is advisable to use the `ListTimeSeries` method when debugging this field.
	Aggregations []*Aggregation `protobuf:"bytes,5,rep,name=aggregations" json:"aggregations,omitempty"`
	// The amount of time that a time series must fail to report new
	// data to be considered failing. Currently, only values that
	// are a multiple of a minute--e.g.  60, 120, or 300
	// seconds--are supported. If an invalid value is given, an
	// error will be returned. The `Duration.nanos` field is
	// ignored.
	Duration *google_protobuf3.Duration `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	// The number/percent of time series for which the comparison must hold
	// in order for the condition to trigger. If unspecified, then the
	// condition will trigger if the comparison is true for any of the
	// time series that have been identified by `filter` and `aggregations`.
	Trigger *AlertPolicy_Condition_Trigger `protobuf:"bytes,3,opt,name=trigger" json:"trigger,omitempty"`
}

func (m *AlertPolicy_Condition_MetricAbsence) Reset()         { *m = AlertPolicy_Condition_MetricAbsence{} }
func (m *AlertPolicy_Condition_MetricAbsence) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_MetricAbsence) ProtoMessage()    {}
func (*AlertPolicy_Condition_MetricAbsence) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 2}
}

func (m *AlertPolicy_Condition_MetricAbsence) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricAbsence) GetAggregations() []*Aggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricAbsence) GetDuration() *google_protobuf3.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricAbsence) GetTrigger() *AlertPolicy_Condition_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertPolicy)(nil), "google.monitoring.v3.AlertPolicy")
	proto.RegisterType((*AlertPolicy_Documentation)(nil), "google.monitoring.v3.AlertPolicy.Documentation")
	proto.RegisterType((*AlertPolicy_Condition)(nil), "google.monitoring.v3.AlertPolicy.Condition")
	proto.RegisterType((*AlertPolicy_Condition_Trigger)(nil), "google.monitoring.v3.AlertPolicy.Condition.Trigger")
	proto.RegisterType((*AlertPolicy_Condition_MetricThreshold)(nil), "google.monitoring.v3.AlertPolicy.Condition.MetricThreshold")
	proto.RegisterType((*AlertPolicy_Condition_MetricAbsence)(nil), "google.monitoring.v3.AlertPolicy.Condition.MetricAbsence")
	proto.RegisterEnum("google.monitoring.v3.AlertPolicy_ConditionCombinerType", AlertPolicy_ConditionCombinerType_name, AlertPolicy_ConditionCombinerType_value)
}

func init() { proto.RegisterFile("google/monitoring/v3/alert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 941 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xeb, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0xe6, 0x76, 0xd2, 0x36, 0xd9, 0xd9, 0xee, 0xae, 0x31, 0x68, 0x95, 0xae, 0x90,
	0x88, 0x40, 0x38, 0x22, 0x01, 0x71, 0x59, 0x81, 0x94, 0x5b, 0x37, 0x11, 0x24, 0xad, 0xa6, 0x69,
	0x91, 0x50, 0x25, 0xcb, 0x71, 0xa6, 0xae, 0x85, 0x3d, 0x63, 0x4d, 0xec, 0xa2, 0xbc, 0x0e, 0x3f,
	0x79, 0x14, 0x1e, 0x81, 0x7f, 0xbc, 0x02, 0xe2, 0x01, 0x90, 0xc7, 0x63, 0xc7, 0xe9, 0xa6, 0xbb,
	0x64, 0xf7, 0x5f, 0xce, 0x9c, 0xef, 0x7c, 0xe7, 0xf6, 0xcd, 0x38, 0xd0, 0xb0, 0x19, 0xb3, 0x5d,
	0xd2, 0xf2, 0x18, 0x75, 0x02, 0xc6, 0x1d, 0x6a, 0xb7, 0xee, 0x3a, 0x2d, 0xd3, 0x25, 0x3c, 0xd0,
	0x7d, 0xce, 0x02, 0x86, 0x8e, 0x63, 0x84, 0xbe, 0x46, 0xe8, 0x77, 0x1d, 0xed, 0x23, 0x19, 0x67,
	0xfa, 0x4e, 0xcb, 0xa4, 0x94, 0x05, 0x66, 0xe0, 0x30, 0xba, 0x8c, 0x63, 0xb4, 0x93, 0xad, 0xac,
	0x16, 0xf3, 0x3c, 0x46, 0x25, 0xe4, 0xd3, 0xad, 0x10, 0x2f, 0x8c, 0x89, 0x0c, 0x4e, 0x2c, 0xc6,
	0x17, 0x12, 0xfb, 0x5c, 0x62, 0x85, 0x35, 0x0f, 0x6f, 0x5a, 0x8b, 0x90, 0x0b, 0xd8, 0x43, 0xfe,
	0xdf, 0xb8, 0xe9, 0xfb, 0x84, 0xcb, 0x72, 0x5e, 0xfc, 0x5d, 0x83, 0x6a, 0x37, 0x6a, 0xe9, 0x9c,
	0xb9, 0x8e, 0xb5, 0x42, 0x08, 0xf6, 0xa9, 0xe9, 0x11, 0x55, 0x69, 0x28, 0xcd, 0x0a, 0x16, 0xbf,
	0xd1, 0x09, 0x1c, 0x2c, 0x9c, 0xa5, 0xef, 0x9a, 0x2b, 0x43, 0xf8, 0x72, 0xc2, 0x57, 0x95, 0x67,
	0xd3, 0x08, 0x72, 0x09, 0x87, 0x0b, 0x66, 0x85, 0x1e, 0xa1, 0x71, 0x91, 0xea, 0x61, 0x43, 0x69,
	0x56, 0xdb, 0x2d, 0x7d, 0xdb, 0x84, 0xf4, 0x4c, 0x42, 0x7d, 0x90, 0x0d, 0xc3, 0x9b, 0x2c, 0x08,
	0x43, 0x35, 0x5c, 0x12, 0x6e, 0xb8, 0xe6, 0x9c, 0xb8, 0x4b, 0xb5, 0xde, 0xc8, 0x37, 0xab, 0xed,
	0x2f, 0xde, 0x4e, 0x7a, 0xb9, 0x24, 0xfc, 0x27, 0x11, 0x33, 0xa4, 0x01, 0x5f, 0x61, 0x08, 0xd3,
	0x03, 0xf4, 0x23, 0x80, 0xc5, 0xe8, 0xc2, 0x11, 0x4b, 0x51, 0x0f, 0x04, 0xe5, 0x67, 0x6f, 0xa7,
	0xec, 0x27, 0x31, 0x38, 0x13, 0x8e, 0x2e, 0xa0, 0x6c, 0x31, 0x6f, 0xee, 0x50, 0xc2, 0xd5, 0x62,
	0x43, 0x69, 0x1e, 0xb5, 0xbf, 0xde, 0x81, 0xaa, 0x2f, 0x43, 0x67, 0x2b, 0x9f, 0xe0, 0x94, 0x08,
	0x7d, 0x09, 0x25, 0x42, 0xcd, 0xb9, 0x4b, 0x16, 0xea, 0x23, 0x31, 0x46, 0x2d, 0xe1, 0x4c, 0xb6,
	0xa8, 0xf7, 0x18, 0x73, 0xaf, 0x4c, 0x37, 0x24, 0x38, 0x81, 0xa2, 0x0e, 0x3c, 0xa1, 0x2c, 0x70,
	0x6e, 0x1c, 0x2b, 0x96, 0x89, 0x75, 0x6b, 0x52, 0x1a, 0x4d, 0xed, 0xa8, 0x91, 0x6f, 0x56, 0xf0,
	0x71, 0xd6, 0xd9, 0x97, 0x3e, 0x34, 0x81, 0x9a, 0xc5, 0x49, 0x56, 0x57, 0x2a, 0x88, 0x94, 0x1f,
	0x6f, 0x6f, 0x63, 0x22, 0x45, 0x88, 0x05, 0x16, 0x1f, 0x25, 0xc1, 0xb1, 0x1d, 0xd1, 0xdd, 0x93,
	0xa9, 0x5a, 0xdd, 0x85, 0xce, 0xdb, 0xb0, 0xb5, 0x53, 0x38, 0xdc, 0x90, 0x07, 0x52, 0xa1, 0x64,
	0x31, 0x1a, 0x10, 0x1a, 0x48, 0x81, 0x26, 0x26, 0xfa, 0x10, 0x2a, 0x9e, 0xe3, 0x11, 0x23, 0x58,
	0xf9, 0x89, 0x40, 0xcb, 0xd1, 0x41, 0x34, 0x5a, 0xed, 0xaf, 0x32, 0x54, 0xd2, 0xa1, 0xa7, 0x12,
	0x3f, 0x78, 0x83, 0xc4, 0x8b, 0xaf, 0x4b, 0x9c, 0xc2, 0xe3, 0x74, 0xf1, 0x46, 0x70, 0xcb, 0xc9,
	0xf2, 0x96, 0xb9, 0x0b, 0x51, 0x47, 0xb5, 0xfd, 0x72, 0x87, 0xad, 0xeb, 0x13, 0x12, 0x70, 0xc7,
	0x9a, 0x25, 0x14, 0xa3, 0x3d, 0x8c, 0x52, 0xe6, 0xf4, 0x14, 0xdd, 0x40, 0x7d, 0x9d, 0xcf, 0x9c,
	0x2f, 0xa3, 0xa6, 0x73, 0x22, 0xd9, 0xb7, 0xbb, 0x27, 0xeb, 0x46, 0xf1, 0x16, 0x19, 0xed, 0xe1,
	0x5a, 0x4a, 0x2a, 0xce, 0x02, 0x6d, 0x08, 0xa5, 0x19, 0x77, 0x6c, 0x9b, 0x70, 0xf4, 0x14, 0x0a,
	0x16, 0x0b, 0xe5, 0x70, 0x0b, 0xa3, 0x3d, 0x1c, 0x9b, 0x48, 0x83, 0x92, 0x4f, 0xb8, 0x95, 0x54,
	0xa0, 0x8c, 0xf6, 0x70, 0x72, 0xd0, 0x2b, 0xc2, 0x7e, 0x34, 0x73, 0xed, 0x9f, 0x3c, 0xd4, 0xee,
	0x35, 0x86, 0x9e, 0x42, 0xf1, 0xc6, 0x71, 0x03, 0xc2, 0xe5, 0x46, 0xa4, 0x85, 0x86, 0x70, 0x60,
	0xda, 0x36, 0x27, 0x76, 0xfc, 0x32, 0xaa, 0x65, 0x71, 0x09, 0x4f, 0x1e, 0x68, 0x6b, 0x8d, 0xc4,
	0x1b, 0x61, 0xe8, 0x73, 0x40, 0x0b, 0x42, 0x99, 0xe7, 0x50, 0x33, 0x60, 0xdc, 0x90, 0xa9, 0x2a,
	0x22, 0xd5, 0xa3, 0x8c, 0xe7, 0x34, 0xce, 0x7a, 0x0d, 0x6a, 0x16, 0xbe, 0x51, 0x01, 0xfc, 0xdf,
	0x0a, 0x9e, 0x65, 0x28, 0xba, 0xd9, 0x62, 0x06, 0xd1, 0xb3, 0xe2, 0xf9, 0x26, 0x77, 0x96, 0x8c,
	0xaa, 0xfb, 0xe2, 0x2d, 0x78, 0x40, 0xf5, 0xfd, 0x14, 0x27, 0x2e, 0x7e, 0x26, 0x0e, 0x7d, 0x02,
	0xb5, 0x54, 0x5a, 0xc6, 0x5d, 0x74, 0xc1, 0xd5, 0x42, 0x34, 0x71, 0x7c, 0x94, 0x1e, 0x8b, 0x6b,
	0x8f, 0xbe, 0x82, 0x72, 0xf2, 0xd2, 0x0b, 0xb1, 0x56, 0xdb, 0x1f, 0xbc, 0xf6, 0x48, 0x0c, 0x24,
	0x00, 0xa7, 0x50, 0x34, 0x81, 0x52, 0x10, 0x2f, 0x5b, 0x2d, 0x89, 0xa8, 0xce, 0x2e, 0x5a, 0x92,
	0x3a, 0xc1, 0x09, 0x87, 0xf6, 0xaf, 0x02, 0x87, 0x1b, 0x02, 0xcb, 0xac, 0x5c, 0x79, 0xe3, 0xca,
	0x0b, 0xef, 0xb6, 0xf2, 0x6c, 0xdb, 0xb9, 0x77, 0x6a, 0x3b, 0xff, 0xfe, 0x6d, 0xf7, 0xaa, 0x50,
	0x49, 0x6f, 0x91, 0xf6, 0x3d, 0xd4, 0xee, 0x7d, 0x6e, 0x50, 0x1d, 0xf2, 0xbf, 0x92, 0x95, 0x9c,
	0x40, 0xf4, 0x13, 0x1d, 0x43, 0x21, 0xde, 0x66, 0x7c, 0x11, 0x62, 0xe3, 0xbb, 0xdc, 0x37, 0xca,
	0x0b, 0x13, 0x9e, 0x6c, 0xfd, 0x1e, 0xa0, 0x67, 0xf0, 0xb8, 0x7f, 0x36, 0xe9, 0x8d, 0xa7, 0x43,
	0xe3, 0x72, 0x7a, 0x71, 0x3e, 0xec, 0x8f, 0x4f, 0xc7, 0xc3, 0x41, 0x7d, 0x0f, 0x95, 0x20, 0xdf,
	0x9d, 0x0e, 0xea, 0x0a, 0x2a, 0x42, 0xee, 0x0c, 0xd7, 0x73, 0xe8, 0x39, 0x68, 0xdd, 0xe9, 0xc0,
	0xf8, 0x79, 0x3c, 0x1b, 0x19, 0x93, 0xee, 0xac, 0x3f, 0x1a, 0x4f, 0x5f, 0x19, 0x78, 0x78, 0x71,
	0x76, 0x89, 0xfb, 0xc3, 0x7a, 0xbe, 0xf7, 0xbb, 0x02, 0xaa, 0xc5, 0xbc, 0xad, 0x2d, 0xf7, 0x20,
	0xee, 0x39, 0x1a, 0xde, 0xb9, 0xf2, 0xcb, 0x0f, 0x12, 0x63, 0x33, 0xd7, 0xa4, 0xb6, 0xce, 0xb8,
	0xdd, 0xb2, 0x09, 0x15, 0xa3, 0x6d, 0xc5, 0x2e, 0xd3, 0x77, 0x96, 0x9b, 0xff, 0x4c, 0x5e, 0xae,
	0xad, 0x3f, 0x72, 0xda, 0xab, 0x98, 0xa0, 0xef, 0xb2, 0x70, 0xa1, 0x4f, 0xd6, 0xa9, 0xae, 0x3a,
	0x7f, 0x26, 0xce, 0x6b, 0xe1, 0xbc, 0x5e, 0x3b, 0xaf, 0xaf, 0x3a, 0xf3, 0xa2, 0x48, 0xd2, 0xf9,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x66, 0xb5, 0x16, 0x64, 0x76, 0x09, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/metrics.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PrometheusTlsConfig_Mode int32

const (
	// control-plane delivers certificates to the prometheus listener.
	// This should be used when prometheus is running inside the Mesh.
	PrometheusTlsConfig_activeMTLSBackend PrometheusTlsConfig_Mode = 0
	// In this way user is resposible for providing certificates to dataplanes.
	// Path for the certificte and the key needs to be provided to the dataplane
	// by environments variables:
	// * DUBBO_DATAPLANE_RUNTIME_METRICS_CERT_PATH
	// * DUBBO_DATAPLANE_RUNTIME_METRICS_KEY_PATH
	PrometheusTlsConfig_providedTLS PrometheusTlsConfig_Mode = 1
	// allows disabling TLS for the prometheus listener.
	PrometheusTlsConfig_disabled PrometheusTlsConfig_Mode = 2
)

// Enum value maps for PrometheusTlsConfig_Mode.
var (
	PrometheusTlsConfig_Mode_name = map[int32]string{
		0: "activeMTLSBackend",
		1: "providedTLS",
		2: "disabled",
	}
	PrometheusTlsConfig_Mode_value = map[string]int32{
		"activeMTLSBackend": 0,
		"providedTLS":       1,
		"disabled":          2,
	}
)

func (x PrometheusTlsConfig_Mode) Enum() *PrometheusTlsConfig_Mode {
	p := new(PrometheusTlsConfig_Mode)
	*p = x
	return p
}

func (x PrometheusTlsConfig_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrometheusTlsConfig_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_mesh_v1alpha1_metrics_proto_enumTypes[0].Descriptor()
}

func (PrometheusTlsConfig_Mode) Type() protoreflect.EnumType {
	return &file_api_mesh_v1alpha1_metrics_proto_enumTypes[0]
}

func (x PrometheusTlsConfig_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrometheusTlsConfig_Mode.Descriptor instead.
func (PrometheusTlsConfig_Mode) EnumDescriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{5, 0}
}

// Metrics defines configuration for metrics that should be collected and
// exposed by dataplanes.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the enabled backend
	EnabledBackend string `protobuf:"bytes,1,opt,name=enabledBackend,proto3" json:"enabledBackend,omitempty"`
	// List of available Metrics backends
	Backends []*MetricsBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metrics) GetEnabledBackend() string {
	if x != nil {
		return x.EnabledBackend
	}
	return ""
}

func (x *Metrics) GetBackends() []*MetricsBackend {
	if x != nil {
		return x.Backends
	}
	return nil
}

// MetricsBackend defines metric backends
type MetricsBackend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the backend, can be then used in Mesh.metrics.enabledBackend
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the backend (Dubbo ships with 'prometheus')
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Configuration of the backend
	Conf *structpb.Struct `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *MetricsBackend) Reset() {
	*x = MetricsBackend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsBackend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsBackend) ProtoMessage() {}

func (x *MetricsBackend) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsBackend.ProtoReflect.Descriptor instead.
func (*MetricsBackend) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *MetricsBackend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricsBackend) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MetricsBackend) GetConf() *structpb.Struct {
	if x != nil {
		return x.Conf
	}
	return nil
}

// PrometheusMetricsBackendConfig defines configuration of Prometheus backend
type PrometheusMetricsBackendConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port on which a dataplane should expose HTTP endpoint with Prometheus
	// metrics.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Path on which a dataplane should expose HTTP endpoint with Prometheus
	// metrics.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Tags associated with an application this dataplane is deployed next to,
	// e.g. service=web, version=1.0.
	// `service` tag is mandatory.
	Tags map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If true then endpoints for scraping metrics won't require mTLS even if mTLS
	// is enabled in Mesh. If nil, then it is treated as false.
	SkipMTLS *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=skipMTLS,proto3" json:"skipMTLS,omitempty"`
	// Map with the configuration of applications which metrics are going to be
	// scrapped by dubbo-dp.
	Aggregate []*PrometheusAggregateMetricsConfig `protobuf:"bytes,5,rep,name=aggregate,proto3" json:"aggregate,omitempty"`
	// Configuration of Envoy's metrics.
	Envoy *PrometheusEnvoyConfig `protobuf:"bytes,6,opt,name=envoy,proto3" json:"envoy,omitempty"`
	// Configuration of TLS for prometheus listener.
	Tls *PrometheusTlsConfig `protobuf:"bytes,7,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *PrometheusMetricsBackendConfig) Reset() {
	*x = PrometheusMetricsBackendConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusMetricsBackendConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusMetricsBackendConfig) ProtoMessage() {}

func (x *PrometheusMetricsBackendConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusMetricsBackendConfig.ProtoReflect.Descriptor instead.
func (*PrometheusMetricsBackendConfig) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *PrometheusMetricsBackendConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PrometheusMetricsBackendConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PrometheusMetricsBackendConfig) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PrometheusMetricsBackendConfig) GetSkipMTLS() *wrapperspb.BoolValue {
	if x != nil {
		return x.SkipMTLS
	}
	return nil
}

func (x *PrometheusMetricsBackendConfig) GetAggregate() []*PrometheusAggregateMetricsConfig {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *PrometheusMetricsBackendConfig) GetEnvoy() *PrometheusEnvoyConfig {
	if x != nil {
		return x.Envoy
	}
	return nil
}

func (x *PrometheusMetricsBackendConfig) GetTls() *PrometheusTlsConfig {
	if x != nil {
		return x.Tls
	}
	return nil
}

// PrometheusAggregateMetricsConfig defines endpoints that should be scrapped
// by dubbo-dp for prometheus metrics.
// Any configuration change require sidecar restart.
type PrometheusAggregateMetricsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name which identify given configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Port on which a service expose HTTP endpoint with Prometheus metrics.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Path on which a service expose HTTP endpoint with Prometheus metrics.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// If false then the application won't be scrapped. If nil, then it is treated
	// as true and dubbo-dp scrapes metrics from the service.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Address on which a service expose HTTP endpoint with Prometheus metrics.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PrometheusAggregateMetricsConfig) Reset() {
	*x = PrometheusAggregateMetricsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusAggregateMetricsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusAggregateMetricsConfig) ProtoMessage() {}

func (x *PrometheusAggregateMetricsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusAggregateMetricsConfig.ProtoReflect.Descriptor instead.
func (*PrometheusAggregateMetricsConfig) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *PrometheusAggregateMetricsConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrometheusAggregateMetricsConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PrometheusAggregateMetricsConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PrometheusAggregateMetricsConfig) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *PrometheusAggregateMetricsConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// PrometheusEnvoyConfig defines filters that should be passed to Envoy
// for filtering.
type PrometheusEnvoyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FilterRegex value that is going to be passed to Envoy for filtering
	// Envoy metrics.
	FilterRegex string `protobuf:"bytes,1,opt,name=filterRegex,proto3" json:"filterRegex,omitempty"`
	// If true then return metrics that Envoy has updated (counters incremented
	// at least once, gauges changed at least once, and histograms added to at
	// least once). If nil, then it is treated as false.
	UsedOnly *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=usedOnly,proto3" json:"usedOnly,omitempty"`
}

func (x *PrometheusEnvoyConfig) Reset() {
	*x = PrometheusEnvoyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusEnvoyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusEnvoyConfig) ProtoMessage() {}

func (x *PrometheusEnvoyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusEnvoyConfig.ProtoReflect.Descriptor instead.
func (*PrometheusEnvoyConfig) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *PrometheusEnvoyConfig) GetFilterRegex() string {
	if x != nil {
		return x.FilterRegex
	}
	return ""
}

func (x *PrometheusEnvoyConfig) GetUsedOnly() *wrapperspb.BoolValue {
	if x != nil {
		return x.UsedOnly
	}
	return nil
}

// PrometheusEnvoyConfig defines Tls configuration for Prometheus listener.
type PrometheusTlsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mode defines how configured is the TLS for Prometheus.
	// Supported values, delegated, disabled, activeMTLSBackend. Default to
	// `activeMTLSBackend`.
	Mode PrometheusTlsConfig_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=dubbo.mesh.v1alpha1.PrometheusTlsConfig_Mode" json:"mode,omitempty"`
}

func (x *PrometheusTlsConfig) Reset() {
	*x = PrometheusTlsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusTlsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusTlsConfig) ProtoMessage() {}

func (x *PrometheusTlsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_metrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusTlsConfig.ProtoReflect.Descriptor instead.
func (*PrometheusTlsConfig) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP(), []int{5}
}

func (x *PrometheusTlsConfig) GetMode() PrometheusTlsConfig_Mode {
	if x != nil {
		return x.Mode
	}
	return PrometheusTlsConfig_activeMTLSBackend
}

var File_api_mesh_v1alpha1_metrics_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_metrics_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x08,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x65, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22,
	0xdf, 0x03, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x51, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x36, 0x0a,
	0x08, 0x73, 0x6b, 0x69, 0x70, 0x4d, 0x54, 0x4c, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x73, 0x6b, 0x69,
	0x70, 0x4d, 0x54, 0x4c, 0x53, 0x12, 0x53, 0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x12, 0x3a, 0x0a, 0x03,
	0x74, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xae, 0x01, 0x0a, 0x20, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x71, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x45, 0x6e, 0x76, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x36, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x54, 0x6c, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x22, 0x3c, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x54, 0x4c, 0x53, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mesh_v1alpha1_metrics_proto_rawDescOnce sync.Once
	file_api_mesh_v1alpha1_metrics_proto_rawDescData = file_api_mesh_v1alpha1_metrics_proto_rawDesc
)

func file_api_mesh_v1alpha1_metrics_proto_rawDescGZIP() []byte {
	file_api_mesh_v1alpha1_metrics_proto_rawDescOnce.Do(func() {
		file_api_mesh_v1alpha1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mesh_v1alpha1_metrics_proto_rawDescData)
	})
	return file_api_mesh_v1alpha1_metrics_proto_rawDescData
}

var file_api_mesh_v1alpha1_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_mesh_v1alpha1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_mesh_v1alpha1_metrics_proto_goTypes = []interface{}{
	(PrometheusTlsConfig_Mode)(0),            // 0: dubbo.mesh.v1alpha1.PrometheusTlsConfig.Mode
	(*Metrics)(nil),                          // 1: dubbo.mesh.v1alpha1.Metrics
	(*MetricsBackend)(nil),                   // 2: dubbo.mesh.v1alpha1.MetricsBackend
	(*PrometheusMetricsBackendConfig)(nil),   // 3: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig
	(*PrometheusAggregateMetricsConfig)(nil), // 4: dubbo.mesh.v1alpha1.PrometheusAggregateMetricsConfig
	(*PrometheusEnvoyConfig)(nil),            // 5: dubbo.mesh.v1alpha1.PrometheusEnvoyConfig
	(*PrometheusTlsConfig)(nil),              // 6: dubbo.mesh.v1alpha1.PrometheusTlsConfig
	nil,                                      // 7: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.TagsEntry
	(*structpb.Struct)(nil),                  // 8: google.protobuf.Struct
	(*wrapperspb.BoolValue)(nil),             // 9: google.protobuf.BoolValue
}
var file_api_mesh_v1alpha1_metrics_proto_depIdxs = []int32{
	2,  // 0: dubbo.mesh.v1alpha1.Metrics.backends:type_name -> dubbo.mesh.v1alpha1.MetricsBackend
	8,  // 1: dubbo.mesh.v1alpha1.MetricsBackend.conf:type_name -> google.protobuf.Struct
	7,  // 2: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.tags:type_name -> dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.TagsEntry
	9,  // 3: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.skipMTLS:type_name -> google.protobuf.BoolValue
	4,  // 4: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.aggregate:type_name -> dubbo.mesh.v1alpha1.PrometheusAggregateMetricsConfig
	5,  // 5: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.envoy:type_name -> dubbo.mesh.v1alpha1.PrometheusEnvoyConfig
	6,  // 6: dubbo.mesh.v1alpha1.PrometheusMetricsBackendConfig.tls:type_name -> dubbo.mesh.v1alpha1.PrometheusTlsConfig
	9,  // 7: dubbo.mesh.v1alpha1.PrometheusAggregateMetricsConfig.enabled:type_name -> google.protobuf.BoolValue
	9,  // 8: dubbo.mesh.v1alpha1.PrometheusEnvoyConfig.usedOnly:type_name -> google.protobuf.BoolValue
	0,  // 9: dubbo.mesh.v1alpha1.PrometheusTlsConfig.mode:type_name -> dubbo.mesh.v1alpha1.PrometheusTlsConfig.Mode
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_metrics_proto_init() }
func file_api_mesh_v1alpha1_metrics_proto_init() {
	if File_api_mesh_v1alpha1_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsBackend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusMetricsBackendConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusAggregateMetricsConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusEnvoyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_mesh_v1alpha1_metrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusTlsConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_mesh_v1alpha1_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_mesh_v1alpha1_metrics_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_metrics_proto_depIdxs,
		EnumInfos:         file_api_mesh_v1alpha1_metrics_proto_enumTypes,
		MessageInfos:      file_api_mesh_v1alpha1_metrics_proto_msgTypes,
	}.Build()
	File_api_mesh_v1alpha1_metrics_proto = out.File
	file_api_mesh_v1alpha1_metrics_proto_rawDesc = nil
	file_api_mesh_v1alpha1_metrics_proto_goTypes = nil
	file_api_mesh_v1alpha1_metrics_proto_depIdxs = nil
}

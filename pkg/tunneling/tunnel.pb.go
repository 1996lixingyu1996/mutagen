// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: tunneling/tunnel.proto

package tunneling

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// TunnelHostCredentials is the format for the tunnel host credentials file.
type TunnelHostCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier is the (unique) tunnel identifier. It is static. It cannot be
	// empty.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version is the tunnel version. It is static.
	Version Version `protobuf:"varint,2,opt,name=version,proto3,enum=tunneling.Version" json:"version,omitempty"`
	// CreationTime is the creation time of the tunnel. It is static. It cannot
	// be nil.
	CreationTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	// CreatingVersionMajor is the major version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionMajor uint32 `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	// CreatingVersionMinor is the minor version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionMinor uint32 `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	// CreatingVersionPatch is the patch version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionPatch uint32 `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	// Token is the API access token for the tunnel endpoint. It is static.
	Token string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	// Secret is the HMAC secret key used for signing and validating offers. It
	// is static.
	Secret []byte `protobuf:"bytes,8,opt,name=secret,proto3" json:"secret,omitempty"`
	// Configuration is the flattened tunnel configuration. It must not be nil.
	// It is static.
	Configuration *Configuration `protobuf:"bytes,9,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *TunnelHostCredentials) Reset() {
	*x = TunnelHostCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunneling_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelHostCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelHostCredentials) ProtoMessage() {}

func (x *TunnelHostCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_tunneling_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelHostCredentials.ProtoReflect.Descriptor instead.
func (*TunnelHostCredentials) Descriptor() ([]byte, []int) {
	return file_tunneling_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *TunnelHostCredentials) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *TunnelHostCredentials) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_Invalid
}

func (x *TunnelHostCredentials) GetCreationTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *TunnelHostCredentials) GetCreatingVersionMajor() uint32 {
	if x != nil {
		return x.CreatingVersionMajor
	}
	return 0
}

func (x *TunnelHostCredentials) GetCreatingVersionMinor() uint32 {
	if x != nil {
		return x.CreatingVersionMinor
	}
	return 0
}

func (x *TunnelHostCredentials) GetCreatingVersionPatch() uint32 {
	if x != nil {
		return x.CreatingVersionPatch
	}
	return 0
}

func (x *TunnelHostCredentials) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TunnelHostCredentials) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *TunnelHostCredentials) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

// Tunnel represents a tunnel client configuration.
type Tunnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier is the (unique) tunnel identifier. It is static. It cannot be
	// empty.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version is the tunnel version. It is static.
	Version Version `protobuf:"varint,2,opt,name=version,proto3,enum=tunneling.Version" json:"version,omitempty"`
	// CreationTime is the creation time of the tunnel. It is static. It cannot
	// be nil.
	CreationTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	// CreatingVersionMajor is the major version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionMajor uint32 `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	// CreatingVersionMinor is the minor version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionMinor uint32 `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	// CreatingVersionPatch is the patch version component of the version of
	// Mutagen which created the tunnel. It is static.
	CreatingVersionPatch uint32 `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	// Token is the API access token for the tunnel endpoint. It is static.
	Token string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	// Secret is the HMAC secret key used for signing and validating offers. It
	// is static.
	Secret []byte `protobuf:"bytes,8,opt,name=secret,proto3" json:"secret,omitempty"`
	// Configuration is the flattened tunnel configuration. It must not be nil.
	// It is static.
	Configuration *Configuration `protobuf:"bytes,9,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Name is a user-friendly name for the tunnel. It may be empty and but is
	// guaranteed to be unique across all tunnels if non-empty. It is only used
	// as a simpler handle for specifying tunnels. It is static.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the tunnel labels. They are static.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not the tunnel is marked as paused.
	Paused bool `protobuf:"varint,12,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *Tunnel) Reset() {
	*x = Tunnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunneling_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tunnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tunnel) ProtoMessage() {}

func (x *Tunnel) ProtoReflect() protoreflect.Message {
	mi := &file_tunneling_tunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tunnel.ProtoReflect.Descriptor instead.
func (*Tunnel) Descriptor() ([]byte, []int) {
	return file_tunneling_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *Tunnel) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Tunnel) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_Invalid
}

func (x *Tunnel) GetCreationTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *Tunnel) GetCreatingVersionMajor() uint32 {
	if x != nil {
		return x.CreatingVersionMajor
	}
	return 0
}

func (x *Tunnel) GetCreatingVersionMinor() uint32 {
	if x != nil {
		return x.CreatingVersionMinor
	}
	return 0
}

func (x *Tunnel) GetCreatingVersionPatch() uint32 {
	if x != nil {
		return x.CreatingVersionPatch
	}
	return 0
}

func (x *Tunnel) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Tunnel) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *Tunnel) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *Tunnel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tunnel) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Tunnel) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

var File_tunneling_tunnel_proto protoreflect.FileDescriptor

var file_tunneling_tunnel_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a,
	0x15, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3e,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe,
	0x04, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69,
	0x6e, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12,
	0x32, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75,
	0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tunneling_tunnel_proto_rawDescOnce sync.Once
	file_tunneling_tunnel_proto_rawDescData = file_tunneling_tunnel_proto_rawDesc
)

func file_tunneling_tunnel_proto_rawDescGZIP() []byte {
	file_tunneling_tunnel_proto_rawDescOnce.Do(func() {
		file_tunneling_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunneling_tunnel_proto_rawDescData)
	})
	return file_tunneling_tunnel_proto_rawDescData
}

var file_tunneling_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tunneling_tunnel_proto_goTypes = []interface{}{
	(*TunnelHostCredentials)(nil), // 0: tunneling.TunnelHostCredentials
	(*Tunnel)(nil),                // 1: tunneling.Tunnel
	nil,                           // 2: tunneling.Tunnel.LabelsEntry
	(Version)(0),                  // 3: tunneling.Version
	(*timestamp.Timestamp)(nil),   // 4: google.protobuf.Timestamp
	(*Configuration)(nil),         // 5: tunneling.Configuration
}
var file_tunneling_tunnel_proto_depIdxs = []int32{
	3, // 0: tunneling.TunnelHostCredentials.version:type_name -> tunneling.Version
	4, // 1: tunneling.TunnelHostCredentials.creationTime:type_name -> google.protobuf.Timestamp
	5, // 2: tunneling.TunnelHostCredentials.configuration:type_name -> tunneling.Configuration
	3, // 3: tunneling.Tunnel.version:type_name -> tunneling.Version
	4, // 4: tunneling.Tunnel.creationTime:type_name -> google.protobuf.Timestamp
	5, // 5: tunneling.Tunnel.configuration:type_name -> tunneling.Configuration
	2, // 6: tunneling.Tunnel.labels:type_name -> tunneling.Tunnel.LabelsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tunneling_tunnel_proto_init() }
func file_tunneling_tunnel_proto_init() {
	if File_tunneling_tunnel_proto != nil {
		return
	}
	file_tunneling_configuration_proto_init()
	file_tunneling_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tunneling_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelHostCredentials); i {
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
		file_tunneling_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tunnel); i {
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
			RawDescriptor: file_tunneling_tunnel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tunneling_tunnel_proto_goTypes,
		DependencyIndexes: file_tunneling_tunnel_proto_depIdxs,
		MessageInfos:      file_tunneling_tunnel_proto_msgTypes,
	}.Build()
	File_tunneling_tunnel_proto = out.File
	file_tunneling_tunnel_proto_rawDesc = nil
	file_tunneling_tunnel_proto_goTypes = nil
	file_tunneling_tunnel_proto_depIdxs = nil
}

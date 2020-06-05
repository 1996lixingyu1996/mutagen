// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: synchronization/core/mode.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
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

// SynchronizationMode specifies the mode for synchronization, encoding both
// directionality and conflict resolution behavior.
type SynchronizationMode int32

const (
	// SynchronizationMode_SynchronizationModeDefault represents an unspecified
	// synchronization mode. It is not valid for use with Reconcile. It should
	// be converted to one of the following values based on the desired default
	// behavior.
	SynchronizationMode_SynchronizationModeDefault SynchronizationMode = 0
	// SynchronizationMode_SynchronizationModeTwoWaySafe represents a
	// bidirectional synchronization mode where automatic conflict resolution is
	// performed only in cases where no data would be lost. Specifically, this
	// means that modified contents are allowed to propagate to the opposite
	// endpoint if the corresponding contents on the opposite endpoint are
	// unmodified or deleted. All other conflicts are left unresolved.
	SynchronizationMode_SynchronizationModeTwoWaySafe SynchronizationMode = 1
	// SynchronizationMode_SynchronizationModeTwoWayResolved is the same as
	// SynchronizationMode_SynchronizationModeTwoWaySafe, but specifies that the
	// alpha endpoint should win automatically in any conflict between alpha and
	// beta, including cases where alpha has deleted contents that beta has
	// modified.
	SynchronizationMode_SynchronizationModeTwoWayResolved SynchronizationMode = 2
	// SynchronizationMode_SynchronizationModeOneWaySafe represents a
	// unidirectional synchronization mode where contents and changes propagate
	// from alpha to beta, but won't overwrite any creations or modifications on
	// beta.
	SynchronizationMode_SynchronizationModeOneWaySafe SynchronizationMode = 3
	// SynchronizationMode_SynchronizationModeOneWayReplica represents a
	// unidirectional synchronization mode where contents on alpha are mirrored
	// (verbatim) to beta, overwriting any conflicting contents on beta and
	// deleting any extraneous contents on beta.
	SynchronizationMode_SynchronizationModeOneWayReplica SynchronizationMode = 4
)

// Enum value maps for SynchronizationMode.
var (
	SynchronizationMode_name = map[int32]string{
		0: "SynchronizationModeDefault",
		1: "SynchronizationModeTwoWaySafe",
		2: "SynchronizationModeTwoWayResolved",
		3: "SynchronizationModeOneWaySafe",
		4: "SynchronizationModeOneWayReplica",
	}
	SynchronizationMode_value = map[string]int32{
		"SynchronizationModeDefault":        0,
		"SynchronizationModeTwoWaySafe":     1,
		"SynchronizationModeTwoWayResolved": 2,
		"SynchronizationModeOneWaySafe":     3,
		"SynchronizationModeOneWayReplica":  4,
	}
)

func (x SynchronizationMode) Enum() *SynchronizationMode {
	p := new(SynchronizationMode)
	*p = x
	return p
}

func (x SynchronizationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SynchronizationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_core_mode_proto_enumTypes[0].Descriptor()
}

func (SynchronizationMode) Type() protoreflect.EnumType {
	return &file_synchronization_core_mode_proto_enumTypes[0]
}

func (x SynchronizationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SynchronizationMode.Descriptor instead.
func (SynchronizationMode) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_core_mode_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_core_mode_proto protoreflect.FileDescriptor

var file_synchronization_core_mode_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0xc8, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12,
	0x21, 0x0a, 0x1d, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x53, 0x61, 0x66, 0x65,
	0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x53, 0x61, 0x66, 0x65, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20,
	0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x10, 0x04, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_mode_proto_rawDescOnce sync.Once
	file_synchronization_core_mode_proto_rawDescData = file_synchronization_core_mode_proto_rawDesc
)

func file_synchronization_core_mode_proto_rawDescGZIP() []byte {
	file_synchronization_core_mode_proto_rawDescOnce.Do(func() {
		file_synchronization_core_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_mode_proto_rawDescData)
	})
	return file_synchronization_core_mode_proto_rawDescData
}

var file_synchronization_core_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_core_mode_proto_goTypes = []interface{}{
	(SynchronizationMode)(0), // 0: core.SynchronizationMode
}
var file_synchronization_core_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_core_mode_proto_init() }
func file_synchronization_core_mode_proto_init() {
	if File_synchronization_core_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_core_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_mode_proto_goTypes,
		DependencyIndexes: file_synchronization_core_mode_proto_depIdxs,
		EnumInfos:         file_synchronization_core_mode_proto_enumTypes,
	}.Build()
	File_synchronization_core_mode_proto = out.File
	file_synchronization_core_mode_proto_rawDesc = nil
	file_synchronization_core_mode_proto_goTypes = nil
	file_synchronization_core_mode_proto_depIdxs = nil
}

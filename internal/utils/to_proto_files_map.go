package utils

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// toProtoFilesMap creates a map of file names to their corresponding FileDescriptorProto objects
// from a CodeGeneratorRequest. This map can be used to quickly access file descriptors by their names.
//
// Parameters:
// req - A pointer to a CodeGeneratorRequest which contains the list of Proto files.
//
// Returns:
// A map where the keys are file names and the values are pointers to FileDescriptorProto objects.
func ToProtoFilesMap(req *pluginpb.CodeGeneratorRequest) map[string]*descriptorpb.FileDescriptorProto {
	files := make(map[string]*descriptorpb.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}
	return files
}

package maptypedefinitions

import (
	"strings"

	"github.com/dclappert/protoc-gen-salesforce-apex/internal/utils"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type MapTypeDefinitions struct {
	typeDefinitionByDefinitionKey map[string]*descriptorpb.DescriptorProto
}

func CreateFrom(codeGeneratorRequest *pluginpb.CodeGeneratorRequest) *MapTypeDefinitions {
	instance := (&MapTypeDefinitions{
		typeDefinitionByDefinitionKey: make(map[string]*descriptorpb.DescriptorProto),
	})
	instance.createMapTypeDefinitionsMap(codeGeneratorRequest)
	return instance
}

func (m *MapTypeDefinitions) Has(key string) bool {
	if _, ok := m.typeDefinitionByDefinitionKey[key]; ok {
		return ok
	}
	return false
}

func (m *MapTypeDefinitions) Get(key string) *descriptorpb.DescriptorProto {
	return m.typeDefinitionByDefinitionKey[key]
}

func (m *MapTypeDefinitions) createMapTypeDefinitionsMap(req *pluginpb.CodeGeneratorRequest) {
	protoFiles := utils.ToProtoFilesMap(req)
	for _, fileName := range req.FileToGenerate {
		protoFile := protoFiles[fileName]
		for _, message := range protoFile.MessageType {
			for _, nestedType := range message.NestedType {
				if !isNestedTypeAMap(nestedType) {
					continue
				}
				m.typeDefinitionByDefinitionKey[fromNestTypeToMapTypeDefinitionKey(protoFile, message, nestedType)] = nestedType
			}
		}
	}
}

// fromNestTypeToMapTypeDefinitionKey generates a unique string key for a map type definition in a Protobuf file.
// The key is composed of the package name, message name, and nested type name (with "Entry" removed),
// all converted to lower case and concatenated with dots.
//
// Parameters:
// protoFile - A pointer to the FileDescriptorProto where the map type is defined.
// message - A pointer to the DescriptorProto for the message that contains the map.
// nestedType - A pointer to the DescriptorProto representing the map's nested type.
//
// Returns:
// A string representing the unique key for the map type definition.
func fromNestTypeToMapTypeDefinitionKey(protoFile *descriptorpb.FileDescriptorProto, message *descriptorpb.DescriptorProto, nestedType *descriptorpb.DescriptorProto) string {
	return strings.ToLower(protoFile.GetPackage()) + "." + strings.ToLower(message.GetName()) + "." + strings.ToLower(strings.Replace(nestedType.GetName(), "Entry", "", -1))
}

// isNestedTypeAMap checks if a given nested type in a Protobuf message represents a map.
// It returns true if the nested type has exactly two fields named 'key' and 'value',
// which is the typical structure for map entries in Protobuf.
//
// Parameters:
// nestedType - A pointer to a DescriptorProto representing the nested type.
//
// Returns:
// A boolean indicating whether the nested type represents a map.
func isNestedTypeAMap(nestedType *descriptorpb.DescriptorProto) bool {
	return len(nestedType.Field) == 2 && nestedType.Field[0].GetName() == "key" && nestedType.Field[1].GetName() == "value"
}

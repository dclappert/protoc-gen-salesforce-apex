package apexclassmetadata

import (
	"bytes"
	"text/template"

	"github.com/dclappert/protoc-gen-salesforce-apex/internal/optionsparser"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const CLASS_METADATA_TEMPLATE_STRING = `<?xml version="1.0" encoding="UTF-8"?>
<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>{{ .ApiVersion }}</apiVersion>
    <status>Active</status>
</ApexClass>`

type ClassMetadataTemplateBind struct {
	ApiVersion *string
}

type ApexClassMetadata struct {
	classMetadataTemplate *template.Template
	params                *FromParams
}

type FromParams struct {
	Message              *descriptorpb.DescriptorProto
	CodeGeneratorRequest *pluginpb.CodeGeneratorRequest
}

func New(params *FromParams) *ApexClassMetadata {
	apexClassMetadata := &ApexClassMetadata{
		params:                params,
		classMetadataTemplate: parseTemplate(),
	}
	return apexClassMetadata
}

func (c *ApexClassMetadata) GetFile() *pluginpb.CodeGeneratorResponse_File {
	options := optionsparser.ParseOptions(c.params.CodeGeneratorRequest.GetParameter())
	b := bytes.NewBuffer([]byte{})
	err := c.classMetadataTemplate.Execute(b, ClassMetadataTemplateBind{ApiVersion: &options.ApiVersion})
	if err != nil {
		panic(err)
	}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(c.params.Message.GetName() + ".cls-meta.xml"),
		Content: proto.String(b.String()),
	}
}

func parseTemplate() *template.Template {
	var err error
	template, err := template.New("classMetadata").Parse(CLASS_METADATA_TEMPLATE_STRING)
	if err != nil {
		panic(err)
	}
	return template
}

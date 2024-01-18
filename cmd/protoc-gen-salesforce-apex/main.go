package main

import (
	"io"
	"os"

	"github.com/dclappert/protoc-gen-salesforce-apex/internal/apexclass"
	"github.com/dclappert/protoc-gen-salesforce-apex/internal/apexclassmetadata"
	"github.com/dclappert/protoc-gen-salesforce-apex/internal/utils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	codeGeneratorRequest := toCodeGeneratorRequest(os.Stdin)
	emitCodeGeneratorResponse(generateApex(codeGeneratorRequest))
}

func toCodeGeneratorRequest(r io.Reader) *pluginpb.CodeGeneratorRequest {
	buf, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	var codeGeneratorRequest pluginpb.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &codeGeneratorRequest); err != nil {
		panic(err)
	}
	return &codeGeneratorRequest
}

func generateApex(req *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	protoFiles := utils.ToProtoFilesMap(req)
	var resp pluginpb.CodeGeneratorResponse
	for _, fileName := range req.FileToGenerate {
		protoFile := protoFiles[fileName]
		for _, message := range protoFile.MessageType {
			resp.File = append(resp.File, apexclass.New(&apexclass.FromParams{
				Message:              message,
				ProtoFile:            protoFile,
				CodeGeneratorRequest: req,
			}).GetFile())
			resp.File = append(resp.File, apexclassmetadata.New(&apexclassmetadata.FromParams{
				Message:              message,
				CodeGeneratorRequest: req,
			}).GetFile())
		}
	}
	return &resp
}

func emitCodeGeneratorResponse(resp *pluginpb.CodeGeneratorResponse) {
	setSupportedFeatures(resp)
	buf, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(buf)
	if err != nil {
		panic(err)
	}
}

func setSupportedFeatures(resp *pluginpb.CodeGeneratorResponse) {
	resp.SupportedFeatures = proto.Uint64(uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL))
}

#!/bin/bash
go install github.com/dclappert/protoc-gen-salesforce-apex/cmd/protoc-gen-salesforce-apex \
  && rm -rf ./examples/target/classes \
  && mkdir -p ./examples/target/classes \
  && protoc \
    --proto_path=./examples/proto \
    --salesforce-apex_out=./examples/target/classes \
    --salesforce-apex_opt=apiVersion="57.0",useProtoFieldNames=true \
    example.proto
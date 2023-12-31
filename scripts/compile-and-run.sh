#!/bin/bash
go install ./cmd/... \
  && rm -rf ./examples/target/classes \
  && mkdir -p ./examples/target/classes \
  && protoc \
    --proto_path=./examples/proto \
    --salesforce-apex_out=./examples/target/classes \
    --salesforce-apex_opt=apiVersion="56.0",useProtoFieldNames=false \
    explorer_note.proto adventure_log.proto
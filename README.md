# protoc-gen-salesforce-apex
A Protoc plugin that generates Salesforce Apex from Protocol Buffer Definitions.

## Installation
To install this Protoc plugin, simply run the following command:
```
go install github.com/dclappert/protoc-gen-salesforce-apex/cmd/protoc-gen-salesforce-apex@v1.0.0
```

## Usage 
To use the plugin, execute protoc with the --salesforce-apex_out flag, specifying your desired output directory:
```
protoc \
    --proto_path=./examples/proto \
    --salesforce-apex_out=./examples/target/classes \
    example.proto
```

## Custom Execution Options
Customize the plugin's behavior by using the `--salesforce-apex_opt` flag. Set options as key-value pairs, separated by commas. For example, to use proto field names (snake_case) instead of JSON field names (lowerCamelCase), include the following flag:
```
protoc \
    --proto_path=./examples/proto \
    --salesforce-apex_out=./examples/target/classes \
    --salesforce-apex_opt=useProtoFieldNames=false \
    example.proto
```

### Available Options

Option | Type | Description
-------|------|------------
useProtoFieldNames | bool | Set to true for proto field names (snake_case), or false for JSON field names (lowerCamelCase).
apiVersion | string | Specifies the Salesforce API version to target.
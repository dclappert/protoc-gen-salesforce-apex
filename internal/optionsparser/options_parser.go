package optionsparser

import (
	"fmt"
	"regexp"
	"strings"
)

type Options struct {
	ApiVersion         string
	UseProtoFieldNames bool
}

func ParseOptions(optsStr string) *Options {
	opts := &Options{}
	setDefaultOptions(opts)
	processOptionsString(optsStr, opts)
	return opts
}

func setDefaultOptions(opts *Options) {
	opts.ApiVersion = "56.0"        // Defaults Salesforce API Version 56.0
	opts.UseProtoFieldNames = false // Uses JSON field names (lowerCamelCase by default)
}

func processOptionsString(optsStr string, opts *Options) {
	for _, opt := range strings.Split(optsStr, ",") {
		keyValue := strings.SplitN(opt, "=", 2)
		if len(keyValue) != 2 {
			continue
		}

		key, value := keyValue[0], keyValue[1]

		switch key {
		case "apiVersion":
			if !validateApiVersion(value) {
				panic(fmt.Errorf("Invalid Salesforce API Version format: \"%s\". Specified version must conform to format: \"XX.0\".", value))
			}
			opts.ApiVersion = value
		case "useProtoFieldNames":
			if value == "true" {
				opts.UseProtoFieldNames = true
			} else if value == "false" {
				opts.UseProtoFieldNames = false
			} else {
				panic(fmt.Errorf("invalid value for useProtoFieldNames: %s", value))
			}
		}
	}
}

func validateApiVersion(version string) bool {
	matched, _ := regexp.MatchString(`^\d{2}\.0$`, version)
	return matched
}

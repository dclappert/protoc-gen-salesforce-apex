package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	OUTPUT_DIR = "./data/output"
	PROTOS     = "explorer_note.proto adventure_log.proto"
)

var fileNames = [4]string{
	"AdventureLog.cls",
	"AdventureLog.cls-meta.xml",
	"ExplorerNote.cls",
	"ExplorerNote.cls-meta.xml",
}

func TestMain(m *testing.M) {
	setup(m)
	code := m.Run()
	teardown(m)
	os.Exit(code)
}

func setup(m *testing.M) {
	os.RemoveAll(OUTPUT_DIR)
	err := os.Mkdir(OUTPUT_DIR, 0755)
	if err != nil {
		panic(err)
	}
}

func teardown(m *testing.M) {
	os.RemoveAll(OUTPUT_DIR)
}

func TestGeneratedIsApexClassIsLowerCamelCaseAndMatchesExpected(t *testing.T) {
	cmd := exec.Command(
		"protoc",
		"--proto_path=../examples/proto",
		"--salesforce-apex_out="+OUTPUT_DIR,
		"explorer_note.proto", "adventure_log.proto")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Running CLI failed: %v", err)
	}

	for _, fileName := range fileNames {
		assertFilesAreEqual(
			fmt.Sprintf("./data/expected-output/using-json-field-names/%s", fileName),
			fmt.Sprintf("%s/%s", OUTPUT_DIR, fileName),
			t,
		)
	}
}

func TestGeneratedIsApexClassIsSnakeCaseAndMatchesExpected(t *testing.T) {
	cmd := exec.Command(
		"protoc",
		"--proto_path=../examples/proto",
		"--salesforce-apex_out="+OUTPUT_DIR,
		"--salesforce-apex_opt=useProtoFieldNames=true",
		"explorer_note.proto", "adventure_log.proto")

	if err := cmd.Run(); err != nil {
		t.Fatalf("Running CLI failed: %v", err)
	}

	for _, fileName := range fileNames {
		assertFilesAreEqual(
			fmt.Sprintf("./data/expected-output/using-proto-field-names/%s", fileName),
			fmt.Sprintf("%s/%s", OUTPUT_DIR, fileName),
			t,
		)
	}
}

func assertFilesAreEqual(expectedApexFilePath string, generatedApexFilePath string, t *testing.T) {
	expectedApex, err := os.ReadFile(filepath.Clean(expectedApexFilePath))
	if err != nil {
		println()
		t.Fatalf("Reading expected Apex file failed: %v", err)
	}

	generatedApex, err := os.ReadFile(filepath.Clean(generatedApexFilePath))
	if err != nil {
		t.Fatalf("Reading generated Apex file failed: %v", err)
	}

	if !reflect.DeepEqual(expectedApex, generatedApex) {
		t.Errorf("Generated Apex does not match expected Apex")
	}
}

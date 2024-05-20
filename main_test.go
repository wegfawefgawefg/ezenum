package main

import (
	"os"
	"strings"
	"testing"

	"github.com/wegfawefgawefg/ezenum/generate"
)

func TestMain(t *testing.T) {
	// Define paths
	testFilePath := "testdata/response_codes/test_response_code_ezenum_gen.go"

	// Clean up before and after test
	cleanUp(t, testFilePath)
	defer cleanUp(t, testFilePath)

	// Run the code generator
	generate.Run()

	// Check if the file exists
	if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
		t.Fatalf("expected file %s to be generated, but it does not exist", testFilePath)
	}

	// Read the generated file
	generatedContent, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("failed to read generated file %s: %v", testFilePath, err)
	}

	// Define the expected content
	expectedContent := `package responsecodes

func (r TestResponseCodes) AsCode() int {
	return int(r)
}

func (r TestResponseCodes) GetDescription() string {
	switch r {
	case Ok:
		return "OK: The request was successful, and the response contains the requested information."
	case Continue:
		return "Continue: The client can continue with the request."
	case SwitchingProtocols:
		return "Switching Protocols: The server understands the request and is asking for a protocol switch to proceed."
	default:
		return "Unknown Response"
	}
}

func IsValidTestResponseCodes(code int) bool {
	switch TestResponseCodes(code) {
	case Continue:
	case SwitchingProtocols:
	case Ok:
		return true
	default:
		return false
	}
}
`

	if !isGeneratedContentCorrect(string(generatedContent), expectedContent) {
		t.Errorf("generated content does not match expected content\nExpected:\n%s\nGot:\n%s", expectedContent, generatedContent)
	}
}

func cleanUp(t *testing.T, filePath string) {
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		t.Fatalf("failed to remove file %s: %v", filePath, err)
	}
}

func isGeneratedContentCorrect(generatedContent, expectedContent string) bool {
	expectedLines := strings.Split(expectedContent, "\n")
	generatedLines := strings.Split(generatedContent, "\n")

	expectedSwitchLines := filterSwitchLines(expectedLines)
	generatedSwitchLines := filterSwitchLines(generatedLines)

	if len(expectedSwitchLines) != len(generatedSwitchLines) {
		return false
	}

	expectedSwitchSet := make(map[string]struct{})
	for _, line := range expectedSwitchLines {
		expectedSwitchSet[line] = struct{}{}
	}

	for _, line := range generatedSwitchLines {
		if _, exists := expectedSwitchSet[line]; !exists {
			return false
		}
	}

	return true
}

func filterSwitchLines(lines []string) []string {
	var switchLines []string
	for _, line := range lines {
		if strings.Contains(line, "case ") || strings.Contains(line, "default") {
			switchLines = append(switchLines, strings.TrimSpace(line))
		}
	}
	return switchLines
}

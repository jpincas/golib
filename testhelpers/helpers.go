package testhelpers

import (
	"fmt"
	"testing"
)

func addPrefixToTestName(prefix, name string) string {
	return fmt.Sprintf("%s : %s", prefix, name)
}

func ErrorWithPrefix(t *testing.T, prefix, name string, expected, got interface{}) {
	Error(t, addPrefixToTestName(prefix, name), expected, got)
}

func Error(t *testing.T, name string, expected, got interface{}) {
	t.Errorf("Testing '%s'.\nExpected: %v\nGot: %v\n", name, expected, got)
}

func FatalWithPrefix(t *testing.T, prefix, name string, expected, got interface{}) {
	Fatal(t, addPrefixToTestName(prefix, name), expected, got)
}

func Fatal(t *testing.T, name string, expected, got interface{}) {
	t.Fatalf("Testing '%s'.\nExpected: %v\nGot: %v\n", name, expected, got)
}

package libs

import (
	"testing"
)

func TestLibs(t *testing.T) {
	result := Libs("works")
	if result != "Libs works" {
		t.Error("Expected Libs to append 'works'")
	}
}

package factorio_test

import (
	"github.com/mcandre/factorio"

	"testing"
)

func TestPlatformMarshalingSymmetric(t *testing.T) {
	platformString := "linux/riscv64"
	platform, err := factorio.ParsePlatform(platformString)

	if err != nil {
		t.Error(err)
	}

	platformString2 := platform.String()

	if platformString2 != platformString {
		t.Errorf("expected symmetric marshaling for platform %v", platformString)
	}
}

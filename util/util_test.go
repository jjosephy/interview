package util

import (
	"testing"
)

func Test_Success_GenUUID(t *testing.T) {
	NewUtil()

	s, e := InstanceUtil.NewUUID()
	if e != nil {
		t.Fatalf("failed to generate the UUID %s", e)
	}
	t.Logf("UUID %s", s)
}

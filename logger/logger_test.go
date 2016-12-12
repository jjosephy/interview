package logger

import (
	"testing"
)

func Test_BasicLogging(t *testing.T) {
	l := NewLogger()

	t.Logf("Here %v", l)

}

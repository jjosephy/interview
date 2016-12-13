package repository

import (
	"testing"
)

func Test_Success_CreateErrorModel(t *testing.T) {
}

func Test_Success_MemoryRepository(t *testing.T) {
	n := NewDBInterviewRepository("uri")

	t.Logf("%v", n)
}

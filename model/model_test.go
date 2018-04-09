package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jjosephy/interview/logger"
	"github.com/jjosephy/interview/util"
)

func Test_Success_CreateErrorModel(t *testing.T) {
	m := ErrorModel{
		Message:   "ErrMessage",
		ErrorCode: 1000,
	}

	if m.Message != "ErrMessage" {
		t.Fatal("failed")
	}
}

func Test_Success_SerializeCommentModel(t *testing.T) {
	// Get a model and translate that
	//TODO: make a helper class to create test models and contracts
	uuid, e := util.InstanceUtil.NewUUID()
	if e != nil {
		logger.LogInstance.LogMsg(
			fmt.Sprintf("Error creating uuid in method model_test.Test_Success_SerializeCommentModel %s", e))
	}

	m := InterviewModel{
		Candidate: "Candidate Name",
		ID:        uuid,
		Comments: Comments{
			CommentModel{Content: "db Content", Interviewer: "interviewer 0"},
			CommentModel{Content: "db Content", Interviewer: "interviewer 1"},
			CommentModel{Content: "db Content", Interviewer: "interviewer 2"},
		},
	}

	_, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Failed Marshal model %v\n", err)
	}

	assertAreEqual(t, m.Candidate == "Candidate Name", "Candiate Names dont match")
}

func assertAreEqual(t *testing.T, c bool, msg string) {
	if c != true {
		t.Fatal(msg)
	}
}

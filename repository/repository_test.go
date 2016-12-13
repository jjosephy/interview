package repository

import (
	"testing"

	"github.com/jjosephy/interview/model"
)

var memRepo *MemoryInterviewRepository

const (
	testID = "1e81b35d-b656-45de-8341-886a670d3cf5"
)

func init() {
	memRepo = NewMemoryRepository()
}

func createModel(uuid string) model.InterviewModel {
	var s string
	if len(uuid) == 0 {
		s = testID
	} else {
		s = uuid
	}

	m := model.InterviewModel{
		Candidate: "Test Candidate",
		ID:        s,
		Comments: model.Comments{
			model.CommentModel{Content: "Test Content_0", Interviewer: "interviewer_0"},
			model.CommentModel{Content: "Test Content_1", Interviewer: "interviewer_1"},
			model.CommentModel{Content: "Test Content_2", Interviewer: "interviewer_2"},
		},
	}

	return m
}

func Test_Success_SaveInterview(t *testing.T) {
	mr, e := memRepo.SaveInterview(createModel(""))
	if e != nil {
		t.Fatalf("Error trying to save interview: %s", e)
	}

	t.Logf("Success: %v", mr)
}

func Test_Success_GetInterview(t *testing.T) {
	mr, e := memRepo.GetInterview(testID, "")
	if e != nil {
		t.Fatalf("Error trying to Fet interview: %s", e)
	}

	if mr.ID != testID {
		t.Fatalf("Test Ids do not match. Expected %s . Actual %s", testID, mr.ID)
	}

	t.Logf("Success: %v", mr)
}

func Test_Success_MemoryRepository(t *testing.T) {
	//n := NewDBInterviewRepository("uri")

	//t.Logf("%v", n)
}

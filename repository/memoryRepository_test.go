package repository

import (
	"encoding/json"
	"testing"

	"github.com/jjosephy/interview/model"
	"gopkg.in/mgo.v2/bson"
)

func Test_Success_CreateSimple(t *testing.T) {
	var err error
	var jsondata []byte
	repo := &MemoryInterviewRepository{}
	id := bson.NewObjectId()

	model := model.InterviewModel{
		Candidate: "Candidate Name",
		ID:        id,
		Comments: model.Comments{
			model.CommentModel{Content: "db Content", Interviewer: "interviewer 0"},
			model.CommentModel{Content: "db Content", Interviewer: "interviewer 1"},
			model.CommentModel{Content: "db Content", Interviewer: "interviewer 2"},
		},
	}

	model, err = repo.SaveInterview(model)
	if err != nil {
		t.Fatal(err.Error())
	}

	model, err = repo.GetInterview(string(id), model.Candidate)
	if err != nil {
		t.Fatal(err.Error())
	}

	jsondata, err = json.Marshal(model)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(string(jsondata))
}

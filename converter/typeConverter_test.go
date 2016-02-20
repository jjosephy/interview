package converter

import (
    "bytes"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "github.com/jjosephy/go/interview/contract/v1"
    "github.com/jjosephy/go/interview/model"
    "io"
    "testing"
)

type nopCloser struct {
    io.Reader
}

func (_ nopCloser) Close() (error) {
    return nil
}

func CreateTestModel() (model.InterviewModel) {
    return model.InterviewModel {
        Candidate: "Candidate Name",
        Id: bson.NewObjectId(),
        Comments: model.Comments {
            model.CommentModel { Content: "db Content", Interviewer: "interviewer 0" },
            model.CommentModel { Content: "db Content", Interviewer: "interviewer 1" },
            model.CommentModel { Content: "db Content", Interviewer: "interviewer 2" },
        },
    }
}

func Test_Success_DecodeContractFromBodyV1(t *testing.T) {
    c := contract.InterviewContractV1 {
        Candidate: "Candidate",
        Id: "2",
        Comments: contract.CommentsV1 {
            contract.CommentContractV1 { Content: "db Content", Interviewer: "interviewer 0", },
            contract.CommentContractV1 { Content: "db Content", Interviewer: "interviewer 1", },
            contract.CommentContractV1 { Content: "db Content", Interviewer: "interviewer 2", },
        },
    }

    s, err := json.Marshal(c)
    if err != nil {
        t.Fatalf("Marshing object failed %v \n", err)
    }

    str := string(s)
    b := nopCloser { bytes.NewBufferString(str) }

    cx, err := DecodeContractFromBodyV1(b)
    if err != nil {
        t.Fatalf("Failed to decode %v \n", err)
    }

    t.Log("Here", cx)
}

func Test_Sucess_ConvertModelToContractV1(t *testing.T) {
    m := CreateTestModel()
    c := ConvertModelToContractV1(m)
    t.Logf("c: %v", c)
}

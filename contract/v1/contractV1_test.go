package contract

import (
    "encoding/json"
    "testing"
)

func IsTrue(t *testing.T, c bool, msg string) {
    if c != true {
        t.Fatal(msg)
    }
}

func CreateTestContract()(InterviewContractV1) {
    return InterviewContractV1 {
        Candidate: "Candidate Name",
        Id: "2",
        Comments: CommentsV1 {
            CommentContractV1 { Content: "db Content", Interviewer: "interviewer 0" },
            CommentContractV1 { Content: "db Content", Interviewer: "interviewer 1" },
            CommentContractV1 { Content: "db Content", Interviewer: "interviewer 2" },
        },
    }
}

func Test_Success_CreateCommentContractV1(t *testing.T) {
    c := CommentContractV1 {
        Content: "This is a comment",
        Interviewer: "Interviewer",
    }

    IsTrue(t, c.Content == "This is a comment", "Comment Content does not match")
    IsTrue(t, c.Interviewer == "Interviewer", "Interviewer does not match")
}

func Test_Success_MarshalContractV1(t *testing.T) {
    c := CreateTestContract()
    s, err := json.Marshal(c)
    if err != nil {
        t.Fatalf("Failed Marshal model %v\n", err)
    }

    str := string(s)
    t.Log(str)
}

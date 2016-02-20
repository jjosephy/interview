package converter

import (
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "github.com/jjosephy/go/interview/contract/v1"
    "github.com/jjosephy/go/interview/model"
    "io"
)

func DecodeContractFromBodyV1(r io.ReadCloser) (contract.InterviewContractV1, error) {
    decoder := json.NewDecoder(r)
    var c contract.InterviewContractV1
    err := decoder.Decode(&c)

    if err != nil {
        return c, err
    }

    return c, nil
}
func ConvertModelToContractV1 (m model.InterviewModel) (c contract.InterviewContractV1) {
    // TODO: validate input

    comments := make([]contract.CommentContractV1, len(m.Comments))

    for i := 0; i < len(m.Comments); i++ {
        cm := contract.CommentContractV1  {
                Content: m.Comments[i].Content,
                Interviewer: m.Comments[i].Interviewer }
        comments[i] = cm
    }

    // TODO: validate success
    return contract.InterviewContractV1 {
        Id: m.Id.Hex(),
        Candidate: m.Candidate,
        Comments: comments,
    }
}

func ConvertContractToModelV1 (c contract.InterviewContractV1) (m model.InterviewModel) {
    // TODO: validate input

    comments := make([]model.CommentModel, len(c.Comments))

    for i := 0; i < len(c.Comments); i++ {
        mc := model.CommentModel  {
                Content: c.Comments[i].Content,
                Interviewer: c.Comments[i].Interviewer }
        comments[i] = mc
    }

    // TODO: validate success
    oid := bson.ObjectId(c.Id)
    return model.InterviewModel {
        Id: oid,
        Candidate: c.Candidate,
        Comments: comments,
    }
}

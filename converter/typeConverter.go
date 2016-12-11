package converter

import (
	"encoding/json"
	"io"

	"github.com/jjosephy/interview/contract/v1"
	"github.com/jjosephy/interview/model"
	"gopkg.in/mgo.v2/bson"
)

// DecodeContractFromBodyV1 decodes a Contract from a Request Body
func DecodeContractFromBodyV1(r io.ReadCloser) (contract.InterviewContractV1, error) {
	decoder := json.NewDecoder(r)
	var c contract.InterviewContractV1
	err := decoder.Decode(&c)

	if err != nil {
		return c, err
	}

	return c, nil
}

// ConvertModelToContractV1 creates a V1 Interview Contract from an Interview Model
func ConvertModelToContractV1(m model.InterviewModel) (c contract.InterviewContractV1) {
	// TODO: validate input

	comments := make([]contract.CommentContractV1, len(m.Comments))

	for i := 0; i < len(m.Comments); i++ {
		cm := contract.CommentContractV1{
			Content:     m.Comments[i].Content,
			Interviewer: m.Comments[i].Interviewer}
		comments[i] = cm
	}

	// TODO: validate success
	return contract.InterviewContractV1{
		ID:        m.ID.Hex(),
		Candidate: m.Candidate,
		Comments:  comments,
	}
}

// ConvertContractToModelV1 creates a interview model from a V1 contract
func ConvertContractToModelV1(c contract.InterviewContractV1) (m model.InterviewModel) {
	// TODO: validate input

	comments := make([]model.CommentModel, len(c.Comments))

	for i := 0; i < len(c.Comments); i++ {
		mc := model.CommentModel{
			Content:     c.Comments[i].Content,
			Interviewer: c.Comments[i].Interviewer}
		comments[i] = mc
	}

	// TODO: validate success
	oid := bson.ObjectId(c.ID)
	return model.InterviewModel{
		ID:        oid,
		Candidate: c.Candidate,
		Comments:  comments,
	}
}

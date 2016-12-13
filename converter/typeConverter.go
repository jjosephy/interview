package converter

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jjosephy/interview/contract/v1"
	"github.com/jjosephy/interview/logger"
	"github.com/jjosephy/interview/model"
	"github.com/jjosephy/interview/util"
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

	return contract.InterviewContractV1{
		ID:        m.ID,
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

	uuid, e := util.InstanceUtil.NewUUID()
	if e != nil {
		logger.LogInstance.LogMsg(
			fmt.Sprintf("Error creating uuid in method typeConverter.ConvertContractToModelV1 %s", e))
	}

	return model.InterviewModel{
		ID:        uuid,
		Candidate: c.Candidate,
		Comments:  comments,
	}
}

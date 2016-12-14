package repository

import (
	"fmt"
	"os"

	"github.com/goamz/goamz/aws"
	"github.com/jjosephy/interview/logger"
	"github.com/jjosephy/interview/model"
	"github.com/jjosephy/interview/util"
)

// InterviewRepository Interface
type InterviewRepository interface {
	GetInterview(id string, name string) (model.InterviewModel, error)
	SaveInterview(model model.InterviewModel) (model.InterviewModel, error)
	CreateUser(m *model.CreateUserModel) (string, error)
}

// DBInterviewRepository Implementation
type DBInterviewRepository struct {
	URI string
}

// NewDBInterviewRepository creates a new instance of DBInterviewRepository
func NewDBInterviewRepository(uri string) *DBInterviewRepository {
	os.Setenv("AWS_ACCESS_KEY_ID", "12312312")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "12312312")

	auth, err := aws.EnvAuth()

	if err != nil {
		fmt.Println("error ", err)
	} else {
		//AccessKey, SecretKey
		fmt.Println("auth success", auth.AccessKey, auth.SecretKey)
	}

	return &DBInterviewRepository{URI: uri}
}

// SaveInterview impl
func (r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {
	return *(r.createModel()), nil
}

// GetInterview impl
func (r *DBInterviewRepository) GetInterview(id string, name string) (m model.InterviewModel, err error) {
	return *(r.createModel()), nil
}

// CreateUser Implements the InterviwRepository Interface
func (r *DBInterviewRepository) CreateUser(m *model.CreateUserModel) (string, error) {

	return "", nil
}

func (r *DBInterviewRepository) createModel() *model.InterviewModel {

	comments := make([]model.CommentModel, 3)
	for i := 0; i < 3; i++ {
		mc := model.CommentModel{
			Content:     "Content",
			Interviewer: "Interviewer"}
		comments[i] = mc
	}

	uuid, e := util.InstanceUtil.NewUUID()
	if e != nil {
		logger.LogInstance.LogMsg(
			fmt.Sprintf("Error creating uuid in method interviewRepository.createModel %s", e))
	}

	return &model.InterviewModel{
		Candidate: "candidate",
		Complete:  false,
		ID:        uuid,
		Comments:  comments,
	}
}

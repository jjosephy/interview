package repository

import (
	"errors"

	"github.com/jjosephy/interview/logger"
	"github.com/jjosephy/interview/model"
	"github.com/jjosephy/interview/util"
)

// MemoryInterviewRepository Implementation
type MemoryInterviewRepository struct {
	m map[string]interface{}
	u map[string]interface{}
}

// NewMemoryRepository creates a new Memory Repositoryu
func NewMemoryRepository() *MemoryInterviewRepository {
	r := &MemoryInterviewRepository{}
	r.m = make(map[string]interface{}, 0)
	r.u = make(map[string]interface{}, 0)

	return r
}

// CreateUser Implements the InterviwRepository Interface
func (r *MemoryInterviewRepository) CreateUser(m *model.CreateUserModel) (string, error) {
	var (
		err  error
		uuid string
	)

	/*
	   mt, ok := r.m[id].(model.InterviewModel)
	   		if !ok {
	   			msg := "interview does not exist by id. repository.GetInterview"
	   			logger.LogInstance.LogError(msg)
	   			return mi, errors.New("not found")
	   		}
	*/

	//TODO: check for existing User
	/*
		for k := range r.u {
			c := r.u[k].(model.CreateUserModel)
			if c.UserName == m.UserName {
				return "", errors.New("User already exists")
			}
		}
	*/

	if uuid, err = util.InstanceUtil.NewUUID(); err != nil {
		return "", err
	}

	r.u[uuid] = m
	return uuid, nil
}

// SaveInterview saves an interview. Implements the Repository interface.
func (r *MemoryInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {
	mi = model.InterviewModel{}

	if r.m == nil {
		msg := "invalid memory, nil repo. repository.SaveInterview"
		// TODO: log error
		logger.LogInstance.LogError(msg)
		return mi, errors.New("msg")
	}

	r.m[m.ID] = m
	return m, nil
}

//GetInterview gets an interview. Implements the Repository interface.
func (r *MemoryInterviewRepository) GetInterview(id string, name string) (mi model.InterviewModel, err error) {
	// test for interview that doesnt exist
	mi = model.InterviewModel{}

	if len(id) == 0 && len(name) == 0 {
		// TODO: log error
		msg := "invalid search query. repository.GetInterview"
		logger.LogInstance.LogError(msg)
		return mi, errors.New(msg)
	}

	if len(id) > 0 {
		mt, ok := r.m[id].(model.InterviewModel)
		if !ok {
			msg := "interview does not exist by id. repository.GetInterview"
			logger.LogInstance.LogError(msg)
			return mi, errors.New("not found")
		}

		mi = mt
	} else if len(name) > 0 {
		for k := range r.m {
			m := r.m[k].(model.InterviewModel)
			if mi.Candidate == name {
				return m, nil
			}
		}
	}

	return mi, nil
}

package repository

import (
	"errors"

	"github.com/jjosephy/interview/model"
)

// MemoryInterviewRepository - InterviewRepository Implementation
type MemoryInterviewRepository struct {
	m map[string]interface{}
}

// NewMemoryRepository creates a new instance of MemoryRepository
func NewMemoryRepository() *MemoryInterviewRepository {
	r := &MemoryInterviewRepository{}
	r.m = make(map[string]interface{}, 0)
	return r
}

// SaveInterview implementation of InterviewRespository Interface
func (r *MemoryInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {

	if &m == nil {
		return m, errors.New("invalid argument {model}")
	}

	// serialize the instance to json
	// store the stirng by id

	mi = m
	return mi, nil
}

// GetInterview implementation of InterviewRespository Interface
func (r *MemoryInterviewRepository) GetInterview(id string, name string) (mi model.InterviewModel, err error) {
	mi = model.InterviewModel{}
	return mi, nil
}

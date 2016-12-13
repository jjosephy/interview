package repository

import (
	"github.com/jjosephy/interview/model"
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

// SaveInterview saves an interview. Implements the Repository interface.
func (r *MemoryInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {

	if r.m == nil {

	}

	mi = model.InterviewModel{}
	return mi, nil
}

//GetInterview gets an interview. Implements the Repository interface.
func (r *MemoryInterviewRepository) GetInterview(id string, name string) (mi model.InterviewModel, err error) {
	mi = model.InterviewModel{}
	return mi, nil
}

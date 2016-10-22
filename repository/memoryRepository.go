package repository

import (
    "github.com/jjosephy/interview/model"
)

// DBInterviewRepository Implementation
type MemoryInterviewRepository struct {
	URI string
    m map[string]interface{}
}

func (r *MemoryInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {
    mi = model.InterviewModel{}
    return mi, nil
}

func (r *MemoryInterviewRepository) GetInterview(id string, name string) (mi model.InterviewModel, err error) {
    mi = model.InterviewModel{}
    return mi, nil
}

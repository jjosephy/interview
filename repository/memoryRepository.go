package repository

import (
    "github.com/jjosephy/interview/model"
)

// DBInterviewRepository Implementation
type MemoryInterviewRepository struct {
    m map[string]interface{}
}

func NewMemoryRepository() *MemoryInterviewRepository {
    r := &MemoryInterviewRepository {}
    r.m = make(map[string]interface{}, 0)

    return r
}

func (r *MemoryInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {

    if r.m == nil {

    }

    mi = model.InterviewModel{}
    return mi, nil
}

func (r *MemoryInterviewRepository) GetInterview(id string, name string) (mi model.InterviewModel, err error) {
    mi = model.InterviewModel{}
    return mi, nil
}

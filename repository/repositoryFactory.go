package repository

import ()

func NewRepository(s string) InterviewRepository {
    if s == "MemoryRepository" {
        return NewMemoryRepository()
    }

    if s == "DBInterviewRepository" {
        return NewDBInterviewRepository("uri")
    }

    return nil
}

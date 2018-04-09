package repository

// NewRepository is a simple factory for creating types for InterviewRepository interface.
func NewRepository(s string) InterviewRepository {
	if s == "MemoryRepository" {
		return NewMemoryRepository()
	}

	if s == "DBInterviewRepository" {
		return NewDBInterviewRepository("uri")
	}

	return nil
}

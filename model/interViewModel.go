package model

// InterviewModel is the internal type to hold Interview Data
type InterviewModel struct {
	Candidate string
	Comments  []CommentModel
	Complete  bool
	ID        string
}

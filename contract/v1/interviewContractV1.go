package contract

// InterviewContractV1 impl
type InterviewContractV1 struct {
	Candidate string              `json:"candidate"`
	Comments  []CommentContractV1 `json:"comments"`
	Complete  bool                `json:"complete"`
	ID        string              `json:"id"`
}

// InterviewsV1 impl
type InterviewsV1 []InterviewContractV1

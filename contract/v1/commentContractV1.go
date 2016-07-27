package contract

// CommentContractV1 impl
type CommentContractV1 struct {
	Content     string `json:"content"`
	Interviewer string `json:"interviewer"`
}

// CommentsV1 impl
type CommentsV1 []CommentContractV1

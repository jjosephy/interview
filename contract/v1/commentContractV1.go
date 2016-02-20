package contract

type CommentContractV1  struct {
    Content             string      `json:"content"`
    Interviewer         string      `json:"interviewer"`
}

type CommentsV1 []CommentContractV1

package model

// CommentModel is the internal type to hold model data
type CommentModel struct {
	Content     string
	Interviewer string
}

// Comments is an array of CommentModel objects
type Comments []CommentModel

package model

import (
	"gopkg.in/mgo.v2/bson"
)

// InterviewModel is the internal type to hold Interview Data
type InterviewModel struct {
	Candidate string
	Comments  []CommentModel
	Complete  bool
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

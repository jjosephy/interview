package model

import (
    "gopkg.in/mgo.v2/bson"
)

type InterviewModel struct {
    Candidate           string
    Comments            []CommentModel
    Complete            bool
    Id                  bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

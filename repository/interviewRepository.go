package repository

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/jjosephy/interview/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InterviewRepository Interface
type InterviewRepository interface {
	GetInterview(id string, name string) (model.InterviewModel, error)
	SaveInterview(model model.InterviewModel) (model.InterviewModel, error)
}

// DBInterviewRepository Implementation
type DBInterviewRepository struct {
	URI string
}

// New creates a new instance of DBInterviewRepository
func New(uri string) *DBInterviewRepository {
	os.Setenv("AWS_ACCESS_KEY_ID", "12312312")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "12312312")

	auth, err := aws.EnvAuth()

	if err != nil {
		fmt.Println("error ", err)
	} else {
		//AccessKey, SecretKey
		fmt.Println("auth success", auth.AccessKey, auth.SecretKey)
	}

	return &DBInterviewRepository{URI: uri}
}

func (r *DBInterviewRepository) getConnection() (s *mgo.Session, err error) {

	if r.URI == "" {
		return nil, errors.New("address for database not configured")
	}

	timeout := 5 * time.Second
	s, err = mgo.DialWithTimeout(r.URI, timeout)
	if err != nil {
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)
	return s, nil
}

// SaveInterview impl
func (r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {
	/*
		var s *mgo.Session
		defer func() {
			if s != nil {
				s.Close()
			}
		}()

		if s, err = r.getConnection(); err != nil {
			return m, err
		}

		m.ID = bson.NewObjectId()
		if err = s.DB("interview").C("interviews").Insert(&m); err != nil {
			return m, err
		}

		defer s.Close()
	*/

	return *(r.createModel()), nil
}

// GetInterview impl
func (r *DBInterviewRepository) GetInterview(id string, name string) (m model.InterviewModel, err error) {

	/*
		var s *mgo.Session
		defer func() {
			if s != nil {
				s.Close()
			}
		}()

		if s, err = r.getConnection(); err != nil {
			return m, err
		}

		if id == "" && name == "" {
			return m, errors.New("invalid search params provided")
		}

		if valid := bson.IsObjectIdHex(id); valid == false {
			return m, errors.New("HexId")
		}

		// TODO: find by candidate name
		m = model.InterviewModel{}
		bid := bson.ObjectIdHex(id)
		err = s.DB("interview").C("interviews").FindId(bid).One(&m)

		if err != nil {
			return m, err
		}
	*/

	// Broken deref
	return *(r.createModel()), nil
}

func (r *DBInterviewRepository) createModel() *model.InterviewModel {

	comments := make([]model.CommentModel, 3)
	for i := 0; i < 3; i++ {
		mc := model.CommentModel{
			Content:     "Content",
			Interviewer: "Interviewer"}
		comments[i] = mc
	}

	return &model.InterviewModel{
		Candidate: "candidate",
		Complete:  false,
		ID:        bson.ObjectId("7DF9B7CD-4A30-4265-86FD-440A3B037976"),
		Comments:  comments,
	}
}

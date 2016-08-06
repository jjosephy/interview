package repository

import (
	"errors"
	"time"

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
	return m, nil
}

// GetInterview impl
func (r *DBInterviewRepository) GetInterview(id string, name string) (m model.InterviewModel, err error) {
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

	return m, nil
}

package repository

import (
    "errors"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/jjosephy/go/interview/model"
    "time"
)

// Interview Repository Interface //
type InterviewRepository interface {
    GetInterview(id string, name string) (model.InterviewModel, error)
    SaveInterview(model model.InterviewModel) (model.InterviewModel, error)
}

// DBInterviewRespository Implementation //
type DBInterviewRepository struct {
    Uri string
}

func (r *DBInterviewRepository) getConnection()(s *mgo.Session, err error) {

    if r.Uri == "" {
        return nil, errors.New("address for database not configured")
    }

    timeout := 5 * time.Second
    s, err = mgo.DialWithTimeout(r.Uri, timeout)
    if err != nil {
        return nil, err
    }

    s.SetMode(mgo.Monotonic, true)
    return  s, nil

    // TODO: Unique index on _id
    /*
    r.DBSession.SetMode(mgo.Monotonic, true)
    index := mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	err = r.DBSession.DB("interview").C("interviews").EnsureIndex(index)
    if err != nil {
        return err
    }
    */
}

func(r *DBInterviewRepository) SaveInterview(m model.InterviewModel) (mi model.InterviewModel, err error) {
    var s *mgo.Session = nil
    defer func() {
        if s != nil {
            s.Close()
        }
    }()

    if s, err = r.getConnection(); err != nil {
        return m, err
    }

    m.Id = bson.NewObjectId()
    if err = s.DB("interview").C("interviews").Insert(&m); err != nil {
        return m, err
    }

    defer s.Close()
    return m, nil
}

func(r *DBInterviewRepository) GetInterview(id string, name string) (m model.InterviewModel, err error) {
    var s *mgo.Session = nil
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

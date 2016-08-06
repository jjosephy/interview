package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jjosephy/interview/converter"
	"github.com/jjosephy/interview/httperror"
	"github.com/jjosephy/interview/model"
	"github.com/jjosephy/interview/repository"
)

// InterviewHandler is the handler for public Interview API
func InterviewHandler(data repository.InterviewRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var version float64
		h := r.Header.Get("api-version")
		if h == "" {
			httperror.NoVersionProvided(w)
			return
		}

		v, err := strconv.ParseFloat(h, 64)
		if err != nil {
			httperror.InvalidVersionProvided(w)
			return
		}
		version = v

		// Switch on Request Method
		switch r.Method {
		case "GET":
			// Get Id or Name for Search
			qID := r.URL.Query()["id"]
			qName := r.URL.Query()["cname"]

			//TODO: scrub input
			if qID == nil && qName == nil {
				httperror.NoQueryParametersProvided(w)
				return
			}

			var id string
			if len(qID) > 0 {
				id = qID[0]
			} else {
				id = ""
			}

			var name string
			if len(qName) > 0 {
				name = qName[0]
			} else {
				name = ""
			}

			// TODO: Find by id or name
			model, err := data.GetInterview(id, name)
			if err != nil {
				switch err.Error() {
				case "not found":
					httperror.InterviewNotFound(w)
					return
				case "HexId":
					httperror.InvalidInterviewId(w)
					return
				default:
					httperror.GetInterviewFailed(w, err)
					return
				}
			}

			switch version {
			case 1.0:
				json.NewEncoder(w).Encode(converter.ConvertModelToContractV1(model))
			default:
				httperror.UnsupportedVersion(w)
				return
			}

		case "POST":
			// check to see if body is null
			var m model.InterviewModel
			switch version {
			case 1.0:
				c, err := converter.DecodeContractFromBodyV1(r.Body)
				if err != nil {
					httperror.FailedDecodingBody(w)
					return
				}
				m = converter.ConvertContractToModelV1(c)
			default:
				httperror.UnsupportedVersion(w)
				return
			}

			m, err := data.SaveInterview(m)
			if err != nil {
				httperror.SaveInterviewFailed(w, err)
				return
			}

			json.NewEncoder(w).Encode(converter.ConvertModelToContractV1(m))

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/contract/v1"
	"github.com/jjosephy/interview/converter"
	"github.com/jjosephy/interview/httperror"
	"github.com/jjosephy/interview/logger"
	"github.com/jjosephy/interview/repository"
)

// UserHandler is the handler for public Interview API
func UserHandler(data repository.InterviewRepository, p authentication.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			version float64
			err     error
		)

		logger.LogInstance.LogMsg("User Handler called")

		// Validate Authorization Header
		h := r.Header.Get("authorization")
		if h == "" {
			httperror.Unauthorized(w)
			return
		}

		res, tErr := p.ValidateToken(h)
		if res == false || tErr != nil {
			httperror.Unauthorized(w)
			return
		}

		// Validate Version Header
		h = r.Header.Get("api-version")
		if h == "" {
			httperror.NoVersionProvided(w)
			return
		}

		version, err = strconv.ParseFloat(h, 64)
		if err != nil {
			httperror.InvalidVersionProvided(w)
			return
		}

		// Switch on Request Method
		switch r.Method {
		case "POST":
			if version == 1.0 {
				decoder := json.NewDecoder(r.Body)
				var c contract.CreateUserContractV1
				err := decoder.Decode(&c)

				if err != nil {
					httperror.GeneralServerError(w, "cant deserialize user")
				}

				m := converter.ConvertCreateUserContractToModelV1(c)
				s, e := data.CreateUser(m)
				if e != nil {
					httperror.GeneralServerError(w, "Create user failed")
				} else {
					w.Write([]byte(s))
				}

			} else {
				httperror.InvalidVersionProvided(w)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/httperror"
)

// TokenHandler is the handler for getting new tokens
func TokenHandler(p authentication.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate the version
		switch r.Method {
		case "POST":
			b, err := ioutil.ReadAll(r.Body)
			if len(b) == 0 || err != nil {
				httperror.NoRequestBody(w)
				return
			}

			parts := strings.Split(string(b), "&")
			if len(parts) != 2 {
				httperror.InvalidRequestBody(w)
				return
			}

			uname := strings.Split(parts[0], "=")
			if len(uname) != 2 {
				httperror.InvalidRequestBody(w)
				return
			}

			pwd := strings.Split(parts[1], "=")
			if len(pwd) != 2 {
				httperror.InvalidRequestBody(w)
				return
			}

			token, err := p.AuthenticateUser(uname[1], pwd[1])
			if err != nil {
				httperror.Unauthorized(w)
			}

			w.Write([]byte(token))

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

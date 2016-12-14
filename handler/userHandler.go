package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jjosephy/interview/authentication"
	"github.com/jjosephy/interview/httperror"
	"github.com/jjosephy/interview/logger"
)

// UserHandler is the handler for public Interview API
func UserHandler(p authentication.Provider) http.HandlerFunc {
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
				fmt.Printf("POST to UserHander %f", version)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

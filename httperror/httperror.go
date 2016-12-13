package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jjosephy/interview/model"
)

const (
	// MsgNoVersionProvided message used when no version is provided
	MsgNoVersionProvided = "No Version Provided"
	//MsgInvalidVersion message
	MsgInvalidVersion = "Invalid Version Provided"
	// MsgNoParametersProvided message
	MsgNoParametersProvided = "No Parameters Provided"
	//MsgUnsupportedVersion message
	MsgUnsupportedVersion = "Unsupported Version Provided"
	// MsgGetInterviewError message
	MsgGetInterviewError = "Error occurred tyring to Get interview details"
	//MsgSaveInterviewError message
	MsgSaveInterviewError = "Error occurred trying to Save the interview"
	// MsgFailedReadingBody message
	MsgFailedReadingBody = "Failed to parse the request body"
	// MsgInterviewNotFound message
	MsgInterviewNotFound = "Interview not found"
	// MsgInterviewInvalid message
	MsgInterviewInvalid = "Invalid Interview Id"
	// MsgUnauthorized message
	MsgUnauthorized = "Unauthorized"
	// MsgNoRequestBody message
	MsgNoRequestBody = "No Request Body Provided"
	//MsgInvalidRequestBody message
	MsgInvalidRequestBody = "Invalid Request Body Provided"
)

const (
	// BadRequestNoInputParameters error
	BadRequestNoInputParameters = 3000
	// BadRequestNoVersion error
	BadRequestNoVersion = 3001
	// BadRequestInvalidVersion error
	BadRequestInvalidVersion = 3002
	// BadRequestUnSupportedVersion error
	BadRequestUnSupportedVersion = 3003
	// BadRequestFailedDecodingRequestBody error
	BadRequestFailedDecodingRequestBody = 3004
	// BadRequestInterviewInvalid error
	BadRequestInterviewInvalid = 3005
	// BadRequestNoRequestBody error
	BadRequestNoRequestBody = 3006
	// BadRequestInvalidRequestBody error
	BadRequestInvalidRequestBody = 3007

	// NotFoundInterviewNotFound error
	NotFoundInterviewNotFound = 4000

	// NoAuthorizedLdapConnectFailure error
	NoAuthorizedLdapConnectFailure = 6000
	// NoAuthorizedFailure error
	NoAuthorizedFailure = 6001

	// ServerErrorGeneral error
	ServerErrorGeneral = 5000
	// ServerErrorGetInterviewFailure error
	ServerErrorGetInterviewFailure = 5001
	// ServerErrorSaveInterviewFailure error
	ServerErrorSaveInterviewFailure = 5002
)

// InvalidRequestBody is used when an invalid request body is passed
func InvalidRequestBody(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestInvalidRequestBody, MsgNoRequestBody)
}

// NoRequestBody is used when no requst body is passed
func NoRequestBody(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestNoRequestBody, MsgNoRequestBody)
}

// Unauthorized is used when the request is unathorized
func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: NoAuthorizedFailure, Message: MsgUnauthorized})
}

// AuthConnectToLDAPFailure is used when an LDAP failure occurs
func AuthConnectToLDAPFailure(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: NoAuthorizedLdapConnectFailure, Message: MsgUnauthorized})
}

// GeneralServerError is used when there is a general server error
func GeneralServerError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: ServerErrorGeneral, Message: msg})
}

// GetInterviewFailed is used when an general error occurs
func GetInterviewFailed(w http.ResponseWriter, e error) {
	s := fmt.Sprint(MsgGetInterviewError, " : ", e)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: ServerErrorGetInterviewFailure, Message: s})
}

// SaveInterviewFailed is used when saving an interview fails
func SaveInterviewFailed(w http.ResponseWriter, e error) {
	s := fmt.Sprint(MsgSaveInterviewError, " : ", e)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: ServerErrorSaveInterviewFailure, Message: s})
}

// InterviewNotFound is used when an interview is not found
func InterviewNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: NotFoundInterviewNotFound, Message: MsgInterviewNotFound})
}

// InvalidInterviewID is used when an invalid interview id is passed
func InvalidInterviewID(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: BadRequestInterviewInvalid, Message: MsgInterviewInvalid})
}

// NoVersionProvided is used when no version header is passed
func NoVersionProvided(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestNoVersion, MsgNoVersionProvided)
}

// InvalidVersionProvided is used when an Invalid version is passed
func InvalidVersionProvided(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestInvalidVersion, MsgInvalidVersion)
}

// NoQueryParametersProvided is used when no query parameter is passed
func NoQueryParametersProvided(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestNoInputParameters, MsgNoParametersProvided)
}

// UnsupportedVersion is used when an unsupported version is passed
func UnsupportedVersion(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestUnSupportedVersion, MsgUnsupportedVersion)
}

// FailedDecodingBody is used when decoding a message body fails
func FailedDecodingBody(w http.ResponseWriter) {
	writeBadRequest(w, BadRequestFailedDecodingRequestBody, MsgFailedReadingBody)
}

func writeBadRequest(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.ErrorModel{ErrorCode: code, Message: msg})
}

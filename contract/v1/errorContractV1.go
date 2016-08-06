package contract

// ErrorContractV1 impl
type ErrorContractV1 struct {
	Message string `json:"message"`
	Code    int    `json:"errorCode"`
}

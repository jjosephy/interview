package contract

// CreateUserContractV1 type for creating users
type CreateUserContractV1 struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

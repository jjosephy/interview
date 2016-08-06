package authentication

// Provider is the Public interface to an authentication Provider type that can be used to implemenent authentication
type Provider interface {
	AuthenticateUser(name string, pwd string) (string, error)
	ValidateToken(token string) (bool, error)
}

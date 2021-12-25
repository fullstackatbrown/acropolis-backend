package auth

// AccessToken is a struct that represents an access token being exchanged for a session cookie.
type AccessToken struct {
	Token string `json:"token"`
}

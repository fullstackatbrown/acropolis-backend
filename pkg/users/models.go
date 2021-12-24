package users

// UserInfo is a collection of standard profile information for a user.
type UserInfo struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhotoURL    string `json:"photoUrl"`
	ID          string `json:"id"`
}

// UserRecord contains metadata associated with a Firebase user account.
type UserRecord struct {
	*UserInfo
	Disabled           bool
	EmailVerified      bool
	CreationTimestamp  int64
	LastLogInTimestamp int64
	// The time at which the user was last active (ID token refreshed), or 0 if
	// the user was never active.
	LastRefreshTimestamp int64
}

// UserToCreate is the parameter struct for the CreateUser function.
type UserToCreate struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

// AccessToken is a struct that represents an access token being exchanged for a session cookie.
type AccessToken struct {
	Token string `json:"token"`
}

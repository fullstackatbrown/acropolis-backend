package users

import (
	"acropolis-backend/pkg/fb"
	firebaseAuth "firebase.google.com/go/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"strings"
)

// authClient is a global variable to hold the initialized Firebase Auth client
var authClient *firebaseAuth.Client

// GetAllUsers gets the user data for all registered users.
func GetAllUsers() ([]*UserRecord, error) {
	var users []*UserRecord
	iter := authClient.Users(fb.FirebaseContext, "")
	for {
		fbUser, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error listing users: %s\n", err)
		}
		user := fbUserToUserRecord(fbUser.UserRecord)

		users = append(users, user)
	}

	return users, nil
}

// GetUser gets the user data corresponding to the specified user ID.
func GetUser(id string) (*UserRecord, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}

	fbUser, err := authClient.GetUser(fb.FirebaseContext, id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v\n", err)
	}

	return fbUserToUserRecord(fbUser), nil
}

// GetUserByEmail gets the user data corresponding to the specified email.
func GetUserByEmail(email string) (*UserRecord, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	fbUser, err := authClient.GetUserByEmail(fb.FirebaseContext, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v\n", err)
	}

	return fbUserToUserRecord(fbUser), nil
}

// CreateUser creates a new user with the specified properties.
func CreateUser(user *UserToCreate) (*UserRecord, error) {
	if err := user.validate(); err != nil {
		return nil, err
	}

	// Create a user in Firebase Auth
	u := (&firebaseAuth.UserToCreate{}).Email(user.Email).Password(user.Password).DisplayName(user.DisplayName)
	fbUser, err := authClient.CreateUser(fb.FirebaseContext, u)
	if err != nil {
		return nil, err
	}

	return fbUserToUserRecord(fbUser), nil
}

// VerifySessionCookie verifies that the given session cookie is valid and returns the associated UserRecord if valid.
func VerifySessionCookie(sessionCookie *http.Cookie) (*UserRecord, error) {
	decoded, err := authClient.VerifySessionCookieAndCheckRevoked(fb.FirebaseContext, sessionCookie.Value)
	if err != nil {
		return nil, fmt.Errorf("error verifying cookie: %v\n", err)
	}

	user, err := GetUser(decoded.UID)
	if err != nil {
		return nil, fmt.Errorf("error getting user from cookie: %v\n", err)
	}

	return user, nil
}

// fbUserToUserRecord converts a Firebase ExportedUserRecord into a UserRecord
func fbUserToUserRecord(fbUser *firebaseAuth.UserRecord) *UserRecord {
	return &UserRecord{
		UserInfo: &UserInfo{
			DisplayName: fbUser.DisplayName,
			Email:       fbUser.Email,
			PhoneNumber: fbUser.PhoneNumber,
			PhotoURL:    fbUser.PhotoURL,
			ID:          fbUser.UID,
		},
		Disabled:             fbUser.Disabled,
		EmailVerified:        fbUser.EmailVerified,
		CreationTimestamp:    fbUser.UserMetadata.CreationTimestamp,
		LastLogInTimestamp:   fbUser.UserMetadata.LastLogInTimestamp,
		LastRefreshTimestamp: fbUser.UserMetadata.LastRefreshTimestamp,
	}
}

// GetUserFromContext gets a UserRecord, if it exists, from an Echo context.
func GetUserFromContext(c echo.Context) *UserRecord {
	user := c.Get("user")
	switch user.(type) {
	case *UserRecord:
		return user.(*UserRecord)
	default:
		return nil
	}
}

// Validators.

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email must be a non-empty string")
	}
	if parts := strings.Split(email, "@"); len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("malformed email string: %q", email)
	}
	return nil
}

func validatePassword(val string) error {
	if len(val) < 6 {
		return fmt.Errorf("password must be a string at least 6 characters long")
	}
	return nil
}

func validateDisplayName(val string) error {
	if val == "" {
		return fmt.Errorf("display name must be a non-empty string")
	}
	return nil
}

func validateID(id string) error {
	if id == "" {
		return fmt.Errorf("id must be a non-empty string")
	}
	if len(id) > 128 {
		return fmt.Errorf("id string must not be longer than 128 characters")
	}
	return nil
}

func init() {
	client, err := fb.FirebaseApp.Auth(fb.FirebaseContext)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	authClient = client
}

package auth

import (
	"acropolis-backend/pkg/firebase"
	"acropolis-backend/pkg/user_mgt"
	firebaseAuth "firebase.google.com/go/auth"
	"fmt"
	"log"
	"net/http"
)

// authClient is a global variable to hold the initialized Firebase Auth client
var authClient *firebaseAuth.Client

// VerifySessionCookie verifies that the given session cookie is valid and returns the associated UserRecord if valid.
func VerifySessionCookie(sessionCookie *http.Cookie) (*user_mgt.UserRecord, error) {
	decoded, err := authClient.VerifySessionCookieAndCheckRevoked(firebase.FirebaseContext, sessionCookie.Value)
	if err != nil {
		return nil, fmt.Errorf("error verifying cookie: %v\n", err)
	}

	user, err := user_mgt.GetUser(decoded.UID)
	if err != nil {
		return nil, fmt.Errorf("error getting user from cookie: %v\n", err)
	}

	return user, nil
}

func init() {
	aClient, err := firebase.FirebaseApp.Auth(firebase.FirebaseContext)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	authClient = aClient
}

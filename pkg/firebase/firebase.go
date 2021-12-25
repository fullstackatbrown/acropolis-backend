package firebase

import (
	"context"
	firebaseSDK "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FirebaseApp is a global variable to hold the initialized Firebase App object
var FirebaseApp *firebaseSDK.App
var FirebaseContext context.Context

func initializeFirebaseApp() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/Users/nathanluu/Downloads/acropolis-dev-729cc-firebase-adminsdk-78nmi-f42ecd2b95.json")
	app, err := firebaseSDK.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err.Error())
	}

	FirebaseApp = app
	FirebaseContext = ctx
}

func init() {
	initializeFirebaseApp()
}

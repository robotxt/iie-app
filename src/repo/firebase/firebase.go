package repo

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"

	"github/robotxt/iie-app/src/logging"

	"google.golang.org/api/option"
)

var googleFirebaseCred string = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
var projectID string = os.Getenv("GCLOUD_PROJECT_ID")

// FirebaseApp struct
type FirebaseApp struct {
	FbApp           *firebase.App
	FbClient        *auth.Client
	FirestoreClient *firestore.Client
}

var (
	// FirebaseAuthClient firebase connection
	FirebaseAuthClient *auth.Client

	// FirebaseAppClient firebase application
	FirebaseAppClient *firebase.App

	// FirebaseAppClient firebase application
	FirestoreClient *firestore.Client
)

// InitializeFirebase APP Auth
func (fb *FirebaseApp) InitializeFirebase(ctx context.Context) *FirebaseApp {
	jsonFile, err := os.Open(googleFirebaseCred)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	opt := option.WithCredentialsJSON([]byte(byteValue))

	fbApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logging.Fatal("error initializing app: %v\n", err)
	}

	fb.FbApp = fbApp
	FirebaseAppClient = fbApp
	return fb
}

// InitializeProdFirebase APP Auth
func (fb *FirebaseApp) InitializeProdFirebase(ctx context.Context) *FirebaseApp {
	conf := &firebase.Config{ProjectID: projectID}
	fbApp, err := firebase.NewApp(ctx, conf)

	if err != nil {
		logging.Fatal("error initializing app: %v\n", err)
	}

	fb.FbApp = fbApp
	FirebaseAppClient = fbApp
	return fb
}

// FirebaseAuthInitialize firebase auth
func (fb *FirebaseApp) FirebaseAuthInitialize(ctx context.Context) *FirebaseApp {
	authclient, err := fb.FbApp.Auth(context.Background())
	if err != nil {
		logging.Fatal("error getting Auth client: %v\n", err)
	}
	fb.FbClient = authclient
	FirebaseAuthClient = authclient
	return fb
}

// FirebaseFireStoreInitialize firebase firestore
func (fb *FirebaseApp) FirebaseFireStoreInitialize(ctx context.Context) *FirebaseApp {
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		logging.Fatal("error getting Auth client: %v\n", err)
	}
	fb.FirestoreClient = firestoreClient
	FirestoreClient = firestoreClient
	return fb
}

func (fb *FirebaseApp) InitializeAllFirebaseService(ctx context.Context) *FirebaseApp {
	// Initialize authclient
	authclient, err := fb.FbApp.Auth(context.Background())
	if err != nil {
		logging.Fatal("error getting Auth client: %v\n", err)
	}
	fb.FbClient = authclient
	FirebaseAuthClient = authclient

	// Initialize firestore
	projectID := "iie-app-52841"
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		logging.Fatal("error getting Auth client: %v\n", err)
	}
	fb.FirestoreClient = firestoreClient
	FirestoreClient = firestoreClient

	return fb
}

// VerifyCustomToken firebase create custom token
func (fb *FirebaseApp) VerifyCustomToken(fbtoken string) (*auth.Token, error) {
	token, err := fb.FbClient.VerifyIDToken(context.Background(), fbtoken)
	if err != nil {
		logging.Debug("error verifying custom token: %v\n", err)
		return nil, err
	}
	logging.Info("Got custom token: %v\n", token)
	return token, err
}

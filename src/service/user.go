package service

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	repo "github/robotxt/iie-app/src/repo/firebase"

	auth "firebase.google.com/go/v4/auth"
)

var userProfileCollection = MyCollections().profile

type UserType struct {
	UID      string
	Email    string
	Password string
	Country  string
	Timezone string
	Age      int
	Group    string
}

// CreateFirebaseUser Create User
func (u *UserType) CreateFirebaseUser(ctx context.Context) (*auth.UserRecord, error) {
	newID := uuid.New()
	fmt.Printf("github.com/google/uuid: %s\n", newID.String())

	userParams := (&auth.UserToCreate{}).
		UID(newID.String()).
		Email(u.Email).
		EmailVerified(false).
		Password(u.Password).
		Disabled(false)

	user, err := repo.FirebaseAuthClient.CreateUser(ctx, userParams)

	return user, err
}

func (u *UserType) GetUserByEmail(ctx context.Context) (*auth.UserRecord, error) {
	user, err := repo.FirebaseAuthClient.GetUserByEmail(ctx, u.Email)
	return user, err
}

func (u *UserType) GetUserByUID(ctx context.Context) (*auth.UserRecord, error) {
	user, err := repo.FirebaseAuthClient.GetUser(ctx, u.UID)
	return user, err
}

func (u *UserType) HashPassword() []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hashedPassword
}

func (u *UserType) CreateUserProfile(ctx context.Context) (*firestore.WriteResult, error) {
	timezone := "Asia/Manila" // default timezone
	if u.Timezone != "" {
		timezone = u.Timezone
	}

	result, err := repo.FirestoreClient.Collection(userProfileCollection).Doc(u.UID).Set(ctx, map[string]interface{}{
		"country":  u.Country,
		"timezone": timezone,
		"age":      u.Age,
		"group":    "CLIENTS",
	})

	return result, err
}

// CreateCustomToken firebase create custom token
func (u *UserType) CreateCustomToken(ctx context.Context) (string, error) {
	token, err := repo.FirebaseAuthClient.CustomToken(ctx, u.UID)
	return token, err
}

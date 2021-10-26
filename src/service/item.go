package service

import (
	"context"
	repo "github/robotxt/iie-app/src/repo/firebase"

	"cloud.google.com/go/firestore"
)

type ItemType struct {
	UID         string
	UserUID     string
	Name        string
	Description string
	Bucket      string
}

func (u *ItemType) CreateItem(ctx context.Context) (*firestore.WriteResult, error) {
	result, err := repo.FirestoreClient.Collection("item").Doc(u.UserUID).Set(ctx, map[string]interface{}{
		"name":        u.Name,
		"description": u.Description,
		"bucket":      u.Bucket,
	})

	return result, err
}

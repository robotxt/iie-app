package service

import (
	"context"
	"encoding/json"
	"github/robotxt/iie-app/src/logging"
	"github/robotxt/iie-app/src/repo"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var ItemCollection = MyCollections().item
var Repo = repo.Repository{}

type ItemType struct {
	UID         string
	UserUID     string
	Name        string
	Description string
	Bucket      string
	Amount      float64
	Tags        string // string separated with commas
}

func (u *ItemType) CreateItem(ctx context.Context) (*firestore.DocumentRef, error) {
	doc, _, err := Repo.Firebase.FirestoreClient.Collection(ItemCollection).Add(ctx, map[string]interface{}{
		"name":        u.Name,
		"description": u.Description,
		"bucket":      u.Bucket,
		"userUID":     u.UserUID,
		"tags":        u.Tags,
		"amount":      u.Amount,
	})

	u.UID = doc.ID
	return doc, err
}

func (u *ItemType) GetUserItems(ctx context.Context) ([]ItemType, error) {
	iter := Repo.Firebase.FirestoreClient.Collection(ItemCollection).Where("userUID", "==", u.UserUID).Documents(ctx)

	allItems := []ItemType{}

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logging.Error("Error fetch data: ", err)
			return allItems, err
		}

		qr, _ := json.Marshal(doc.Data())
		var item ItemType
		err = json.Unmarshal(qr, &item)

		allItems = append(allItems, item)
	}

	return allItems, nil
}

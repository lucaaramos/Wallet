package repository

import (
	"context"

	"wallet/services/accounts/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	collection *mongo.Collection
}

func NewAccountRepository(db *mongo.Database) *AccountRepository {
	return &AccountRepository{
		collection: db.Collection("accounts"),
	}
}

func (r *AccountRepository) Save(acc models.Account) error {
	_, err := r.collection.InsertOne(context.Background(), acc)
	return err
}

func (r *AccountRepository) FindByID(id string) (*models.Account, error) {
	var acc models.Account

	err := r.collection.FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&acc)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (r *AccountRepository) Update(acc models.Account) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": acc.ID},
		bson.M{"$set": acc},
	)

	return err
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"user_id" json:"user_id"`
	Balance int64              `bson:"balance" json:"balance"`
	Amount  int64              `bson:"amount" json:"amount"`
}

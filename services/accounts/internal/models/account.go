package models

type Account struct {
	ID      string  `bson:"_id,omitempty" json:"id"`
	UserID  string  `bson:"user_id" json:"user_id"`
	Balance float64 `bson:"balance" json:"balance"`
	Amount  float64 `bson:"amount" json:"amount"`
}

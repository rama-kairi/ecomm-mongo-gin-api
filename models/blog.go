package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseSchema struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
type User struct {
	BaseSchema `bson:",inline"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" gorm:"primary_key"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `gorm:"unique" json:"email"`
	Password  string             `json:"password" gorm:"not null"`
}

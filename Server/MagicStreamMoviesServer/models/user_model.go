package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"

)

type User struct {
	ID 				bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID 			string			`json:"user_id" bson:"user_id"`
	FirstName 		string		    `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName 		string	    	`json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	UserName 		string	    	`json:"username" bson:"username"`
	Email 			string			`json:"email" bson:"email" validate:"email,required"`
	Password 		string			`json:"password" bson:"password" validate:"required,min=6"`
	Role 			string			`json:"role" bson:"role" validade:"oneof=ADMIN USER"`
	CreatedAt 		time.Time 	    `json:"created_at" bson:"created_at"`
	UpdatedAt 		time.Time 	    `json:"updated_at" bson:"updated_at"`
	Token 			string			`json:"token" bson:"token"`
	RefreshToken 	string 			`json:"refresh_token" bson:"refresh_token"`
	FavouriteGenres []Genre 		`json:"favourite_genres" bson:"favourite_genres" validate:"required,dive"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`	
}

type UserResponse struct {
	UserID 			string			`json:"user_id"`
	FirstName 		string		    `json:"first_name"`
	LastName 		string	    	`json:"last_name"`
	Email 			string			`json:"email"`
	Role 			string			`json:"role"`
	CreatedAt 		time.Time 	    `json:"created_at"`
	FavouriteGenres []Genre 		`json:"favourite_genres"`
}
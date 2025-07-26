package models

import (
	"time"
)

// Payload represents the JWT token payload structure
type Payload struct {
	Iss  string    `json:"iss,omitempty"`
	Aud  string    `json:"aud,omitempty"`
	Iat  time.Time `json:"iat,omitempty"`
	Nbf  time.Time `json:"nbf,omitempty"`
	Exp  time.Time `json:"exp,omitempty"`
	User string    `json:"user,omitempty"`
	Role string    `json:"role,omitempty"`
}

type User struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	Nama         string    `json:"nama" bson:"nama"`
	Email        string    `json:"email" bson:"email"`
	Password     string    `json:"password" bson:"password"`
	Role         string    `json:"role,omitempty" bson:"role,omitempty"`
	Status       string    `json:"status,omitempty" bson:"status,omitempty"`
	LastActive   time.Time `json:"lastActive,omitempty" bson:"lastActive,omitempty"`
	Bio          string    `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfileImage string    `json:"profileImage,omitempty" bson:"profileImage,omitempty"`
}

// UserResponse is a model without password for returning to clients
type UserResponse struct {
	ID           string    `json:"id,omitempty"`
	Nama         string    `json:"nama"`
	Email        string    `json:"email"`
	Role         string    `json:"role,omitempty"`
	Status       string    `json:"status,omitempty"`
	LastActive   time.Time `json:"lastActive,omitempty"`
	Bio          string    `json:"bio,omitempty"`
	ProfileImage string    `json:"profileImage,omitempty"`
}

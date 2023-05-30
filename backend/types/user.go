package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Username    string `json:"username" db:"username"`
	Hash        string `json:"-" db:"hash"` // shouldn't be serialized
	DateCreated string `json:"date_created" db:"date_created"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}

func NewUser(email, username, hash string) *User {
	return &User{
		ID:          uuid.New().String(),
		Email:       email,
		Username:    username,
		Hash:        hash,
		DateCreated: time.Now().Format(time.RFC3339),
		DateUpdated: time.Now().Format(time.RFC3339),
	}
}

// Add avatars to the profile type later
type Profile struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	ProfileName string `json:"profile_name" db:"profile_name"`
	DateCreated string `json:"date_created" db:"date_created"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}

func NewProfile(userID, profileName string) *Profile {
	return &Profile{
		ID:          uuid.New().String(),
		UserID:      userID,
		ProfileName: profileName,
		DateCreated: time.Now().Format(time.RFC3339),
		DateUpdated: time.Now().Format(time.RFC3339),
	}
}

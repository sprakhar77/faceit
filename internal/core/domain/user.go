package domain

import (
	"strings"
	"time"
)

// User represents the User domain model
type User struct {
	ID        int64     `json:"id" db:"id""`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	NickName  string    `json:"nickname" db:"nickname"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Country   string    `json:"country" db:"country"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u User) TrimSpaces() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Country = strings.TrimSpace(u.Country)
}

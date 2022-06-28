package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Employee struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Admin       string    `json:"admin" db:"admin"`
	Employee    string    `json:"employee" db:"employee"`
	LastUpdated string    `json:"last_updated" db:"last_updated"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	Name            string `form:"name"`
	Email           string
	Telephone       string `form:"work-telephone"`
	MobileTelephone string `form:"mobile-telephone"`
	Fax             string `form:"fax-number"`
	ContacType      string `form:"contac-type"`
	InternalAdmin   string `form:"internal-admin"`
	Employer        string `form:"employer"`
	AccsessClient   string `form:"accsess-client"`
}

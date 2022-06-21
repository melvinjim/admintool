package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Employee struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"file_name" db:"file_name"`
	Email       string    `json:"email" db:"email"`
	Admin       string    `json:"admin" db:"admin"`
	Employer    string    `json:"employer" db:"employer"`
	LastUpdated string    `json:"last_updated" db:"last_updated"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

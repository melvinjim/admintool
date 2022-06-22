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

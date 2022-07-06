package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Employer struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	Phone     string    `json:"phone" db:"phone"`
	CreateAt  time.Time `json:"create_at" db:"create_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

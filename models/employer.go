package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Employer struct {
	ID        uuid.UUID `form:"id" db:"id"`
	Name      string    `form:"name" db:"name"`
	Address   string    `form:"address" db:"address"`
	Phone     string    `form:"phone" db:"phone"`
	CreateAt  time.Time `form:"created_at" db:"created_at"`
	UpdatedAt time.Time `form:"updated_at" db:"updated_at"`
	
}

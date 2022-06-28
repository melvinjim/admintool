package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Employee struct {
	ID              uuid.UUID `form:"id" db:"id"`
	Name            string    `form:"name" db:"name"`
	Email           string    `form:"email" db:"email"`
	WorkTelephone   int       `form:"work_telephone" db:"work_telephone"`
	MobileTelephone int       `form:"mobile_telephone" db:"mobile_telephone"`
	Fax             string    `form:"fax_number" db:"fax"`
	ContacType      string    `form:"contac_type" db:"contac_type"`
	InternalAdmin   string    `form:"internal_admin" db:"internal_admin"`
	Employer        string    `form:"employer" db:"employer"`
	AccsessClient   string    `form:"accsess_client" db:"accsess_client"`
	Admin           string    `form:"admin" db:"admin"`
	LastUpdated     string    `form:"last_updated" db:"last_updated"`
	CreatedAt       time.Time `form:"created_at" db:"created_at"`
	UpdatedAt       time.Time `form:"updated_at" db:"updated_at"`
}

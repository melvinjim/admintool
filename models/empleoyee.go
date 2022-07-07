package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Employee struct {
	ID              uuid.UUID `form:"id" db:"id"`
	Name            string    `form:"name" db:"name"`
	Email           string    `form:"email" db:"email"`
	WorkTelephone   string    `form:"work_telephone" db:"work_telephone"`
	MobileTelephone string    `form:"mobile_telephone" db:"mobile_telephone"`
	Fax             string    `form:"fax_number" db:"fax"`
	ContactType     string    `form:"contac_type" db:"contac_type"`
	InternalAdmin   string    `form:"internal_admin" db:"internal_admin"`
	AccessClient    string    `form:"access_client" db:"access_client"`
	Admin           string    `form:"admin" db:"admin"`
	LastUpdated     string    `form:"last_updated" db:"last_updated"`
	EmployerID      uuid.UUID `form:"employer_id" db:"employer_id"`
	CreatedAt       time.Time `form:"created_at" db:"created_at"`
	UpdatedAt       time.Time `form:"updated_at" db:"updated_at"`
}

func (e *Employee) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Title"},
		&validators.StringIsPresent{Field: e.Email, Name: "FirstName"},
		&validators.StringIsPresent{Field: e.WorkTelephone, Name: "LastName"},
	), nil
}

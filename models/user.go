package models

import "github.com/gofrs/uuid"

type User struct {
	ID         uuid.UUID
	EmployerID Employer `json:"employer_id,omitempty" has_many:"employer" fk_id:"employer_id"`
}

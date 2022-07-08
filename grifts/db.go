package grifts

import (
	"admintool/models"

	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		return nil
	})

})

var _ = grift.Add("db:seed", func(c *grift.Context) error {
	e := []models.Employer{
		{ID: uuid.UUID{}, Name: "Acme LLC", Address: "street 84", Phone: "3004800947"},
		{ID: uuid.UUID{}, Name: "Magnus LLC", Address: "street 32", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "RMC123", Address: "street 49", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "Ramon", Address: "street 31", Phone: "3004800947"},
		{ID: uuid.UUID{}, Name: "Pedro", Address: "street 30", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "Juan", Address: "street 8", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "Camilo", Address: "street 89", Phone: "3004800947"},
		{ID: uuid.UUID{}, Name: "Luis", Address: "street 45", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "Manuel", Address: "street 59", Phone: "3006547676"},
		{ID: uuid.UUID{}, Name: "Esteban", Address: "street 87", Phone: "3004800947"},
	}
	err := models.DB.Create(e)
	return err
})

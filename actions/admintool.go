package actions

import (
	"admintool/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ListEmployees(c buffalo.Context) error {
	employees := []models.Employee{}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.All(&employees)
	if err != nil {
		return err
	}

	c.Set("employees", employees)

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func AddUser(c buffalo.Context) error {
	employee := models.Employee{}
	c.Set("employee", employee)

	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func Create(c buffalo.Context) error {
	employee := models.Employee{}
	if err := c.Bind(&employee); err != nil {
		return err
	}

	tx := c.Value("tx").(*pop.Connection)

	verrs, err := employee.Validate(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("employee", employee)
		c.Set("errors", verrs)
		return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
	}

	errCreate := tx.Create(&employee)
	if errCreate != nil {
		return errCreate
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

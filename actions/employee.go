package actions

import (
	"admintool/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func ListEmployees(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	employees := []models.Employee{}

	filter := &models.Filter{}
	if err := c.Bind(filter); err != nil {
		return err
	}

	q := tx.Where(`employees.name like ?`, `%%`)
	err := q.All(&employees)
	if err != nil {
		return err
	}

	employers := []models.Employer{}
	err = tx.All(&employers)
	if err != nil {
		return err
	}

	c.Set("employers", employers)

	c.Set("employees", employees)

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func AddNewEmployees(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	employee := models.Employee{}
	c.Set("employee", employee)

	employers := []models.Employer{}
	if err := tx.All(&employers); err != nil {
		return err
	}

	employersMap := map[string]uuid.UUID{}
	for _, e := range employers {
		employersMap[e.Name] = e.ID
	}

	c.Set("employersMap", employersMap)

	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func CreateEmployees(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)

	employee := models.Employee{}
	if err := c.Bind(&employee); err != nil {
		return err
	}

	verrs, err := employee.Validate(tx)
	if err != nil {
		return err
	}

	if employee.InternalAdmin == "True" {
		employee.Admin = "Administrator"
	} else {
		employee.Admin = "User"
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

func InfoEmployer(c buffalo.Context) error {
	employee := models.Employee{}
	employeeID := c.Param("employee_id")

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Find(&employee, employeeID)
	if err != nil {
		return err
	}

	employers := models.Employer{}
	err = tx.Find(&employers, employee.EmployerID)
	if err != nil {
		return err
	}

	c.Set("employers", employers)

	c.Set("employee", employee)

	return c.Render(http.StatusOK, r.HTML("users/info_user.plush.html"))
}

func Edit(c buffalo.Context) error {
	employee := models.Employee{}
	employeeID := c.Param("employee_id")

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Find(&employee, employeeID)
	if err != nil {
		return err
	}

	employersMaps := []models.Employer{}
	if err := tx.All(&employersMaps); err != nil {
		return err
	}

	employersMap := map[string]uuid.UUID{}
	for _, e := range employersMaps {
		employersMap[e.Name] = e.ID
	}

	employers := models.Employer{}
	err = tx.Find(&employers, employee.EmployerID)
	if err != nil {
		return err
	}

	c.Set("employers", employers)

	c.Set("employersMap", employersMap)

	c.Set("employee", employee)

	return c.Render(http.StatusOK, r.HTML("users/edit.plush.html"))
}

func EmployerEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	employee := models.Employee{}
	employeeID := c.Param("employee_id")

	err := tx.Find(&employee, employeeID)
	if err != nil {
		return err
	}

	if err := c.Bind(&employee); err != nil {
		return err
	}

	if employee.InternalAdmin == "True" {
		employee.Admin = "Administrator"
	} else {
		employee.Admin = "User"
	}

	err = tx.Update(&employee)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func DeleteEmployees(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	employee := models.Employee{}
	employeeID := c.Param("employee_id")

	err := tx.Find(&employee, employeeID)
	if err != nil {
		return err
	}

	err = tx.Destroy(&employee)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func Filter(c buffalo.Context) error {

	filter := &models.Filter{}
	if err := c.Bind(filter); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

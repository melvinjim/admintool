package actions

import (
	"admintool/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Admintool(c buffalo.Context) error {
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
	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func ReceiveData(c buffalo.Context) error {

	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("Nombre:", u.Name)
	fmt.Println("Email:", u.Email)
	fmt.Println("Work telephone:", u.Telephone)
	fmt.Println("Mobile Telephone:", u.MobileTelephone)
	fmt.Println("Fax:", u.Fax)
	fmt.Println("Contact Type:", u.ContacType)
	fmt.Println("Is Internal Admin:", u.InternalAdmin)
	fmt.Println("Employer:", u.Employer)
	fmt.Println("Accsess Client:", u.AccsessClient)

	return c.Redirect(http.StatusSeeOther, "/users/new")
}

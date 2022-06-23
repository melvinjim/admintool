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
	_, fileInfo, err := c.Request().FormFile("information")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo)

	return c.Redirect(http.StatusSeeOther, "/users/new")
}

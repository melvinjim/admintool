package actions

import (
	"admintool/models"
	"log"
	"net/http"

	v "github.com/gobuffalo/validate"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func main() {
	u := models.Employee{Name: "", Email: "", WorkTelephone: ""}
	errors := v.Validate(&u)
	log.Println(errors.Errors)
}

func (u *models.Employee) IsValid(errors *v.Errors) {
	if u.Name == "" {
		errors.Add("name", "Name must not be blank!")
	}
	if u.Email == "" {
		errors.Add("email", "Email must not be blank!")
	}
	if u.WorkTelephone == "" {
		errors.Add("work telephone", "work telephone must not be blank!")
	}
}

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
	u := models.Employee{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Create(&u)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Admintool(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("admintool/index.plush.html"))
}

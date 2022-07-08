package grifts

import (
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("db", func() {

	Desc("seed", "Task Description")
	Add("seed", func(c *Context) error {
		return nil
	})

})

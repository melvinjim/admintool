package grifts

import (
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("foo", func() {

	Desc("bar", "Task Description")
	Add("bar", func(c *Context) error {
		return nil
	})

})

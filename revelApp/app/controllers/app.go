package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type Application struct {
	GorpController
}

func (c Application) Index()  revel.Result{
	fmt.Println("Application.Index")
	return c.Render()
}

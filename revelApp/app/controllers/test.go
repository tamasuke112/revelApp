package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"revelApp/app/util"
	"revelApp/app/models"
)

type Test struct {
	Application
}

func (c Test) Index() revel.Result{
	fmt.Println("hello testIndex")
	return c.Render()
}

func (c Test) List() revel.Result{
	fmt.Println("hello testList")
	
	query := util.GetSql("getTestList")
	var results []models.Test
	rows, err := c.Txn.Query(query)
	
	if err != nil{
		panic(err)
	}
	
	for rows.Next() {
		result := models.Test{}
		rows.Scan(&result.Name)
		results = append(results, result)
	}
	fmt.Println("testList start")
	fmt.Println(results)
	fmt.Println("testList end")
	return c.Render(results)	
}

func (c Test) Create() revel.Result{
	fmt.Println("hello testCreate")
	return c.Render()
}

func (c Test) SaveUser() revel.Result{
	query := util.GetSql("insert")
	_, err := c.Txn.Exec(query, "test0909")
	if err != nil{
		panic(err)
	}
	return c.Redirect(c.List())	
}
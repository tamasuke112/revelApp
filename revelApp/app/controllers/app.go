package controllers

import (
	//"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	//"github.com/revel/samples/booking/app/models"
	//"github.com/revel/samples/booking/app/routes"
	"fmt"
	//"revelApp/app/util"
	//"revelApp/app/models"
	//"revelApp/app/routes"
	//"database/sql"
	//"log"
)

type Application struct {
	GorpController
}
/*
type Test struct{
	name string	
}
*/
/*
func (c Application) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}
*/
/*
func (c Application) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}
*/
/*
func (c Application) getUser(username string) *models.User {
	users, err := c.Txn.Select(models.User{}, `select * from User where Username = ?`, username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.User)
}

func (c Application) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Hotels.Index())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}
*/
func (c Application) Index()  revel.Result{
	/*
	fmt.Println("hello index")
	fmt.Println("getTest")
	result, err := c.getTest()
	if err != nil {
		panic(err)
	}
	fmt.Println("result start")
	fmt.Println(result)
	fmt.Println("result end")
	fmt.Println("inser start")
	c.setTest()
	fmt.Println("inser end")
	*/
	fmt.Println("Application.Index")
	return c.Render()
}
/*
func (c Application) getTest() (models.Test, error){
	query := util.GetSql("test")
	b := "han9"
	tempData :=models.Test{}
	err := c.Txn.QueryRow(query, b).Scan(&tempData.Name)
	
	switch{	
		case err == sql.ErrNoRows:
			return tempData, fmt.Errorf("err", b)
		case err != nil:
			return tempData, err
	}
	return tempData,nil
}

func (c Application) setTest() revel.Result{
	query := util.GetSql("insert")
	_, err := c.Txn.Exec(query, "test0909")
	if err != nil{
		panic(err)
	}
	return c.Redirect(routes.Application.Index())
}
*/
/*
func (c Application) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Application.Register())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.Username
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(routes.Hotels.Index())
}
*/
/*
func (c Application) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Hotels.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.Index())
}
*/
/*
func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}
*/
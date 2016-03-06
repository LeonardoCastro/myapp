package controllers

import (
	"github.com/LeonardoCastro/myapp/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// App revel.controller + sqlx.transaction
type App struct {
	SqlxController
}

// Index function rendering greeting at Index page
func (c App) Index() revel.Result {
	greeting := "¡Aloha marajá!"
	return c.Render(greeting)
}

// Register NOT SURE WHAT IT DOES
func (c App) Register() revel.Result {
	return c.Render()
}

// AddUser checks if user is connected and sends it to RenderArgs NOT SURE WHAT IT DOES
func (c App) AddUser() revel.Result {
	if user := c.Connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

// Connected checks if a user is connected
func (c App) Connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username)
	}
	return nil
}

// Login connects a user
func (c App) Login(username, password string, remember bool) revel.Result {
	user := c.GetUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			return c.Render(username)
		}
	}
	c.Flash.Out["user"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(App.Index)
}

// GetUser chek if a user exists on the database
func (c App) GetUser(username string) *models.User {
	users := []models.User{}
	err := c.Txn.Select(&users, "select * from Users where Username = $1 ", username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return &users[0] //.(*models.User)
}

// Logout disconnects a user
func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}

// SaveUser on the database
func (c App) SaveUser(user models.User, verifyPassword string) revel.Result {

	var userRegex = regexp.MustCompile("^\\w*$")

	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password must match")

	c.Validation.Check(user.Username,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{15},
		revel.Match{userRegex},
	)

	c.Validation.Check(user.Password,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{20},
	)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Register)
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)

	c.Txn.MustExec("INSERT INTO Users (username, password, hashedpassword) VALUES ($1, $2, $3)", user.Username, user.Password, user.HashedPassword)
	c.Commit(c.Txn)
	username := user.Username
	return c.Render(username)
}

// Hello test function
func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}

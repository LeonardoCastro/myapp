package controllers

import (
	 "github.com/revel/revel"
	 "github.com/LeonardoCastro/myapp/models"
	 "golang.org/x/crypto/bcrypt"
 )


type App struct {
	SqlxController
}


func (c App) Index() revel.Result {
	greeting := "¡Aloha marajá!"
	return c.Render(greeting)
}


func (c App) Register() revel.Result {
	return c.Render()
}


func (c App) AddUser() revel.Result {
	if user :=c.Connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}


func (c App) Connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username)
	}
	return nil
}


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
		c.Render(username)
		}
	}
	c.Flash.Out["user"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(App.Index)
}


func (c App) GetUser(username string) *models.User {
	users := []models.User{}
	err := c.Txn.Select(&users, "select * from Users where Username = $1 ", username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	} else {
		return &users[0]//.(*models.User)
	}
}


func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}


func (c App) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
				Message("Password must match")
	user.Validate(c.Validation)

	// c.Validation.Required(username).Message("A username is required!")
	// c.Validation.Length(username, 8).Message("Username must be 8 characters.")
	// c.Validation.Match(username, regexp.MustCompile(`[A-Za-z0-9_]`)).Message("Invalid character.")
	//
	// c.Validation.Required(email).Message("A valid email is required")
	// c.Validation.Email(email).Message("invalid email.")
	//
	// c.Validation.Required(pwd).Message("Password is required.")
	// c.Validation.Required(pwdConf).Message("Please confirm password.")
	// c.Validation.Match(pwd, regexp.MustCompile(`[A-Za-z0-9_]`)).Message("Invalid character.")
	// c.Validation.Match(pwdConf, regexp.MustCompile(`[A-Za-z0-9_]`)).Message("Invalid character.")
	// c.Validation.Match(pwd, regexp.MustCompile(pwdConf)).Message("Passwords must match.")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)

	c.Txn.MustExec("INSERT INTO Users VALUES ($1, $2, $3, $4)", 1, user.Username, user.Password, user.HashedPassword)
	username := user.Username
	return c.Render(username)
}

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

package models

import (
	"github.com/revel/revel"
	"regexp"
)

// User type used for linking to the database
type User struct {
	ID             int
	Username       string
	Password       string
	HashedPassword []byte
}

var userRegex = regexp.MustCompile("^\\w*$")

// Validate if user inputs are correct
func (user *User) Validate(v *revel.Validation, min, max int) *revel.ValidationResult {
	return v.Check(user.Username,
		revel.Required{},
		revel.MinSize{min},
		revel.MaxSize{max},
		//revel.Match{userRegex},
	)
	//user.ValidatePassword(v, user.Password)
}

//ValidatePassword to check password and it's validation
func (user *User) ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{20},
	)
}

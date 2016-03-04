package models

import (
  "github.com/revel/revel"
  "regexp"
)

type User struct {
  Id                int
  Username          string
  Password          string
  HashedPassword    []byte
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
  v.Check(user.Username,
    revel.Required{},
    revel.MinSize{4},
    revel.MaxSize{15},
    revel.Match{userRegex},
  )

  ValidatePassword(v, user.Password)

}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
  return v.Check(password,
      revel.Required{},
      revel.MinSize{8},
      revel.MaxSize{20},
    )
}

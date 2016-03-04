package controllers

import "github.com/revel/revel"

func init() {
  	revel.OnAppStart(InitDB)
  	revel.InterceptMethod((*SqlxController).Begin, revel.BEFORE)
    revel.InterceptMethod(App.AddUser, revel.BEFORE)
}

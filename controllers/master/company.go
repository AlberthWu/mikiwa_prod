package controllers

import (
	"fmt"
	base "mikiwa_prod/controllers"
)

type CompanyController struct {
	base.BaseControllers
}

func (c *CompanyController) Post() {
	data := fmt.Sprintln("test")
	c.Data["json"] = data
	c.ServeJSON()

}

func (c *CompanyController) Put() {

}

func (c *CompanyController) Delete() {

}

func (c *CompanyController) GetOne() {

}

func (c *CompanyController) GetAll() {
	data := fmt.Sprintln("get All")
	c.Data["json"] = data
	c.ServeJSON()
}

package controllers

import (
	"fmt"
	base "mikiwa_prod/controllers"
)

type CompanyController struct {
	base.BaseController
}

func (c *CompanyController) Post() {

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

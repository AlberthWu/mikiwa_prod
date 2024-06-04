package controllers

import (
	"fmt"
	base "mikiwa_prod/controllers"
)

type CompanyController struct {
	base.BaseControllers
}

func (c *CompanyController) Post() {
	fmt.Println("tes")

}

func (c *CompanyController) Put() {

}

func (c *CompanyController) Delete() {

}

func (c *CompanyController) GetOne() {

}

func (c *CompanyController) GetAll() {
	fmt.Println("get")

}

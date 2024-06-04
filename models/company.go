package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type (
	Company struct {
		Id              int       `json:"id" orm:"column(id);auto;pk"`
		Position        int8      `json:"position" orm:"column(position)"`
		TypeId          int       `json:"type_id" orm:"column(type_id)"`
		ParentId        int       `json:"parent_id" orm:"column(parent_id)"`
		Code            string    `json:"code" orm:"column(code)"`
		Name            string    `json:"name" orm:"column(name)"`
		Email           string    `json:"email" orm:"column(email)"`
		Phone           string    `json:"phone" orm:"column(phone)"`
		Fax             string    `json:"fax" orm:"column(fax)"`
		Npwp            string    `json:"npwp" orm:"column(npwp)"`
		NpwpName        string    `json:"npwp_name" orm:"column(npwp_name)"`
		NpwpAddress     string    `json:"npwp_address" orm:"column(npwp_address)"`
		Terms           int       `json:"terms" orm:"column(terms)"`
		Credit          float64   `json:"credit" orm:"column(credit);digits(20);decimals(2);default(0)"`
		IsTax           int8      `json:"is_tax" orm:"column(is_tax)"`
		Address         string    `json:"address" orm:"column(address)"`
		DistrictId      int       `json:"district_id" orm:"column(district_id)"`
		CityId          int       `json:"city_id" orm:"column(city_id)"`
		StateId         int       `json:"state_id" orm:"column(state_id)"`
		Zip             int       `json:"zip" orm:"column(zip)"`
		BankId          int       `json:"bank_id" orm:"column(bank_id)"`
		BankAccountName string    `json:"bank_account_name" orm:"column(bank_account_name)"`
		BankNo          string    `json:"bank_no" orm:"column(bank_no)"`
		BankName        string    `json:"bank_name" orm:"column(bank_name)"`
		BankBranch      string    `json:"bank_branch" orm:"column(bank_branch)"`
		Status          int8      `json:"status" orm:"column(status)"`
		IsPo            int8      `json:"is_po" orm:"column(is_po)"`
		IsCash          int8      `json:"is_cash" orm:"column(is_cash)"`
		IsReceipt       int8      `json:"is_receipt" orm:"column(is_receipt)"`
		PriceMethod     int8      `json:"price_method" orm:"column(price_method)"`
		CreatedAt       time.Time `json:"created_at" orm:"column(created_at);type(timestamp);auto_now_add"`
		UpdatedAt       time.Time `json:"updated_at" orm:"column(updated_at);type(timestamp);auto_now"`
		DeletedAt       time.Time `json:"deleted_at" orm:"column(deleted_at);type(timestamp);null"`
	}

	Plant struct {
		Id          int       `json:"id" orm:"column(id);auto;pk"`
		CompanyId   int       `json:"company_id" orm:"column(company_id)"`
		Name        string    `json:"name" orm:"column(name)"`
		Pic         string    `json:"pic" orm:"column(pic)"`
		Phone       string    `json:"phone" orm:"column(phone)"`
		Fax         string    `json:"fax" orm:"column(fax)"`
		Address     string    `json:"address" orm:"column(address)"`
		DistrictId  int       `json:"district_id" orm:"column(district_id)"`
		CityId      int       `json:"city_id" orm:"column(city_id)"`
		StateId     int       `json:"state_id" orm:"column(state_id)"`
		Zip         int       `json:"zip" orm:"column(zip)"`
		IsDo        int8      `json:"is_do" orm:"column(is_do)"`
		IsPo        int8      `json:"is_po" orm:"column(is_po)"`
		IsSchedule  int8      `json:"is_schedule" orm:"column(is_schedule)"`
		IsReceipt   int8      `json:"is_receipt" orm:"column(is_receipt)"`
		PriceMethod int8      `json:"price_method" orm:"column(price_method)"`
		Status      int8      `json:"status" orm:"column(status)"`
		CreatedAt   time.Time `json:"created_at" orm:"column(created_at);type(timestamp);auto_now_add"`
		UpdatedAt   time.Time `json:"updated_at" orm:"column(updated_at);type(timestamp);auto_now"`
		DeletedAt   time.Time `json:"deleted_at" orm:"column(deleted_at);type(timestamp);null"`
	}
)

func (t *Company) TableName() string {
	return "companies"
}

func Companies() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Company))
}

func (t *Plant) TableName() string {
	return "plants"
}

func Plants() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Plant))
}

func init() {
	orm.RegisterModel(new(Company), new(Plant))
}

func (t *Company) Insert(m Company) (*Company, error) {
	o := orm.NewOrm()

	if _, err := o.Insert(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (t *Company) Update(fields ...string) error {
	o := orm.NewOrm()
	if _, err := o.Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Plant) Insert(m Plant) (*Plant, error) {
	o := orm.NewOrm()

	if _, err := o.Insert(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (t *Plant) Update(fields ...string) error {
	o := orm.NewOrm()
	if _, err := o.Update(t, fields...); err != nil {
		return err
	}
	return nil
}

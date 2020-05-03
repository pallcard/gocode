package service

import (
	"project05/model"
)

type CustomerService struct {
	customers [] model.Customer
	CustomerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService {}
	customerService.CustomerNum = 1
	customer := model.NewCustomer(1, "lk", "man", 20, "11111", "lk@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}


func (this *CustomerService) List() [] model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.CustomerNum++
	customer.Id = this.CustomerNum
	this.customers = append(this.customers, customer)
	return true
}
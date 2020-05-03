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

func (this *CustomerService) FindIndexById(id int) int {

	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) FindByIndex(index int) model.Customer {

	return this.customers[index]
}


func (this *CustomerService) Delete(id int) bool {
	index := this.FindIndexById(id)
	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...) 
	return true
}

func (this *CustomerService) Update(id int, customer model.Customer) bool {
	index := this.FindIndexById(id)
	if index == -1 {
		return false
	}
	this.customers[index] = customer
	return true
}
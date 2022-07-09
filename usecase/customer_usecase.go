package usecase

import (
	"livecode3/model"
	"livecode3/repo"
)

type CustomerUseCase interface {
	CustomerRegistration(name string, mobile string) (model.Customer, error)
}

type customerUseCase struct {
	repo repo.CustomerRepo
}

func (c *customerUseCase) CustomerRegistration(name string, mobile string) (model.Customer, error) {
	customer := model.Customer{
		CustomerName:  name,
		MobilePhoneNo: mobile,
	}
	// ngisi meja
	// if table.IsAvailable {
	// 	table.IsAvailable = false
	// }
	err := c.repo.Create(&customer)
	return customer, err
}

func NewCustomerUseCase(repo repo.CustomerRepo) CustomerUseCase {
	custRepo := new(customerUseCase)
	custRepo.repo = repo
	return custRepo
}

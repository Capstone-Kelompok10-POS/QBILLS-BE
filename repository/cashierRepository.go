package repository

import (
	"qbills/models/domain"

	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type CashierRepository interface {
	Create(cashier *domain.Cashier) (*domain.Cashier, error)
	FindByUsername(username string) (*domain.Cashier, error)
	FindById(id int) (*domain.Cashier, error)
	Update(cashier *domain.Cashier, id int) (*domain.Cashier, error)
}

type CashierRepositoryImpl struct {
	DB *gorm.DB
}

func NewCashierRepository(DB *gorm.DB) CashierRepository {
	return &CashierRepositoryImpl{DB: DB}
}

func (repository *CashierRepositoryImpl) Create(cashier *domain.Cashier) (*domain.Cashier, error) {
	cashierDB := req.CashierDomainintoCashierSchema(*cashier)
	result := repository.DB.Create(&cashierDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.CashierSchemaToCashierDomain(cashierDB)

	return results, nil
}

func (repository *CashierRepositoryImpl) FindById(id int) (*domain.Cashier, error) {
	cashier := domain.Cashier{}

	result := repository.DB.First(&cashier, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cashier, nil
}

func (repository *CashierRepositoryImpl) FindByUsername(username string) (*domain.Cashier, error) {
	cashier := domain.Cashier{}

	result := repository.DB.Where("username = ?", username).First(&cashier)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cashier, nil
}

func (repository *CashierRepositoryImpl) Update(cashier *domain.Cashier, id int) (*domain.Cashier, error) {
	result := repository.DB.Table("cashier").Where("id = ?", id).Updates(domain.Cashier{
		Fullname: cashier.Fullname,
		Username: cashier.Username,
		Password: cashier.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return cashier, nil
}
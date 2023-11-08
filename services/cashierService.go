package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CashierService interface {
	LoginCashier(ctx echo.Context, request web.CashierLoginRequest) (*domain.Cashier, error)
	CreateCashier(ctx echo.Context, request web.CashierCreateRequest) (*domain.Cashier, error)
	UpdateCashier(ctx echo.Context, request web.CashierUpdateRequest, id int) (*domain.Cashier, error)
	FindById(ctx echo.Context, id int) (*domain.Cashier, error)
}

type CashierServiceImpl struct {
	CashierRepository repository.CashierRepository
	Validate          *validator.Validate
}


func NewCashierService(cashierRepository repository.CashierRepository, validate *validator.Validate) *CashierServiceImpl {
	return &CashierServiceImpl{
		CashierRepository: cashierRepository,
		Validate:          validate,
	}
}

func (service *CashierServiceImpl) CreateCashier(ctx echo.Context, request web.CashierCreateRequest) (*domain.Cashier, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingCashier, _ := service.CashierRepository.FindByUsername(request.Username)
	if existingCashier != nil {
		return nil, fmt.Errorf("email already exists")
	}
	cashier := req.CashierCreateRequestToCashierDomain(request)

	cashier.Password = helpers.HashPassword(cashier.Password)
	result, err := service.CashierRepository.Create(cashier)

	if err != nil {
		return nil, fmt.Errorf("error creating admin %s", err.Error())
	}

	return result, nil
}

func (service *CashierServiceImpl) LoginCashier(ctx echo.Context, request web.CashierLoginRequest) (*domain.Cashier, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingCashier, err := service.CashierRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("infalid email or password")
	}

	cashier := req.CashierLoginRequestToCashierDomain(request)

	err = helpers.ComparePassword(existingCashier.Password, cashier.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingCashier, nil
}

func (service *CashierServiceImpl) UpdateCashier(ctx echo.Context, request web.CashierUpdateRequest, id int) (*domain.Cashier, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingCashier, _ := service.CashierRepository.FindById(id)
	if existingCashier == nil {
		return nil, fmt.Errorf("cashier not found")
	}

	cashier := req.CashierUpdateRequestToCashierDomain(request)
	cashier.Password = helpers.HashPassword(cashier.Password)
	result, err := service.CashierRepository.Update(cashier, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data cashier: %s", err.Error())
	}

	return result, nil
}

func (service *CashierServiceImpl) FindById(ctx echo.Context, id int) (*domain.Cashier, error) {
	existingCashier, _ := service.CashierRepository.FindById(id)
	if existingCashier == nil {
		return nil, fmt.Errorf("cashier not found")
	}

	return existingCashier, nil
}

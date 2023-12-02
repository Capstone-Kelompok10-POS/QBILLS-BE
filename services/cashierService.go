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
	FindAll(ctx echo.Context) ([]domain.Cashier, error)
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	FindByName(ctx echo.Context, name string) (*domain.Cashier, error)
=======
	FindByUsername(ctx echo.Context, name string) (*domain.Cashier, error)
>>>>>>> Stashed changes
=======
	FindByName(ctx echo.Context, name string) (*domain.Cashier, error)
>>>>>>> Stashed changes
	DeleteCashier(ctx echo.Context, id int) error
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
		return nil, fmt.Errorf("username already exists")
	}
	cashier := req.CashierCreateRequestToCashierDomain(request)

	cashier.Password = helpers.HashPassword(cashier.Password)
	result, err := service.CashierRepository.Create(cashier)

	if err != nil {
		return nil, fmt.Errorf("error creating cashier %s", err.Error())
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
		return nil, fmt.Errorf("invalid username or password")
	}

	cashier := req.CashierLoginRequestToCashierDomain(request)

	err = helpers.ComparePassword(existingCashier.Password, cashier.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
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
<<<<<<< Updated upstream

	cashier := req.CashierUpdateRequestToCashierDomain(request)
=======
	cashier := req.CashierUpdateRequestToCashierDomain(request)
	existingCashierUsername, _ := service.CashierRepository.FindByUsername(cashier.Username)
	if existingCashierUsername != nil {
		return nil, fmt.Errorf("username already exists")
	}
>>>>>>> Stashed changes
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

func (service *CashierServiceImpl) FindAll(ctx echo.Context) ([]domain.Cashier, error) {
	cashiers, err := service.CashierRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("cashier not found")
	}

	return cashiers, nil
}

<<<<<<< Updated upstream
<<<<<<< Updated upstream
func (service *CashierServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Cashier, error) {
	cashier, _ := service.CashierRepository.FindByName(name)
=======
func (service *CashierServiceImpl) FindByUsername(ctx echo.Context, name string) (*domain.Cashier, error) {
	cashier, _ := service.CashierRepository.FindByUsername(name)
	fmt.Println(cashier)
>>>>>>> Stashed changes
=======
func (service *CashierServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Cashier, error) {
	cashier, _ := service.CashierRepository.FindByName(name)
>>>>>>> Stashed changes
	if cashier == nil {
		return nil, fmt.Errorf("cashier not found")
	}

	return cashier, nil
}

<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
func (service *CashierServiceImpl) DeleteCashier(ctx echo.Context, id int) error {
	existingCashier, _ := service.CashierRepository.FindById(id)
	if existingCashier == nil {
		return fmt.Errorf("cashier not found")
	}

	err := service.CashierRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting cashier: %s", err)
	}

	return nil
}

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

type SuperAdminService interface {
	LoginSuperAdmin(ctx echo.Context, request web.SuperAdminLoginRequest) (*domain.SuperAdmin, error)
	FindById(ctx echo.Context, id int) (*domain.SuperAdmin, error)
	FindAll(ctx echo.Context) ([]domain.SuperAdmin, error)
}

type SuperAdminServiceImpl struct {
	SuperAdminRepository repository.SuperAdminRepository
	Validate        *validator.Validate
}

func NewSuperAdminService(SuperAdminRepository repository.SuperAdminRepository, validate *validator.Validate) *SuperAdminServiceImpl {
	return &SuperAdminServiceImpl{
		SuperAdminRepository: SuperAdminRepository,
		Validate:        validate,
	}
}


func (service *SuperAdminServiceImpl) LoginSuperAdmin(ctx echo.Context, request web.SuperAdminLoginRequest) (*domain.SuperAdmin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingSuperAdmin, err := service.SuperAdminRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	SuperAdmin := req.SuperAdminLoginRequestToSuperAdminDomain(request)

	err = helpers.ComparePassword(existingSuperAdmin.Password, SuperAdmin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingSuperAdmin, nil

}



func (service *SuperAdminServiceImpl) FindById(ctx echo.Context, id int) (*domain.SuperAdmin, error) {
	existingSuperAdmin, _ := service.SuperAdminRepository.FindById(id)
	if existingSuperAdmin == nil {
		return nil, fmt.Errorf("SuperAdmin not found")
	}

	return existingSuperAdmin, nil
}

func (service *SuperAdminServiceImpl) FindAll(ctx echo.Context) ([]domain.SuperAdmin, error) {
	SuperAdmins, err := service.SuperAdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("SuperAdmins not found")
	}

	return SuperAdmins, nil
}



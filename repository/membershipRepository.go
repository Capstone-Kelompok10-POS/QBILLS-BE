package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"

	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership *domain.Membership) (*domain.Membership, error)
	Update(membership *domain.Membership, id int) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, error)
	FindByPhoneNumber(phone_number string) (*domain.Membership, error)
=======
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)	
	FindAll() ([]domain.Membership, error)
	FindByTelephone(telephone string) (*domain.Membership, error)
=======
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, error)
	FindByPhoneNumber(phone_number string) (*domain.Membership, error)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Delete(id int) error
}

type MembershipRepositoryImpl struct {
	DB *gorm.DB
}

func NewMembershipRepository(DB *gorm.DB) MembershipRepository {
	return &MembershipRepositoryImpl{DB: DB}
}

func (repository *MembershipRepositoryImpl) Create(membership *domain.Membership) (*domain.Membership, error) {
	membershipDB := req.MembershipDomainintoMembershipSchema(*membership)
	result := repository.DB.Create(&membershipDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.MembershipSchemaToMembershipDomain(membershipDB)

	return results, nil
}

func (repository *MembershipRepositoryImpl) Update(membership *domain.Membership, id int) (*domain.Membership, error) {
	result := repository.DB.Table("memberships").Where("id = ?", id).Updates(domain.Membership{
<<<<<<< Updated upstream
		Name:         membership.Name,
		Phone_Number: membership.Phone_Number})
=======
<<<<<<< Updated upstream
		Name: membership.Name,
		Telephone: membership.Telephone})
=======
		Name:         membership.Name,
		Phone_Number: membership.Phone_Number})
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if result.Error != nil {
		return nil, result.Error
	}

	return membership, nil
}

func (repository *MembershipRepositoryImpl) FindById(id int) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.First(&membership, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &membership, nil
}

<<<<<<< Updated upstream
func (repository *MembershipRepositoryImpl) FindByPhoneNumber(phone_number string) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("phone_number = ?", phone_number).First(&membership)
=======
<<<<<<< Updated upstream
func (repository *MembershipRepositoryImpl) FindByTelephone(telephone string) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("telephone = ?", telephone).First(&membership)
=======
func (repository *MembershipRepositoryImpl) FindByPhoneNumber(phone_number string) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("phone_number = ?", phone_number).Where("deleted_at IS NULL").First(&membership)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if result.Error != nil {
		return nil, result.Error
	}

	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindByName(name string) (*domain.Membership, error) {
	membership := domain.Membership{}

<<<<<<< Updated upstream
	result := repository.DB.Where("name = ?", name).First(&membership)
=======
	result := repository.DB.Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").Find(&membership)
>>>>>>> Stashed changes
	if result.Error != nil {
		return nil, result.Error
	}

	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindAll() ([]domain.Membership, error) {
	membership := []domain.Membership{}

	result := repository.DB.Find(&membership)
	if result.Error != nil {
		return nil, result.Error
	}
	return membership, nil
}

func (repository *MembershipRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Membership{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
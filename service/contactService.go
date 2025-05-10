package service

import (
	"graph/domain"
	"graph/dto"
	"graph/errs"
)

type ContactService struct {
	repo *domain.ContactPostgresRepo
}

func NewContactService(repo *domain.ContactPostgresRepo) *ContactService {
	return &ContactService{repo: repo}
}

func (c *ContactService) MakeOne(coDto dto.ContactRequestDto) (*domain.Contact, *errs.AppError) {
	phones := make([]string, 0)
	for _, phone := range coDto.PhoneNumbers {
		phones = append(phones, phone.Number)
	}
	if existed, err := PhonesChecker(c, phones); err != nil {
		return nil, err
	} else if !existed {
		return c.repo.Create(coDto)
	}
	return nil, errs.NewBadRequestError("the phone number is existed")
}
func (c *ContactService) Search(coSeaDto dto.ContactSearchRequestDto) ([]domain.Contact, *errs.AppError) {
	return c.repo.FindWithField(coSeaDto)
}

func (c *ContactService) Update(cUpdDto dto.ContactUpdateRequestDto) (*domain.Contact, *errs.AppError) {
	phones := make([]string, 0)
	for _, phone := range cUpdDto.PhoneNumbers {
		phones = append(phones, phone.Number)
	}
	if existed, err := PhonesChecker(c, phones); err != nil {
		return nil, err
	} else if !existed {
		return c.repo.Update(cUpdDto)
	}
	return nil, errs.NewBadRequestError("the phone number is existed")
}

func PhonesChecker(c *ContactService, phoneNumber []string) (bool, *errs.AppError) {
	for _, phone := range phoneNumber {
		if existed, err := c.repo.CheckPhoneExistence(phone); err != nil {
			return false, err
		} else if existed {
			return true, nil
		}
	}
	return false, nil
}

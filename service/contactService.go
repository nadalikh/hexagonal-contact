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
	return c.repo.Create(coDto)
}
func (c *ContactService) Search(coSeaDto dto.ContactSearchRequestDto) ([]domain.Contact, *errs.AppError) {
	return c.repo.FindWithField(coSeaDto)
}

func (c *ContactService) Update(cUpdDto dto.ContactUpdateRequestDto) (*domain.Contact, *errs.AppError) {
	return c.repo.Update(cUpdDto)
}

package service

import (
	"graph/domain"
	"graph/dto"
	"graph/errs"
)

type PhoneNumberService struct {
	repo *domain.PhoneNumberPostgresRepo
}

func NewPhoneNumberService(repo *domain.PhoneNumberPostgresRepo) *PhoneNumberService {
	return &PhoneNumberService{repo: repo}
}

func (c *PhoneNumberService) AddOne(phNumDto dto.AddPhoneNumberRequestDto) *errs.AppError {
	return c.repo.AddToContact(phNumDto)
}

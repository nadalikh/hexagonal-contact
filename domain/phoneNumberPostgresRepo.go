package domain

import (
	"errors"
	"gorm.io/gorm"
	"graph/dto"
	"graph/errs"
	"graph/logger"
)

type PhoneNumberPostgresRepo struct {
	db *gorm.DB
}

func NewPhoneNumberPostgresRepo(db *gorm.DB) *PhoneNumberPostgresRepo {
	if err := db.AutoMigrate(&PhoneNumber{}); err != nil {
		logger.Error(err.Error())
	}
	return &PhoneNumberPostgresRepo{db: db}
}

func (r *PhoneNumberPostgresRepo) AddToContact(phNumDto dto.AddPhoneNumberRequestDto) *errs.AppError {
	var contact Contact
	if err := r.db.Where("id = ?", phNumDto.ContactId).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("contact not found")
		}
		return errs.NewUnexpectedError("failed to fetch contact: " + err.Error())
	}
	phoneNumber := PhoneNumber{Number: phNumDto.Number}
	if err :=
		r.db.Model(&contact).
			Association("PhoneNumbers").
			Append(&phoneNumber); err != nil {
		return errs.NewUnexpectedError("failed to a Association: " + err.Error())
	}
	return nil
}

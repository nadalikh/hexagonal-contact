package domain

import (
	"gorm.io/gorm"
	"graph/dto"
	"graph/errs"
	"graph/logger"
)

type ContactPostgresRepo struct {
	db *gorm.DB
}

func NewContactPostgresRepo(db *gorm.DB) *ContactPostgresRepo {
	if err := db.AutoMigrate(&Contact{}); err != nil {
		logger.Error(err.Error())
	}
	return &ContactPostgresRepo{db: db}
}

func (c *ContactPostgresRepo) Create(coDto dto.ContactRequestDto) (*Contact, *errs.AppError) {
	contact := new(Contact)
	contact.FromDto(&coDto)

	for i, phNum := range coDto.PhoneNumbers {
		contact.PhoneNumbers[i] = PhoneNumber{Number: phNum.Number}
	}

	if err := c.db.Create(contact).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return contact, nil
}

func (c *ContactPostgresRepo) FindWithField(coSeaDto dto.ContactSearchRequestDto) ([]Contact, *errs.AppError) {
	var contacts []Contact
	param := "%" + coSeaDto.Param + "%"
	err := c.db.
		Joins("LEFT JOIN phone_numbers ON phone_numbers.contact_id = contacts.id").
		Where("contacts.first_name ILIKE ? OR contacts.last_name ILIKE ? OR phone_numbers.number ILIKE ?", param, param, param).
		Preload("PhoneNumbers").
		Distinct().
		Find(&contacts).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return contacts, nil
}

func (c *ContactPostgresRepo) Update(coUpdDto dto.ContactUpdateRequestDto) (*Contact, *errs.AppError) {
	tx := c.db.Begin()
	if tx.Error != nil {
		return nil, errs.NewUnexpectedError("failed to start transaction: " + tx.Error.Error())
	}

	if err := tx.Model(&Contact{}).
		Where("id = ?", coUpdDto.ContactId).
		Updates(map[string]interface{}{
			"first_name": coUpdDto.FirstName,
			"last_name":  coUpdDto.LastName,
		}).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewUnexpectedError("failed to update contact: " + err.Error())
	}

	for _, ph := range coUpdDto.PhoneNumbers {
		if err := tx.Model(&PhoneNumber{}).
			Where("id = ? AND contact_id = ?", ph.PhoneId, coUpdDto.ContactId).
			Update("number", ph.Number).Error; err != nil {
			tx.Rollback()
			return nil, errs.NewUnexpectedError("failed to update phone number: " + err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errs.NewUnexpectedError("transaction commit failed: " + err.Error())
	}

	var updatedContact Contact
	if err := c.db.Preload("PhoneNumbers").
		First(&updatedContact, "id = ?", coUpdDto.ContactId).Error; err != nil {
		return nil, errs.NewUnexpectedError("failed to retrieve updated contact: " + err.Error())
	}

	return &updatedContact, nil
}
func (r *ContactPostgresRepo) CheckPhoneExistence(number string) (bool, *errs.AppError) {
	var count int64
	if err := r.db.Model(&PhoneNumber{}).Where("number = ?", number).Count(&count).Error; err != nil {
		return false, errs.NewUnexpectedError(err.Error())
	}
	return count > 0, nil
}

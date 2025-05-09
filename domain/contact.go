package domain

import (
	"graph/dto"
	"graph/errs"
)

type Contact struct {
	BaseModel
	FirstName    string
	LastName     string
	PhoneNumbers []PhoneNumber `gorm:"foreignKey:ContactID;constraint:OnDelete:CASCADE"`
}
type ContactRepo interface {
	Create(coDto dto.ContactRequestDto) (*Contact, *errs.AppError)
	FindWithField(coDto dto.ContactSearchRequestDto) ([]Contact, *errs.AppError)
	Update(coUpdDto dto.ContactUpdateRequestDto) (*Contact, *errs.AppError)
}
type ContactServicePrototype interface {
	MakeOne(coDto dto.ContactRequestDto) (*Contact, *errs.AppError)
	Search(coSeaDto dto.ContactSearchRequestDto) ([]Contact, *errs.AppError)
	Update(cUpdDto dto.ContactUpdateRequestDto) (*Contact, *errs.AppError)
}

func (c *Contact) FromDto(coDto *dto.ContactRequestDto) {
	c.FirstName = coDto.FirstName
	c.LastName = coDto.LastName
	c.PhoneNumbers = make([]PhoneNumber, len(coDto.PhoneNumbers))
}

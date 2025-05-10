package domain

import (
	"graph/dto"
	"graph/errs"
)

type PhoneNumber struct {
	BaseModel
	Number    string
	Contact   *Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ContactID string
}
type PhoneNumberRepo interface {
	AddToContact(dto dto.AddPhoneNumberRequestDto) *errs.AppError
}
type PhoneNumberServicePrototype interface {
	AddOne(coDto dto.AddPhoneNumberRequestDto) *errs.AppError
}

func (p *PhoneNumber) ToDto() dto.PhoneNumberResponseDto {
	return dto.PhoneNumberResponseDto{
		BaseDtoResponse: dto.BaseDtoResponse{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			DeletedAt: p.DeletedAt,
		},
		Number: p.Number,
	}
}

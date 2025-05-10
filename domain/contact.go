package domain

import (
	"fmt"
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
	CheckPhoneExistence(number string) (bool, *errs.AppError)
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
func (c *Contact) ToDto() dto.ContactResponseDto {
	phoneNumbers := make([]dto.PhoneNumberResponseDto, 0)
	for _, phoneNumber := range c.PhoneNumbers {
		fmt.Println("phonen:", phoneNumber.ID)
		phoneNumbers = append(phoneNumbers, phoneNumber.ToDto())
	}
	return dto.ContactResponseDto{
		BaseDtoResponse: dto.BaseDtoResponse{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: c.DeletedAt,
		},
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		PhoneNumbers: phoneNumbers,
	}
}
func ListToDto(contacts []Contact) []dto.ContactResponseDto {
	fmt.Println("len of contancts: ", len(contacts))
	dtoContacts := make([]dto.ContactResponseDto, 0)
	for _, contact := range contacts {
		fmt.Println("id list to do", contact.ID)
		dtoContacts = append(dtoContacts, contact.ToDto())
	}
	fmt.Println("len of res: ", len(dtoContacts))
	return dtoContacts
}

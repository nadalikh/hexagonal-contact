package dto

type ContactRequestDto struct {
	FirstName    string                  `json:"first_name"`
	LastName     string                  `json:"last_name"`
	PhoneNumbers []PhoneNumberRequestDto `json:"phone_numbers" binding:"required,dive,required"`
}
type ContactSearchRequestDto struct {
	Param string `json:"param"`
}
type ContactUpdateRequestDto struct {
	ContactId    string                        `json:"contact_id"`
	PhoneNumbers []PhoneNumberUpdateRequestDto `json:"phone_numbers" binding:"dive"`
	FirstName    string                        `json:"first_name"`
	LastName     string                        `json:"last_name"`
}
type ContactResponseDto struct {
	BaseDtoResponse
	FirstName    string                   `json:"first_name"`
	LastName     string                   `json:"last_name"`
	PhoneNumbers []PhoneNumberResponseDto `json:"phone_numbers"`
}

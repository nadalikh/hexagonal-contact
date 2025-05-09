package dto

type ContactRequestDto struct {
	FirstName    string                  `json:"first_name"`
	LastName     string                  `json:"last_name"`
	PhoneNumbers []PhoneNumberRequestDto `json:"phone_numbers"`
}
type ContactSearchRequestDto struct {
	Param string `json:"param"`
}
type ContactUpdateRequestDto struct {
	ContactId    string                     `json:"contact_id"`
	PhoneNumbers []PhoneNumberUpdateRequest `json:"phone_numbers"`
	FirstName    string                     `json:"first_name"`
	LastName     string                     `json:"last_name"`
}

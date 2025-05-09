package dto

type PhoneNumberRequestDto struct {
	Number string `json:"number" binding:"required,phone"`
}
type AddPhoneNumberRequestDto struct {
	Number    string `json:"number" binding:"required,phone"`
	ContactId string `json:"contact_id"`
}
type PhoneNumberUpdateRequest struct {
	Number  string `json:"number" binding:"phone"`
	PhoneId string `json:"phone_id"`
}

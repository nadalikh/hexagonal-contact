package app

import (
	"github.com/gin-gonic/gin"
	"graph/dto"
	"graph/service"
	"net/http"
)

type PhoneNumberRestHandler struct {
	phoneNumberService *service.PhoneNumberService
}

func NewPhoneNumberRestHandler(phNumSer *service.PhoneNumberService) *PhoneNumberRestHandler {
	return &PhoneNumberRestHandler{phoneNumberService: phNumSer}
}

func (phNumResHan *PhoneNumberRestHandler) AddToContact(c *gin.Context) {
	var addPhoneNumberRequestDto dto.AddPhoneNumberRequestDto
	if err := c.ShouldBindJSON(&addPhoneNumberRequestDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := phNumResHan.phoneNumberService.AddOne(addPhoneNumberRequestDto); err != nil {
		c.AbortWithStatusJSON(err.Code, err.Message)
		return
	} else {
		c.JSON(http.StatusCreated, Response{Message: SUCCESSFULL_ADDING_PHONE})
	}
}

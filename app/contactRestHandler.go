package app

import (
	"github.com/gin-gonic/gin"
	"graph/dto"
	"graph/service"
	"net/http"
)

type ContactRestHandler struct {
	contactService *service.ContactService
}

func NewContactRestHandler(coSer *service.ContactService) *ContactRestHandler {
	return &ContactRestHandler{contactService: coSer}
}

func (coResHan *ContactRestHandler) CreateOne(c *gin.Context) {
	var contactRequest dto.ContactRequestDto
	if err := c.ShouldBindJSON(&contactRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: err.Error()})
		return
	}
	if contact, err := coResHan.contactService.MakeOne(contactRequest); err != nil {
		c.AbortWithStatusJSON(err.Code, Response{Message: err.Message})
		return
	} else {
		c.JSON(http.StatusCreated, Response{Message: SUCCESSFULL_ADDING_CONATACT, Data: contact})
	}
}
func (coResHan *ContactRestHandler) Search(c *gin.Context) {
	contactSearchReq := dto.ContactSearchRequestDto{Param: c.Query("param")}
	if contact, err := coResHan.contactService.Search(contactSearchReq); err != nil {
		c.AbortWithStatusJSON(err.Code, Response{Message: err.Message})
		return
	} else {
		c.JSON(http.StatusOK, Response{Data: contact})
	}
}

func (coResHan *ContactRestHandler) Update(c *gin.Context) {
	var contactRequest dto.ContactUpdateRequestDto
	if err := c.ShouldBindJSON(&contactRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: err.Error()})
		return
	}
	if contact, err := coResHan.contactService.Update(contactRequest); err != nil {
		c.AbortWithStatusJSON(err.Code, Response{Message: err.Message})
		return
	} else {
		c.JSON(http.StatusOK, Response{Message: SUCCESSFULL_UPDATE_CONATACT, Data: contact})
	}
}

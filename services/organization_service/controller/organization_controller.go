package controller

import (
	"net/http"
	"organization_service/model"
	"organization_service/saga"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	s *saga.OrganizationSaga
}

func NewOrganizationController(s *saga.OrganizationSaga) *OrganizationController {
	return &OrganizationController{s: s}
}

func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
	var organization *OrganizationDTO
	if err := ctx.ShouldBind(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o := model.Organization{Name: organization.OrganizationName}
	u := model.User{Name: organization.UserName, Email: organization.Email}

	if err := c.s.Execute(o, u); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute organization saga"})
		return
	}

}

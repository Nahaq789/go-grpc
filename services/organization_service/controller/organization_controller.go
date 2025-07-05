package controller

import (
	"net/http"
	"organization_service/model"
	"organization_service/repository"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	r *repository.OrganizationRepository
}

func NewOrganizationController(r *repository.OrganizationRepository) *OrganizationController {
	return &OrganizationController{r: r}
}

func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
	var organization *OrganizationDTO
	if err := ctx.ShouldBind(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o := model.Organization{Name: organization.OrganizationName}
	u := model.User{Name: organization.UserName, Email: organization.Email}
	if err := c.r.CreateOrganization(o, u); err != nil {
		return
	}
}

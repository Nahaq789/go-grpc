package controller

import (
	"organization_service/model"
	"organization_service/repository"
)

type OrganizationController struct {
	r *repository.OrganizationRepository
}

func NewOrganizationController(r *repository.OrganizationRepository) *OrganizationController {
	return &OrganizationController{r: r}
}

func (c *OrganizationController) CreateOrganization(name string) error {
	o := &model.Organization{Name: name}
	if err := c.r.CreateOrganization(o); err != nil {
		return err
	}
	return nil
}

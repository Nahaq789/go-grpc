package controller

type OrganizationDTO struct {
	OrganizationName string `json:"organization_name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	UserName         string `json:"user_name" binding:"required"`
}

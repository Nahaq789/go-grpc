package saga

import (
	"log"
	"organization_service/model"
	"organization_service/repository"
)

type OrganizationSaga struct {
	c *repository.UserServiceClient
}

func NewOrganizationSaga(c *repository.UserServiceClient) *OrganizationSaga {
	return &OrganizationSaga{
		c: c,
	}
}

func (s *OrganizationSaga) ExecuteCreateUser(u model.User) error {
	err := s.c.CallUserServiceCreateUser(u)
	if err != nil {
		log.Printf("Error creating user in user service: %v", err)
		return err
	}
	return nil
}

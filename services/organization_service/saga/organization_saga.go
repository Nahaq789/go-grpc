package saga

import (
	"log"
	"organization_service/model"
	"organization_service/repository"
)

type OrganizationSaga struct {
	r *repository.OrganizationRepository
	c *repository.UserServiceClient
}

func NewOrganizationSaga(r *repository.OrganizationRepository, c *repository.UserServiceClient) *OrganizationSaga {
	return &OrganizationSaga{
		r: r,
		c: c,
	}
}

func (s *OrganizationSaga) Execute(o model.Organization, u model.User) error {
	id, err := s.r.CreateOrganization(o)
	if err != nil {
		log.Println("failed to create organization")
		return err
	}

	if err := s.c.CallUserServiceCreateUser(u); err != nil {
		log.Println("failed to create user")
		log.Println("rollback organization create")
		if compErr := s.CompCreateOrganization(id); compErr != nil {
			log.Println("failed to rollback organization")
			return compErr
		}
		return err
	}

	return nil
}

func (s *OrganizationSaga) CompCreateOrganization(id int64) error {
	err := s.r.CompCreateOrganization(id)
	if err != nil {
		log.Printf("Error rolling back organization creation: %v", err)
		return err
	}
	log.Println("Organization creation rolled back successfully")
	return nil
}

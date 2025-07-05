package repository

import (
	"context"
	"fmt"
	"organization_service/model"
	"organization_service/proto"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db     *gorm.DB
	client proto.UserServiceClient
}

func NewOrganizationRepository(db *gorm.DB, client proto.UserServiceClient) *OrganizationRepository {
	return &OrganizationRepository{db: db, client: client}
}

func (r *OrganizationRepository) CreateOrganization(o model.Organization, u model.User) error {
	if err := r.db.Debug().Create(o).Error; err != nil {
		return err
	}
	r.callUserServiceCreateUser(u)
	return nil
}

func (r *OrganizationRepository) callUserServiceCreateUser(u model.User) error {
	req := &proto.CreateUserRequest{
		Name:  u.Name,
		Email: u.Email,
	}
	res, err := r.client.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Printf("User created with ID: %s\n", res.Message)
	return nil
}

package repository

import (
	"organization_service/model"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) CreateOrganization(o model.Organization, u model.User) error {
	if err := r.db.Debug().Create(o).Error; err != nil {
		return err
	}
	return nil
}

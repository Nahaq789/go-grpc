package repository

import (
	"log"
	"organization_service/model"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) CreateOrganization(o model.Organization) (int64, error) {
	if err := r.db.Create(&o).Error; err != nil {
		return 0, err
	}
	log.Printf("organization id: %v \n", o.Id)
	return o.Id, nil
}

func (r *OrganizationRepository) CompCreateOrganization(id int64) error {
	if err := r.db.Delete(&model.Organization{}, id).Error; err != nil {
		return err
	}
	log.Println("Organization deleted successfully, rolling back user creation")
	return nil
}

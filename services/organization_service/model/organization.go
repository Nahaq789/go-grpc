package model

type Organization struct {
	Name string `json:"name" gorm:"size:100;not null"`
}

func NewOrganization(name string) *Organization {
	return &Organization{Name: name}
}

package model

type Organization struct {
	Id   int64
	Name string `json:"name" gorm:"size:100;not null"`
}

func NewOrganization(name string) *Organization {
	return &Organization{Name: name}
}

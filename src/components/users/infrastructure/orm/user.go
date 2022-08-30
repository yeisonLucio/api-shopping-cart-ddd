package user_orm

import (
	"encoding/json"

	"gorm.io/gorm"
)

type UserGorm struct {
	gorm.Model
	Email        string `json:"email" gorm:"size:50;uniqueIndex:idx_email"`
	Password     string `json:"password" gorm:"size:250"`
	FirstName    string `json:"first_name" gorm:"size:50"`
	LastName     string `json:"last_name" gorm:"size:50"`
	Phone        string `json:"phone" gorm:"size:50;uniqueIndex:idx_phone"`
	Token        string `json:"token" gorm:"size:250"`
	RefreshToken string `json:"refresh_token" gorm:"size:250"`
	UserType     string `json:"user_type" gorm:"size:50"`
}

func (p *UserGorm) Parse(structure interface{}) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &structure)
}

func (UserGorm) TableName() string {
	return "users"
}

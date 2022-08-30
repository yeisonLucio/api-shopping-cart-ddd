package user_infrastructure

import (
	"errors"

	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	user_orm "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/orm"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
	"gorm.io/gorm"
)

type MysqlUser struct {
	Db *gorm.DB
}

func NewMysqlUser(db *gorm.DB) *MysqlUser {
	return &MysqlUser{db}
}

func (mu *MysqlUser) Create(user user_domain.User) (user_domain.User, error) {

	var userGorm user_orm.UserGorm

	if err := user.Parse(&userGorm); err != nil {
		return user, err
	}

	if result := connections.DB.Create(&userGorm); result.Error != nil {
		return user, result.Error
	}

	user.Password = ""
	user.ID = userGorm.ID

	return user, nil
}

func (mu *MysqlUser) FindById(id string) (user_domain.User, error) {

	var user user_domain.User

	if err := connections.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	user.Password = ""

	return user, nil
}

func (mu *MysqlUser) Exists(column string, value string) bool {

	var user user_orm.UserGorm

	result := connections.DB.First(&user, column+" = ?", value)

	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (mu *MysqlUser) Update(user user_domain.User) error {

	var userToUpdate user_orm.UserGorm

	if err := user.Parse(&userToUpdate); err != nil {
		return err
	}

	if err := connections.DB.UpdateColumns(&userToUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (mu *MysqlUser) FindByUsername(username string) (user_domain.User, error) {

	var user user_domain.User

	err := connections.DB.
		Where("email = ?", username).
		Or("phone = ?", username).
		First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

package mysql

import (
	"Software-Market-Go-API/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (u *MySQLDB) GetUserById(id string) (*entity.User, error) {

	user := new(entity.User)
	if err := u.Db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, errors.New("record not found")
		}

		return nil, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil
}

func (u *MySQLDB) GetUserByName(name string) (*entity.User, error) {

	user := new(entity.User)
	if err := u.Db.Where("name = ?", name).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, errors.New("record not found")
		}

		return nil, fmt.Errorf("unexpected error: " + err.Error())
	}

	return user, nil
}

func (u *MySQLDB) SaveUser(id string, user *entity.User) error {

	if id == "" {
		if err := u.Db.Create(user).Error; err != nil {

			return fmt.Errorf("unexpected error: %w", err)
		}
	}

	if err := u.Db.Save(user).Error; err != nil {

		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

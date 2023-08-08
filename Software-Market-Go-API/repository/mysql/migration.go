package mysql

import (
	"Software-Market-Go-API/entity"
	"fmt"
)

func (m *MySQLDB) Migrate() {

	err := m.Db.AutoMigrate(&entity.Software{})
	if err != nil {

		panic(fmt.Errorf("can not migrate software : %w", err))
	}

	err = m.Db.AutoMigrate(&entity.Rate{})
	if err != nil {

		panic(fmt.Errorf("can not migrate  rate : %w", err))
	}

	err = m.Db.AutoMigrate(&entity.Review{})
	if err != nil {

		panic(fmt.Errorf("can not migrate review : %w", err))
	}

	err = m.Db.AutoMigrate(&entity.User{})
	if err != nil {

		panic(fmt.Errorf("can not migrate  user : %w", err))
	}

}

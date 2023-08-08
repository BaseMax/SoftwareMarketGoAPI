package mysql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	Db *gorm.DB
}

func ConnectDatabase() *MySQLDB {

	var err error

	err = godotenv.Load("local.env")
	if err != nil {

		panic(fmt.Errorf("some error occured. Err: %s", err))
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("Failed to connect to database")
	}

	// admin := new(entity.User)
	// admin.ID = uuid.New()
	// admin.Name = "admin1"
	// admin.Email = "admin1@gmail.com"
	// admin.Password, _ = hashpassword.HashPassword("admin123")
	// admin.IsAdmin = true
	// admin.CreatedAt = time.Now()

	// err = DB.Create(admin).Error
	// if err != nil {
	// 	panic("Failed create admin")
	// }

	return &MySQLDB{Db: DB}
}

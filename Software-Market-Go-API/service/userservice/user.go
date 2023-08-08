package userservice

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/pkg/hashpassword"
	jwt2 "Software-Market-Go-API/pkg/jwt"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserById(id string) (*entity.User, error)
	SaveUser(id string, user *entity.User) error
	GetUserByName(name string) (*entity.User, error)
}

type UserService struct {
	repo UserRepo
}

func New(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) IsAdmin(claims jwt.MapClaims) bool {

	m := make(map[string]string, 0)

	m = jwt2.GetClaimsValue(claims)

	if m["role"] != "admin" {

		return false
	}

	user := new(entity.User)
	user, err := u.repo.GetUserById(m["userId"])
	if err != nil {

		return false
	}

	if !user.IsAdmin {

		return false
	}

	return true
}

func (u *UserService) CreateUser(r *RgisterRequest) (*entity.User, error) {

	user := new(entity.User)

	user.ID = uuid.New()
	user.Name = r.Name
	user.Email = r.Email
	user.Password, _ = hashpassword.HashPassword(r.Password)
	user.IsAdmin = r.IsAdmin
	user.CreatedAt = time.Now()

	if err := u.repo.SaveUser("", user); err != nil {

		return nil, fmt.Errorf("can not save user : %w", err)
	}

	return user, nil
}

func (u *UserService) IsUserExist(id string) (bool, error) {

	_, err := u.repo.GetUserById(id)
	if err != nil {

		return false, err
	}

	return true, nil
}

func (u *UserService) UpdateUser(id string, r *RgisterRequest) (*entity.User, error) {

	user := new(entity.User)

	user, err := u.repo.GetUserById(id)
	if err != nil {

		return nil, err
	}

	user.Name = r.Name
	user.Email = r.Email
	user.Password, _ = hashpassword.HashPassword(r.Password)
	user.IsAdmin = r.IsAdmin
	user.UpdatedAt = time.Now()

	if err := u.repo.SaveUser(id, user); err != nil {

		return nil, fmt.Errorf("can not save user : %w", err)
	}

	return user, nil
}

func (u *UserService) LoadUserByName(name string) (*entity.User, error) {

	user := new(entity.User)

	user, err := u.repo.GetUserByName(name)
	if err != nil {
		return nil, err

	}

	return user, nil
}

func (u *UserService) CheckPassword(pass, hashPass string) error {

	if !hashpassword.CheckPasswordHash(pass, hashPass) {

		return fmt.Errorf("user name or password is incorrect")
	}

	return nil
}

func (u *UserService) LoadUserById(id string) (*entity.User, error) {

	user := new(entity.User)

	user, err := u.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

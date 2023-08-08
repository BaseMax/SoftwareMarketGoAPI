package panel

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/pkg/customresponse"
	"Software-Market-Go-API/pkg/jwt"
	"Software-Market-Go-API/service/userservice"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func (h Handler) LoginAdmin(ctx echo.Context) error {

	lr := new(userservice.LoginRequest)
	if err := ctx.Bind(lr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := lr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	user, err := h.userSvc.LoadUserByName(lr.Name)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := h.userSvc.CheckPassword(lr.Password, user.Password); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	token, err := jwt.GenerateJWTAccessToken(user.ID.String(), user.IsAdmin)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "Login successfully"
	resp.Result = token

	return ctx.JSON(http.StatusOK, resp)
}

func (h Handler) UpdateAdmin(ctx echo.Context) error {

	err := godotenv.Load("local.env")
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("some error occured. Err: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	tokenString := ctx.Get("token-string").(string)

	claims, err := jwt.VerifyToken(tokenString, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("invalid token : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if !h.userSvc.IsAdmin(claims) {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = "Access denied"
		resp.Result = nil

		return ctx.JSON(http.StatusForbidden, resp)
	}

	id := ctx.Param("id")

	isExist, err := h.userSvc.IsUserExist(id)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if !isExist {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = "user not found"
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	rr := new(userservice.RgisterRequest)
	if err := ctx.Bind(rr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := rr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	user := new(entity.User)
	if user, err = h.userSvc.UpdateUser(id, rr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not update new admin : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "updated successfully"
	resp.Result = user

	return ctx.JSON(http.StatusCreated, resp)
}

func (h Handler) AddAdmin(ctx echo.Context) error {

	err := godotenv.Load("local.env")
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("some error occured. Err: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	tokenString := ctx.Get("token-string").(string)

	claims, err := jwt.VerifyToken(tokenString, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("invalid token : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if !h.userSvc.IsAdmin(claims) {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = "Access denied"
		resp.Result = nil

		return ctx.JSON(http.StatusForbidden, resp)
	}

	rr := new(userservice.RgisterRequest)
	if err := ctx.Bind(rr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err = rr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	user := new(entity.User)
	user, err = h.userSvc.CreateUser(rr)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not create new user : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "add successfully"
	resp.Result = user

	return ctx.JSON(http.StatusCreated, resp)
}

package api

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/pkg/customresponse"
	"Software-Market-Go-API/pkg/jwt"
	"Software-Market-Go-API/service/softwareservice"
	"Software-Market-Go-API/service/userservice"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// GET /users/:id/software
// Retrieves all software listings associated with a specific user. The :id parameter represents the ID of the user.
func (h Handler) GetUsersSoftwareByAssociated(ctx echo.Context) error {

	return nil
}

//-------------------------

// GET /users/:id/reviews
// Retrieves all reviews posted by a specific user. The :id parameter represents the ID of the user.
func (h Handler) GetUsersReviews(ctx echo.Context) error {

	sort := ctx.QueryParam("sort")
	limit := ctx.QueryParam("limit")

	pfilter, err := softwareservice.MakePublicFilter(sort, limit)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("invalid filter value: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	err = godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
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

	claimValue := make(map[string]string)
	claimValue = jwt.GetClaimsValue(claims)

	isExist, err := h.userSvc.IsUserExist(claimValue["userId"])
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

	reviews := make(entity.Reviews, 0)
	reviews, err = h.softwareSvc.LoadReviewsByUserId(claimValue["userId"], pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)

	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load reviews successfully"
	resp.Result = reviews

	return ctx.JSON(http.StatusOK, resp)
}

//-------------------------

func (h Handler) Register(ctx echo.Context) error {

	rr := new(userservice.RgisterRequest)
	if err := ctx.Bind(rr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	rr.IsAdmin = false

	if err := rr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	u := new(entity.User)
	u, err := h.userSvc.LoadUserByName(rr.Name)
	if err != nil && err.Error() != "record not found" {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if u != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = "user is already registered"
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
	resp.Message = "registered user successfully"
	resp.Result = user

	return ctx.JSON(http.StatusOK, resp)

}

func (h Handler) Login(ctx echo.Context) error {

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

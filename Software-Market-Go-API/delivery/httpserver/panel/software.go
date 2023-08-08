package panel

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/pkg/customresponse"
	"Software-Market-Go-API/pkg/jwt"
	"Software-Market-Go-API/service/softwareservice"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// POST /software
// Creates a new software listing in the marketplace. Requires the following JSON fields in the request body:

// name (string): Name of the software.
// description (string): Description of the software.
// price (float): Price of the software.
func (h Handler) AddSoftware(ctx echo.Context) error {

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

	rsr := new(softwareservice.RgisterSoftwareRequest)
	if err := ctx.Bind(rsr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := rsr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	software := new(entity.Software)
	software, err = h.softwareSvc.CreateSoftware(rsr)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not create new software : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "create software successfully"
	resp.Result = software

	return ctx.JSON(http.StatusCreated, resp)
}

//---------------------------------------------------------
// PUT /software/:id
// Updates an existing software listing in the marketplace. Requires the following JSON fields in the request body:

// name (string): Updated name of the software.
// description (string): Updated description of the software.
// price (float): Updated price of the software.
func (h Handler) UpdateSoftware(ctx echo.Context) error {

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

	isExist, err := h.softwareSvc.IsSofwareExist(id)
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
		resp.Message = "software not found"
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	rsr := new(softwareservice.RgisterSoftwareRequest)
	if err := ctx.Bind(rsr); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := rsr.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("validation error : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	software := new(entity.Software)
	software, err = h.softwareSvc.UpdateSoftware(id, rsr)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not update new software : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "update software successfully"
	resp.Result = software

	return ctx.JSON(http.StatusCreated, resp)
}

// ----------------------------------------------------------------

// DELETE /software/:id
// Deletes a software listing from the marketplace based on its ID.
func (h Handler) DeleteSoftware(ctx echo.Context) error {

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

	isExist, err := h.softwareSvc.IsSofwareExist(id)
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
		resp.Message = "software not found"
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := h.softwareSvc.RemoveSoftware(id); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not remove software : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "remove software successfully"
	resp.Result = nil

	return ctx.JSON(http.StatusOK, resp)
}

//------------------------------------------------------------------

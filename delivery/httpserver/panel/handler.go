package panel

import (
	"Software-Market-Go-API/service/softwareservice"
	"Software-Market-Go-API/service/userservice"
)

type Handler struct {
	userSvc     userservice.UserService
	softwareSvc softwareservice.SoftwareService
}

func NewHandler(userSvc userservice.UserService, softwareSvc softwareservice.SoftwareService) Handler {
	return Handler{
		userSvc:     userSvc,
		softwareSvc: softwareSvc,
	}
}

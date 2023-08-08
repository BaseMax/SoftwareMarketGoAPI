package httpserver

import (
	"Software-Market-Go-API/delivery/httpserver/api"
	"Software-Market-Go-API/delivery/httpserver/panel"

	"github.com/labstack/echo/v4"
)

func Routes(apiHandler api.Handler, panelHandler panel.Handler, e *echo.Echo) {

	panelHandler.PanelRoute(e)
	apiHandler.ApiRoute(e)
}

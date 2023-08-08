package panel

import "github.com/labstack/echo/v4"

func (h Handler) PanelRoute(a *echo.Echo) {

	a1 := a.Group("/panel") // TODO add middleware if needed

	as := a1.Group("/software", GetTokenFromHeader)
	as.POST("/", h.AddSoftware)
	as.PUT("/:id", h.UpdateSoftware)
	as.DELETE("/:id", h.DeleteSoftware)

	au := a1.Group("/users")
	au.POST("/admin", h.AddAdmin, GetTokenFromHeader)
	au.POST("/login", h.LoginAdmin)
	au.PUT("/admin/:id", h.UpdateAdmin, GetTokenFromHeader)

}

package api

import "github.com/labstack/echo/v4"

func (h Handler) ApiRoute(a *echo.Echo) {

	a1 := a.Group("/api") // TODO add middleware if needed

	as := a1.Group("/software")
	as.GET("/", h.GetSoftwares)
	as.GET("/:id", h.GetSoftware)
	as.GET("/category/:category", h.GetSoftwaresByCategory)
	as.GET("/search/:query", h.SearchSoftware)
	as.GET("/top-rated", h.GetTopRatedSoftware)
	as.GET("/recommended", h.GetRecommendedSoftware, GetTokenFromHeader)
	as.GET("/:id/reviews", h.GetSoftwareReviews)
	as.POST("/:id/ratings", h.AddRateToSoftware, GetTokenFromHeader)
	as.POST("/:id/reviews", h.AddReviewToSoftware, GetTokenFromHeader)

	au := a1.Group("/users")
	au.GET("/reviews", h.GetUsersReviews, GetTokenFromHeader) // change /:id/reviews to /reviews get user id from jwt token
	au.GET("/:id/software", h.GetUsersSoftwareByAssociated)
	au.POST("/login", h.Login)
	au.POST("/register", h.Register)
}

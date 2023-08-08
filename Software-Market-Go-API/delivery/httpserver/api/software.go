package api

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/pkg/customresponse"
	"Software-Market-Go-API/pkg/jwt"
	"Software-Market-Go-API/service/softwareservice"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// GET /software
// Retrieves all software available in the marketplace.
func (h Handler) GetSoftwares(ctx echo.Context) error {

	sort := ctx.QueryParam("sort")
	limit := ctx.QueryParam("limit")

	pfilter, err := softwareservice.MakePublicFilter(sort, limit)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("Invalid filter value :  %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	softwares := make(entity.Softwares, 0)
	softwares, err = h.softwareSvc.LoadSoftwares(pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load softwares successfully"
	resp.Result = softwares

	return ctx.JSON(http.StatusOK, resp)
}

//---------------------------

// GET /software/:id
// Retrieves a specific software by its ID.
func (h Handler) GetSoftware(ctx echo.Context) error {

	id := ctx.Param("id")

	software := new(entity.Software)
	software, err := h.softwareSvc.LoadSoftwareById(id)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load softwares successfully"
	resp.Result = software

	return ctx.JSON(http.StatusOK, resp)
}

//---------------------------

// GET /software/category/:category
// Retrieves all software listings in a specific category. The :category parameter represents the category of the software.
func (h Handler) GetSoftwaresByCategory(ctx echo.Context) error {

	category := ctx.Param("category")
	sort := ctx.QueryParam("sort")
	limit := ctx.QueryParam("limit")

	pfilter, err := softwareservice.MakePublicFilter(sort, limit)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("Invalid filter value: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	softwares := make(entity.Softwares, 0)
	softwares, err = h.softwareSvc.LoadSoftwaresByCategory(category, pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load softwares successfully"
	resp.Result = softwares

	return ctx.JSON(http.StatusOK, resp)
}

//---------------------------

// GET /software/search?query=:searchQuery
// Searches for software listings based on a given search query. The :searchQuery parameter represents the search query string.
func (h Handler) SearchSoftware(ctx echo.Context) error {

	searchKey := ctx.Param("query")
	fmt.Println(searchKey)
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

	softwares := make(entity.Softwares, 0)
	softwares, err = h.softwareSvc.LoadSoftwaresBySearchKey(searchKey, pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load softwares successfully"
	resp.Result = softwares

	return ctx.JSON(http.StatusOK, resp)
}

//---------------------------

// GET /software/top-rated
// Retrieves the top-rated software listings in the marketplace based on user ratings.
func (h Handler) GetTopRatedSoftware(ctx echo.Context) error {

	limit := ctx.QueryParam("limit")

	pfilter, err := softwareservice.MakePublicFilter("", limit)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("invalid filter value: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	softwares := make(entity.Softwares, 0)
	softwares, err = h.softwareSvc.LoadTopRateSoftwares(pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load softwares successfully"
	resp.Result = softwares

	return ctx.JSON(http.StatusOK, resp)
}

// --------------------------
// // this func needs to curren user id
// GET /software/recommended
// Retrieves a list of recommended software listings for the current user. The recommendations can be based on user preferences, previous purchases, or any other relevant criteria.
func (h Handler) GetRecommendedSoftware(ctx echo.Context) error {

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

	user := new(entity.User)
	user, err = h.userSvc.LoadUserById(claimValue["userId"])
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	favoriteCategories := user.FavoriteCategories

	rates := make(entity.Rates, 0)
	rates, err = h.softwareSvc.LoadRatesByUserId(user.ID.String(), pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	topRateSoftwareCategories := make([]string, 0)

	for _, rate := range rates {

		software := new(entity.Software)
		software, err := h.softwareSvc.LoadSoftwareById(rate.SoftwareId.String())
		if err != nil {

			resp := new(customresponse.Response)
			resp.Success = false
			resp.Message = err.Error()
			resp.Result = nil

			return ctx.JSON(http.StatusBadRequest, resp)
		}

		topRateSoftwareCategories = append(topRateSoftwareCategories, software.Category)
	}

	topratecategories := strings.Join(topRateSoftwareCategories, ",")
	allTop := topratecategories + favoriteCategories

	allTopcategory := strings.Split(allTop, ",")

	counts := make(map[string]int)
	var topString string
	var maxCount int

	for _, str := range allTopcategory {
		counts[str]++
		if counts[str] > maxCount {
			maxCount = counts[str]
			topString = str
		}
	}

	softwares := make(entity.Softwares, 0)
	softwares, err = h.softwareSvc.LoadSoftwaresByCategory(topString, pfilter)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "load recommended softwares "
	resp.Result = softwares

	return ctx.JSON(http.StatusBadRequest, resp)
}

// ---------------------------

// POST /software/:id/ratings
// Adds a user rating for a specific software listing. Requires the following JSON fields in the request body:
// rating (integer): The user's rating for the software (e.g., a value between 1 and 5).
func (h Handler) AddRateToSoftware(ctx echo.Context) error {

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

	rateReq := new(softwareservice.RateRequersForm)
	if err := ctx.Bind(rateReq); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind: %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := rateReq.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
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

	rateReq.UserId = claimValue["userId"]

	rateReq.SoftwareId = id

	rate := new(entity.Rate)
	rate, err = h.softwareSvc.AddRate(rateReq)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "rated successfully"
	resp.Result = rate

	return ctx.JSON(http.StatusCreated, resp)
}

//---------------------------

// GET /software/:id/reviews
// Retrieves all reviews for a specific software listing.
func (h Handler) GetSoftwareReviews(ctx echo.Context) error {

	id := ctx.Param("id")
	sort := ctx.QueryParam("sort")
	limit := ctx.QueryParam("limit")

	pfilter, err := softwareservice.MakePublicFilter(sort, limit)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("invalid filter value : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

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

	reviews := make(entity.Reviews, 0)
	reviews, err = h.softwareSvc.LoadReviews(id, pfilter)
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

// ---------------------------

// POST /software/:id/reviews
// Adds a new review for a specific software listing. Requires the following JSON fields in the request body:
//
//	title (string): Title of the review.
//	content (string): Content of the review.
//	user_id (integer): ID of the user who posted the review.
func (h Handler) AddReviewToSoftware(ctx echo.Context) error {

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

	reviewReq := new(softwareservice.ReviewRequersForm)
	if err := ctx.Bind(reviewReq); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = fmt.Errorf("can not bind : %w", err).Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err := reviewReq.Validate(); err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
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
	reviewReq.UserId = claimValue["userId"]
	reviewReq.SoftwareId = id

	review := new(entity.Review)
	review, err = h.softwareSvc.AddReview(reviewReq)
	if err != nil {

		resp := new(customresponse.Response)
		resp.Success = false
		resp.Message = err.Error()
		resp.Result = nil

		return ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := new(customresponse.Response)
	resp.Success = true
	resp.Message = "rated successfully"
	resp.Result = review

	return ctx.JSON(http.StatusCreated, review)
}

//--------------------------

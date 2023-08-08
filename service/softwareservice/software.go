package softwareservice

import (
	"Software-Market-Go-API/entity"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SoftwareRepo interface {
	GetSoftwareById(id string) (*entity.Software, error)
	SaveSoftware(id string, software *entity.Software) error
	DeleteSoftware(id string) error
	GetSoftwares(pfilter PublicFilter) (entity.Softwares, error)
	GetSoftwaresByCategory(category string, pfilter PublicFilter) (entity.Softwares, error)
	GetSoftwaresBySearchKey(searchKey string, pfilter PublicFilter) (entity.Softwares, error)
	GetTopRateSoftwares(pfilter PublicFilter) (entity.Softwares, error)
	RateSoftware(rate *entity.Rate) error
	GetReviewsBySoftwareId(id string, pfilter PublicFilter) (entity.Reviews, error)
	ReviewSoftware(review *entity.Review) error
	GetReviewsByUserId(id string, pfilter PublicFilter) (entity.Reviews, error)
	GetRatesByUserId(id string, pfilter PublicFilter) (entity.Rates, error) // filter rates by rate value more than 3
}

type SoftwareService struct {
	repo SoftwareRepo
}

func New(repo SoftwareRepo) *SoftwareService {
	return &SoftwareService{
		repo: repo,
	}
}

func (s *SoftwareService) CreateSoftware(rsr *RgisterSoftwareRequest) (*entity.Software, error) {

	software := new(entity.Software)

	software.ID = uuid.New()
	software.Name = rsr.Name
	software.Description = rsr.Description
	software.Price = rsr.Price
	software.Creator = rsr.Creator
	software.Version = rsr.Version
	software.Category = strings.Join(rsr.Category, ",")
	software.CreatedAt = time.Now()

	if err := s.repo.SaveSoftware("", software); err != nil {

		return nil, fmt.Errorf("can not save software : %w", err)
	}

	return software, nil
}

func (s *SoftwareService) IsSofwareExist(id string) (bool, error) {

	_, err := s.repo.GetSoftwareById(id)
	if err != nil {

		return false, err
	}

	return true, nil
}

func (s *SoftwareService) UpdateSoftware(id string, rsr *RgisterSoftwareRequest) (*entity.Software, error) {

	software := new(entity.Software)

	software, err := s.repo.GetSoftwareById(id)
	if err != nil {

		return nil, err
	}

	software.Name = rsr.Name
	software.Description = rsr.Description
	software.Price = rsr.Price
	software.Creator = rsr.Creator
	software.Version = rsr.Version
	software.Category = strings.Join(rsr.Category, ",")
	software.UpdatedAt = time.Now()

	if err := s.repo.SaveSoftware(id, software); err != nil {

		return nil, fmt.Errorf("can not update software : %w", err)
	}

	return software, nil
}

func (s *SoftwareService) RemoveSoftware(id string) error {

	if err := s.repo.DeleteSoftware(id); err != nil {

		return err
	}

	return nil
}

func (s *SoftwareService) LoadSoftwares(pfilter PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)

	softwares, err := s.repo.GetSoftwares(pfilter)
	if err != nil {

		return entity.Softwares{}, err
	}

	if len(softwares) == 0 {

		return entity.Softwares{}, fmt.Errorf("not found any software")
	}

	return softwares, nil
}

func (s *SoftwareService) LoadSoftwareById(id string) (*entity.Software, error) {

	_, err := uuid.Parse(id)
	if err != nil {

		return nil, fmt.Errorf("invalid id")
	}

	software := new(entity.Software)
	software, err = s.repo.GetSoftwareById(id)
	if err != nil {

		return nil, fmt.Errorf("software not found : %w", err)
	}

	return software, nil
}

func (s *SoftwareService) LoadSoftwaresByCategory(category string, pfilter PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	softwares, err := s.repo.GetSoftwaresByCategory(category, pfilter)
	if err != nil {

		return entity.Softwares{}, fmt.Errorf("softwares not found : %w", err)
	}

	if len(softwares) == 0 {

		return entity.Softwares{}, fmt.Errorf("not found any software")
	}

	return softwares, nil
}

func (s *SoftwareService) LoadSoftwaresBySearchKey(searchKey string, pfilter PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	softwares, err := s.repo.GetSoftwaresBySearchKey(searchKey, pfilter)
	if err != nil {

		return entity.Softwares{}, fmt.Errorf("softwares not found : %w", err)
	}

	if len(softwares) == 0 {

		return entity.Softwares{}, fmt.Errorf("not found any software")
	}

	return softwares, nil
}

func (s *SoftwareService) LoadTopRateSoftwares(pfilter PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	softwares, err := s.repo.GetTopRateSoftwares(pfilter)
	if err != nil {

		return entity.Softwares{}, fmt.Errorf("softwares not found : %w", err)
	}

	if len(softwares) == 0 {

		return entity.Softwares{}, fmt.Errorf("not found any software")
	}

	return softwares, nil
}

func (s *SoftwareService) AddRate(rateReq *RateRequersForm) (*entity.Rate, error) {

	rate := new(entity.Rate)

	rate.ID = uuid.New()
	rate.Value = rateReq.Value
	rate.SoftwareId, _ = uuid.Parse(rateReq.SoftwareId)
	rate.UserId, _ = uuid.Parse(rateReq.UserId)
	rate.RatedAt = time.Now()

	if err := s.repo.RateSoftware(rate); err != nil {

		return nil, err
	}

	return rate, nil
}

func (s *SoftwareService) LoadReviews(id string, pfilter PublicFilter) (entity.Reviews, error) {

	reviews := make(entity.Reviews, 0)
	reviews, err := s.repo.GetReviewsBySoftwareId(id, pfilter)
	if err != nil {

		return entity.Reviews{}, fmt.Errorf("reviews not found : %w", err)
	}

	if len(reviews) == 0 {

		return entity.Reviews{}, fmt.Errorf("not found any reviews")
	}

	return reviews, nil
}

func (s *SoftwareService) AddReview(reviewReq *ReviewRequersForm) (*entity.Review, error) {

	review := new(entity.Review)

	review.ID = uuid.New()
	review.Title = reviewReq.Title
	review.Content = reviewReq.Content
	review.SoftwareId, _ = uuid.Parse(reviewReq.SoftwareId)
	review.UserId, _ = uuid.Parse(reviewReq.UserId)
	review.CreatedAt = time.Now()

	if err := s.repo.ReviewSoftware(review); err != nil {

		return nil, err
	}

	return review, nil
}

func (s *SoftwareService) LoadReviewsByUserId(id string, pfilter PublicFilter) (entity.Reviews, error) {

	reviews := make(entity.Reviews, 0)
	reviews, err := s.repo.GetReviewsByUserId(id, pfilter)
	if err != nil {

		return entity.Reviews{}, fmt.Errorf("reviews not found : %w", err)
	}

	if len(reviews) == 0 {

		return entity.Reviews{}, fmt.Errorf("not found any reviews")
	}

	return reviews, nil
}

func (s *SoftwareService) LoadRatesByUserId(id string, pfilter PublicFilter) (entity.Rates, error) {

	rates := make(entity.Rates, 0)
	rates, err := s.repo.GetRatesByUserId(id, pfilter)
	if err != nil {

		return entity.Rates{}, fmt.Errorf("reviews not found : %w", err)
	}

	if len(rates) == 0 {

		return entity.Rates{}, fmt.Errorf("not found any reviews")
	}

	return rates, nil
}

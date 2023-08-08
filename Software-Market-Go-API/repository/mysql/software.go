package mysql

import (
	"Software-Market-Go-API/entity"
	"Software-Market-Go-API/service/softwareservice"
	"fmt"

	"gorm.io/gorm"
)

func (s *MySQLDB) GetSoftwareById(id string) (*entity.Software, error) {

	software := new(entity.Software)
	if err := s.Db.Where("id = ?", id).First(software).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return nil, fmt.Errorf("record not found : %w", err)
		}

		return nil, fmt.Errorf("unexpected error : %w", err)
	}

	return software, nil
}

func (s *MySQLDB) SaveSoftware(id string, software *entity.Software) error {

	if id == "" {
		if err := s.Db.Create(software).Error; err != nil {

			return fmt.Errorf("unexpected error: %w", err)
		}
	}

	if err := s.Db.Save(software).Error; err != nil {

		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil

}

func (s *MySQLDB) DeleteSoftware(id string) error {

	software := new(entity.Software)
	if err := s.Db.Where("id = ?", id).Delete(software).Error; err != nil {

		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

func (s *MySQLDB) GetSoftwares(pfilter softwareservice.PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	if err := s.Db.Order(pfilter.Sort).Limit(pfilter.Limit).Find(&softwares).Error; err != nil {

		return entity.Softwares{}, fmt.Errorf("unexpected error: %w", err)
	}

	return softwares, nil
}

func (s *MySQLDB) GetSoftwaresByCategory(category string, pfilter softwareservice.PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	d := "%"
	category = fmt.Sprintf("%s%s%s", d, category, d)
	if err := s.Db.Order(pfilter.Sort).Limit(pfilter.Limit).Where("category LIKE ?", category).Find(&softwares).Error; err != nil {

		return nil, fmt.Errorf("unexpected error : %w", err)
	}

	return softwares, nil
}

func (s *MySQLDB) GetSoftwaresBySearchKey(searchKey string, pfilter softwareservice.PublicFilter) (entity.Softwares, error) {

	softwares := make(entity.Softwares, 0)
	d := "%"
	searchKeyC := fmt.Sprintf("%s%s%s", d, searchKey, d)
	if err := s.Db.Order(pfilter.Sort).Limit(pfilter.Limit).Where("name = ?", searchKey).Or("category LIKE ?", searchKeyC).Or("creator = ?", searchKey).Find(&softwares).Error; err != nil {

		return nil, fmt.Errorf("unexpected error : %w", err)
	}

	return softwares, nil
}

func (s *MySQLDB) GetTopRateSoftwares(pfilter softwareservice.PublicFilter) (entity.Softwares, error) {

	type result struct {
		Value      float64 `json:"value"` // value can be between 1 and 5
		SoftwareId string  `json:"software_id"`
	}

	type results []result

	softwares := make(entity.Softwares, 0)

	re := make(results, 0)
	if err := s.Db.Table("rates").Select("software_id ,avg(value) as value").Group("software_id").Find(&re).Error; err != nil {
		return nil, fmt.Errorf("unexpected error : %w", err)
	}

	for _, r := range re {

		software, err := s.GetSoftwareById(r.SoftwareId)
		if err != nil {
			return nil, err
		}

		softwares = append(softwares, *software)
	}

	return softwares, nil
}

func (s *MySQLDB) RateSoftware(rate *entity.Rate) error {

	if err := s.Db.Create(rate).Error; err != nil {

		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

func (s *MySQLDB) GetReviewsBySoftwareId(id string, pfilter softwareservice.PublicFilter) (entity.Reviews, error) {

	reviews := make(entity.Reviews, 0)
	if err := s.Db.Order(pfilter.Sort).Limit(pfilter.Limit).Where("software_id = ?", id).Find(&reviews).Error; err != nil {

		return entity.Reviews{}, fmt.Errorf("unexpected error: %w", err)
	}

	return reviews, nil

}

func (s *MySQLDB) ReviewSoftware(review *entity.Review) error {

	if err := s.Db.Create(review).Error; err != nil {

		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

func (s *MySQLDB) GetReviewsByUserId(id string, pfilter softwareservice.PublicFilter) (entity.Reviews, error) {

	reviews := make(entity.Reviews, 0)
	if err := s.Db.Order(pfilter.Sort).Limit(pfilter.Limit).Where("user_id = ?", id).Find(&reviews).Error; err != nil {

		return entity.Reviews{}, fmt.Errorf("unexpected error: %w", err)
	}

	return reviews, nil
}

func (s *MySQLDB) GetRatesByUserId(id string, pfilter softwareservice.PublicFilter) (entity.Rates, error) {

	rates := make(entity.Rates, 0)
	if err := s.Db.Limit(pfilter.Limit).Where("user_id = ?  AND value >= ?", id, 3).Find(&rates).Error; err != nil {

		return entity.Rates{}, fmt.Errorf("unexpected error: %w", err)
	}

	return rates, nil
} // filter rates by rate value more than 3

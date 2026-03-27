package repository

import (
	"35home/models"

	"github.com/jinzhu/gorm"
)

type VoteRepository interface {
	Create(vote *models.Vote) error
	FindByUserID(userID uint64) (*models.Vote, error)
	CountByStatus(status string) (int, error)
	CountByCity() (map[string]int, error)
	CountByIndustry() (map[string]int, error)
	CountByOccupation() (map[string]int, error)
	CountByAge() (map[string]int, error)
	CountByGender() (map[string]int, error)
	TotalVotes() (int, error)
}

type voteRepository struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepository{db: db}
}

func (r *voteRepository) Create(vote *models.Vote) error {
	return r.db.Create(vote).Error
}

func (r *voteRepository) FindByUserID(userID uint64) (*models.Vote, error) {
	var vote models.Vote
	err := r.db.Where("user_id = ?", userID).First(&vote).Error
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *voteRepository) CountByStatus(status string) (int, error) {
	var count int
	err := r.db.Model(&models.Vote{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (r *voteRepository) CountByCity() (map[string]int, error) {
	type Result struct {
		City  string
		Count int
	}

	var results []Result
	err := r.db.Model(&models.Vote{}).Select("city, count(*) as count").Where("status = ?", "unemployed").Group("city").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.City] = result.Count
	}

	return stats, nil
}

func (r *voteRepository) CountByIndustry() (map[string]int, error) {
	type Result struct {
		Industry string
		Count    int
	}

	var results []Result
	err := r.db.Model(&models.Vote{}).Select("industry, count(*) as count").Where("status = ?", "unemployed").Group("industry").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.Industry] = result.Count
	}

	return stats, nil
}

func (r *voteRepository) CountByOccupation() (map[string]int, error) {
	type Result struct {
		Occupation string
		Count      int
	}

	var results []Result
	err := r.db.Model(&models.Vote{}).Select("occupation, count(*) as count").Where("status = ?", "unemployed").Group("occupation").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.Occupation] = result.Count
	}

	return stats, nil
}

func (r *voteRepository) CountByAge() (map[string]int, error) {
	type Result struct {
		AgeGroup string
		Count    int
	}

	var results []Result
	sql := "CASE WHEN age BETWEEN 35 AND 39 THEN '35-39岁' WHEN age BETWEEN 40 AND 44 THEN '40-44岁' WHEN age BETWEEN 45 AND 49 THEN '45-49岁' ELSE '50岁以上' END as age_group, count(*) as count"
	err := r.db.Model(&models.Vote{}).Select(sql).Where("status = ?", "unemployed").Group("age_group").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.AgeGroup] = result.Count
	}

	return stats, nil
}

func (r *voteRepository) CountByGender() (map[string]int, error) {
	type Result struct {
		Gender string
		Count  int
	}

	var results []Result
	err := r.db.Model(&models.Vote{}).Select("gender, count(*) as count").Where("status = ?", "unemployed").Group("gender").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.Gender] = result.Count
	}

	return stats, nil
}

func (r *voteRepository) TotalVotes() (int, error) {
	var count int
	err := r.db.Model(&models.Vote{}).Count(&count).Error
	return count, err
}

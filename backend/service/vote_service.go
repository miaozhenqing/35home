package service

import (
	"errors"

	"35home/common"
	"35home/dto"
	"35home/models"
	"35home/repository"
)

type VoteService interface {
	SubmitVote(userID uint, req *dto.VoteRequest) (*dto.VoteResponse, error)
	GetVoteStats() (*dto.VoteStatsResponse, error)
}

type voteService struct {
	voteRepo repository.VoteRepository
}

func NewVoteService(voteRepo repository.VoteRepository) VoteService {
	return &voteService{voteRepo: voteRepo}
}

func (s *voteService) SubmitVote(userID uint, req *dto.VoteRequest) (*dto.VoteResponse, error) {
	// 检查用户是否已经投过票
	existingVote, err := s.voteRepo.FindByUserID(userID)
	if err == nil && existingVote != nil {
		return nil, errors.New(common.ErrorUserAlreadyVoted)
	}

	// 创建新投票
	vote := &models.Vote{
		UserID:     userID,
		Status:     req.Status,
		City:       req.City,
		Industry:   req.Industry,
		Occupation: req.Occupation,
		Age:        req.Age,
		Gender:     req.Gender,
	}

	if err := s.voteRepo.Create(vote); err != nil {
		return nil, errors.New(common.ErrorFailedToSubmitVote)
	}

	// 构建响应
	response := &dto.VoteResponse{
		Vote: dto.VoteInfo{
			ID:         vote.ID,
			Status:     vote.Status,
			City:       vote.City,
			Industry:   vote.Industry,
			Occupation: vote.Occupation,
			Age:        vote.Age,
			Gender:     vote.Gender,
			CreatedAt:  vote.CreatedAt,
		},
	}

	return response, nil
}

func (s *voteService) GetVoteStats() (*dto.VoteStatsResponse, error) {
	// 获取总投票数
	totalVotes, err := s.voteRepo.TotalVotes()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetTotalVotes)
	}

	// 获取工作状态投票数
	employedCount, err := s.voteRepo.CountByStatus("employed")
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetEmployedCount)
	}

	// 获取失业状态投票数
	unemployedCount, err := s.voteRepo.CountByStatus("unemployed")
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetUnemployedCount)
	}

	// 获取城市统计
	cityStats, err := s.voteRepo.CountByCity()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetCityStats)
	}

	// 获取行业统计
	industryStats, err := s.voteRepo.CountByIndustry()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetIndustryStats)
	}

	// 获取职业统计
	occupationStats, err := s.voteRepo.CountByOccupation()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetOccupationStats)
	}

	// 获取年龄统计
	ageStats, err := s.voteRepo.CountByAge()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetAgeStats)
	}

	// 获取性别统计
	genderStats, err := s.voteRepo.CountByGender()
	if err != nil {
		return nil, errors.New(common.ErrorFailedToGetGenderStats)
	}

	// 构建响应
	response := &dto.VoteStatsResponse{
		TotalVotes:      totalVotes,
		EmployedCount:   employedCount,
		UnemployedCount: unemployedCount,
		CityStats:       cityStats,
		IndustryStats:   industryStats,
		OccupationStats: occupationStats,
		AgeStats:        ageStats,
		GenderStats:     genderStats,
	}

	return response, nil
}

package dto

import "time"

type VoteRequest struct {
	Status     string `json:"status" binding:"required,oneof=employed unemployed"`
	City       string `json:"city" binding:"required"`
	Industry   string `json:"industry" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Age        int    `json:"age" binding:"required,min=18,max=80"`
	Gender     string `json:"gender" binding:"required,oneof=male female"`
}

type VoteResponse struct {
	Vote VoteInfo `json:"vote"`
}

type VoteInfo struct {
	ID         uint      `json:"id"`
	Status     string    `json:"status"`
	City       string    `json:"city"`
	Industry   string    `json:"industry"`
	Occupation string    `json:"occupation"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	CreatedAt  time.Time `json:"createdAt"`
}

type VoteStatsResponse struct {
	TotalVotes      int            `json:"totalVotes"`
	EmployedCount   int            `json:"employedCount"`
	UnemployedCount int            `json:"unemployedCount"`
	CityStats       map[string]int `json:"cityStats"`
	IndustryStats   map[string]int `json:"industryStats"`
	OccupationStats map[string]int `json:"occupationStats"`
	AgeStats        map[string]int `json:"ageStats"`
	GenderStats     map[string]int `json:"genderStats"`
}

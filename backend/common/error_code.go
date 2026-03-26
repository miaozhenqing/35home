package common

// 错误码定义
const (
	// 通用错误
	ErrorInternalServer = "INTERNAL_SERVER_ERROR"
	ErrorInvalidRequest = "INVALID_REQUEST"
	
	// 用户相关错误
	ErrorEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	ErrorInvalidEmailOrPassword = "INVALID_EMAIL_OR_PASSWORD"
	ErrorPasswordsDoNotMatch = "PASSWORDS_DO_NOT_MATCH"
	ErrorFailedToHashPassword = "FAILED_TO_HASH_PASSWORD"
	ErrorFailedToCreateUser = "FAILED_TO_CREATE_USER"
	
	// 投票相关错误
	ErrorUserAlreadyVoted = "USER_ALREADY_VOTED"
	ErrorFailedToSubmitVote = "FAILED_TO_SUBMIT_VOTE"
	ErrorFailedToGetTotalVotes = "FAILED_TO_GET_TOTAL_VOTES"
	ErrorFailedToGetEmployedCount = "FAILED_TO_GET_EMPLOYED_COUNT"
	ErrorFailedToGetUnemployedCount = "FAILED_TO_GET_UNEMPLOYED_COUNT"
	ErrorFailedToGetCityStats = "FAILED_TO_GET_CITY_STATS"
	ErrorFailedToGetIndustryStats = "FAILED_TO_GET_INDUSTRY_STATS"
	ErrorFailedToGetOccupationStats = "FAILED_TO_GET_OCCUPATION_STATS"
	ErrorFailedToGetAgeStats = "FAILED_TO_GET_AGE_STATS"
	ErrorFailedToGetGenderStats = "FAILED_TO_GET_GENDER_STATS"
)

// 错误信息映射
var ErrorMessages = map[string]string{
	ErrorInternalServer: "Internal server error",
	ErrorInvalidRequest: "Invalid request parameters",
	ErrorEmailAlreadyRegistered: "Email already registered",
	ErrorInvalidEmailOrPassword: "Invalid email or password",
	ErrorPasswordsDoNotMatch: "Passwords do not match",
	ErrorFailedToHashPassword: "Failed to hash password",
	ErrorFailedToCreateUser: "Failed to create user",
	ErrorUserAlreadyVoted: "User has already voted",
	ErrorFailedToSubmitVote: "Failed to submit vote",
	ErrorFailedToGetTotalVotes: "Failed to get total votes",
	ErrorFailedToGetEmployedCount: "Failed to get employed count",
	ErrorFailedToGetUnemployedCount: "Failed to get unemployed count",
	ErrorFailedToGetCityStats: "Failed to get city stats",
	ErrorFailedToGetIndustryStats: "Failed to get industry stats",
	ErrorFailedToGetOccupationStats: "Failed to get occupation stats",
	ErrorFailedToGetAgeStats: "Failed to get age stats",
	ErrorFailedToGetGenderStats: "Failed to get gender stats",
}

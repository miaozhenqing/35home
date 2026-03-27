package common

type ErrorCode int

const (
	ErrSuccess ErrorCode = 0

	ErrInternalServer ErrorCode = 1001
	ErrInvalidRequest ErrorCode = 1002

	ErrEmailAlreadyRegistered ErrorCode = 2001
	ErrInvalidEmailOrPassword ErrorCode = 2002
	ErrPasswordsDoNotMatch ErrorCode = 2003
	ErrFailedToHashPassword ErrorCode = 2004
	ErrFailedToCreateUser ErrorCode = 2005

	ErrUserAlreadyVoted ErrorCode = 3001
	ErrFailedToSubmitVote ErrorCode = 3002
	ErrFailedToGetTotalVotes ErrorCode = 3003
	ErrFailedToGetEmployedCount ErrorCode = 3004
	ErrFailedToGetUnemployedCount ErrorCode = 3005
	ErrFailedToGetCityStats ErrorCode = 3006
	ErrFailedToGetIndustryStats ErrorCode = 3007
	ErrFailedToGetOccupationStats ErrorCode = 3008
	ErrFailedToGetAgeStats ErrorCode = 3009
	ErrFailedToGetGenderStats ErrorCode = 3010
)

var ErrorMessages = map[ErrorCode]string{
	ErrSuccess: "Success",
	ErrInternalServer: "Internal server error",
	ErrInvalidRequest: "Invalid request parameters",
	ErrEmailAlreadyRegistered: "Email already registered",
	ErrInvalidEmailOrPassword: "Invalid email or password",
	ErrPasswordsDoNotMatch: "Passwords do not match",
	ErrFailedToHashPassword: "Failed to hash password",
	ErrFailedToCreateUser: "Failed to create user",
	ErrUserAlreadyVoted: "User has already voted",
	ErrFailedToSubmitVote: "Failed to submit vote",
	ErrFailedToGetTotalVotes: "Failed to get total votes",
	ErrFailedToGetEmployedCount: "Failed to get employed count",
	ErrFailedToGetUnemployedCount: "Failed to get unemployed count",
	ErrFailedToGetCityStats: "Failed to get city stats",
	ErrFailedToGetIndustryStats: "Failed to get industry stats",
	ErrFailedToGetOccupationStats: "Failed to get occupation stats",
	ErrFailedToGetAgeStats: "Failed to get age stats",
	ErrFailedToGetGenderStats: "Failed to get gender stats",
}

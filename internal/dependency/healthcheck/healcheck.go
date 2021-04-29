package healthcheck

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type Result interface {
	// Identity returns the identity
	Identity() int
	// GetOperationID returns the operation id
	GetOperationID() int
	// GetWeightedAverageScore returns the weighted average score
	GetWeightedAverageScore() int
	// GetDBConfigScore returns the database configuration score
	GetDBConfigScore() int
	// GetDBConfigData returns the database configuration data
	GetDBConfigData() string
	// GetDBConfigAdvice returns the database configuration advice
	GetDBConfigAdvice() string
	// GetCPUUsageScore returns the cpu usage score
	GetCPUUsageScore() int
	// GetCPUUsageData returns the cpu usage data
	GetCPUUsageData() string
	// GetIOUtilScore returns the io util score
	GetIOUtilScore() int
	// GetIOUtilData returns the io util data
	GetIOUtilData() string
	// GetDiskCapacityUsageScore returns the disk capacity usage score
	GetDiskCapacityUsageScore() int
	// GetDiskCapacityUsageData returns the disk capacity usage data
	GetDiskCapacityUsageData() string
	// GetConnectionUsageScore returns the connection usage score
	GetConnectionUsageScore() int
	// GetConnectionUsageData returns the connection usage data
	GetConnectionUsageData() string
	// GetAverageActiveSessionNumScore returns the average active session number score
	GetAverageActiveSessionNumScore() int
	// GetAverageActiveSessionNumData returns the average active session number data
	GetAverageActiveSessionNumData() string
	// GetCacheHitRatioScore returns the cache hit ratio score
	GetCacheHitRatioScore() int
	// GetCacheHitRatioData returns the cache hit ratio data
	GetCacheHitRatioData() string
	// GetSlowQueryScore returns the slow query score
	GetSlowQueryScore() int
	// GetSlowQueryData returns the slow query data
	GetSlowQueryData() string
	// GetSlowQueryAdvice returns the slow query advice
	GetSlowQueryAdvice() string
	// GetAccurateReview returns the accurate review
	GetAccurateReview() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// MarshalJSON marshals Result to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSON marshals only specified field of the Result to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// InitOperation initiates the operation
	InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error)
	// GetResultByOperationID returns the result
	GetResultByOperationID(operationID int) (Result, error)
	// UpdateAccurateReviewByOperationID updates the accurate review
	UpdateAccurateReviewByOperationID(operationID int, review int) error
}

type Service interface {
	// GetResult returns the result
	GetResult() Result
	// GetResultByOperationID gets the result by operation id from the middleware
	GetResultByOperationID(id int) error
	// Check checks the server health status
	Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error
	// ReviewAccurate reviews the accurate of the check
	ReviewAccurate(id, review int) error
	// MarshalJSON marshals Service to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSON marshals only specified field of the Service to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type Engine interface {
	// Run checks the server health status
	Run()
}

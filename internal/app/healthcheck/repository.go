package healthcheck

import (
	"errors"
	"fmt"
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

var _ healthcheck.Repository = (*Repository)(nil)

type Repository struct {
	Database middleware.Pool
}

// NewRepository returns *Repository with given middleware.Pool
func NewRepository(db middleware.Pool) *Repository {
	return &Repository{Database: db}
}

// NewRepository returns *Repository with global mysql pool
func NewRepositoryWithGlobal() *Repository {
	return NewRepository(global.DASMySQLPool)
}

// Execute executes given command and placeholders on the middleware
func (r *Repository) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := r.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("healthcheck Repository.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (r *Repository) Transaction() (middleware.Transaction, error) {
	return r.Database.Transaction()
}

// GetResultByOperationID gets a Result by the operationID from the middleware
func (r *Repository) GetResultByOperationID(operationID int) (healthcheck.Result, error) {
	sql := `
		select id, operation_id, weighted_average_score, db_config_score, db_config_data, 
		db_config_advice, cpu_usage_score, cpu_usage_data, cpu_usage_high, io_util_score,
		io_util_data, io_util_high, disk_capacity_usage_score, disk_capacity_usage_data, 
		disk_capacity_usage_high, connection_usage_score, connection_usage_data, 
		connection_usage_high, average_active_session_num_score, average_active_session_num_data,
		average_active_session_num_high, cache_miss_ratio_score, cache_miss_ratio_data, 
		cache_miss_ratio_high, table_size_score, table_size_data, table_size_high, slow_query_score,
		slow_query_data, slow_query_advice, accurate_review, del_flag, create_time, last_update_time
		from t_hc_result
		where del_flag = 0
		and operation_id = ? 
		order by id;
	`
	log.Debugf("healthCheck Repository.GetResultByOperationID select sql: \n%s\nplaceholders: %s", sql, operationID)

	result, err := r.Execute(sql, operationID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("healthCheck Repository.GetResultByOperationID(): data does not exists, operation_id: %d", operationID))
	case 1:
		hcInfo := NewEmptyResultWithRepo(r)
		// map to struct
		err = result.MapToStructByRowIndex(hcInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return hcInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("healthCheck Repository.GetResultByOperationID(): duplicate key exists, operation_id: %d", operationID))
	}
}

// IsRunning gets status by the mysqlServerID from the middleware
func (r *Repository) IsRunning(mysqlServerID int) (bool, error) {
	sql := `select count(1) from t_hc_operation_info where del_flag = 0 and mysql_server_id = ? and status = 1;`
	log.Debugf("healthCheck Repository.IsRunning() select sql: \n%s\nplaceholders: %s", sql, mysqlServerID)

	result, err := r.Execute(sql, mysqlServerID)
	if err != nil {
		return false, err
	}
	count, _ := result.GetInt(constant.ZeroInt, constant.ZeroInt)

	return count != 0, nil
}

// InitOperation creates a operationInfo in the middleware
func (r *Repository) InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	startTimeStr := startTime.Format(constant.TimeLayoutSecond)
	endTimeStr := endTime.Format(constant.TimeLayoutSecond)
	stepInt := int(step.Seconds())

	sql := `insert into t_hc_operation_info(mysql_server_id, start_time, end_time, step) values(?, ?, ?, ?);`
	log.Debugf("healthCheck Repository.InitOperation() insert sql: \n%s\nplaceholders: %s, %s, %s, %s", sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)

	_, err := r.Execute(sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)
	if err != nil {
		return constant.ZeroInt, err
	}

	sql = `
		select id from t_hc_operation_info where del_flag = 0 and 
		mysql_server_id = ? and start_time = ? and end_time = ? and step = ?;
	`
	log.Debugf("healthCheck Repository.InitOperation() select sql: \n%s\nplaceholders: %s, %s, %s, %s", sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)

	result, err := r.Execute(sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// UpdateOperationStatus updates the status and message by the operationID in the middleware
func (r *Repository) UpdateOperationStatus(operationID int, status int, message string) error {
	sql := `update t_hc_operation_info set status = ?, message = ? where id = ?;`
	log.Debugf("healthCheck Repository.UpdateOperationStatus() update sql: \n%s\nplaceholders: %s, %s, %s", sql, operationID, status, message)
	_, err := r.Execute(sql, status, message, operationID)

	return err
}

// SaveResult saves the result in the middleware
func (r *Repository) SaveResult(result healthcheck.Result) error {
	sql := `insert into t_hc_result(operation_id, weighted_average_score, db_config_score, db_config_data, 
		db_config_advice, cpu_usage_score, cpu_usage_data, cpu_usage_high, io_util_score,
		io_util_data, io_util_high, disk_capacity_usage_score, disk_capacity_usage_data, 
		disk_capacity_usage_high, connection_usage_score, connection_usage_data, 
		connection_usage_high, average_active_session_num_score, average_active_session_num_data,
		average_active_session_num_high, cache_miss_ratio_score, cache_miss_ratio_data, 
		cache_miss_ratio_high, table_size_score, table_size_data, table_size_high, slow_query_score,
		slow_query_data, slow_query_advice, accurate_review) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	log.Debugf("healthCheck Repository.SaveResult() insert sql: \n%s\nplaceholders: %s, %s, %s, %s, %s, "+
		"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		sql, result.GetOperationID(), result.GetWeightedAverageScore(), result.GetDBConfigScore(), result.GetDBConfigData(),
		result.GetDBConfigAdvice(), result.GetCPUUsageScore(), result.GetCPUUsageData(), result.GetCPUUsageHigh(),
		result.GetIOUtilScore(), result.GetIOUtilData(), result.GetIOUtilHigh(), result.GetDiskCapacityUsageScore(),
		result.GetDiskCapacityUsageData(), result.GetDiskCapacityUsageHigh(), result.GetConnectionUsageScore(),
		result.GetConnectionUsageData(), result.GetConnectionUsageHigh(), result.GetAverageActiveSessionNumScore(),
		result.GetAverageActiveSessionNumData(), result.GetAverageActiveSessionNumHigh(), result.GetCacheMissRatioScore(),
		result.GetCacheMissRatioData(), result.GetCacheMissRatioHigh(), result.GetTableSizeScore(), result.GetTableSizeData(),
		result.GetTableSizeHigh(), result.GetSlowQueryScore(), result.GetSlowQueryData(), result.GetSlowQueryAdvice(),
		result.GetAccurateReview())

	// execute
	_, err := r.Execute(sql, result.GetOperationID(), result.GetWeightedAverageScore(), result.GetDBConfigScore(),
		result.GetDBConfigData(), result.GetDBConfigAdvice(), result.GetCPUUsageScore(), result.GetCPUUsageData(),
		result.GetCPUUsageHigh(), result.GetIOUtilScore(), result.GetIOUtilData(), result.GetIOUtilHigh(),
		result.GetDiskCapacityUsageScore(), result.GetDiskCapacityUsageData(), result.GetDiskCapacityUsageHigh(),
		result.GetConnectionUsageScore(), result.GetConnectionUsageData(), result.GetConnectionUsageHigh(),
		result.GetAverageActiveSessionNumScore(), result.GetAverageActiveSessionNumData(), result.GetAverageActiveSessionNumHigh(),
		result.GetCacheMissRatioScore(), result.GetCacheMissRatioData(), result.GetCacheMissRatioHigh(),
		result.GetTableSizeScore(), result.GetTableSizeData(), result.GetTableSizeHigh(), result.GetSlowQueryScore(),
		result.GetSlowQueryData(), result.GetSlowQueryAdvice(), result.GetAccurateReview())

	return err
}

// UpdateAccurateReviewByOperationID updates the accurateReview by the operationID in the middleware
func (r *Repository) UpdateAccurateReviewByOperationID(operationID int, review int) error {
	sql := `update t_hc_result set accurate_review = ? where operation_id = ?;`
	log.Debugf("healthCheck Repository.UpdateAccurateReviewByOperationID() update sql: \n%s\nplaceholders: %s, %s", sql, operationID, review)

	_, err := r.Execute(sql, review, operationID)
	return err
}

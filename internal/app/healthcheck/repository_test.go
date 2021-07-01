package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	// modify these connection information
	dbAddr   = "127.0.0.1:3306"
	dbDBName = "das"
	dbDBUser = "root"
	dbDBPass = "mysql123"

	defaultResultOperationID                  = 1
	defaultResultWeightedAverageScore         = 1
	defaultResultDBConfigScore                = 1
	defaultResultDBConfigData                 = ""
	defaultResultDBConfigAdvice               = ""
	defaultResultCPUUsageScore                = 1
	defaultResultCPUUsageData                 = ""
	defaultResultCPUUsageHigh                 = ""
	defaultResultIOUtilScore                  = 1
	defaultResultIOUtilData                   = ""
	defaultResultIOUtilHigh                   = ""
	defaultResultDiskCapacityUsageScore       = 1
	defaultResultDiskCapacityUsageData        = ""
	defaultResultDiskCapacityUsageHigh        = ""
	defaultResultConnectionUsageScore         = 1
	defaultResultConnectionUsageData          = ""
	defaultResultConnectionUsageHigh          = ""
	defaultResultAverageActiveSessionNumScore = 1
	defaultResultAverageActiveSessionNumData  = ""
	defaultResultAverageActiveSessionNumHigh  = ""
	defaultResultCacheMissRatioScore          = 1
	defaultResultCacheMissRatioData           = 1.00
	defaultResultCacheMissRatioHigh           = 1.00
	defaultResultTableSizeScore               = 1
	defaultResultTableSizeData                = ""
	defaultResultTableSizeHigh                = ""
	defaultResultSlowQueryScore               = 1
	defaultResultSlowQueryData                = ""
	defaultResultSlowQueryAdvice              = ""
	defaultResultAccurateReview               = 0

	defaultResultMysqlServerID = 1
	defaultResultStartTime     = "2021-05-01 10:00:00.000000"
	defaultResultEndTime       = "2021-05-01 13:00:00.000000"
	defaultResultStep          = 10
	newResultStatus            = 1
	AccurateReviewStruct       = "AccurateReview"
	newResultAccurateReview    = 1
)

var repository = initRepository()

func initRepository() *Repository {
	pool, err := mysql.NewPoolWithDefault(dbAddr, dbDBName, dbDBUser, dbDBPass)
	log.Infof("pool: %v, error: %v", pool, err)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}

	return NewRepository(pool)
}

func createResult() error {
	hcInfo := NewResultWithDefault(defaultResultOperationID, defaultResultWeightedAverageScore, defaultResultDBConfigScore,
		defaultResultCPUUsageScore, defaultResultIOUtilScore, defaultResultDiskCapacityUsageScore, defaultResultConnectionUsageScore,
		defaultResultAverageActiveSessionNumScore, defaultResultCacheMissRatioScore, defaultResultTableSizeScore, defaultResultSlowQueryScore, defaultResultAccurateReview)
	err := repository.SaveResult(hcInfo)

	return err
}

func deleteResultByID(id int) error {
	sql := `delete from t_hc_result where id = ?`
	_, err := repository.Execute(sql, id)
	return err
}

func deleteOperationInfoByID(id int) error {
	sql := `delete from t_hc_operation_info where id = ?`
	_, err := repository.Execute(sql, id)
	return err
}

func TestRepositoryAll(t *testing.T) {
	TestRepository_Execute(t)
	TestRepository_GetResultByOperationID(t)
	TestRepository_IsRunning(t)
	TestRepository_InitOperation(t)
	TestRepository_UpdateOperationStatus(t)
	TestRepository_SaveResult(t)
	TestRepository_UpdateAccurateReviewByOperationID(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := "select 1;"
	result, err := repository.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestRepository_Transaction(t *testing.T) {
	asst := assert.New(t)

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
	tx, err := repository.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultResultOperationID, defaultResultWeightedAverageScore, defaultResultDBConfigScore,
		defaultResultDBConfigData, defaultResultDBConfigAdvice, defaultResultCPUUsageScore, defaultResultCPUUsageData,
		defaultResultCPUUsageHigh, defaultResultIOUtilScore, defaultResultIOUtilData, defaultResultIOUtilHigh,
		defaultResultDiskCapacityUsageScore, defaultResultDiskCapacityUsageData, defaultResultDiskCapacityUsageHigh,
		defaultResultConnectionUsageScore, defaultResultConnectionUsageData, defaultResultConnectionUsageHigh,
		defaultResultAverageActiveSessionNumScore, defaultResultAverageActiveSessionNumData, defaultResultAverageActiveSessionNumHigh,
		defaultResultCacheMissRatioScore, defaultResultCacheMissRatioData, defaultResultCacheMissRatioHigh,
		defaultResultTableSizeScore, defaultResultTableSizeData, defaultResultTableSizeHigh, defaultResultSlowQueryScore,
		defaultResultSlowQueryData, defaultResultSlowQueryAdvice, defaultResultAccurateReview)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select operation_id from t_hc_result where operation_id = ?`
	result, err := tx.Execute(sql, defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	operationID, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if operationID != defaultResultOperationID {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entity, err := repository.GetResultByOperationID(defaultResultOperationID)
	if entity != nil {
		asst.Fail("test Transaction() failed")
	}
}

func TestRepository_GetResultByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	result, err := repository.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	operationID := result.GetOperationID()
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	asst.Equal(defaultResultOperationID, operationID, "test GetResultByOperationID() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

func TestRepository_IsRunning(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_hc_operation_info(mysql_server_id, start_time, end_time, step) values(?, ?, ?, ?);`
	_, err := repository.Execute(sql, defaultResultMysqlServerID, defaultResultStartTime, defaultResultEndTime, defaultResultStep)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	result, err := repository.IsRunning(defaultResultMysqlServerID)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	asst.False(result, "test IsRunning() failed")
	// delete
	sql = `select id from t_hc_operation_info order by id desc limit 0,1`
	resultID, err := repository.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	id, err := resultID.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
}

func TestRepository_InitOperation(t *testing.T) {
	asst := assert.New(t)

	startTime, _ := time.ParseInLocation(constant.TimeLayoutSecond, defaultResultStartTime, time.Local)
	endTime, _ := time.ParseInLocation(constant.TimeLayoutSecond, defaultResultEndTime, time.Local)
	step := time.Duration(int64(defaultResultStep))

	id, err := repository.InitOperation(defaultResultMysqlServerID, startTime, endTime, step)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	sql := `select mysql_server_id from t_hc_operation_info where id = ?;`
	result, err := repository.Execute(sql, id)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	mysqlServerID, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	asst.Equal(defaultResultMysqlServerID, mysqlServerID, "test InitOperation() failed")
	// delete
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
}

func TestRepository_UpdateOperationStatus(t *testing.T) {
	asst := assert.New(t)

	startTime, _ := time.ParseInLocation(constant.TimeLayoutSecond, defaultResultStartTime, time.Local)
	endTime, _ := time.ParseInLocation(constant.TimeLayoutSecond, defaultResultEndTime, time.Local)
	step := time.Duration(int64(defaultResultStep))

	id, err := repository.InitOperation(defaultResultMysqlServerID, startTime, endTime, step)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	err = repository.UpdateOperationStatus(id, newResultStatus, "")
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	sql := `select status from t_hc_operation_info where id = ?;`
	result, err := repository.Execute(sql, id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	status, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	asst.Equal(newResultStatus, status, "test UpdateOperationStatus() failed")
	// delete
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
}

func TestRepository_SaveResult(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	result, err := repository.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	asst.Equal(defaultResultOperationID, result.GetOperationID(), "test SaveResult() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
}

func TestRepository_UpdateAccurateReviewByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccurateReviewByOperationID() failed", err))
	result, err := repository.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccurateReviewByOperationID() failed", err))
	err = result.Set(map[string]interface{}{AccurateReviewStruct: newResultAccurateReview})
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccurateReviewByOperationID() failed", err))
	err = repository.UpdateAccurateReviewByOperationID(result.GetOperationID(), newResultAccurateReview)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccurateReviewByOperationID() failed", err))
	asst.Equal(newResultAccurateReview, result.GetAccurateReview(), "test UpdateAccurateReviewByOperationID() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccurateReviewByOperationID() failed", err))
}

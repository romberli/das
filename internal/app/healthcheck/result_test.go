package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	resultDBAddress = "127.0.0.1:3306"
	resultHostIP    = "127.0.0.1"
	resultPortNum   = 3306
	resultDBName    = "das"
	resultDBUser    = "root"
	resultDBPass    = "root"

	resultID                           = 1
	resultOperationID                  = 1
	resultWeightedAverageScore         = 1
	resultDBConfigScore                = 1
	resultDBConfigData                 = "db config data"
	resultDBConfigAdvice               = "db config advice"
	resultCPUUsageScore                = 80
	resultCPUUsageData                 = "cpu usage data"
	resultCPUUsageHigh                 = "cpu usage high"
	resultIOUtilScore                  = 80
	resultIOUtilData                   = "io util data"
	resultIOUtilHigh                   = "io util high"
	resultDiskCapacityUsageScore       = 80
	resultDiskCapacityUsageData        = "disk capacity usage data"
	resultDiskCapacityUsageHigh        = "disk capacity usage high"
	resultConnectionUsageScore         = 80
	resultConnectionUsageData          = "connection usage data"
	resultConnectionUsageHigh          = "connection usage high"
	resultAverageActiveSessionNumScore = 80
	resultAverageActiveSessionNumData  = "average active session num data"
	resultAverageActiveSessionNumHigh  = "average active session num high"
	resultCacheMissRatioScore          = 80
	resultCacheMissRatioData           = 0.8
	resultCacheMissRatioHigh           = 0.8
	resultTableSizeScore               = 80
	resultTableSizeData                = "table size data"
	resultTableSizeHigh                = "table size high"
	resultSlowQueryScore               = 80
	resultSlowQueryData                = "slow query data"
	resultSlowQueryAdvice              = "slow query advice"
	resultAccurateReview               = 0
	resultDelFlag                      = 0
)

func rInitRepository() *Repository {
	pool, err := mysql.NewPoolWithDefault(resultDBAddress, resultDBName, resultDBUser, resultDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}
	return NewRepository(pool)
}

var rRepo = rInitRepository()

func rCreateService() (*Service, error) {
	var result = NewResult(rRepo, resultOperationID, resultWeightedAverageScore, resultDBConfigScore, resultDBConfigData, resultDBConfigAdvice, resultCPUUsageScore, resultCPUUsageData, resultCPUUsageHigh, resultIOUtilScore, resultIOUtilData, resultIOUtilHigh, resultDiskCapacityUsageScore, resultDiskCapacityUsageData, resultDiskCapacityUsageHigh, resultConnectionUsageScore, resultConnectionUsageData, resultConnectionUsageHigh, resultAverageActiveSessionNumScore, resultAverageActiveSessionNumData, resultAverageActiveSessionNumHigh, resultCacheMissRatioScore, resultCacheMissRatioData, resultCacheMissRatioHigh, resultTableSizeScore, resultTableSizeData, resultTableSizeHigh, resultSlowQueryScore, resultSlowQueryData, resultSlowQueryAdvice)
	err := rRepo.SaveResult(result)
	if err != nil {
		return nil, err
	}
	return &Service{
		Repository: rRepo,
		Result:     result,
	}, nil
}

func rDeleteHCResultByOperationID(operationID int) error {
	sql := `delete from t_hc_result where operation_id = ?`
	_, err := rRepo.Execute(sql, operationID)
	return err
}

func TestResultAll(t *testing.T) {
	TestResult_Identity(t)
	TestResult_GetOperationID(t)
	TestResult_GetWeightedAverageScore(t)
	TestResult_GetDBConfigScore(t)
	TestResult_GetDBConfigData(t)
	TestResult_GetDBConfigAdvice(t)
	TestResult_GetCPUUsageScore(t)
	TestResult_GetCPUUsageData(t)
	TestResult_GetCPUUsageHigh(t)
	TestResult_GetIOUtilScore(t)
	TestResult_GetIOUtilData(t)
	TestResult_GetIOUtilHigh(t)
	TestResult_GetDiskCapacityUsageScore(t)
	TestResult_GetDiskCapacityUsageData(t)
	TestResult_GetDiskCapacityUsageHigh(t)
	TestResult_GetConnectionUsageScore(t)
	TestResult_GetConnectionUsageData(t)
	TestResult_GetConnectionUsageHigh(t)
	TestResult_GetAverageActiveSessionNumScore(t)
	TestResult_GetAverageActiveSessionNumData(t)
	TestResult_GetAverageActiveSessionNumHigh(t)
	TestResult_GetCacheMissRatioScore(t)
	TestResult_GetCacheMissRatioData(t)
	TestResult_GetCacheMissRatioHigh(t)
	TestResult_GetTableSizeScore(t)
	TestResult_GetTableSizeData(t)
	TestResult_GetTableSizeHigh(t)
	TestResult_GetSlowQueryScore(t)
	TestResult_GetSlowQueryData(t)
	TestResult_GetSlowQueryAdvice(t)
	TestResult_GetAccurateReview(t)
	TestResult_GetDelFlag(t)
	TestResult_GetCreateTime(t)
	TestResult_GetLastUpdateTime(t)
	TestResult_Set(t)
	TestResult_MarshalJSON(t)
	TestResult_MarshalJSONWithFields(t)
}

func TestResult_Identity(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
	result := service.GetResult()
	id := result.Identity()
	asst.IsType(resultID, id, "test Identity() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
}

func TestResult_GetOperationID(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
	result := service.GetResult()
	operationID := result.GetOperationID()
	asst.Equal(resultOperationID, operationID, "test GetOperationID() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
}

func TestResult_GetWeightedAverageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
	result := service.GetResult()
	weightedAverageScore := result.GetWeightedAverageScore()
	asst.Equal(resultWeightedAverageScore, weightedAverageScore, "test GetWeightedAverageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
}

func TestResult_GetDBConfigScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
	result := service.GetResult()
	dbConfigScore := result.GetDBConfigScore()
	asst.Equal(resultDBConfigScore, dbConfigScore, "test GetDBConfigScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
}

func TestResult_GetDBConfigData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
	result := service.GetResult()
	dbConfigData := result.GetDBConfigData()
	asst.Equal(resultDBConfigData, dbConfigData, "test GetDBConfigData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
}

func TestResult_GetDBConfigAdvice(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
	result := service.GetResult()
	dbConfigAdvice := result.GetDBConfigAdvice()
	asst.Equal(resultDBConfigAdvice, dbConfigAdvice, "test GetDBConfigAdvice() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
}

func TestResult_GetCPUUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
	result := service.GetResult()
	cpuUsageScore := result.GetCPUUsageScore()
	asst.Equal(resultCPUUsageScore, cpuUsageScore, "test GetCPUUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
}

func TestResult_GetCPUUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
	result := service.GetResult()
	cpuUsageData := result.GetCPUUsageData()
	asst.Equal(resultCPUUsageData, cpuUsageData, "test GetCPUUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
}

func TestResult_GetCPUUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
	result := service.GetResult()
	cpuUsageHigh := result.GetCPUUsageHigh()
	asst.Equal(resultCPUUsageHigh, cpuUsageHigh, "test GetCPUUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
}

func TestResult_GetIOUtilScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	result := service.GetResult()
	ioUtilScore := result.GetIOUtilScore()
	asst.Equal(resultIOUtilScore, ioUtilScore, "test GetIOUtilScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
}

func TestResult_GetIOUtilData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	result := service.GetResult()
	ioUtilData := result.GetIOUtilData()
	asst.Equal(resultIOUtilData, ioUtilData, "test GetIOUtilData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
}

func TestResult_GetIOUtilHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
	result := service.GetResult()
	ioUtilHigh := result.GetIOUtilHigh()
	asst.Equal(resultIOUtilHigh, ioUtilHigh, "test GetIOUtilHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
}

func TestResult_GetDiskCapacityUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	result := service.GetResult()
	diskCapacityUsageScore := result.GetDiskCapacityUsageScore()
	asst.Equal(resultDiskCapacityUsageScore, diskCapacityUsageScore, "test GetDiskCapacityUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
}

func TestResult_GetDiskCapacityUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	result := service.GetResult()
	diskCapacityUsageData := result.GetDiskCapacityUsageData()
	asst.Equal(resultDiskCapacityUsageData, diskCapacityUsageData, "test GetDiskCapacityUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
}

func TestResult_GetDiskCapacityUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
	result := service.GetResult()
	diskCapacityUsageHigh := result.GetDiskCapacityUsageHigh()
	asst.Equal(resultDiskCapacityUsageHigh, diskCapacityUsageHigh, "test GetDiskCapacityUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
}

func TestResult_GetConnectionUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
	result := service.GetResult()
	connectionUsageScore := result.GetConnectionUsageScore()
	asst.Equal(resultConnectionUsageScore, connectionUsageScore, "test GetConnectionUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
}

func TestResult_GetConnectionUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
	result := service.GetResult()
	connectionUsageData := result.GetConnectionUsageData()
	asst.Equal(resultConnectionUsageData, connectionUsageData, "test GetConnectionUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
}

func TestResult_GetConnectionUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
	result := service.GetResult()
	connectionUsageHigh := result.GetConnectionUsageHigh()
	asst.Equal(resultConnectionUsageHigh, connectionUsageHigh, "test GetConnectionUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
}

func TestResult_GetAverageActiveSessionNumScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumScore() failed", err))
	result := service.GetResult()
	averageActiveSessionNumScore := result.GetAverageActiveSessionNumScore()
	asst.Equal(resultAverageActiveSessionNumScore, averageActiveSessionNumScore, "test GetAverageActiveSessionNumScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumScore() failed", err))
}

func TestResult_GetAverageActiveSessionNumData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumData() failed", err))
	result := service.GetResult()
	averageActiveSessionNumData := result.GetAverageActiveSessionNumData()
	asst.Equal(resultAverageActiveSessionNumData, averageActiveSessionNumData, "test GetAverageActiveSessionNumData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumData() failed", err))
}

func TestResult_GetAverageActiveSessionNumHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumHigh() failed", err))
	result := service.GetResult()
	averageActiveSessionNumHigh := result.GetAverageActiveSessionNumHigh()
	asst.Equal(resultAverageActiveSessionNumHigh, averageActiveSessionNumHigh, "test GetAverageActiveSessionNumHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionNumHigh() failed", err))
}

func TestResult_GetCacheMissRatioScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
	result := service.GetResult()
	cacheMissRatioScore := result.GetCacheMissRatioScore()
	asst.Equal(resultCacheMissRatioScore, cacheMissRatioScore, "test GetCacheMissRatioScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
}

func TestResult_GetCacheMissRatioData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
	result := service.GetResult()
	cacheMissRatioData := result.GetCacheMissRatioData()
	asst.Equal(resultCacheMissRatioData, cacheMissRatioData, "test GetCacheMissRatioData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
}

func TestResult_GetCacheMissRatioHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
	result := service.GetResult()
	cacheMissRatioHigh := result.GetCacheMissRatioHigh()
	asst.Equal(resultCacheMissRatioHigh, cacheMissRatioHigh, "test GetCacheMissRatioHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
}

func TestResult_GetTableSizeScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
	result := service.GetResult()
	tableSizeScore := result.GetTableSizeScore()
	asst.Equal(resultTableSizeScore, tableSizeScore, "test GetTableSizeScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
}

func TestResult_GetTableSizeData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
	result := service.GetResult()
	tableSizeData := result.GetTableSizeData()
	asst.Equal(resultTableSizeData, tableSizeData, "test GetTableSizeData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
}

func TestResult_GetTableSizeHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
	result := service.GetResult()
	tableSizeHigh := result.GetTableSizeHigh()
	asst.Equal(resultTableSizeHigh, tableSizeHigh, "test GetTableSizeHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
}

func TestResult_GetSlowQueryScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
	result := service.GetResult()
	slowQueryScore := result.GetSlowQueryScore()
	asst.Equal(resultSlowQueryScore, slowQueryScore, "test GetSlowQueryScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
}

func TestResult_GetSlowQueryData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
	result := service.GetResult()
	slowQueryData := result.GetSlowQueryData()
	asst.Equal(resultSlowQueryData, slowQueryData, "test GetSlowQueryData() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
}

func TestResult_GetSlowQueryAdvice(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
	result := service.GetResult()
	slowQueryAdvice := result.GetSlowQueryAdvice()
	asst.Equal(resultSlowQueryAdvice, slowQueryAdvice, "test GetSlowQueryAdvice() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
}

func TestResult_GetAccurateReview(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAccurateReview() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAccurateReview() failed", err))
	result := service.GetResult()
	accurateReview := result.GetAccurateReview()
	asst.Equal(resultAccurateReview, accurateReview, "test GetAccurateReview() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAccurateReview() failed", err))
}

func TestResult_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
	result := service.GetResult()
	delFlag := result.GetDelFlag()
	asst.Equal(resultDelFlag, delFlag, "test GetDelFlag() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
}

func TestResult_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
	result := service.GetResult()
	createTime := result.GetCreateTime()
	asst.IsType(time.Now(), createTime, "test GetCreateTime() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
}

func TestResult_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
	result := service.GetResult()
	lastUpdateTime := result.GetLastUpdateTime()
	asst.IsType(time.Now(), lastUpdateTime, "test GetLastUpdateTime() failed")
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
}

func TestResult_Set(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	result := service.GetResult()

	fields := make(map[string]interface{})
	fields["ID"] = resultID
	fields["OperationID"] = resultOperationID

	err = result.Set(fields)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))

	// field XX does not exist
	fields["XX"] = 100
	err = result.Set(fields)
	asst.NotNil(err, common.CombineMessageWithError("test Set() failed", err))

	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
}

func TestResult_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	result := service.GetResult()
	_, err = result.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
}

func TestResult_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	err = service.GetResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	result := service.GetResult()
	_, err = result.MarshalJSONWithFields("ID", "OperationID", "WeightedAverageScore", "DBConfigScore", "DBConfigData", "DBConfigAdvice", "CPUUsageScore", "CPUUsageData", "CPUUsageHigh", "IOUtilScore", "IOUtilData", "IOUtilHigh", "DiskCapacityUsageScore", "DiskCapacityUsageData", "DiskCapacityUsageHigh", "ConnectionUsageScore", "ConnectionUsageData", "ConnectionUsageHigh", "AverageActiveSessionNumScore", "AverageActiveSessionNumData", "AverageActiveSessionNumHigh", "CacheMissRatioScore", "CacheMissRatioData", "CacheMissRatioHigh", "TableSizeScore", "TableSizeData", "TableSizeHigh", "SlowQueryScore", "SlowQueryData", "SlowQueryAdvice")
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	// delete
	err = rDeleteHCResultByOperationID(resultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
}

// go test ./result_test.go ./result.go ./repository.go ./service.go ./default_engine.go

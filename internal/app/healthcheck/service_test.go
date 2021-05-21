package healthcheck

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	serviceDBAddress = "127.0.0.1:3306"
	serviceHostIP    = "127.0.0.1"
	servicePortNum   = 3306
	serviceDBName    = "das"
	serviceDBUser    = "root"
	serviceDBPass    = "root"

	serviceID                           = 1
	serviceOperationID                  = 1
	serviceWeightedAverageScore         = 1
	serviceDBConfigScore                = 1
	serviceDBConfigData                 = "db config data"
	serviceDBConfigAdvice               = "db config advice"
	serviceCPUUsageScore                = 80
	serviceCPUUsageData                 = "cpu usage data"
	serviceCPUUsageHigh                 = "cpu usage high"
	serviceIOUtilScore                  = 80
	serviceIOUtilData                   = "io util data"
	serviceIOUtilHigh                   = "io util high"
	serviceDiskCapacityUsageScore       = 80
	serviceDiskCapacityUsageData        = "disk capacity usage data"
	serviceDiskCapacityUsageHigh        = "disk capacity usage high"
	serviceConnectionUsageScore         = 80
	serviceConnectionUsageData          = "connection usage data"
	serviceConnectionUsageHigh          = "connection usage high"
	serviceAverageActiveSessionNumScore = 80
	serviceAverageActiveSessionNumData  = "average active session num data"
	serviceAverageActiveSessionNumHigh  = "average active session num high"
	serviceCacheMissRatioScore          = 80
	serviceCacheMissRatioData           = 0.8
	serviceCacheMissRatioHigh           = 0.8
	serviceTableSizeScore               = 80
	serviceTableSizeData                = "table size data"
	serviceTableSizeHigh                = "table size high"
	serviceSlowQueryScore               = 80
	serviceSlowQueryData                = "slow query data"
	serviceSlowQueryAdvice              = "slow query advice"
	serviceAccurateReview               = 0
	serviceDelFlag                      = 0

	serviceMysqlServerID   = 1
	serviceStartTime       = "2021-05-18 12:00:00.000000"
	serviceEndTime         = "2021-05-18 15:00:00.000000"
	serviceStep            = 10
	servicenewResultStatus = 1
)

func sInitRepository() *Repository {
	pool, err := mysql.NewPoolWithDefault(serviceDBAddress, serviceDBName, serviceDBUser, serviceDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}
	return NewRepository(pool)
}

var sRepo = sInitRepository()

func sCreateService() (*Service, error) {
	var result = NewResult(sRepo, serviceOperationID, serviceWeightedAverageScore, serviceDBConfigScore, serviceDBConfigData, serviceDBConfigAdvice, serviceCPUUsageScore, serviceCPUUsageData, serviceCPUUsageHigh, serviceIOUtilScore, serviceIOUtilData, serviceIOUtilHigh, serviceDiskCapacityUsageScore, serviceDiskCapacityUsageData, serviceDiskCapacityUsageHigh, serviceConnectionUsageScore, serviceConnectionUsageData, serviceConnectionUsageHigh, serviceAverageActiveSessionNumScore, serviceAverageActiveSessionNumData, serviceAverageActiveSessionNumHigh, serviceCacheMissRatioScore, serviceCacheMissRatioData, serviceCacheMissRatioHigh, serviceTableSizeScore, serviceTableSizeData, serviceTableSizeHigh, serviceSlowQueryScore, serviceSlowQueryData, serviceSlowQueryAdvice)
	err := sRepo.SaveResult(result)
	if err != nil {
		return nil, err
	}
	return &Service{
		Repository: sRepo,
		Result:     result,
	}, nil
}

func sDeleteHCResultByOperationID(operationID int) error {
	sql := `delete from t_hc_result where operation_id = ?`
	_, err := sRepo.Execute(sql, operationID)
	return err
}

func TestServiceAll(t *testing.T) {
	TestService_GetResult(t)
	TestService_GetResultByOperationID(t)
	TestService_Check(t)
	TestService_ReviewAccurate(t)
	TestService_MarshalJSON(t)
	TestService_MarshalJSONWithFields(t)
}

func TestService_GetResult(t *testing.T) {
	asst := assert.New(t)

	service, err := sCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetResult() failed", err))
	result := service.GetResult()
	asst.Equal(serviceOperationID, result.GetOperationID(), common.CombineMessageWithError("test GetResult() failed", err))
	asst.Equal(serviceWeightedAverageScore, result.GetWeightedAverageScore(), common.CombineMessageWithError("test GetResult() failed", err))
	// delete
	err = sDeleteHCResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResult() failed", err))
}

func TestService_GetResultByOperationID(t *testing.T) {
	asst := assert.New(t)

	service, err := sCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	err = service.GetResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	result := service.GetResult()
	asst.Equal(serviceOperationID, result.GetOperationID(), common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	asst.Equal(serviceWeightedAverageScore, result.GetWeightedAverageScore(), common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	// delete
	err = sDeleteHCResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

// bug
func TestService_Check(t *testing.T) {
	// asst := assert.New(t)

	// service, err := sCreateService()
	// asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))

	// startTime, _ := now.Parse(serviceStartTime)
	// endTime, _ := now.Parse(serviceEndTime)
	// step := time.Duration(serviceStep) * time.Second

	// err = service.Check(serviceMysqlServerID, startTime, endTime, step)
	// asst.Nil(err, common.CombineMessageWithError("test Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) failed", err))

	// // delete
	// err = sDeleteHCResultByOperationID(serviceOperationID)
	// asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

// bug
func TestService_CheckByHostInfo(t *testing.T) {
	// asst := assert.New(t)

	// service, err := sCreateService()
	// asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) failed", err))

	// startTime, _ := now.Parse(serviceStartTime)
	// endTime, _ := now.Parse(serviceEndTime)
	// step := time.Duration(serviceStep) * time.Second

	// err = service.CheckByHostInfo(serviceHostIP, servicePortNum, startTime, endTime, step)
	// asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) failed", err))

	// // delete
	// err = sDeleteHCResultByOperationID(serviceOperationID)
	// asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) failed", err))
}

func TestService_ReviewAccurate(t *testing.T) {
	asst := assert.New(t)

	service, err := sCreateService()
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccurate(id, review int) failed", err))
	review := 2
	err = service.ReviewAccurate(serviceOperationID, review)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccurate(id, review int) failed", err))
	err = service.GetResultByOperationID(serviceOperationID)
	result := service.GetResult()
	reviewed := result.GetAccurateReview()
	asst.Equal(review, reviewed, common.CombineMessageWithError("test ReviewAccurate(id, review int) failed", err))
	// delete
	err = sDeleteHCResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccurate(id, review int) failed", err))
}

func TestService_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	service, err := sCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	_, err = service.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	// delete
	err = sDeleteHCResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
}

func TestService_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	service, err := sCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields(fields ...string) failed", err))
	_, err = service.MarshalJSONWithFields("ID", "OperationID", "WeightedAverageScore", "DBConfigScore", "DBConfigData", "DBConfigAdvice", "CPUUsageScore", "CPUUsageData", "CPUUsageHigh", "IOUtilScore", "IOUtilData", "IOUtilHigh", "DiskCapacityUsageScore", "DiskCapacityUsageData", "DiskCapacityUsageHigh", "ConnectionUsageScore", "ConnectionUsageData", "ConnectionUsageHigh", "AverageActiveSessionNumScore", "AverageActiveSessionNumData", "AverageActiveSessionNumHigh", "CacheMissRatioScore", "CacheMissRatioData", "CacheMissRatioHigh", "TableSizeScore", "TableSizeData", "TableSizeHigh", "SlowQueryScore", "SlowQueryData", "SlowQueryAdvice")
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields(fields ...string) failed", err))
	// delete
	err = sDeleteHCResultByOperationID(serviceOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields(fields ...string) failed", err))
}

// go test ./service_test.go ./service.go ./repository.go ./default_engine.go ./result.go

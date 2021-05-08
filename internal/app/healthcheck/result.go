package healthcheck

import (
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

type Result struct {
	ID                           int       `middleware:"id" json:"id"`
	OperationID                  int       `middleware:"operation_id" json:"operation_id"`
	WeightedAverageScore         int       `middleware:"weighted_average_score" json:"weighted_average_score"`
	DBConfigScore                int       `middleware:"db_config_score" json:"db_config_score"`
	DBConfigData                 string    `middleware:"db_config_data" json:"db_config_data"`
	DBConfigAdvice               string    `middleware:"db_config_advice" json:"db_config_advice"`
	CPUUsageScore                int       `middleware:"cpu_usage_score" json:"cpu_usage_score"`
	CPUUsageData                 string    `middleware:"cpu_usage_data" json:"cpu_usage_data"`
	CPUUsageHigh                 string    `middleware:"cpu_usage_high" json:"cpu_usage_high"`
	IOUtilScore                  int       `middleware:"io_util_score" json:"io_util_score"`
	IOUtilData                   string    `middleware:"io_util_data" json:"io_util_data"`
	IOUtilHigh                   string    `middleware:"io_util_high" json:"io_util_high"`
	DiskCapacityUsageScore       int       `middleware:"disk_capacity_usage_score" json:"disk_capacity_usage_score"`
	DiskCapacityUsageData        string    `middleware:"disk_capacity_usage_data" json:"disk_capacity_usage_data"`
	DiskCapacityUsageHigh        string    `middleware:"disk_capacity_usage_high" json:"disk_capacity_usage_high"`
	ConnectionUsageScore         int       `middleware:"connection_usage_score" json:"connection_usage_score"`
	ConnectionUsageData          string    `middleware:"connection_usage_data" json:"connection_usage_data"`
	ConnectionUsageHigh          string    `middleware:"connection_usage_high" json:"connection_usage_high"`
	AverageActiveSessionNumScore int       `middleware:"average_active_session_num_score" json:"average_active_session_num_score"`
	AverageActiveSessionNumData  string    `middleware:"average_active_session_num_data" json:"average_active_session_num_data"`
	AverageActiveSessionNumHigh  string    `middleware:"average_active_session_num_high" json:"average_active_session_num_high"`
	CacheMissRatioScore          int       `middleware:"cache_miss_ratio_score" json:"cache_miss_ratio_score"`
	CacheMissRatioData           string    `middleware:"cache_miss_ratio_data" json:"cache_miss_ratio_data"`
	CacheMissRatioHigh           string    `middleware:"cache_miss_ratio_high" json:"cache_miss_ratio_high"`
	SlowQueryScore               int       `middleware:"slow_query_score" json:"slow_query_score"`
	SlowQueryData                string    `middleware:"slow_query_data" json:"slow_query_data"`
	SlowQueryAdvice              string    `middleware:"slow_query_advice" json:"slow_query_advice"`
	AccurateReview               int       `middleware:"accurate_review" json:"accurate_review"`
	DelFlag                      int       `middleware:"del_flag" json:"del_flag"`
	CreateTime                   time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime               time.Time `middleware:"last_update_time" json:"last_update_time"`
}

func NewResult(operationID int, weightedAverageScore int, dbConfigScore int, dbConfigData string, dbConfigAdvice string,
	cpuUsageScore int, cpuUsageData string, cpuUsageHigh string, ioUtilScore int, ioUtilData string, ioUtilHigh string,
	diskCapacityUsageScore int, diskCapacityUsageData string, diskCapacityUsageHigh string,
	connectionUsageScore int, connectionUsageData string, connectionUsageHigh string,
	averageActiveSessionNumScore int, averageActiveSessionNumData string, averageActiveSessionNumHigh string,
	cacheMissRatioScore int, cacheMissRatioData string, cacheMissRatioHigh string,
	slowQueryScore int, slowQueryData string, slowQueryAdvice string) *Result {
	return &Result{
		OperationID:                  operationID,
		WeightedAverageScore:         weightedAverageScore,
		DBConfigScore:                dbConfigScore,
		DBConfigData:                 dbConfigData,
		DBConfigAdvice:               dbConfigAdvice,
		CPUUsageScore:                cpuUsageScore,
		CPUUsageData:                 cpuUsageData,
		CPUUsageHigh:                 cpuUsageHigh,
		IOUtilScore:                  ioUtilScore,
		IOUtilData:                   ioUtilData,
		IOUtilHigh:                   ioUtilHigh,
		DiskCapacityUsageScore:       diskCapacityUsageScore,
		DiskCapacityUsageData:        diskCapacityUsageData,
		DiskCapacityUsageHigh:        diskCapacityUsageHigh,
		ConnectionUsageScore:         connectionUsageScore,
		ConnectionUsageData:          connectionUsageData,
		ConnectionUsageHigh:          connectionUsageHigh,
		AverageActiveSessionNumScore: averageActiveSessionNumScore,
		AverageActiveSessionNumData:  averageActiveSessionNumData,
		AverageActiveSessionNumHigh:  averageActiveSessionNumHigh,
		CacheMissRatioScore:          cacheMissRatioScore,
		CacheMissRatioData:           cacheMissRatioData,
		CacheMissRatioHigh:           cacheMissRatioHigh,
		SlowQueryScore:               slowQueryScore,
		SlowQueryData:                slowQueryData,
		SlowQueryAdvice:              slowQueryAdvice,
	}
}

func NewEmptyResult() *Result {
	return &Result{}
}

func (r *Result) Identity() int {
	return r.ID
}

func (r *Result) GetOperationID() int {
	return r.OperationID
}

func (r *Result) GetWeightedAverageScore() int {
	return r.WeightedAverageScore
}

func (r *Result) GetDBConfigScore() int {
	return r.DBConfigScore
}

func (r *Result) GetDBConfigData() string {
	return r.DBConfigData
}

func (r *Result) GetDBConfigAdvice() string {
	return r.DBConfigAdvice
}

func (r *Result) GetCPUUsageScore() int {
	return r.CPUUsageScore
}

func (r *Result) GetCPUUsageData() string {
	return r.CPUUsageData
}

func (r *Result) GetCPUUsageHigh() string {
	return r.CPUUsageHigh
}

func (r *Result) GetIOUtilScore() int {
	return r.IOUtilScore
}

func (r *Result) GetIOUtilData() string {
	return r.IOUtilData
}

func (r *Result) GetIOUtilHigh() string {
	return r.IOUtilHigh
}

func (r *Result) GetDiskCapacityUsageScore() int {
	return r.DiskCapacityUsageScore
}

func (r *Result) GetDiskCapacityUsageData() string {
	return r.DiskCapacityUsageData
}

func (r *Result) GetDiskCapacityUsageHigh() string {
	return r.DiskCapacityUsageHigh
}

func (r *Result) GetConnectionUsageScore() int {
	return r.ConnectionUsageScore
}

func (r *Result) GetConnectionUsageData() string {
	return r.ConnectionUsageData
}

func (r *Result) GetConnectionUsageHigh() string {
	return r.ConnectionUsageHigh
}

func (r *Result) GetAverageActiveSessionNumScore() int {
	return r.AverageActiveSessionNumScore
}

func (r *Result) GetAverageActiveSessionNumData() string {
	return r.AverageActiveSessionNumData
}

func (r *Result) GetAverageActiveSessionNumHigh() string {
	return r.AverageActiveSessionNumHigh
}

func (r *Result) GetCacheMissRatioScore() int {
	return r.CacheMissRatioScore
}

func (r *Result) GetCacheMissRatioData() string {
	return r.CacheMissRatioData
}

func (r *Result) GetCacheMissRatioHigh() string {
	return r.CacheMissRatioHigh
}

func (r *Result) GetSlowQueryScore() int {
	return r.SlowQueryScore
}

func (r *Result) GetSlowQueryData() string {
	return r.SlowQueryData
}

func (r *Result) GetSlowQueryAdvice() string {
	return r.SlowQueryAdvice
}

func (r *Result) GetAccurateReview() int {
	return r.AccurateReview
}

func (r *Result) GetDelFlag() int {
	return r.DelFlag
}

func (r *Result) GetCreateTime() time.Time {
	return r.CreateTime
}

func (r *Result) GetLastUpdateTime() time.Time {
	return r.LastUpdateTime
}

// Set sets DB with given fields, key is the field name and value is the relevant value of the key
func (r *Result) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(r, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Result) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(r, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the DB to json string
func (r *Result) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(r, fields...)
}

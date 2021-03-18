package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MonitorSystemInfo)(nil)

type MonitorSystemInfo struct {
	dependency.Repository
	ID                       int       `middleware:"id" json:"id"`
	MonitorSystemName        string    `middleware:"system_name" json:"system_name"`
	MonitorSystemType        int       `middleware:"system_type" json:"system_type"`
	MonitorSystemHostIP      string    `middleware:"host_ip" json:"host_ip"`
	MonitorSystemPortNum     int       `middleware:"port_num" json:"port_num"`
	MonitorSystemPortNumSlow int       `middleware:"port_num_slow" json:"port_num_slow"`
	BaseUrl                  string    `middleware:"base_url" json:"base_url"`
	DelFlag                  int       `middleware:"del_flag" json:"del_flag"`
	CreateTime               time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime           time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMonitorSystemInfo returns a new *MonitorSystemInfo
func NewMonitorSystemInfo(repo *MonitorSystemRepo, id int, systemName string, systemType int, hostIP string, portNum int, portNumSlow int, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		repo,
		id,
		systemName,
		systemType,
		hostIP,
		portNum,
		portNumSlow,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMonitorSystemInfo returns a new *MonitorSystemInfo with default MonitorSystemRepo
func NewMonitorSystemInfoWithGlobal(id int, monitorSystemName string, systemType int, hostIP string, portNum int, portNumSlow int, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		NewMonitorSystemRepoWithGlobal(),
		id,
		monitorSystemName,
		systemType,
		hostIP,
		portNum,
		portNumSlow,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMonitorSystemInfoWithGlobal returns a new *MonitorSystemInfo with global repository
func NewEmptyMonitorSystemInfoWithGlobal() *MonitorSystemInfo {
	return &MonitorSystemInfo{Repository: NewMonitorSystemRepoWithGlobal()}
}

// NewMonitorSystemInfoWithDefault returns a new *MonitorSystemInfo with default MonitorSystemRepo
func NewMonitorSystemInfoWithDefault(monitorSystemName string, systemType int, hostIP string, portNum int, portNumSlow int, baseUrl string) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		Repository:               NewMonitorSystemRepoWithGlobal(),
		MonitorSystemName:        monitorSystemName,
		MonitorSystemType:        systemType,
		MonitorSystemHostIP:      hostIP,
		MonitorSystemPortNum:     portNum,
		MonitorSystemPortNumSlow: portNumSlow,
		BaseUrl:                  baseUrl,
	}
}

// NewMonitorSystemInfoWithMapAndRandom returns a new *MonitorSystemInfo with given map
func NewMonitorSystemInfoWithMapAndRandom(fields map[string]interface{}) (*MonitorSystemInfo, error) {
	msi := &MonitorSystemInfo{}
	err := common.SetValuesWithMapAndRandom(msi, fields)
	if err != nil {
		return nil, err
	}

	return msi, nil
}

// Identity returns ID of entity
func (msi *MonitorSystemInfo) Identity() string {
	return strconv.Itoa(msi.ID)
}

// IsDeleted checks if delete flag had been set
func (msi *MonitorSystemInfo) IsDeleted() bool {
	return msi.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (msi *MonitorSystemInfo) GetCreateTime() time.Time {
	return msi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (msi *MonitorSystemInfo) GetLastUpdateTime() time.Time {
	return msi.LastUpdateTime
}

// Get returns value of given field
func (msi *MonitorSystemInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(msi, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (msi *MonitorSystemInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(msi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (msi *MonitorSystemInfo) Delete() {
	msi.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (msi *MonitorSystemInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(msi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (msi *MonitorSystemInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(msi, fields...)
}

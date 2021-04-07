package metadata

import (
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.MonitorSystem = (*MonitorSystemInfo)(nil)

type MonitorSystemInfo struct {
	metadata.MonitorSystemRepo
	ID                       int       `middleware:"id" json:"id"`
	MonitorSystemName        string    `middleware:"system_name" json:"system_name"`
	MonitorSystemType        int       `middleware:"system_type" json:"system_type"`
	MonitorSystemHostIP      string    `middleware:"host_ip" json:"host_ip"`
	MonitorSystemPortNum     int       `middleware:"port_num" json:"port_num"`
	MonitorSystemPortNumSlow int       `middleware:"port_num_slow" json:"port_num_slow"`
	BaseURL                  string    `middleware:"base_url" json:"base_url"`
	EnvID                    int       `middleware:"env_id" json:"env_id"`
	DelFlag                  int       `middleware:"del_flag" json:"del_flag"`
	CreateTime               time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime           time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMonitorSystemInfo returns a new *MonitorSystemInfo
func NewMonitorSystemInfo(repo *MonitorSystemRepo, id int, systemName string, systemType int,
	hostIP string, portNum int, portNumSlow int, baseURL string, envID int, delFlag int,
	createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		repo,
		id,
		systemName,
		systemType,
		hostIP,
		portNum,
		portNumSlow,
		baseURL,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMonitorSystemInfoWithGlobal NewMonitorSystemInfo returns a new MonitorSystemInfo with default MonitorSystemRepo
func NewMonitorSystemInfoWithGlobal(id int, monitorSystemName string, systemType int, hostIP string, portNum int,
	portNumSlow int, baseURL string, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		NewMonitorSystemRepoWithGlobal(),
		id,
		monitorSystemName,
		systemType,
		hostIP,
		portNum,
		portNumSlow,
		baseURL,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMonitorSystemInfoWithGlobal return a new MonitorSystemInfo
func NewEmptyMonitorSystemInfoWithGlobal() *MonitorSystemInfo {
	return &MonitorSystemInfo{MonitorSystemRepo: NewMonitorSystemRepoWithGlobal()}
}

// NewMonitorSystemInfoWithDefault returns a new *MonitorSystemInfo with default MonitorSystemRepo
func NewMonitorSystemInfoWithDefault(monitorSystemName string, systemType int, hostIP string, portNum int,
	portNumSlow int, baseURL string, envID int) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		MonitorSystemRepo:        NewMonitorSystemRepoWithGlobal(),
		MonitorSystemName:        monitorSystemName,
		MonitorSystemType:        systemType,
		MonitorSystemHostIP:      hostIP,
		MonitorSystemPortNum:     portNum,
		MonitorSystemPortNumSlow: portNumSlow,
		BaseURL:                  baseURL,
		EnvID:                    envID,
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

// Identity returns the identity
func (msi *MonitorSystemInfo) Identity() int {
	return msi.ID
}

// GetSystemName returns the monitor system name
func (msi *MonitorSystemInfo) GetSystemName() string {
	return msi.MonitorSystemName
}

// GetSystemType returns the monitor system type
func (msi *MonitorSystemInfo) GetSystemType() int {
	return msi.MonitorSystemType
}

// GetHostIP returns the monitor system hostIP
func (msi *MonitorSystemInfo) GetHostIP() string {
	return msi.MonitorSystemHostIP
}

// GetPortNum returns the monitor system portNum
func (msi *MonitorSystemInfo) GetPortNum() int {
	return msi.MonitorSystemPortNum
}

// GetPortNumSlow returns the monitor system portNumSlow
func (msi *MonitorSystemInfo) GetPortNumSlow() int {
	return msi.MonitorSystemPortNumSlow
}

// GetBaseURL returns the monitor system baseUrl
func (msi *MonitorSystemInfo) GetBaseURL() string {
	return msi.BaseURL
}

// GeEnvID returns the monitor system envID
func (msi *MonitorSystemInfo) GetEnvID() int {
	return msi.EnvID
}

// GetDelFlag returns the delete flag
func (msi *MonitorSystemInfo) GetDelFlag() int {
	return msi.DelFlag
}

// GetCreateTime returns created time of entity
func (msi *MonitorSystemInfo) GetCreateTime() time.Time {
	return msi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (msi *MonitorSystemInfo) GetLastUpdateTime() time.Time {
	return msi.LastUpdateTime
}

// Set sets MonitorSystem with given fields, key is the field name and value is the relevant value of the key
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

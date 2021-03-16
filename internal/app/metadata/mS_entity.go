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
	ID             int       `middleware:"id" json:"id"`
	MSName         string    `middleware:"system_name" json:"system_name"`
	SystemType     string    `middleware:"system_type" json:"system_type"`
	HostIp         string    `middleware:"host_ip" json:"host_ip"`
	PortNum        string    `middleware:"port_num" json:"port_num"`
	PortNumSlow    string    `middleware:"port_num_slow" json:"port_num_slow"`
	BaseUrl        string    `middleware:"base_url" json:"base_url"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMSInfo returns a new MonitorSystemInfo
func NewMSInfo(repo *MSRepo, id int, systemName string, systemType string, hostIp string, portNum string, portNumSlow string, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		repo,
		id,
		systemName,
		systemType,
		hostIp,
		portNum,
		portNumSlow,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMSInfo returns a new MonitorSystemInfo with default MSRepo
func NewMSInfoWithGlobal(id int, mSName, systemType, hostIp, portNum, portNumSlow, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		NewMSRepoWithGlobal(),
		id,
		mSName,
		systemType,
		hostIp,
		portNum,
		portNumSlow,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyMSInfoWithGlobal() *MonitorSystemInfo {
	return &MonitorSystemInfo{Repository: NewMSRepoWithGlobal()}
}

// NewMSInfoWithDefault returns a new MonitorSystemInfo with default MSRepo
func NewMSInfoWithDefault(mSName, systemType, hostIp, portNum, portNumSlow, baseUrl string) *MonitorSystemInfo {
	return &MonitorSystemInfo{
		Repository:  NewMSRepoWithGlobal(),
		MSName:      mSName,
		SystemType:  systemType,
		HostIp:      hostIp,
		PortNum:     portNum,
		PortNumSlow: portNumSlow,
		BaseUrl:     baseUrl,
	}
}

// Identity returns ID of entity
func (ei *MonitorSystemInfo) Identity() string {
	return strconv.Itoa(ei.ID)
}

// IsDeleted checks if delete flag had been set
func (ei *MonitorSystemInfo) IsDeleted() bool {
	return ei.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (ei *MonitorSystemInfo) GetCreateTime() time.Time {
	return ei.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (ei *MonitorSystemInfo) GetLastUpdateTime() time.Time {
	return ei.LastUpdateTime
}

// Get returns value of given field
func (ei *MonitorSystemInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(ei, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (ei *MonitorSystemInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ei, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (ei *MonitorSystemInfo) Delete() {
	ei.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (ei *MonitorSystemInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ei, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (ei *MonitorSystemInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ei, fields...)
}

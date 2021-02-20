package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MSInfo)(nil)

type MSInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	MSName        string    `middleware:"system_name" json:"system_name"`
	HostIp string `middleware:"host_ip" json:"host_ip"`
	PortNum string `middleware:"port_num" json:"port_num"`
	BaseUrl string `middleware:"base_url" json:"base_url"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMSInfo returns a new MSInfo
func NewMSInfo(repo *MSRepo, id int, systemName string, hostIp string, portNum string, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MSInfo {
	return &MSInfo{
		repo,
		id,
		systemName,
		hostIp,
		portNum,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMSInfo returns a new MSInfo with default MSRepo
func NewMSInfoWithGlobal(id int, systemName string, hostIp string, portNum string, baseUrl string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MSInfo {
	return &MSInfo{
		NewMSRepoWithGlobal(),
		id,
		systemName,
		hostIp,
		portNum,
		baseUrl,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyMSInfoWithGlobal() *MSInfo {
	return &MSInfo{Repository: NewMSRepoWithGlobal()}
}

// NewMSInfoWithDefault returns a new MSInfo with default MSRepo
func NewMSInfoWithDefault(mSName, hostIp, baseUrl string, portNum string) *MSInfo {
	return &MSInfo{
		Repository: NewMSRepoWithGlobal(),
		MSName:    mSName,
		HostIp:  hostIp,
		PortNum: portNum,
		BaseUrl: baseUrl,
	}
}

// Identity returns ID of entity
func (ei *MSInfo) Identity() string {
	return strconv.Itoa(ei.ID)
}

// IsDeleted checks if delete flag had been set
func (ei *MSInfo) IsDeleted() bool {
	return ei.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (ei *MSInfo) GetCreateTime() time.Time {
	return ei.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (ei *MSInfo) GetLastUpdateTime() time.Time {
	return ei.LastUpdateTime
}

// Get returns value of given field
func (ei *MSInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(ei, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (ei *MSInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ei, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (ei *MSInfo) Delete() {
	ei.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (ei *MSInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ei, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (ei *MSInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ei, fields...)
}

package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MYSQLServerInfo)(nil)

// MYSQLServerInfo is a struct map to table in the database
type MYSQLServerInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	ClusterID      int       `middleware:"cluster_id" json:"cluster_id"`
	HostIP         string    `middleware:"host_ip" json:"host_ip"`
	PortNum        int       `middleware:"port_num" json:"port_num"`
	DeploymentType int       `middleware:"deployment_type" json:"deployment_type"`
	Version        string    `middleware:"version" json:"version"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMYSQLServerInfo returns a new MYSQLServerInfo
func NewMYSQLServerInfo(repo *MYSQLServerRepo,
	id int,
	clusterID int,
	hostIP string,
	portNum int,
	deploymentType int,
	version string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MYSQLServerInfo {
	return &MYSQLServerInfo{
		repo,
		id,
		clusterID,
		hostIP,
		portNum,
		deploymentType,
		version,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMYSQLServerInfoWithGlobal returns a new MYSQLServerInfo with default MYSQLServerRepo
func NewMYSQLServerInfoWithGlobal(
	id int,
	clusterID int,
	hostIP string,
	portNum int,
	deploymentType int,
	version string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MYSQLServerInfo {
	return &MYSQLServerInfo{
		NewMYSQLServerRepoWithGlobal(),
		id,
		clusterID,
		hostIP,
		portNum,
		deploymentType,
		version,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMYSQLServerInfoWithGlobal returns a new MYSQLServerInfo with default MYSQLServerRepo
func NewEmptyMYSQLServerInfoWithGlobal() *MYSQLServerInfo {
	return &MYSQLServerInfo{Repository: NewMYSQLServerRepoWithGlobal()}
}

// NewMYSQLServerInfoWithDefault returns a new MYSQLServerInfo with default MYSQLServerRepo
func NewMYSQLServerInfoWithDefault(hostIP string, portNum int) *MYSQLServerInfo {
	return &MYSQLServerInfo{
		Repository: NewMYSQLServerRepoWithGlobal(),
		HostIP:     hostIP,
		PortNum:    portNum,
	}
}

// Identity returns ID of entity
func (mci *MYSQLServerInfo) Identity() string {
	return strconv.Itoa(mci.ID)
}

// IsDeleted checks if delete flag had been set
func (mci *MYSQLServerInfo) IsDeleted() bool {
	return mci.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (mci *MYSQLServerInfo) GetCreateTime() time.Time {
	return mci.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (mci *MYSQLServerInfo) GetLastUpdateTime() time.Time {
	return mci.LastUpdateTime
}

// Get returns value of given field
func (mci *MYSQLServerInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(mci, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (mci *MYSQLServerInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(mci, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (mci *MYSQLServerInfo) Delete() {
	mci.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (mci *MYSQLServerInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(mci, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (mci *MYSQLServerInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mci, fields...)
}

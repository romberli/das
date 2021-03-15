package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MySQLServerInfo)(nil)

// MySQLServerInfo is a struct map to table in the database
type MySQLServerInfo struct {
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

// NewMySQLServerInfo returns a new MySQLServerInfo
func NewMySQLServerInfo(repo *MySQLServerRepo,
	id int,
	clusterID int,
	hostIP string,
	portNum int,
	deploymentType int,
	version string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MySQLServerInfo {
	return &MySQLServerInfo{
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

// NewMySQLServerInfoWithGlobal returns a new MySQLServerInfo with default MySQLServerRepo
func NewMySQLServerInfoWithGlobal(
	id int,
	clusterID int,
	hostIP string,
	portNum int,
	deploymentType int,
	version string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MySQLServerInfo {
	return &MySQLServerInfo{
		NewMySQLServerRepoWithGlobal(),
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

// NewEmptyMySQLServerInfoWithGlobal returns a new MySQLServerInfo with default MySQLServerRepo
func NewEmptyMySQLServerInfoWithGlobal() *MySQLServerInfo {
	return &MySQLServerInfo{Repository: NewMySQLServerRepoWithGlobal()}
}

// NewMySQLServerInfoWithDefault returns a new MySQLServerInfo with default MySQLServerRepo
func NewMySQLServerInfoWithDefault(hostIP string, portNum int) *MySQLServerInfo {
	return &MySQLServerInfo{
		Repository: NewMySQLServerRepoWithGlobal(),
		HostIP:     hostIP,
		PortNum:    portNum,
	}
}

// Identity returns ID of entity
func (msi *MySQLServerInfo) Identity() string {
	return strconv.Itoa(msi.ID)
}

// IsDeleted checks if delete flag had been set
func (msi *MySQLServerInfo) IsDeleted() bool {
	return msi.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (msi *MySQLServerInfo) GetCreateTime() time.Time {
	return msi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (msi *MySQLServerInfo) GetLastUpdateTime() time.Time {
	return msi.LastUpdateTime
}

// Get returns value of given field
func (msi *MySQLServerInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(msi, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (msi *MySQLServerInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(msi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (msi *MySQLServerInfo) Delete() {
	msi.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (msi *MySQLServerInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(msi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (msi *MySQLServerInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(msi, fields...)
}

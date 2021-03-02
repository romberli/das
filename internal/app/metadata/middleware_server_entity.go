package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MiddlewareServerInfo)(nil)

type MiddlewareServerInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	ClusterID      int       `middleware:"cluster_id" json:"cluster_id"`
	ServerName     string    `middleware:"server_name" json:"server_name"`
	MiddlewareRole int       `middleware:"middleware_role" json:"middleware_role"`
	HostIP         string    `middleware:"host_ip" json:"host_ip"`
	PortNum        int       `middleware:"port_num" json:"port_num"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMiddlewareServerInfo returns a new MiddlewareServerInfo
func NewMiddlewareServerInfo(repo *MiddlewareServerRepo, id int, clusterID int, serverName string, middlewareRole int, hostIP string, portNum int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareServerInfo {
	return &MiddlewareServerInfo{
		repo,
		id,
		clusterID,
		serverName,
		middlewareRole,
		hostIP,
		portNum,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMiddlewareServerInfo returns a new MiddlewareServerInfo with default MiddlewareServerRepo
func NewMiddlewareServerInfoWithGlobal(id int, clusterID int, serverName string, middlewareRole int, hostIP string, portNum int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareServerInfo {
	return &MiddlewareServerInfo{
		NewMiddlewareServerRepoWithGlobal(),
		id,
		clusterID,
		serverName,
		middlewareRole,
		hostIP,
		portNum,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyMiddlewareServerInfoWithGlobal() *MiddlewareServerInfo {
	return &MiddlewareServerInfo{Repository: NewMiddlewareServerRepoWithGlobal()}
}

// NewMiddlewareServerInfoWithDefault returns a new MiddlewareServerInfo with default MiddlewareServerRepo
func NewMiddlewareServerInfoWithDefault(clusterID int, serverName string, middlewareRole int, hostIP string, portNum int) *MiddlewareServerInfo {
	return &MiddlewareServerInfo{
		Repository:     NewMiddlewareServerRepoWithGlobal(),
		ClusterID:      clusterID,
		ServerName:     serverName,
		MiddlewareRole: middlewareRole,
		HostIP:         hostIP,
		PortNum:        portNum,
	}
}

// Identity returns ID of entity
func (msi *MiddlewareServerInfo) Identity() string {
	return strconv.Itoa(msi.ID)
}

// IsDeleted checks if delete flag had been set
func (msi *MiddlewareServerInfo) IsDeleted() bool {
	return msi.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (msi *MiddlewareServerInfo) GetCreateTime() time.Time {
	return msi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (msi *MiddlewareServerInfo) GetLastUpdateTime() time.Time {
	return msi.LastUpdateTime
}

// Get returns value of given field
func (msi *MiddlewareServerInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(msi, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (msi *MiddlewareServerInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(msi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (msi *MiddlewareServerInfo) Delete() {
	msi.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (msi *MiddlewareServerInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(msi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (msi *MiddlewareServerInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(msi, fields...)
}

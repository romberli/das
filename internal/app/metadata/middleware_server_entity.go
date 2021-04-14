package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	middlewareServerClusterIDStruct      = "ClusterID"
	middlewareServerNameStruct           = "ServerName"
	middlewareServerMiddlewareRoleStruct = "MiddlewareRole"
	middlewareServerHostIPStruct         = "HostIP"
	middlewareServerPortNumStruct        = "PortNum"
)

var _ metadata.MiddlewareServer = (*MiddlewareServerInfo)(nil)

type MiddlewareServerInfo struct {
	metadata.MiddlewareServerRepo
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
func NewMiddlewareServerInfo(repo metadata.MiddlewareServerRepo, id int, clusterID int, serverName string, middlewareRole int, hostIP string, portNum int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareServerInfo {
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
	return &MiddlewareServerInfo{MiddlewareServerRepo: NewMiddlewareServerRepoWithGlobal()}
}

// NewMiddlewareServerInfoWithDefault returns a new MiddlewareServerInfo with default MiddlewareServerRepo
func NewMiddlewareServerInfoWithDefault(clusterID int, serverName string, middlewareRole int, hostIP string, portNum int) *MiddlewareServerInfo {
	return &MiddlewareServerInfo{
		MiddlewareServerRepo: NewMiddlewareServerRepoWithGlobal(),
		ClusterID:            clusterID,
		ServerName:           serverName,
		MiddlewareRole:       middlewareRole,
		HostIP:               hostIP,
		PortNum:              portNum,
	}
}

// NewMiddlewareServerInfoWithMapAndRandom returns a new *MiddlewareServerInfo with given map
func NewMiddlewareServerInfoWithMapAndRandom(fields map[string]interface{}) (*MiddlewareServerInfo, error) {
	msi := &MiddlewareServerInfo{}
	err := common.SetValuesWithMapAndRandom(msi, fields)
	if err != nil {
		return nil, err
	}

	return msi, nil
}

// Identity returns ID of entity
func (msi *MiddlewareServerInfo) Identity() int {
	return msi.ID
}

// GetClusterID returns the middleware cluster id
func (msi *MiddlewareServerInfo) GetClusterID() int {
	return msi.ClusterID
}

// GetServerName returns the server name
func (msi *MiddlewareServerInfo) GetServerName() string {
	return msi.ServerName
}

// GetMiddlewareRole returns the middleware role
func (msi *MiddlewareServerInfo) GetMiddlewareRole() int {
	return msi.MiddlewareRole
}

// GetHostIP returns the host ip
func (msi *MiddlewareServerInfo) GetHostIP() string {
	return msi.HostIP
}

// GetPortNum returns the port number
func (msi *MiddlewareServerInfo) GetPortNum() int {
	return msi.PortNum
}

// IsDeleted checks if delete flag had been set
func (msi *MiddlewareServerInfo) GetDelFlag() int {
	return msi.DelFlag
}

// GetCreateTime returns created time of entity
func (msi *MiddlewareServerInfo) GetCreateTime() time.Time {
	return msi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (msi *MiddlewareServerInfo) GetLastUpdateTime() time.Time {
	return msi.LastUpdateTime
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

package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*MYSQLClusterInfo)(nil)

// MYSQLClusterInfo is a struct map to table in the database
type MYSQLClusterInfo struct {
	dependency.Repository
	ID                  int       `middleware:"id" json:"id"`
	ClusterName         string    `middleware:"cluster_name" json:"cluster_name"`
	MiddlewareClusterID int       `middleware:"middleware_cluster_id" json:"middleware_cluster_id"`
	MonitorSystemID     int       `middleware:"monitor_system_id" json:"monitor_system_id"`
	OwnerID             int       `middleware:"owner_id" json:"owner_id"`
	OwnerGroup          string    `middleware:"owner_group" json:"owner_group"`
	EnvID               int       `middleware:"env_id" json:"env_id"`
	DelFlag             int       `middleware:"del_flag" json:"del_flag"`
	CreateTime          time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime      time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMYSQLClusterInfo returns a new MYSQLClusterInfo
func NewMYSQLClusterInfo(repo *MYSQLClusterRepo,
	id int,
	clusterName string,
	middlewareClusterID int,
	monitorSystemID int,
	ownerID int, ownerGroup string,
	envID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MYSQLClusterInfo {
	return &MYSQLClusterInfo{
		repo,
		id,
		clusterName,
		middlewareClusterID,
		monitorSystemID,
		ownerID,
		ownerGroup,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMYSQLClusterInfoWithGlobal returns a new MYSQLClusterInfo with default MYSQLClusterRepo
func NewMYSQLClusterInfoWithGlobal(
	id int,
	clusterName string,
	middlewareClusterID int,
	monitorSystemID int,
	ownerID int, ownerGroup string,
	envID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MYSQLClusterInfo {
	return &MYSQLClusterInfo{
		NewMYSQLClusterRepoWithGlobal(),
		id,
		clusterName,
		middlewareClusterID,
		monitorSystemID,
		ownerID,
		ownerGroup,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMYSQLClusterInfoWithGlobal returns a new MYSQLClusterInfo with default MYSQLClusterRepo
func NewEmptyMYSQLClusterInfoWithGlobal() *MYSQLClusterInfo {
	return &MYSQLClusterInfo{Repository: NewMYSQLClusterRepoWithGlobal()}
}

// NewMYSQLClusterInfoWithDefault returns a new MYSQLClusterInfo with default MYSQLClusterRepo
func NewMYSQLClusterInfoWithDefault(clusterName string) *MYSQLClusterInfo {
	return &MYSQLClusterInfo{
		Repository:  NewMYSQLClusterRepoWithGlobal(),
		ClusterName: clusterName,
	}
}

// Identity returns ID of entity
func (mci *MYSQLClusterInfo) Identity() string {
	return strconv.Itoa(mci.ID)
}

// IsDeleted checks if delete flag had been set
func (mci *MYSQLClusterInfo) IsDeleted() bool {
	return mci.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (mci *MYSQLClusterInfo) GetCreateTime() time.Time {
	return mci.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (mci *MYSQLClusterInfo) GetLastUpdateTime() time.Time {
	return mci.LastUpdateTime
}

// Get returns value of given field
func (mci *MYSQLClusterInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(mci, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (mci *MYSQLClusterInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(mci, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (mci *MYSQLClusterInfo) Delete() {
	mci.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (mci *MYSQLClusterInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(mci, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (mci *MYSQLClusterInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mci, fields...)
}

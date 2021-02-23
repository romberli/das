package metadata

import (
	"github.com/romberli/das/internal/dependency"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"strconv"
	"time"
)

var _ dependency.Entity = (*MiddlewareClusterInfo)(nil)

type MiddlewareClusterInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	ClusterName    string    `middleware:"cluster_name" json:"cluster_name"`
	EnvID          int       `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMiddlewareClusterInfo returns a new MiddlewareClusterInfo
func NewMiddlewareClusterInfo(repo *MiddlewareClusterRepo, id int, middlewareClusterName string, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		repo,
		id,
		middlewareClusterName,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMiddlewareClusterInfo returns a new MiddlewareClusterInfo with default MiddlewareClusterRepo
func NewMiddlewareClusterInfoWithGlobal(id int, middlewareClusterName string, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		NewMiddlewareClusterRepoWithGlobal(),
		id,
		middlewareClusterName,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyMiddlewareClusterInfoWithGlobal() *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{Repository: NewMiddlewareClusterRepoWithGlobal()}
}

// NewMiddlewareClusterInfoWithDefault returns a new MiddlewareClusterInfo with default MiddlewareClusterRepo
func NewMiddlewareClusterInfoWithDefault(middlewareClusterName string, envID int) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		Repository:  NewMiddlewareClusterRepoWithGlobal(),
		ClusterName: middlewareClusterName,
		EnvID:       envID,
	}
}

// Identity returns ID of entity
func (mci *MiddlewareClusterInfo) Identity() string {
	return strconv.Itoa(mci.ID)
}

// IsDeleted checks if delete flag had been set
func (mci *MiddlewareClusterInfo) IsDeleted() bool {
	return mci.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (mci *MiddlewareClusterInfo) GetCreateTime() time.Time {
	return mci.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (mci *MiddlewareClusterInfo) GetLastUpdateTime() time.Time {
	return mci.LastUpdateTime
}

// Get returns value of given field
func (mci *MiddlewareClusterInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(mci, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (mci *MiddlewareClusterInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(mci, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (mci *MiddlewareClusterInfo) Delete() {
	mci.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (mci *MiddlewareClusterInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(mci, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (mci *MiddlewareClusterInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mci, fields...)
}

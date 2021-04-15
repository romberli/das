package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	middlewareClusterNameStruct    = "ClusterName"
	middlewareClusterOwnerIDStruct = "OwnerID"
	middlewareClusterEnvIDStruct   = "EnvID"
)

var _ metadata.MiddlewareCluster = (*MiddlewareClusterInfo)(nil)

type MiddlewareClusterInfo struct {
	metadata.MiddlewareClusterRepo
	ID             int       `middleware:"id" json:"id"`
	ClusterName    string    `middleware:"cluster_name" json:"cluster_name"`
	OwnerID        int       `middleware:"owner_id" json:"owner_id"`
	EnvID          int       `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMiddlewareClusterInfo returns a new MiddlewareClusterInfo
func NewMiddlewareClusterInfo(repo metadata.MiddlewareClusterRepo, id int, middlewareClusterName string, ownerID int, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		repo,
		id,
		middlewareClusterName,
		ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMiddlewareClusterInfo returns a new MiddlewareClusterInfo with default MiddlewareClusterRepo
func NewMiddlewareClusterInfoWithGlobal(id int, middlewareClusterName string, ownerID int, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		NewMiddlewareClusterRepoWithGlobal(),
		id,
		middlewareClusterName,
		ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMiddlewareClusterInfoWithGlobal return a new MiddlewareClusterInfo
func NewEmptyMiddlewareClusterInfoWithGlobal() *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{MiddlewareClusterRepo: NewMiddlewareClusterRepoWithGlobal()}
}

// NewMiddlewareClusterInfoWithDefault returns a new MiddlewareClusterInfo with default MiddlewareClusterRepo
func NewMiddlewareClusterInfoWithDefault(middlewareClusterName string, ownerID, envID int) *MiddlewareClusterInfo {
	return &MiddlewareClusterInfo{
		MiddlewareClusterRepo: NewMiddlewareClusterRepoWithGlobal(),
		ClusterName:           middlewareClusterName,
		OwnerID:               ownerID,
		EnvID:                 envID,
	}
}

// NewMiddlewareClusterInfoWithMapAndRandom returns a new *MiddlewareClusterInfoInfo with given map
func NewMiddlewareClusterInfoWithMapAndRandom(fields map[string]interface{}) (*MiddlewareClusterInfo, error) {
	mci := &MiddlewareClusterInfo{}
	err := common.SetValuesWithMapAndRandom(mci, fields)
	if err != nil {
		return nil, err
	}

	return mci, nil
}

// Identity returns ID of entity
func (mci *MiddlewareClusterInfo) Identity() int {
	return mci.ID
}

// GetClusterName returns the cluster name
func (mci *MiddlewareClusterInfo) GetClusterName() string {
	return mci.ClusterName
}

// GetOwnerID returns the owner id
func (mci *MiddlewareClusterInfo) GetOwnerID() int {
	return mci.OwnerID
}

// GetEnvID returns the env id
func (mci *MiddlewareClusterInfo) GetEnvID() int {
	return mci.EnvID
}

// GetDelFlag returns the delete flag
func (mci *MiddlewareClusterInfo) GetDelFlag() int {
	return mci.DelFlag
}

// GetCreateTime returns created time of entity
func (mci *MiddlewareClusterInfo) GetCreateTime() time.Time {
	return mci.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (mci *MiddlewareClusterInfo) GetLastUpdateTime() time.Time {
	return mci.LastUpdateTime
}

// GetMiddlewareServerIDList gets the middleware server id list of this cluster
func (mci *MiddlewareClusterInfo) GetMiddlewareServerIDList() ([]int, error) {
	return mci.MiddlewareClusterRepo.GetMiddlewareServerIDList(mci.Identity())
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

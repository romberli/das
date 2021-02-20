package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*DbInfo)(nil)

type DbInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	DbName        string    `middleware:"db_name" json:"db_name"`
	ClusterId int `middleware:"cluster_id" json:"cluster_id"`
	ClusterType int `middleware:"cluster_type" json:"cluster_type"`
	OwnerId string `middleware:"owner_id" json:"owner_id"`
	OwnerGroup string `middleware:"owner_group" json:"owner_group"`
	EnvId string `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDbInfo returns a new DbInfo
func NewDbInfo(repo *DbRepo, id int, dbName string, clusterId int, clusterType int, ownerId string, ownerGroup string, envId string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DbInfo {
	return &DbInfo{
		repo,
		id,
		dbName,
		clusterId,
		clusterType,
		ownerId,
		ownerGroup,
		envId,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewDbInfo returns a new DbInfo with default DbRepo
func NewDbInfoWithGlobal(id int, dbName string, clusterId int, clusterType int, ownerId string, ownerGroup string, envId string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DbInfo {
	return &DbInfo{
		NewDbRepoWithGlobal(),
		id,
		dbName,
		clusterId,
		clusterType,
		ownerId,
		ownerGroup,
		envId,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyDbInfoWithGlobal() *DbInfo {
	return &DbInfo{Repository: NewDbRepoWithGlobal()}
}

// NewDbInfoWithDefault returns a new DbInfo with default DbRepo
func NewDbInfoWithDefault(dbName, ownerId, envId string) *DbInfo {
	return &DbInfo{
		Repository: NewDbRepoWithGlobal(),
		DbName:    dbName,
		OwnerId:   ownerId,
		EnvId:     envId,
	}
}

// Identity returns ID of entity
func (ei *DbInfo) Identity() string {
	return strconv.Itoa(ei.ID)
}

// IsDeleted checks if delete flag had been set
func (ei *DbInfo) IsDeleted() bool {
	return ei.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (ei *DbInfo) GetCreateTime() time.Time {
	return ei.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (ei *DbInfo) GetLastUpdateTime() time.Time {
	return ei.LastUpdateTime
}

// Get returns value of given field
func (ei *DbInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(ei, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (ei *DbInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ei, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (ei *DbInfo) Delete() {
	ei.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (ei *DbInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ei, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (ei *DbInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ei, fields...)
}

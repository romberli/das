package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*DBInfo)(nil)

type DBInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	DBName         string    `middleware:"db_name" json:"db_name"`
	ClusterID      string    `middleware:"cluster_id" json:"cluster_id"`
	ClusterType    string    `middleware:"cluster_type" json:"cluster_type"`
	OwnerID        string    `middleware:"owner_id" json:"owner_id"`
	OwnerGroup     string    `middleware:"owner_group" json:"owner_group"`
	EnvID          string    `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDBInfo returns a new DBInfo
func NewDBInfo(repo *DBRepo, id int, dbName string, clusterID string, clusterType string, ownerID string, ownerGroup string, envID string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		repo,
		id,
		dbName,
		clusterID,
		clusterType,
		ownerID,
		ownerGroup,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewDBInfo returns a new DBInfo with default DBRepo
func NewDBInfoWithGlobal(id int, dbName string, clusterID string, clusterType string, ownerID string, ownerGroup string, envID string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		NewDBRepoWithGlobal(),
		id,
		dbName,
		clusterID,
		clusterType,
		ownerID,
		ownerGroup,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

func NewEmptyDBInfoWithGlobal() *DBInfo {
	return &DBInfo{Repository: NewDBRepoWithGlobal()}
}

// NewDBInfoWithDefault returns a new DBInfo with default DBRepo
func NewDBInfoWithDefault(dbName, clusterID, ownerID, ownerGroup, envID string) *DBInfo {
	return &DBInfo{
		Repository: NewDBRepoWithGlobal(),
		DBName:     dbName,
		ClusterID:  constant.Default,
		OwnerID:    ownerID,
		OwnerGroup: ownerGroup,
		EnvID:      envID,
	}
}

// Identity returns ID of entity
func (dbi *DBInfo) Identity() string {
	return strconv.Itoa(dbi.ID)
}

// IsDeleted checks if delete flag had been set
func (dbi *DBInfo) IsDeleted() bool {
	return dbi.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (dbi *DBInfo) GetCreateTime() time.Time {
	return dbi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (dbi *DBInfo) GetLastUpdateTime() time.Time {
	return dbi.LastUpdateTime
}

// Get returns value of given field
func (dbi *DBInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(dbi, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (dbi *DBInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(dbi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (dbi *DBInfo) Delete() {
	dbi.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (dbi *DBInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(dbi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (dbi *DBInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(dbi, fields...)
}

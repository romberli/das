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
	ClusterID      int       `middleware:"cluster_id" json:"cluster_id"`
	ClusterType    int       `middleware:"cluster_type" json:"cluster_type"`
	OwnerID        int       `middleware:"owner_id" json:"owner_id"`
	OwnerGroup     string    `middleware:"owner_group" json:"owner_group"`
	EnvID          int       `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDBInfo returns a new *DBInfo
func NewDBInfo(repo *DBRepo, id int, dbName string, clusterID int, clusterType int, ownerID int, ownerGroup string, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
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

// NewDBInfo returns a new *DBInfo with default DBRepo
func NewDBInfoWithGlobal(id int, dbName string, clusterID int, clusterType int, ownerID int, ownerGroup string, envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
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

// NewEmptyDBInfoWithGlobal returns a new *DBInfo with global repository
func NewEmptyDBInfoWithGlobal() *DBInfo {
	return &DBInfo{Repository: NewDBRepoWithGlobal()}
}

// NewDBInfoWithDefault returns a new *DBInfo with default DBRepo
func NewDBInfoWithDefault(dbName string, clusterID, clusterType, envID int) *DBInfo {
	return &DBInfo{
		Repository:  NewDBRepoWithGlobal(),
		DBName:      dbName,
		ClusterID:   clusterID,
		ClusterType: clusterType,
		OwnerID:     constant.DefaultRandomInt,
		OwnerGroup:  constant.DefaultRandomString,
		EnvID:       envID,
	}
}

// NewDBInfoWithMapAndRandom returns a new *DBInfo with given map
func NewDBInfoWithMapAndRandom(fields map[string]interface{}) (*DBInfo, error) {
	dbi := &DBInfo{}
	err := common.SetValuesWithMapAndRandom(dbi, fields)
	if err != nil {
		return nil, err
	}

	return dbi, nil
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

package metadata

import (
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	dbDBNameStruct      = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbEnvIDStruct       = "EnvID"
	db
)

var _ metadata.DB = (*DBInfo)(nil)

type DBInfo struct {
	metadata.DBRepo
	ID             int       `middleware:"id" json:"id"`
	DBName         string    `middleware:"db_name" json:"db_name"`
	ClusterID      int       `middleware:"cluster_id" json:"cluster_id"`
	ClusterType    int       `middleware:"cluster_type" json:"cluster_type"`
	OwnerID        int       `middleware:"owner_id" json:"owner_id"`
	EnvID          int       `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDBInfo returns a new *DBInfo
func NewDBInfo(repo *DBRepo, id int, dbName string, clusterID int, clusterType int, ownerID int,
	envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		repo,
		id,
		dbName,
		clusterID,
		clusterType,
		ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewDBInfoWithGlobal NewDBInfo returns a new DBInfo with default DBRepo
func NewDBInfoWithGlobal(id int, dbName string, clusterID, clusterType, ownerID, envID, delFlag int,
	createTime, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		NewDBRepoWithGlobal(),
		id,
		dbName,
		clusterID,
		clusterType,
		ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyDBInfoWithRepo return a new DBInfo
func NewEmptyDBInfoWithRepo(repo *DBRepo) *DBInfo {
	return &DBInfo{DBRepo: repo}
}

// NewEmptyDBInfoWithGlobal return a new DBInfo
func NewEmptyDBInfoWithGlobal() *DBInfo {
	return NewEmptyDBInfoWithRepo(NewDBRepoWithGlobal())
}

// NewDBInfoWithDefault returns a new *DBInfo with default DBRepo
func NewDBInfoWithDefault(dbName string, clusterID, clusterType, envID int) *DBInfo {
	return &DBInfo{
		DBRepo:      NewDBRepoWithGlobal(),
		DBName:      dbName,
		ClusterID:   clusterID,
		ClusterType: clusterType,
		OwnerID:     constant.DefaultRandomInt,
		EnvID:       envID,
	}
}

// NewDBInfoWithMapAndRandom returns a new *DBInfo with given map
func NewDBInfoWithMapAndRandom(fields map[string]interface{}) (*DBInfo, error) {
	di := &DBInfo{}
	err := common.SetValuesWithMapAndRandom(di, fields)
	if err != nil {
		return nil, err
	}

	return di, nil
}

// Identity returns the identity
func (di *DBInfo) Identity() int {
	return di.ID
}

// GetDBName returns the db name
func (di *DBInfo) GetDBName() string {
	return di.DBName
}

// GetClusterID returns the cluster id
func (di *DBInfo) GetClusterID() int {
	return di.ClusterID
}

// GetClusterType returns the cluster type
func (di *DBInfo) GetClusterType() int {
	return di.ClusterType
}

// GetOwnerID returns the owner id
func (di *DBInfo) GetOwnerID() int {
	return di.OwnerID
}

// GetEnvID returns the env id
func (di *DBInfo) GetEnvID() int {
	return di.EnvID
}

// GetDelFlag returns the delete flag
func (di *DBInfo) GetDelFlag() int {
	return di.DelFlag
}

// GetCreateTime returns the create time
func (di *DBInfo) GetCreateTime() time.Time {
	return di.CreateTime
}

// GetLastUpdateTime returns the last update time
func (di *DBInfo) GetLastUpdateTime() time.Time {
	return di.LastUpdateTime
}

// GetAppIDList gets app identity list that uses this db
func (di *DBInfo) GetAppIDList() ([]int, error) {
	return di.DBRepo.GetAppIDList(di.ID)
}

// Set sets DB with given fields, key is the field name and value is the relevant value of the key
func (di *DBInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(di, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (di *DBInfo) Delete() {
	di.DelFlag = 1
}

// AddApp adds a new map of application system and database in the middleware
func (di *DBInfo) AddApp(appID int) error {
	return di.DBRepo.AddApp(di.ID, appID)
}

// DeleteApp delete the map of application system and database in the middleware
func (di *DBInfo) DeleteApp(appID int) error {
	return di.DBRepo.DeleteApp(di.ID, appID)
}

// MarshalJSON marshals DB to json string
func (di *DBInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(di, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the DB to json string
func (di *DBInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(di, fields...)
}

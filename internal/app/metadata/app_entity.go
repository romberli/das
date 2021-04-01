package metadata

import (
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	appAppNameStruct = "AppName"
	appLevelStruct   = "Level"
)

var _ metadata.App = (*AppInfo)(nil)

type AppInfo struct {
	metadata.AppRepo
	ID             int       `middleware:"id" json:"id"`
	AppName        string    `middleware:"app_name" json:"app_name"`
	Level          int       `middleware:"level" json:"level"`
	OwnerID        int       `middleware:"owner_id" json:"owner_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewAppInfo returns a new AppInfo
func NewAppInfo(repo metadata.AppRepo, id int, appName string, level int, ownerID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *AppInfo {
	return &AppInfo{
		repo,
		id,
		appName,
		level,
		ownerID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewAppInfoWithGlobal NewAppInfo returns a new AppInfo with default AppRepo
func NewAppInfoWithGlobal(id int, appName string, level int, ownerID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *AppInfo {
	return &AppInfo{
		NewAppRepoWithGlobal(),
		id,
		appName,
		level,
		ownerID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyAppInfoWithGlobal return a new AppInfo
func NewEmptyAppInfoWithGlobal() *AppInfo {
	return &AppInfo{AppRepo: NewAppRepoWithGlobal()}
}

// NewAppInfoWithDefault returns a new AppInfo with default value
func NewAppInfoWithDefault(appName string, level int) *AppInfo {
	return &AppInfo{
		AppRepo: NewAppRepoWithGlobal(),
		AppName: appName,
		Level:   level,
		OwnerID: constant.DefaultRandomInt,
	}
}

// NewAppInfoWithMapAndRandom returns a new *AppInfoInfo with given map
func NewAppInfoWithMapAndRandom(fields map[string]interface{}) (*AppInfo, error) {
	asi := &AppInfo{}
	err := common.SetValuesWithMapAndRandom(asi, fields)
	if err != nil {
		return nil, err
	}

	return asi, nil
}

// Identity returns the identity
func (asi *AppInfo) Identity() int {
	return asi.ID
}

// GetSystemName returns the app name
func (asi *AppInfo) GetAppName() string {
	return asi.AppName
}

// GetLevel returns the level
func (asi *AppInfo) GetLevel() int {
	return asi.Level
}

// GetOwnerID returns the owner id
func (asi *AppInfo) GetOwnerID() int {
	return asi.OwnerID
}

// GetDelFlag returns the delete flag
func (asi *AppInfo) GetDelFlag() int {
	return asi.DelFlag
}

// GetCreateTime returns the create time
func (asi *AppInfo) GetCreateTime() time.Time {
	return asi.CreateTime
}

// GetLastUpdateTime returns the last update time
func (asi *AppInfo) GetLastUpdateTime() time.Time {
	return asi.LastUpdateTime
}

// GetDBIDList gets database identity list that the app uses
func (asi *AppInfo) GetDBIDList() ([]int, error) {
	return asi.AppRepo.GetDBIDList(asi.Identity())
}

// Set sets App with given fields, key is the field name and value is the relevant value of the key
func (asi *AppInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(asi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (asi *AppInfo) Delete() {
	asi.DelFlag = 1
}

// AddDB adds a new map of the app and database in the middleware
func (asi *AppInfo) AddDB(dbID int) error {
	return asi.AppRepo.AddDB(asi.Identity(), dbID)
}

// DeleteDB deletes the map of the app and database in the middleware
func (asi *AppInfo) DeleteDB(dbID int) error {
	return asi.AppRepo.DeleteDB(asi.Identity(), dbID)
}

// MarshalJSON marshals App to json bytes
func (asi *AppInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(asi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified fields of the App to json string
func (asi *AppInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(asi, fields...)
}

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

// func NewEmptyAppInfoWithRepo(repo *metadata.AppRepo) *AppInfo{
// 	return &AppInfo{AppRepo: repo}
// }

// // NewEmptyAppInfoWithGlobal return a new AppInfo
// func NewEmptyAppInfoWithGlobal() *AppInfo {
// 	return NewEmptyAppInfoWithRepo(NewAppRepoWithGlobal())
// }

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
	ai := &AppInfo{}
	err := common.SetValuesWithMapAndRandom(ai, fields)
	if err != nil {
		return nil, err
	}

	return ai, nil
}

// Identity returns the identity
func (ai *AppInfo) Identity() int {
	return ai.ID
}

// GetSystemName returns the app name
func (ai *AppInfo) GetAppName() string {
	return ai.AppName
}

// GetLevel returns the level
func (ai *AppInfo) GetLevel() int {
	return ai.Level
}

// GetOwnerID returns the owner id
func (ai *AppInfo) GetOwnerID() int {
	return ai.OwnerID
}

// GetDelFlag returns the delete flag
func (ai *AppInfo) GetDelFlag() int {
	return ai.DelFlag
}

// GetCreateTime returns the create time
func (ai *AppInfo) GetCreateTime() time.Time {
	return ai.CreateTime
}

// GetLastUpdateTime returns the last update time
func (ai *AppInfo) GetLastUpdateTime() time.Time {
	return ai.LastUpdateTime
}

// GetDBIDList gets database identity list that the app uses
func (ai *AppInfo) GetDBIDList() ([]int, error) {
	return ai.AppRepo.GetDBIDList(ai.Identity())
}

// Set sets App with given fields, key is the field name and value is the relevant value of the key
func (ai *AppInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ai, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (ai *AppInfo) Delete() {
	ai.DelFlag = 1
}

// AddDB adds a new map of the app and database in the middleware
func (ai *AppInfo) AddDB(dbID int) error {
	return ai.AppRepo.AddDB(ai.Identity(), dbID)
}

// DeleteDB deletes the map of the app and database in the middleware
func (ai *AppInfo) DeleteDB(dbID int) error {
	return ai.AppRepo.DeleteDB(ai.Identity(), dbID)
}

// MarshalJSON marshals App to json bytes
func (ai *AppInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ai, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified fields of the App to json string
func (ai *AppInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ai, fields...)
}

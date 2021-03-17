package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*AppSystemInfo)(nil)

// AppSystemInfo store systeminfo
type AppSystemInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	AppSystemName  string    `middleware:"system_name" json:"system_name"`
	Level          int       `middleware:"level" json:"level"`
	OwnerID        int       `middleware:"owner_id" json:"owner_id"`
	OwnerGroup     string    `middleware:"owner_group" json:"owner_group"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewAppSystemInfo returns a new AppSystemInfo
func NewAppSystemInfo(repo *AppSystemRepo, id int, appSystemName string, level int, ownerID int, ownerGroup string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *AppSystemInfo {
	return &AppSystemInfo{
		repo,
		id,
		appSystemName,
		level,
		ownerID,
		ownerGroup,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewAppSystemInfoWithGlobal NewAppSystemInfo returns a new AppSystemInfo with default AppSystemRepo
func NewAppSystemInfoWithGlobal(id int, appSystemName string, level int, ownerID int, ownerGroup string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *AppSystemInfo {
	return &AppSystemInfo{
		NewAppSystemRepoWithGlobal(),
		id,
		appSystemName,
		level,
		ownerID,
		ownerGroup,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyAppSystemInfoWithGlobal return AppSystemInfo
func NewEmptyAppSystemInfoWithGlobal() *AppSystemInfo {
	return &AppSystemInfo{Repository: NewAppSystemRepoWithGlobal()}
}

// NewAppSystemInfoWithDefault returns a new AppSystemInfo with default AppSystemRepo
func NewAppSystemInfoWithDefault(appSystemName string, level int, ownerID int, ownerGroup string) *AppSystemInfo {
	return &AppSystemInfo{
		Repository:    NewAppSystemRepoWithGlobal(),
		AppSystemName: appSystemName,
		Level:         level,
		OwnerID:       ownerID,
		OwnerGroup:    ownerGroup,
	}
}

// Identity returns ID of entity
func (asi *AppSystemInfo) Identity() string {
	return strconv.Itoa(asi.ID)
}

// IsDeleted checks if delete flag had been set
func (asi *AppSystemInfo) IsDeleted() bool {
	return asi.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (asi *AppSystemInfo) GetCreateTime() time.Time {
	return asi.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (asi *AppSystemInfo) GetLastUpdateTime() time.Time {
	return asi.LastUpdateTime
}

// Get returns value of given field
func (asi *AppSystemInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(asi, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (asi *AppSystemInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(asi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (asi *AppSystemInfo) Delete() {
	asi.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (asi *AppSystemInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(asi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (asi *AppSystemInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(asi, fields...)
}

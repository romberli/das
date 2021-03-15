package metadata

import (
	"strconv"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency"
)

var _ dependency.Entity = (*EnvInfo)(nil)

type EnvInfo struct {
	dependency.Repository
	ID             int       `middleware:"id" json:"id"`
	EnvName        string    `middleware:"env_name" json:"env_name"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewEnvInfo returns a new *EnvInfo
func NewEnvInfo(repo *EnvRepo, id int, envName string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *EnvInfo {
	return &EnvInfo{
		repo,
		id,
		envName,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEnvInfo returns a new *EnvInfo with default EnvRepo
func NewEnvInfoWithGlobal(id int, envName string, delFlag int, createTime time.Time, lastUpdateTime time.Time) *EnvInfo {
	return &EnvInfo{
		NewEnvRepoWithGlobal(),
		id,
		envName,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyEnvInfoWithGlobal returns a new *EnvInfo with global repository
func NewEmptyEnvInfoWithGlobal() *EnvInfo {
	return &EnvInfo{Repository: NewEnvRepoWithGlobal()}
}

// NewEnvInfoWithDefault returns a new *EnvInfo with default EnvRepo
func NewEnvInfoWithDefault(envName string) *EnvInfo {
	return &EnvInfo{
		Repository: NewEnvRepoWithGlobal(),
		EnvName:    envName,
	}
}

// NewEnvInfoWithMapAndRandom returns a new *EnvInfo with given map
func NewEnvInfoWithMapAndRandom(fields map[string]interface{}) (*EnvInfo, error) {
	ei := &EnvInfo{}
	err := common.SetValuesWithMapAndRandom(ei, fields)
	if err != nil {
		return nil, err
	}

	return ei, nil
}

// Identity returns ID of entity
func (ei *EnvInfo) Identity() string {
	return strconv.Itoa(ei.ID)
}

// IsDeleted checks if delete flag had been set
func (ei *EnvInfo) IsDeleted() bool {
	return ei.DelFlag != constant.ZeroInt
}

// GetCreateTime returns created time of entity
func (ei *EnvInfo) GetCreateTime() time.Time {
	return ei.CreateTime
}

// GetLastUpdateTime returns last updated time of entity
func (ei *EnvInfo) GetLastUpdateTime() time.Time {
	return ei.LastUpdateTime
}

// Get returns value of given field
func (ei *EnvInfo) Get(field string) (interface{}, error) {
	return common.GetValueOfStruct(ei, field)
}

// Set sets entity with given fields, key is the field name and value is the relevant value of the key
func (ei *EnvInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(ei, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (ei *EnvInfo) Delete() {
	ei.DelFlag = 1
}

// MarshalJSON marshals entity to json string, it only marshals fields that has default tag
func (ei *EnvInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ei, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of entity to json string
func (ei *EnvInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ei, fields...)
}

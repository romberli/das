package dependency

import (
	"time"
)

type Entity interface {
	// Identity returns identity of entity
	Identity() string
	// IsDeleted checks if delete flag had been set
	IsDeleted() bool
	// GetCreateTime returns created time of entity
	GetCreateTime() time.Time
	// GetLastUpdateTime returns last updated time of entity
	GetLastUpdateTime() time.Time
	// Get returns value of given field
	Get(field string) (interface{}, error)
	// Set sets entity with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete set DelFlag to true
	Delete()
	// MarshalJSON marshals entity to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of entity to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

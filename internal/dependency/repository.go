package dependency

import (
	"github.com/romberli/go-util/middleware"
)

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// SelectAll returns all entities
	GetAll() ([]Entity, error)
	// Select returns an entity of the given id
	GetByID(id string) (Entity, error)
	// Create creates data with given entity in the middleware
	Create(entity Entity) (Entity, error)
	// Update updates data with given entity in the middleware
	Update(entity Entity) error
	// Delete deletes data in the middleware, it's recommended to use soft deletion
	Delete(id string) error
}

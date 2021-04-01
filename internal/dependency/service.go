package dependency

type Service interface {
	// GetEntities returns entities of the service
	GetEntities() []Entity
	// GetAll gets all entities from the middleware
	GetAll() error
	// GetID gets identity of an entity with given fields
	// GetID(fields map[string]interface{}) (string, error)
	// GetByID gets an entity that contains the given id from the middleware
	GetByID(id string) error
	// Create creates a new entity and insert it into the middleware
	Create(fields map[string]interface{}) error
	// Update gets an entity that contains the given id from the middleware,
	// and then update its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id string, fields map[string]interface{}) error
	// Delete deletes an entity by given id
	Delete(id string) error
	// Marshal marshals service.Envs
	Marshal() ([]byte, error)
	// Marshal marshals service.Envs with given fields
	MarshalWithFields(fields ...string) ([]byte, error)
}

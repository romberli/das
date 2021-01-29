package dependency

type Service interface {
	// GetEntities returns entities of the service
	GetEntities() []Entity
	// GetAll gets all environment entities from the middleware
	GetAll() error
	// GetByID gets an environment entity that contains the given id from the middleware
	GetByID(id string) error
	// Create creates a new environment entity and insert it into the middleware
	Create(fields map[string]interface{}) error
	// Update gets an environment entity that contains the given id from the middleware,
	// and then update its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middlewareUpdate(id string, fields map[string]interface{}) message
	Delete(id string) error
	// Marshal marshals service.Entities
	Marshal() ([]byte, error)
	// Marshal marshals service.Entities with given fields
	MarshalWithFields(fields ...string) ([]byte, error)
}

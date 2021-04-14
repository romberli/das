package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"
)

const middlewareServersStruct = "MiddlewareServers"

var _ metadata.MiddlewareServerService = (*MiddlewareServerService)(nil)

type MiddlewareServerService struct {
	metadata.MiddlewareServerRepo
	MiddlewareServers []metadata.MiddlewareServer `json:"middleware_servers"`
}

// NewMiddlewareServerService returns a new *MiddlewareServerService
func NewMiddlewareServerService(repo metadata.MiddlewareServerRepo) *MiddlewareServerService {
	return &MiddlewareServerService{repo, []metadata.MiddlewareServer{}}
}

// NewMiddlewareServerServiceWithDefault returns a new *MiddlewareServerService with default repository
func NewMiddlewareServerServiceWithDefault() *MiddlewareServerService {
	return NewMiddlewareServerService(NewMiddlewareServerRepoWithGlobal())
}

// GetMiddlewareServers returns middleware servers of the service
func (mss *MiddlewareServerService) GetMiddlewareServers() []metadata.MiddlewareServer {
	return mss.MiddlewareServers
}

// GetAll gets all middleware servers from the middleware
func (mss *MiddlewareServerService) GetAll() error {
	var err error
	mss.MiddlewareServers, err = mss.MiddlewareServerRepo.GetAll()

	return err
}

// GetByClusterID gets middleware servers with given cluster id
func (mss *MiddlewareServerService) GetByClusterID(clusterID int) error {
	var err error
	mss.MiddlewareServers, err = mss.MiddlewareServerRepo.GetByClusterID(clusterID)
	return err
}

// GetByID gets a middleware server by the identity from the middleware
func (mss *MiddlewareServerService) GetByID(id int) error {
	middlewareServer, err := mss.MiddlewareServerRepo.GetByID(id)
	if err != nil {
		return err
	}

	mss.MiddlewareServers = append(mss.MiddlewareServers, middlewareServer)

	return err
}

// GetByHostInfo gets a middleware server with given host ip and port number
func (mss *MiddlewareServerService) GetByHostInfo(hostIP string, portNum int) error {
	middlewareServer, err := mss.MiddlewareServerRepo.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}
	mss.MiddlewareServers = append(mss.MiddlewareServers, middlewareServer)
	return err
}

// Create creates a middleware server in the middleware
func (mss *MiddlewareServerService) Create(fields map[string]interface{}) error {
	// generate new map
	_, clusterIDExists := fields[middlewareServerClusterIDStruct]
	_, serverNameExists := fields[middlewareServerNameStruct]
	_, middlewareRoleExists := fields[middlewareServerMiddlewareRoleStruct]
	_, hostIPExists := fields[middlewareServerHostIPStruct]
	_, portNumExists := fields[middlewareServerPortNumStruct]
	if !clusterIDExists || !serverNameExists || !middlewareRoleExists || !hostIPExists || !portNumExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s, %s, %s, %s and %s", middlewareServerClusterIDStruct, middlewareServerNameStruct, middlewareServerMiddlewareRoleStruct, middlewareServerHostIPStruct, middlewareServerPortNumStruct))
	}
	// create a new entity
	middlewareServerInfo, err := NewMiddlewareServerInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	middlewareServer, err := mss.MiddlewareServerRepo.Create(middlewareServerInfo)
	if err != nil {
		return err
	}

	mss.MiddlewareServers = append(mss.MiddlewareServers, middlewareServer)
	return nil
}

// Update gets a middleware server of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MiddlewareServerService) Update(id int, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.MiddlewareServers[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.MiddlewareServerRepo.Update(mss.MiddlewareServers[constant.ZeroInt])
}

// Delete deletes the middleware server of given id in the middleware
func (mss *MiddlewareServerService) Delete(id int) error {
	return mss.MiddlewareServerRepo.Delete(id)
}

// Marshal marshals MiddlewareServerService.MiddlewareServers to json bytes
func (mss *MiddlewareServerService) Marshal() ([]byte, error) {
	return mss.MarshalWithFields(middlewareServersStruct)
}

// MarshalWithFields marshals only specified fields of the MiddlewareServerService to json bytes
func (mss *MiddlewareServerService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mss, fields...)
}

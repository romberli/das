package metadata

import (
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const (
	clusterIDStruct      = "ClusterID"
	serverNameStruct     = "ServerName"
	hostIPStruct         = "HostIP"
	portNumStruct        = "PortNum"
	deploymentTypeStruct = "DeploymentType"
	versionStruct        = "Version"
)

const msMySQLServersStruct = "MySQLServers"

var _ metadata.MySQLServerService = (*MySQLServerService)(nil)

// MySQLServerService implements Service interface
type MySQLServerService struct {
	MySQLServerRepo metadata.MySQLServerRepo
	MySQLServers    []metadata.MySQLServer
}

// NewMySQLServerService returns a new *MySQLServerService
func NewMySQLServerService(repo metadata.MySQLServerRepo) *MySQLServerService {
	return &MySQLServerService{repo, []metadata.MySQLServer{}}
}

// NewMySQLServerServiceWithDefault returns a new *MySQLServerService with default repository
func NewMySQLServerServiceWithDefault() *MySQLServerService {
	return NewMySQLServerService(NewMySQLServerRepoWithGlobal())
}

// GetMySQLServers returns entities of the service
func (mss *MySQLServerService) GetMySQLServers() []metadata.MySQLServer {
	entityList := make([]metadata.MySQLServer, len(mss.MySQLServers))
	for i := range entityList {
		entityList[i] = mss.MySQLServers[i]
	}

	return entityList
}

// GetAll gets all mysql server entities from the middleware
func (mss *MySQLServerService) GetAll() error {
	var err error
	mss.MySQLServers, err = mss.MySQLServerRepo.GetAll()

	return err
}

// GetByClusterID gets mysql servers with given cluster id
func (mss *MySQLServerService) GetByClusterID(clusterID int) error {
	mysqlServers, err := mss.MySQLServerRepo.GetByClusterID(clusterID)
	if err != nil {
		return err
	}

	mss.MySQLServers = mysqlServers

	return nil
}

// GetByID gets an mysql server entity that contains the given id from the middleware
func (mss *MySQLServerService) GetByID(id int) error {
	entity, err := mss.MySQLServerRepo.GetByID(id)
	if err != nil {
		return err
	}

	mss.MySQLServers = append(mss.MySQLServers, entity)

	return err
}

// GetByHostInfo gets a mysql server with given host ip and port number
func (mss *MySQLServerService) GetByHostInfo(hostIP string, portNum int) error {
	mysqlServer, err := mss.MySQLServerRepo.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	mss.MySQLServers = append(mss.MySQLServers, mysqlServer)

	return nil
}

// Create creates a new mysql server entity and insert it into the middleware
func (mss *MySQLServerService) Create(fields map[string]interface{}) error {
	// generate new map
	_, clusterIDExists := fields[clusterIDStruct]
	_, serverNameExists := fields[serverNameStruct]
	_, hostIPExists := fields[hostIPStruct]
	_, portNumExists := fields[portNumStruct]
	_, deploymentTypeExists := fields[deploymentTypeStruct]

	if !clusterIDExists || !serverNameExists || !hostIPExists || !portNumExists ||
		!deploymentTypeExists {
		return message.NewMessage(
			message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s and %s and %s and %s",
				clusterIDStruct,
				serverNameStruct,
				hostIPStruct,
				portNumStruct,
				deploymentTypeStruct))
	}

	// create a new entity
	mysqlServerInfo, err := NewMySQLServerInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mss.MySQLServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return err
	}

	mss.MySQLServers = append(mss.MySQLServers, entity)
	return nil
}

// Update gets an mysql server entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MySQLServerService) Update(id int, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.MySQLServers[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.MySQLServerRepo.Update(mss.MySQLServers[constant.ZeroInt])
}

// Delete deletes the mysql server entity that contains the given id in the middleware
func (mss *MySQLServerService) Delete(id int) error {
	return mss.MySQLServerRepo.Delete(id)
}

// Marshal marshals service.Envs
func (mss *MySQLServerService) Marshal() ([]byte, error) {
	return mss.MarshalWithFields(msMySQLServersStruct)
}

// MarshalWithFields marshals service.Envs with given fields
func (mss *MySQLServerService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mss, fields...)
}

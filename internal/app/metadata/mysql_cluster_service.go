package metadata

import (
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const (
	clusterNameStruct         = "ClusterName"
	middlewareClusterIDStruct = "MiddlewareClusterID"
	monitorSystemIDStruct     = "MonitorSystemID"
	ownerIDStruct             = "OwnerID"
	envIDStruct               = "EnvID"
)

const mcMySQLClustersStruct = "MySQLClusters"

var _ metadata.MySQLClusterService = (*MySQLClusterService)(nil)

// MySQLClusterService implements Service interface
type MySQLClusterService struct {
	MySQLClusterRepo  metadata.MySQLClusterRepo
	MySQLClusters     []metadata.MySQLCluster `json:"mysql_clusters"`
	MySQLServerIDList []int                   `json:"mysql_server_id_list"`
}

// NewMySQLClusterService returns a new *MySQLClusterService
func NewMySQLClusterService(repo metadata.MySQLClusterRepo) *MySQLClusterService {
	return &MySQLClusterService{repo, []metadata.MySQLCluster{}, []int{}}
}

// NewMySQLClusterServiceWithDefault returns a new *MySQLClusterService with default repository
func NewMySQLClusterServiceWithDefault() *MySQLClusterService {
	return NewMySQLClusterService(NewMySQLClusterRepoWithGlobal())
}

// GetMySQLClusters returns entities of the service
func (mcs *MySQLClusterService) GetMySQLClusters() []metadata.MySQLCluster {
	mysqlClusterList := make([]metadata.MySQLCluster, len(mcs.MySQLClusters))
	for i := range mysqlClusterList {
		mysqlClusterList[i] = mcs.MySQLClusters[i]
	}

	return mysqlClusterList
}

// GetAll gets all mysql cluster entities from the middleware
func (mcs *MySQLClusterService) GetAll() error {
	var err error
	mcs.MySQLClusters, err = mcs.MySQLClusterRepo.GetAll()

	return err
}

// GetByEnv gets mysql clusters of given env id
func (mcs *MySQLClusterService) GetByEnv(envID int) error {
	var err error

	mcs.MySQLClusters, err = mcs.MySQLClusterRepo.GetByEnv(envID)

	return err
}

// GetByID gets an mysql cluster entity that contains the given id from the middleware
func (mcs *MySQLClusterService) GetByID(id int) error {
	mysqlCluster, err := mcs.MySQLClusterRepo.GetByID(id)
	if err != nil {
		return err
	}

	mcs.MySQLClusters = append(mcs.MySQLClusters, mysqlCluster)

	return err
}

// GetByName gets a mysql cluster of given cluster name
func (mcs *MySQLClusterService) GetByName(clusterName string) error {
	mysqlCluster, err := mcs.MySQLClusterRepo.GetByName(clusterName)

	mcs.MySQLClusters = append(mcs.MySQLClusters, mysqlCluster)

	return err
}

// GetMySQLServerIDList gets the mysql server id list of given cluster id
func (mcs *MySQLClusterService) GetMySQLServerIDList(clusterID int) error {
	mysqlCluster, err := mcs.MySQLClusterRepo.GetByID(clusterID)
	if err != nil {
		return err
	}

	mcs.MySQLServerIDList, err = mysqlCluster.GetMySQLServerIDList()

	return err
}

// Create creates a new mysql cluster entity and insert it into the middleware
func (mcs *MySQLClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	_, clusterNameExists := fields[clusterNameStruct]
	_, envIDExists := fields[envIDStruct]

	if !clusterNameExists || !envIDExists {
		return message.NewMessage(
			message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s",
				clusterNameStruct,
				envIDStruct))
	}

	// create a new entity
	mysqlClusterInfo, err := NewMySQLClusterInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	mysqlCluster, err := mcs.MySQLClusterRepo.Create(mysqlClusterInfo)
	if err != nil {
		return err
	}

	mcs.MySQLClusters = append(mcs.MySQLClusters, mysqlCluster)
	return nil
}

// Update gets an mysql cluster entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mcs *MySQLClusterService) Update(id int, fields map[string]interface{}) error {
	err := mcs.GetByID(id)
	if err != nil {
		return err
	}
	err = mcs.MySQLClusters[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mcs.MySQLClusterRepo.Update(mcs.MySQLClusters[constant.ZeroInt])
}

// Delete deletes the mysql cluster entity that contains the given id in the middleware
func (mcs *MySQLClusterService) Delete(id int) error {
	return mcs.MySQLClusterRepo.Delete(id)
}

// Marshal marshals service.Envs
func (mcs *MySQLClusterService) Marshal() ([]byte, error) {
	return mcs.MarshalWithFields(mcMySQLClustersStruct)
}

// MarshalWithFields marshals service.Envs with given fields
func (mcs *MySQLClusterService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mcs, fields...)
}

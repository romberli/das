package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

var _ metadata.DBService = (*DBService)(nil)

type DBService struct {
	metadata.DBRepo
	DBs       []metadata.DB
	AppIDList []int
}

// NewDBService returns a new *DBService
func NewDBService(repo metadata.DBRepo) *DBService {
	return &DBService{repo, []metadata.DB{}, []int{}}
}

// NewDBServiceWithDefault returns a new *DBService with default repository
func NewDBServiceWithDefault() *DBService {
	return NewDBService(NewDBRepoWithGlobal())
}

// GetDBs returns databases of the service
func (ds *DBService) GetDBs() []metadata.DB {
	return ds.DBs
}

// GetAll gets all databases from the middleware
func (ds *DBService) GetAll() error {
	var err error

	ds.DBs, err = ds.DBRepo.GetAll()

	return err
}

// GetByEnv gets all databases of given env_id
func (ds *DBService) GetByEnv(envID int) error {
	var err error

	ds.DBs, err = ds.DBRepo.GetByEnv(envID)

	return err
}

// GetByID gets an database of the given id from the middleware
func (ds *DBService) GetByID(id int) error {
	db, err := ds.DBRepo.GetByID(id)
	if err != nil {
		return err
	}

	ds.DBs = append(ds.DBs, db)

	return err
}

// GetAppIDList gets an app identity list that uses this db
func (ds *DBService) GetAppIDList(dbID int) error {
	db, err := ds.DBRepo.GetByID(dbID)
	if err != nil {
		return err
	}

	ds.AppIDList, err = db.GetAppIDList()

	return err
}

// Create creates an new database in the middleware
func (ds *DBService) Create(fields map[string]interface{}) error {
	// generate new map
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	if !dbNameExists || !clusterIDExists || !clusterTypeExists || !envIDExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s",
			dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct))
	}
	// create a new entity
	dbInfo, err := NewDBInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	db, err := ds.DBRepo.Create(dbInfo)
	if err != nil {
		return err
	}

	ds.DBs = append(ds.DBs, db)

	return nil
}

// Update gets a database of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (ds *DBService) Update(id int, fields map[string]interface{}) error {
	err := ds.GetByID(id)
	if err != nil {
		return err
	}
	err = ds.DBs[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return ds.DBRepo.Update(ds.DBs[constant.ZeroInt])
}

// Delete deletes the database of given id in the middleware
func (ds *DBService) Delete(id int) error {
	return ds.DBRepo.Delete(id)
}

// AddApp adds a new map of app and database in the middleware
func (ds *DBService) AddApp(dbID, appID int) error {
	err := ds.DBRepo.AddApp(dbID, appID)
	if err != nil {
		return err
	}

	return ds.GetAppIDList(dbID)
}

// DeleteApp deletes the map of app and database in the middleware
func (ds *DBService) DeleteApp(dbID, appID int) error {
	err := ds.DBRepo.DeleteApp(dbID, appID)
	if err != nil {
		return err
	}

	return ds.GetAppIDList(dbID)
}

// Marshal marshals DBService.DBs to json bytes
func (ds *DBService) Marshal() ([]byte, error) {
	return json.Marshal(ds.DBs)
}

// MarshalWithFields marshals only specified fields of the DBService to json bytes
func (ds *DBService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(ds.DBs))
	for i := range interfaceList {
		dbInfo, err := common.CopyStructWithFields(ds.DBs[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = dbInfo
	}

	return json.Marshal(interfaceList)
}

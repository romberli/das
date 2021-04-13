package metadata

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const appAppsStruct = "Apps"

var _ metadata.AppService = (*AppService)(nil)

type AppService struct {
	metadata.AppRepo
	Apps     []metadata.App `json:"apps"`
	DBIDList []int          `json:"db_id_list"`
}

// NewAppService returns a new *AppService
func NewAppService(repo metadata.AppRepo) *AppService {
	return &AppService{repo, []metadata.App{}, []int{}}
}

// NewAppServiceWithDefault returns a new *AppService with default repository
func NewAppServiceWithDefault() *AppService {
	return NewAppService(NewAppRepoWithGlobal())
}

// GetApps returns apps of the service
func (as *AppService) GetApps() []metadata.App {
	return as.Apps
}

// GetAll gets all apps from the middleware
func (as *AppService) GetAll() error {
	var err error

	as.Apps, err = as.AppRepo.GetAll()

	return err
}

// GetByID gets an app of the given id from the middleware
func (as *AppService) GetByID(id int) error {
	entity, err := as.AppRepo.GetByID(id)
	if err != nil {
		return err
	}

	as.Apps = append(as.Apps, entity)

	return err
}

// GetAppByName gets App from the middleware by name
func (as *AppService) GetAppByName(appName string) error {
	app, err := as.AppRepo.GetAppByName(appName)
	if err != nil {
		return err
	}

	as.Apps = append(as.Apps, app)

	return nil
}

// GetDBIDList gets a database identity list that the app uses
func (as *AppService) GetDBIDList(id int) error {
	app, err := as.AppRepo.GetByID(id)
	if err != nil {
		return err
	}

	as.DBIDList, err = app.GetDBIDList()

	return err
}

// Create creates an app in the middleware
func (as *AppService) Create(fields map[string]interface{}) error {
	// generate new map
	_, ok := fields[appAppNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appAppNameStruct)
	}
	_, ok = fields[appLevelStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, appLevelStruct)
	}

	// create a new entity
	appInfo, err := NewAppInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}

	// insert into middleware
	app, err := as.AppRepo.Create(appInfo)
	if err != nil {
		return err
	}

	as.Apps = append(as.Apps, app)
	return nil
}

// Update gets the app of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (as *AppService) Update(id int, fields map[string]interface{}) error {
	err := as.GetByID(id)
	if err != nil {
		return err
	}
	err = as.Apps[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return as.AppRepo.Update(as.Apps[constant.ZeroInt])
}

// Delete deletes the app of given id in the middleware
func (as *AppService) Delete(id int) error {
	return as.AppRepo.Delete(id)
}

// AddDB adds a new map of app and database in the middleware
func (as *AppService) AddDB(appID, dbID int) error {
	err := as.AppRepo.AddDB(appID, dbID)
	if err != nil {
		return err
	}

	return as.GetDBIDList(appID)
}

// DeleteDB deletes the map of app and database in the middleware
func (as *AppService) DeleteDB(appID, dbID int) error {
	err := as.AppRepo.DeleteDB(appID, dbID)
	if err != nil {
		return err
	}

	return as.GetDBIDList(appID)
}

// Marshal marshals AppService.Apps to json bytes
func (as *AppService) Marshal() ([]byte, error) {
	return as.MarshalWithFields(appAppsStruct)
}

// MarshalWithFields marshals only specified fields of the AppService to json bytes
func (as *AppService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(as, fields...)
}

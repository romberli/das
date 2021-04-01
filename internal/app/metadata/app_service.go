package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

var _ metadata.AppService = (*AppService)(nil)

type AppService struct {
	metadata.AppRepo
	Apps     []metadata.App
	DBIDList []int
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
func (ass *AppService) GetApps() []metadata.App {
	return ass.Apps
}

// GetAll gets all apps from the middleware
func (ass *AppService) GetAll() error {
	var err error
	ass.Apps, err = ass.AppRepo.GetAll()

	return err
}

// GetByID gets an app of the given id from the middleware
func (ass *AppService) GetByID(id int) error {
	entity, err := ass.AppRepo.GetByID(id)
	if err != nil {
		return err
	}

	ass.Apps = append(ass.Apps, entity)

	return err
}

// GetAppByName gets App from the middleware by name
func (ass *AppService) GetAppByName(appName string) error {
	app, err := ass.AppRepo.GetAppByName(appName)
	if err != nil {
		return err
	}

	ass.Apps = append(ass.Apps, app)

	return nil
}

// GetDBIDList gets a database identity list that the app uses
func (ass *AppService) GetDBIDList(id int) error {
	app, err := ass.AppRepo.GetByID(id)
	if err != nil {
		return err
	}

	ass.DBIDList, err = app.GetDBIDList()

	return err
}

// Create creates an app in the middleware
func (ass *AppService) Create(fields map[string]interface{}) error {
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
	app, err := ass.AppRepo.Create(appInfo)
	if err != nil {
		return err
	}

	ass.Apps = append(ass.Apps, app)
	return nil
}

// Update gets the app of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (ass *AppService) Update(id int, fields map[string]interface{}) error {
	err := ass.GetByID(id)
	if err != nil {
		return err
	}
	err = ass.Apps[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return ass.AppRepo.Update(ass.Apps[constant.ZeroInt])
}

// Delete deletes the app of given id in the middleware
func (ass *AppService) Delete(id int) error {
	return ass.AppRepo.Delete(id)
}

// AddDB adds a new map of app and database in the middleware
func (ass *AppService) AddDB(appID, dbID int) error {
	err := ass.AppRepo.AddDB(appID, dbID)
	if err != nil {
		return err
	}

	return ass.GetDBIDList(appID)
}

// DeleteDB deletes the map of app and database in the middleware
func (ass *AppService) DeleteDB(appID, dbID int) error {
	err := ass.AppRepo.DeleteDB(appID, dbID)
	if err != nil {
		return err
	}

	return ass.GetDBIDList(appID)
}

// Marshal marshals AppService.Apps to json bytes
func (ass *AppService) Marshal() ([]byte, error) {
	return json.Marshal(ass.Apps)
}

// MarshalWithFields marshals only specified fields of the AppService to json bytes
func (ass *AppService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(ass.Apps))
	for i := range interfaceList {
		app, err := common.CopyStructWithFields(ass.Apps[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = app
	}

	return json.Marshal(interfaceList)
}

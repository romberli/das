package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const (
	monitorSystemNameStruct        = "MonitorSystemName"
	monitorSystemTypeStruct        = "MonitorSystemType"
	monitorSystemHostIPStruct      = "MonitorSystemHostIP"
	monitorSystemPortNumStruct     = "MonitorSystemPortNum"
	monitorSystemPortNumSlowStruct = "MonitorSystemPortNumSlow"
	monitorSystemBaseUrlStruct     = "BaseURL"
	monitorSystemEnvIDStruct       = "EnvID"
)

var _ metadata.MonitorSystemService = (*MonitorSystemService)(nil)

type MonitorSystemService struct {
	metadata.MonitorSystemRepo
	MonitorSystems []metadata.MonitorSystem
}

// NewMonitorSystemService returns a new *MonitorSystemService
func NewMonitorSystemService(repo metadata.MonitorSystemRepo) *MonitorSystemService {
	return &MonitorSystemService{repo, []metadata.MonitorSystem{}}
}

// NewMonitorSystemServiceWithDefault returns a new *MonitorSystemService with default repository
func NewMonitorSystemServiceWithDefault() *MonitorSystemService {
	return NewMonitorSystemService(NewMonitorSystemRepoWithGlobal())
}

// GetMonitorSystems returns monitor systems of the service
func (mss *MonitorSystemService) GetMonitorSystems() []metadata.MonitorSystem {
	return mss.MonitorSystems
}

// GetAll gets all monitor systems from the middleware
func (mss *MonitorSystemService) GetAll() error {
	var err error

	mss.MonitorSystems, err = mss.MonitorSystemRepo.GetAll()

	return err
}

// GetByID gets an monitor system of the given id from the middleware
func (mss *MonitorSystemService) GetByID(id int) error {
	monitorSystem, err := mss.MonitorSystemRepo.GetByID(id)
	if err != nil {
		return err
	}

	mss.MonitorSystems = append(mss.MonitorSystems, monitorSystem)

	return err
}

// GetByEnv gets all monitor systems from the middleware by env_id
func (mss *MonitorSystemService) GetByEnv(envID int) error {
	var err error

	mss.MonitorSystems, err = mss.MonitorSystemRepo.GetByEnv(envID)

	return err
}

// GetByHostInfo gets monitor system from the middleware by host_info
func (mss *MonitorSystemService) GetByHostInfo(hostIP string, portNum int) error {
	monitorSystem, err := mss.MonitorSystemRepo.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	mss.MonitorSystems = append(mss.MonitorSystems, monitorSystem)

	return err
}

// Create creates an new monitor system in the middleware
func (mss *MonitorSystemService) Create(fields map[string]interface{}) error {
	// generate new map
	_, monitorSystemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIPExists := fields[monitorSystemHostIPStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	_, envIDExists := fields[monitorSystemEnvIDStruct]
	if !monitorSystemNameExists || !systemTypeExists || !hostIPExists || !portNumExists || !portNumSlowExists || !baseUrlExists || !envIDExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s",
			monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIPStruct, monitorSystemPortNumStruct,
			monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct, monitorSystemEnvIDStruct))
	}
	// create a new entity
	monitorSystemInfo, err := NewMonitorSystemInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	monitorSystem, err := mss.MonitorSystemRepo.Create(monitorSystemInfo)
	if err != nil {
		return err
	}
	mss.MonitorSystems = append(mss.MonitorSystems, monitorSystem)

	return nil
}

// Update gets the monitor system of the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MonitorSystemService) Update(id int, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.MonitorSystems[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.MonitorSystemRepo.Update(mss.MonitorSystems[constant.ZeroInt])
}

// Delete deletes the monitor system of given id in the middleware
func (mss *MonitorSystemService) Delete(id int) error {
	return mss.MonitorSystemRepo.Delete(id)
}

// Marshal marshals service.Envs
func (mss *MonitorSystemService) Marshal() ([]byte, error) {
	return json.Marshal(mss.MonitorSystems)
}

// Marshal marshals service.Envs with given fields
func (mss *MonitorSystemService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(mss.MonitorSystems))
	for i := range interfaceList {
		monitorSystemInfo, err := common.CopyStructWithFields(mss.MonitorSystems[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = monitorSystemInfo
	}

	return json.Marshal(interfaceList)
}

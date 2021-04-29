package healthcheck

import (
	"time"

	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/common"
)

const (
	resultStruct = "Result"
	defaultStep  = time.Minute
)

var (
	defaultStartTime = time.Now().Add(-7 * 24 * time.Hour)
	defaultEndTime   = time.Now()
)

var _ healthcheck.Service = (*Service)(nil)

type OperationInfo struct {
	OperationID   int
	MysqlServerID int
	StartTime     time.Time
	EndTime       time.Time
	Step          time.Duration
}

func NewOperationInfo(operationID, mysqlServerID int, startTime, endTime time.Time, step time.Duration) *OperationInfo {
	return &OperationInfo{
		OperationID:   operationID,
		MysqlServerID: mysqlServerID,
		StartTime:     startTime,
		EndTime:       endTime,
		Step:          step,
	}
}

type Service struct {
	healthcheck.Repository
	Engine healthcheck.Engine
	Result *Result `json:"result"`
}

func NewService(repo healthcheck.Repository) *Service {
	return newService(repo)
}

func NewServiceWithDefault() healthcheck.Service {
	return newService(NewRepositoryWithGlobal())

}

func newService(repo healthcheck.Repository) *Service {
	return &Service{
		Repository: repo,
		Result:     NewEmptyResult(),
	}
}

func (s *Service) GetResult() healthcheck.Result {
	return s.Result
}

func (s *Service) GetResultByOperationID(id int) error {

}

func (s *Service) Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error {
	// init
	err := s.init(mysqlServerID, startTime, endTime, step)
	if err != nil {
		return err
	}
	// run
	go s.Engine.Run()

	return nil
}

func (s *Service) init(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error {
	// check if operation with the same mysql server id is still running

	// insert operation message

	// init engine

}

func (s *Service) ReviewAccurate(id, review int) error {

}

func (s *Service) MarshalJSON() ([]byte, error) {
	return s.MarshalJSONWithFields(resultStruct)
}

func (s *Service) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(s, fields...)
}

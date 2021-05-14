package healthcheck

import (
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

var _ healthcheck.Repository = (*Repository)(nil)

type Repository struct {
	Database middleware.Pool
}

// NewDBRepo returns *DBRepo with given middleware.Pool
func NewRepository(db middleware.Pool) *Repository {
	return &Repository{Database: db}
}

// NewDBRepo returns *DBRepo with global mysql pool
func NewRepositoryWithGlobal() *Repository {
	return NewRepository(global.MySQLPool)
}

// Execute executes given command and placeholders on the middleware
func (r *Repository) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := r.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("healthcheck Repository.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (r *Repository) Transaction() (middleware.Transaction, error) {
	return r.Database.Transaction()
}

func (r *Repository) GetResultByOperationID(operationID int) (healthcheck.Result, error) {
	return nil, nil
}

func (r *Repository) IsRunning(mysqlServerID int) (bool, error) {
	return false, nil
}

func (r *Repository) InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	return constant.ZeroInt, nil
}

func (r *Repository) UpdateOperationStatus(operationID int, status int, message string) error {
	return nil
}

func (r *Repository) SaveResult(result healthcheck.Result) error {
	return nil
}

func (r *Repository) UpdateAccurateReviewByOperationID(operationID int, review int) error {
	return nil
}

func (r *Repository) GetEngineConfig() (DefaultEngineConfig, error) {
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, 
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healthcheck Repository.GetEngineConfig() sql: \n%s", sql)

	result, err := r.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init DefaultEngineConfig
	defaultEngineConfig := NewEmptyDefaultEngineConfig()
	// map to struct
	err = result.MapToStructSlice(defaultEngineConfig, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	return defaultEngineConfig, nil
}

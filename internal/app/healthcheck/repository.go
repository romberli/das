package healthcheck

import (
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/healthcheck"
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

}

func (r *Repository) IsRunning(mysqlServerID int) (bool, error) {

}

func (r *Repository) InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {

}

func (r *Repository) UpdateOperationStatus(operationID int, status int, message string) error {

}

func (r *Repository) SaveResult(result healthcheck.Result) error {

}

func (r *Repository) UpdateAccurateReviewByOperationID(operationID int, review int) error {

}

package sqladvisor

import (
	"github.com/romberli/go-util/middleware"
)

type Advisor interface {
	GetFingerprint() string
	GetSQLID() string
	Advise() (string, error)
}

type Repository interface {
	Execute(command string, args ...interface{}) (middleware.Result, error)
	Transaction() (middleware.Transaction, error)
	Save(dbID int, sqlText, result, message string) error
}

type Service interface {
	GetFingerprint() string
	GetSQLID() string
	Advise() (string, error)
}

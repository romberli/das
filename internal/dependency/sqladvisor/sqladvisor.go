package sqladvisor

import (
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/sql/parser"
)

type Advisor interface {
	// GetParser returns the parser
	GetParser() *parser.Parser
	// GetFingerprint returns the fingerprint of the sql text
	GetFingerprint(sqlText string) string
	// GetSQLID returns the identity of the sql text
	GetSQLID(sqlText string) string
	// Advise parses the sql text and returns the tuning advice
	Advise(dbID int, sqlText string) (string, string, error)
}

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// Save saves sql tuning advice into the middleware
	Save(dbID int, sqlText, advice, message string) error
}

type Service interface {
	// GetFingerprint returns the fingerprint of the sql text
	GetFingerprint(sqlText string) string
	// GetSQLID returns the identity of the sql text
	GetSQLID(sqlText string) string
	// Advise parses the sql text and returns the tuning advice,
	// note that only the first sql statement in the sql text will be advised
	Advise(dbID int, sqlText string) (string, error)
}

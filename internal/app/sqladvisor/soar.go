package sqladvisor

import (
	"errors"
	"fmt"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/sqladvisor"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware/sql/parser"
	"github.com/spf13/viper"
)

var _ sqladvisor.Advisor = (*DefaultAdvisor)(nil)

type SQL struct {
	ID          string
	Text        string
	Fingerprint string
}

type DefaultAdvisor struct {
	parser     *parser.Parser
	sqlText    string
	dbID       int
	soarBin    string
	configFile string
}

func NewDefaultAdvisor(sqlText string, dbID int, soarBin, configFile string) *DefaultAdvisor {
	return newDefaultAdvisor(sqlText, dbID, soarBin, configFile)
}

func NewDefaultAdvisorWithDefault(sqlText string, dbID int) *DefaultAdvisor {
	soarBin := viper.GetString(config.SQLAdvisorSoarBin)
	configFile := viper.GetString(config.SQLAdvisorSoarConfig)

	return newDefaultAdvisor(sqlText, dbID, soarBin, configFile)
}

func newDefaultAdvisor(sqlText string, dbID int, soarBin, configFile string) *DefaultAdvisor {
	return &DefaultAdvisor{
		parser:     parser.NewParserWithDefault(),
		sqlText:    sqlText,
		dbID:       dbID,
		soarBin:    soarBin,
		configFile: configFile,
	}
}

func (da *DefaultAdvisor) GetSQLText() string {
	return da.sqlText
}

func (da *DefaultAdvisor) GetDBID() int {
	return da.dbID
}

func (da *DefaultAdvisor) GetFingerprint() string {
	return da.parser.GetFingerprint(da.GetSQLText())
}

func (da *DefaultAdvisor) GetSQLID() string {
	return da.parser.GetSQLID(da.GetSQLText())
}

func (da *DefaultAdvisor) Advise() (string, error) {
	dsn, err := da.getOnlineDSN()
	if err != nil {
		return constant.EmptyString, nil
	}

	command := fmt.Sprintf("%s -config=%s -online-dsn=%s -query=%s", da.soarBin, da.configFile, dsn, da.sqlText)

	result, err := linux.ExecuteCommand(command)
	if err != nil {
		return constant.EmptyString, err
	}

	return da.parseResult(result)
}

func (da *DefaultAdvisor) getOnlineDSN() (string, error) {
	// get db service
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetByID(da.dbID)
	if err != nil {
		return constant.EmptyString, nil
	}
	// get db
	db := dbService.DBs[constant.ZeroInt]
	clusterID := db.GetClusterID()
	dbName := db.GetDBName()
	// get mysql server service
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByClusterID(clusterID)
	if err != nil {
		return constant.EmptyString, err
	}
	// get mysql server
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]

	hostIP := mysqlServer.GetHostIP()
	portNum := mysqlServer.GetPortNum()
	mysqlUser := viper.GetString(config.DBSoarMySQLUserKey)
	mysqlPass := viper.GetString(config.DBSoarMySQLPassKey)

	return fmt.Sprintf("%s:%s@%s:%d/%s", mysqlUser, mysqlPass, hostIP, portNum, dbName), nil
}

func (da *DefaultAdvisor) parseResult(result string) (string, error) {
	var (
		advice  string
		message string
	)

	return advice, errors.New(message)
}

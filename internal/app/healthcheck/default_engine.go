package healthcheck

import (
	"time"

	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
)

var _ healthcheck.Engine = (*defaultEngine)(nil)

type defaultEngineConfig struct {
	ID               int       `middleware:"id" json:"id"`
	ItemName         string    `middleware:"item_name" json:"item_name"`
	ItemWeight       int       `middleware:"item_weight" json:"item_weight"`
	MinThreshold     float64   `middleware:"min_threshold" json:"min_threshold"`
	MaxThreshold     float64   `middleware:"max_threshold" json:"max_threshold"`
	GoodBasicScore   int       `middleware:"good_basic_score" json:"good_basic_score"`
	MediumBasicScore int       `middleware:"medium_basic_score" json:"medium_basic_score"`
	BadBasicScore    int       `middleware:"bad_basic_score" json:"bad_basic_score"`
	Unit             float64   `middleware:"unit" json:"unit"`
	ScorePerUnit     int       `middleware:"score_per_unit" json:"score_per_unit"`
	DelFlag          int       `middleware:"del_flag" json:"del_flag"`
	CreateTime       time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime   time.Time `middleware:"last_update_time" json:"last_update_time"`
}

func NewEmptyDefaultEngineConfig() *defaultEngineConfig {
	return &defaultEngineConfig{}
}

type defaultEngine struct {
	operationInfo         *OperationInfo
	repo                  healthcheck.Repository
	applicationMysqlConn  mysql.Conn
	monitorPrometheusConn prometheus.Conn
	monitorClickhouseConn clickhouse.Conn
	monitorMysqlConn      mysql.Conn
	engineConfig          *defaultEngineConfig
	result                *Result
}

func NewDefaultEngine(operationInfo *OperationInfo, repo healthcheck.Repository, applicationMySQLConn mysql.Conn,
	monitorPrometheusConn prometheus.Conn, monitorClickhouseConn clickhouse.Conn, monitorMySQLConn mysql.Conn) *defaultEngine {
	return &defaultEngine{
		operationInfo:         operationInfo,
		repo:                  repo,
		applicationMysqlConn:  applicationMySQLConn,
		monitorPrometheusConn: monitorPrometheusConn,
		monitorClickhouseConn: monitorClickhouseConn,
		monitorMysqlConn:      monitorMySQLConn,
		engineConfig:          NewEmptyDefaultEngineConfig(),
		result:                NewEmptyResult(),
	}
}

func (de *defaultEngine) loadConfig() error {

}

func (de *defaultEngine) Run() {
	// load config

	// check db config

	// check cpu usage

	// check io util

	// check disk capacity usage

	// check connection usage

	// check active session number

	// check cache hit ratio

	// check slow query

	// summarize

	// save result

	//

}

func (de *defaultEngine) summarize() int {

}

func (de *defaultEngine) checkDBConfig() error {

}

func (de *defaultEngine) checkCPUUsage() error {

}

func (de *defaultEngine) checkIOUtil() error {

}

func (de *defaultEngine) checkDiskCapacityUsage() error {

}

func (de *defaultEngine) checkConnectionUsage() error {

}

func (de *defaultEngine) checkActiveSessionNum() error {

}

func (de *defaultEngine) checkCacheHitRatio() error {

}

func (de *defaultEngine) checkSlowQuery() error {

}

func (de *defaultEngine) postRun(message string) error {

}

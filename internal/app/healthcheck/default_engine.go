package healthcheck

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
	"github.com/romberli/log"
)

const (
	defaultDBConfigScore                   = 5
	defaultMinScore                        = 0
	defaultMaxScore                        = 100.0
	defaultDBConfigItemName                = "db_config"
	defaultCPUUsageItemName                = "cpu_usage"
	defaultIOUtilItemName                  = "io_util"
	defaultDiskCapacityUsageItemName       = "disk_capacity_usage"
	defaultConnectionUsageItemName         = "connection_usage"
	defaultAverageActiveSessionNumItemName = "average_active_session_num"
	defaultCacheMissRatioItemName          = "cache_miss_ratio"
	defaultTableRowsItemName               = "table_rows"
	defaultTableSizeItemName               = "table_size"
	defaultSlowQueryItemName               = "slow_query"
)

var _ healthcheck.Engine = (*DefaultEngine)(nil)

type DefaultItemConfig struct {
	ID                          int       `middleware:"id" json:"id"`
	ItemName                    string    `middleware:"item_name" json:"item_name"`
	ItemWeight                  int       `middleware:"item_weight" json:"item_weight"`
	LowWatermark                float64   `middleware:"low_watermark" json:"low_watermark"`
	HighWatermark               float64   `middleware:"high_watermark" json:"high_watermark"`
	Unit                        float64   `middleware:"unit" json:"unit"`
	ScoreDeductionPerUnitHigh   float64   `middleware:"score_deduction_per_unit_high" json:"score_deduction_per_unit_high"`
	MaxScoreDeductionHigh       float64   `middleware:"max_score_deduction_high" json:"max_score_deduction_high"`
	ScoreDeductionPerUnitMedium float64   `middleware:"score_deduction_per_unit_medium" json:"score_deduction_per_unit_medium"`
	MaxScoreDeductionMedium     float64   `middleware:"max_score_deduction_medium" json:"max_score_deduction_medium"`
	DelFlag                     int       `middleware:"del_flag" json:"del_flag"`
	CreateTime                  time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime              time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewEmptyDefaultItemConfig returns a new *DefaultItemConfig
func NewEmptyDefaultItemConfig() *DefaultItemConfig {
	return &DefaultItemConfig{}
}

type DefaultEngineConfig map[string]*DefaultItemConfig

// NewEmptyDefaultEngineConfig returns a new empty *DefaultItemConfig
func NewEmptyDefaultEngineConfig() DefaultEngineConfig {
	return map[string]*DefaultItemConfig{}
}

// getItemConfig returns *DefaultItemConfig with given item name
func (dec DefaultEngineConfig) getItemConfig(item string) *DefaultItemConfig {
	return dec[item]
}

// Validate validates if engine configuration is valid
func (dec DefaultEngineConfig) Validate() bool {
	return true
}

type DefaultEngine struct {
	healthcheck.Repository
	operationInfo         *OperationInfo
	applicationMysqlConn  *mysql.Conn
	monitorPrometheusConn *prometheus.Conn
	monitorClickhouseConn *clickhouse.Conn
	monitorMysqlConn      *mysql.Conn
	engineConfig          DefaultEngineConfig
	result                *Result
}

// NewDefaultEngine returns a new *DefaultEngine
func NewDefaultEngine(repo healthcheck.Repository, operationInfo *OperationInfo, applicationMySQLConn *mysql.Conn,
	monitorPrometheusConn *prometheus.Conn, monitorClickhouseConn *clickhouse.Conn, monitorMySQLConn *mysql.Conn) *DefaultEngine {
	return &DefaultEngine{
		Repository:            repo,
		operationInfo:         operationInfo,
		applicationMysqlConn:  applicationMySQLConn,
		monitorPrometheusConn: monitorPrometheusConn,
		monitorClickhouseConn: monitorClickhouseConn,
		monitorMysqlConn:      monitorMySQLConn,
		engineConfig:          NewEmptyDefaultEngineConfig(),
		result:                NewEmptyResult(),
	}
}

// getItemConfig returns *DefaultItemConfig with given item name
func (de *DefaultEngine) getItemConfig(item string) *DefaultItemConfig {
	return de.engineConfig.getItemConfig(item)
}

// Run runs healthcheck
func (de *DefaultEngine) Run() {
	// run
	err := de.run()

	if err != nil {
		log.Error(message.NewMessage(msghc.ErrHealthcheckDefaultEngineRun, err.Error()).Error())
		// update status
		updateErr := de.Repository.UpdateOperationStatus(de.operationInfo.OperationID, defaultFailedStatus, err.Error())
		if updateErr != nil {
			log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
		}
	}

	// update operation status
	msg := fmt.Sprintf("healthcheck completed successfully. engine: default, operation_id: %d", de.operationInfo.OperationID)
	updateErr := de.Repository.UpdateOperationStatus(de.operationInfo.OperationID, defaultSuccessStatus, msg)
	if updateErr != nil {
		log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
	}
}

// run runs healthcheck
func (de *DefaultEngine) run() error {
	// pre run
	err := de.preRun()
	if err != nil {
		return err
	}
	// check db config
	err = de.checkDBConfig()
	if err != nil {
		return err
	}
	// check cpu usage
	err = de.checkCPUUsage()
	if err != nil {
		return err
	}
	// check io util
	err = de.checkIOUtil()
	if err != nil {
		return err
	}
	// check disk capacity usage
	err = de.checkDiskCapacityUsage()
	if err != nil {
		return err
	}
	// check connection usage
	err = de.checkConnectionUsage()
	if err != nil {
		return err
	}
	// check active session number
	err = de.checkActiveSessionNum()
	if err != nil {
		return err
	}
	// check cache miss ratio
	err = de.checkCacheMissRatio()
	if err != nil {
		return err
	}
	// check slow query
	err = de.checkSlowQuery()
	if err != nil {
		return err
	}
	// summarize
	de.summarize()
	// post run
	return de.postRun()
}

// preRun performs pre-run actions, for now, it only loads engine config
func (de *DefaultEngine) preRun() error {
	return de.loadEngineConfig()
}

// loadEngineConfig loads engine config
func (de *DefaultEngine) loadEngineConfig() error {
	// load config

	// validate config

	return nil
}

// checkDBConfig checks database configuration
func (de *DefaultEngine) checkDBConfig() error {
	// max_user_connection

	// log_bin

	// binlog_format

	// binlog_row_image

	// sync_binlog

	// innodb_flush_log_at_trx_commit

	// gtid_mode

	// enforce_gtid_consistency

	// slave-parallel-type

	// slave-parallel-workers

	// master_info_repository

	// relay_log_info_repository

	// report_host

	// report_port

	// innodb_buffer_pool_chunk_size

	// innodb_flush_method

	// innodb_monitor_enable

	// innodb_print_all_deadlocks

	// slow_query_log

	// performance_schema

	return nil
}

// checkCPUUsage checks cpu usage
func (de *DefaultEngine) checkCPUUsage() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	query := fmt.Sprintf(`
		sum(avg by (node_name,mode) (clamp_max(((avg by (mode,node_name) ((
		clamp_max(rate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[20s]),1)) or
		(clamp_max(irate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[5m]),1)) ))*100 or
		(avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[20s]) or
		avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[5m]))),100)))
	`, serverName, serverName, serverName, serverName)
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	cpuUsageConfig := de.getItemConfig(defaultCPUUsageItemName)

	var (
		cpuUsage            float64
		cpuUsageHighSum     float64
		cpuUsageHighCount   int
		cpuUsageMediumSum   float64
		cpuUsageMediumCount int

		cpuUsageHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		cpuUsage, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case cpuUsage >= cpuUsageConfig.HighWatermark:
			cpuUsageHigh = append(cpuUsageHigh, rowData)

			cpuUsageHighSum += rowData[constant.ZeroInt].(float64)
			cpuUsageHighCount++
		case cpuUsage >= cpuUsageConfig.LowWatermark:
			cpuUsageMediumSum += rowData[constant.ZeroInt].(float64)
			cpuUsageMediumCount++
		}
	}

	// cpu usage data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.CPUUsageData = string(jsonBytesTotal)
	// cpu usage high
	jsonBytesHigh, err := json.Marshal(cpuUsageHigh)
	if err != nil {
		return nil
	}
	de.result.CPUUsageData = string(jsonBytesHigh)

	// cpu usage high score deduction
	cpuUsageScoreDeductionHigh := (cpuUsageHighSum/float64(cpuUsageHighCount) - cpuUsageConfig.HighWatermark) / cpuUsageConfig.Unit * cpuUsageConfig.ScoreDeductionPerUnitHigh
	if cpuUsageScoreDeductionHigh > cpuUsageConfig.MaxScoreDeductionHigh {
		cpuUsageScoreDeductionHigh = cpuUsageConfig.MaxScoreDeductionHigh
	}
	// cpu usage medium score deduction
	cpuUsageScoreDeductionMedium := (cpuUsageMediumSum/float64(cpuUsageMediumCount) - cpuUsageConfig.LowWatermark) / cpuUsageConfig.Unit * cpuUsageConfig.ScoreDeductionPerUnitMedium
	if cpuUsageScoreDeductionMedium > cpuUsageConfig.MaxScoreDeductionMedium {
		cpuUsageScoreDeductionMedium = cpuUsageConfig.MaxScoreDeductionMedium
	}
	// cpu usage score
	de.result.CPUUsageScore = int(defaultMaxScore - cpuUsageScoreDeductionHigh - cpuUsageScoreDeductionMedium)
	if de.result.CPUUsageScore < constant.ZeroInt {
		de.result.CPUUsageScore = constant.ZeroInt
	}

	return nil
}

// checkIOUtil check io util
func (de *DefaultEngine) checkIOUtil() error {
	return nil
}

// checkDiskCapacityUsage checks disk capacity usage
func (de *DefaultEngine) checkDiskCapacityUsage() error {
	return nil
}

// checkConnectionUsage checks connection usage
func (de *DefaultEngine) checkConnectionUsage() error {
	return nil
}

// checkActiveSessionNum check active session number
func (de *DefaultEngine) checkActiveSessionNum() error {
	return nil
}

// checkCacheMissRatio checks cache miss ratio
func (de *DefaultEngine) checkCacheMissRatio() error {
	return nil
}

// checkTableSize checks table size
func (de *DefaultEngine) checkTableSize() error {
	// check table rows

	// check table size

	return nil
}

// checkSlowQuery checks slow query
func (de *DefaultEngine) checkSlowQuery() error {
	return nil
}

// summarize summarizes all item scores with weight
func (de *DefaultEngine) summarize() {
	de.result.WeightedAverageScore = (de.result.DBConfigScore*de.getItemConfig(defaultDBConfigItemName).ItemWeight +
		de.result.CPUUsageScore*de.getItemConfig(defaultCPUUsageItemName).ItemWeight +
		de.result.IOUtilScore*de.getItemConfig(defaultIOUtilItemName).ItemWeight +
		de.result.DiskCapacityUsageScore*de.getItemConfig(defaultDiskCapacityUsageItemName).ItemWeight +
		de.result.ConnectionUsageScore*de.getItemConfig(defaultConnectionUsageItemName).ItemWeight +
		de.result.AverageActiveSessionNumScore*de.getItemConfig(defaultAverageActiveSessionNumItemName).ItemWeight +
		de.result.CacheMissRatioScore*de.getItemConfig(defaultCacheMissRatioItemName).ItemWeight +
		de.result.TableSizeScore*(de.getItemConfig(defaultTableRowsItemName).ItemWeight+de.getItemConfig(defaultTableSizeItemName).ItemWeight) +
		de.result.SlowQueryScore*de.getItemConfig(defaultSlowQueryItemName).ItemWeight) / constant.MaxPercentage

	if de.result.WeightedAverageScore < defaultMinScore {
		de.result.WeightedAverageScore = defaultMinScore
	}
}

// postRun performs post-run actions, for now, it ony saves healthcheck result to the middleware
func (de *DefaultEngine) postRun() error {
	// save result
	return de.Repository.SaveResult(de.result)
}

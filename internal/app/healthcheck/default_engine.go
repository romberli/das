package healthcheck

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/pingcap/errors"

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
	defaultSlowQueryExecutionTimeItemName  = "slow_query_execution_time"
	defaultSlowQueryRowsExaminedItemName   = "slow_query_rows_examined"

	dbConfigMaxUserConnection         = "max_user_connection"
	dbConfigLogBin                    = "log_bin"
	dbConfigBinlogFormat              = "binglog_format"
	dbConfigBinlogRowImage            = "binlog_row_image"
	dbConfigSyncBinlog                = "sync_binlog"
	dbConfigInnodbFlushLogAtTrxCommit = "innodb_flush_log_at_trx_commit"
	dbConfigGtidMode                  = "gtid_mode"
	dbConfigEnforceGtidConsistency    = "enforce_gtid_consistency"
	dbConfigSlaveParallelType         = "slave_parallel_type"
	dbConfigSlaveParallelWorkers      = "slave_parallel_workers"
	dbConfigMasterInfoRepository      = "master_info_repository"
	dbConfigRelayLogInfoRepository    = "relay_log_info_repository"
	dbConfigReportHost                = "report_host"
	dbConfigReportPort                = "report_port"
	dbConfigInnodbBufferPoolSize      = "innodb_buffer_pool_size"
	dbConfigInnodbBufferPoolChunkSize = "innodb_buffer_pool_chunk_size"
	dbConfigInnodbFlushMethod         = "innodb_flush_method"
	dbConfigInnodbMonitorEnable       = "innodb_monitor_enable"
	dbConfigInnodbPrintAllDeadlocks   = "innodb_print_all_deadlocks"
	dbConfigSlowQueryLog              = "slow_query_log"
	dbConfigPerformanceSchema         = "performance_schema"

	dbConfigMaxUserConnectionValid         = "2000"
	dbConfigLogBinValid                    = "ON"
	dbConfigBinlogFormatValid              = "ROW"
	dbConfigBinlogRowImageValid            = "FULL"
	dbConfigSyncBinlogValid                = "1"
	dbConfigInnodbFlushLogAtTrxCommitValid = "1"
	dbConfigGtidModeValid                  = "ON"
	dbConfigEnforceGtidConsistencyValid    = "ON"
	dbConfigSlaveParallelTypeValid         = "LOGICAL_CLOCK"
	dbConfigSlaveParallelWorkersValid      = "16"
	dbConfigMasterInfoRepositoryValid      = "TABLE"
	dbConfigRelayLogInfoRepositoryValid    = "TABLE"
	dbConfigInnodbFlushMethodValid         = "O_DIRECT"
	dbConfigInnodbMonitorEnableValid       = "all"
	dbConfigInnodbPrintAllDeadlocksValid   = "ON"
	dbConfigSlowQueryLogValid              = "ON"
	dbConfigPerformanceSchemaValid         = "ON"
)

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

var _ healthcheck.Engine = (*DefaultEngine)(nil)

type GlobalVariables struct {
	VariableName  string `middleware:"variable_name" json:"variable_name"`
	VariableValue string `middleware:"variable_value" json:"variable_value"`
}

// NewEmptyGlobalVariables returns a new *GlobalVariables
func NewEmptyGlobalVariables() *GlobalVariables {
	return &GlobalVariables{}
}

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
func (dec DefaultEngineConfig) Validate() error {
	itemWeightCount := constant.ZeroInt
	// validate defaultEngineConfig exits items
	if len(dec) == constant.ZeroInt {
		log.Errorf("default engine config doesn't have content.")
		return nil
	}
	for itemName, defaultItemConfig := range dec {
		// validate item weight
		if defaultItemConfig.ItemWeight > 100 || defaultItemConfig.ItemWeight < 0 {
			log.Errorf("item name: %s item weight is invalid.", itemName)
			return nil
		}
		// validate low watermark
		if defaultItemConfig.LowWatermark < 0 {
			log.Errorf("item name: %s low watermark is invalid.", itemName)
			return nil
		}
		// validate high watermark
		if defaultItemConfig.HighWatermark <= defaultItemConfig.LowWatermark {
			log.Errorf("item name: %s high watermark is invalid.", itemName)
			return nil
		}
		// validate unit
		if defaultItemConfig.Unit < 0 {
			log.Errorf("item name: %s unit is invalid.", itemName)
			return nil
		}
		// validate score deduction per unit high
		if defaultItemConfig.ScoreDeductionPerUnitHigh > 100 || defaultItemConfig.ScoreDeductionPerUnitHigh < 0 || defaultItemConfig.ScoreDeductionPerUnitHigh > defaultItemConfig.MaxScoreDeductionHigh {
			log.Errorf("item name: %s score deduction per unit high is invalid.", itemName)
			return nil
		}
		// validate max score deduction high
		if defaultItemConfig.MaxScoreDeductionHigh > 100 || defaultItemConfig.MaxScoreDeductionHigh < 0 {
			log.Errorf("item name: %s max score deduction high is invalid.", itemName)
			return nil
		}
		// validate score deduction per unit medium
		if defaultItemConfig.ScoreDeductionPerUnitMedium > 100 || defaultItemConfig.ScoreDeductionPerUnitMedium < 0 || defaultItemConfig.ScoreDeductionPerUnitMedium > defaultItemConfig.MaxScoreDeductionMedium {
			log.Errorf("item name: %s score deduction per unit medium is invalid.", itemName)
			return nil
		}
		// validate max score deduction medium
		if defaultItemConfig.MaxScoreDeductionMedium > 100 || defaultItemConfig.MaxScoreDeductionMedium < 0 {
			log.Errorf("item name: %s max score deduction medium is invalid.", itemName)
			return nil
		}
		itemWeightCount += defaultItemConfig.ItemWeight
	}
	// validate item weigh count is 100
	if itemWeightCount != 100 {
		log.Errorf("all items weight weight count is not 100.")
		return nil
	}
	return nil
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

//
func NewDefaultItemConfig(itemName string, itemWeight int, lowWatermark float64, highWatermark float64, unit float64,
	scoreDeductionPerUnitHigh float64, maxScoreDeductionHigh float64, scoreDeductionPerUnitMedium float64, maxScoreDeductionMedium float64) *DefaultItemConfig {
	return &DefaultItemConfig{
		ItemName:                    itemName,
		ItemWeight:                  itemWeight,
		LowWatermark:                lowWatermark,
		HighWatermark:               highWatermark,
		Unit:                        unit,
		ScoreDeductionPerUnitHigh:   scoreDeductionPerUnitHigh,
		MaxScoreDeductionHigh:       maxScoreDeductionHigh,
		ScoreDeductionPerUnitMedium: scoreDeductionPerUnitMedium,
		MaxScoreDeductionMedium:     maxScoreDeductionMedium,
	}
}

// getItemConfig returns *DefaultItemConfig with given item name
func (de *DefaultEngine) getItemConfig(item string) *DefaultItemConfig {
	return de.engineConfig.getItemConfig(item)
}

func (de *DefaultEngine) getPmmVersion() (int, error) {
	prometheusInfo, err := de.monitorPrometheusConn.API.Buildinfo(context.Background())
	if err != nil {
		return 0, err
	}
	prometheusVersion, err := strconv.Atoi(prometheusInfo.Version[0:1])
	if err != nil {
		return 0, err
	}
	return prometheusVersion, nil
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
	// check table size
	err = de.checkTableSize()
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
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healcheck Repository.loadEngineConfig() sql: \n%s\nplaceholders: %s", sql)
	result, err := de.Repository.Execute(sql)
	if err != nil {
		return nil
	}
	// init []*DefaultItemConfig
	defaultEngineConfigList := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultEngineConfigList {
		defaultEngineConfigList[i] = NewEmptyDefaultItemConfig()
	}
	// map to struct
	err = result.MapToStructSlice(defaultEngineConfigList, constant.DefaultMiddlewareTag)
	if err != nil {
		return err
	}
	entityList := NewEmptyDefaultEngineConfig()
	for i := range defaultEngineConfigList {
		itemName := defaultEngineConfigList[i].ItemName
		entityList[itemName] = defaultEngineConfigList[i]
	}
	// validate config
	validate := entityList.Validate()
	if validate == nil {
		return errors.New("default engine config formant is invalid.")
	}
	return nil
}

// checkDBConfig checks database configuration
func (de *DefaultEngine) checkDBConfig() error {
	// load database config
	sql := `select variable_name, variable_value
		from global_variables;`
	result, err := de.result.Execute(sql)
	if err != nil {
		return err
	}
	variableList := make([]*GlobalVariables, result.RowNumber())
	for i := range variableList {
		variableList[i] = NewEmptyGlobalVariables()
	}
	// map to struct
	err = result.MapToStructSlice(variableList, constant.DefaultMiddlewareTag)
	if err != nil {
		return err
	}

	dbConfigConfig := de.getItemConfig(defaultDBConfigItemName)

	var (
		dbConfigCount   int
		dbConfigInvalid []GlobalVariables
		dbConfigAdvice  []GlobalVariables
		advice          string
	)

	// max_user_connection
	if variableMap[dbConfigMaxUserConnection] != dbConfigMaxUserConnectionValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigMaxUserConnection,
			variableMap[dbConfigMaxUserConnection],
		})
		advice = "It is recommended that " + dbConfigMaxUserConnection + " be set to " + dbConfigMaxUserConnectionValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// log_bin
	if variableMap[dbConfigLogBin] != dbConfigLogBinValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigLogBin,
			variableMap[dbConfigLogBin],
		})
		advice = "It is recommended that " + dbConfigLogBin + " be set to " + dbConfigLogBinValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// binlog_format
	if variableMap[dbConfigBinlogFormat] != dbConfigBinlogFormatValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigBinlogFormat,
			variableMap[dbConfigBinlogFormat],
		})
		advice = "It is recommended that " + dbConfigBinlogFormat + " be set to " + dbConfigBinlogFormatValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// binlog_row_image
	if variableMap[dbConfigBinlogRowImage] != dbConfigBinlogRowImageValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigBinlogRowImage,
			variableMap[dbConfigBinlogRowImage],
		})
		advice = "It is recommended that " + dbConfigBinlogRowImage + " be set to " + dbConfigBinlogRowImageValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// sync_binlog
	if variableMap[dbConfigSyncBinlog] != dbConfigSyncBinlogValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSyncBinlog,
			variableMap[dbConfigSyncBinlog],
		})
		advice = "It is recommended that " + dbConfigSyncBinlog + " be set to " + dbConfigSyncBinlogValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_flush_log_at_trx_commit
	if variableMap[dbConfigInnodbFlushLogAtTrxCommit] != dbConfigInnodbFlushLogAtTrxCommitValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbFlushLogAtTrxCommit,
			variableMap[dbConfigInnodbFlushLogAtTrxCommit],
		})
		advice = "It is recommended that " + dbConfigInnodbFlushLogAtTrxCommit + " be set to " + dbConfigInnodbFlushLogAtTrxCommitValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// gtid_mode
	if variableMap[dbConfigGtidMode] != dbConfigGtidModeValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigGtidMode,
			variableMap[dbConfigGtidMode],
		})
		advice = "It is recommended that " + dbConfigGtidMode + " be set to " + dbConfigGtidModeValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// enforce_gtid_consistency
	if variableMap[dbConfigEnforceGtidConsistency] != dbConfigEnforceGtidConsistencyValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigEnforceGtidConsistency,
			variableMap[dbConfigEnforceGtidConsistency],
		})
		advice = "It is recommended that " + dbConfigEnforceGtidConsistency + " be set to " + dbConfigEnforceGtidConsistencyValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slave-parallel-type
	if variableMap[dbConfigSlaveParallelType] != dbConfigSlaveParallelTypeValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlaveParallelType,
			variableMap[dbConfigSlaveParallelType],
		})
		advice = "It is recommended that " + dbConfigSlaveParallelType + " be set to " + dbConfigSlaveParallelTypeValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slave-parallel-workers
	if variableMap[dbConfigSlaveParallelWorkers] != dbConfigSlaveParallelWorkersValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlaveParallelWorkers,
			variableMap[dbConfigSlaveParallelWorkers],
		})
		advice = "It is recommended that " + dbConfigSlaveParallelWorkers + " be set to " + dbConfigSlaveParallelWorkersValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// master_info_repository
	if variableMap[dbConfigMasterInfoRepository] != dbConfigMasterInfoRepositoryValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigMasterInfoRepository,
			variableMap[dbConfigMasterInfoRepository],
		})
		advice = "It is recommended that " + dbConfigMasterInfoRepository + " be set to " + dbConfigMasterInfoRepositoryValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// relay_log_info_repository
	if variableMap[dbConfigRelayLogInfoRepository] != dbConfigRelayLogInfoRepositoryValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigRelayLogInfoRepository,
			variableMap[dbConfigRelayLogInfoRepository],
		})
		advice = "It is recommended that " + dbConfigRelayLogInfoRepository + " be set to " + dbConfigRelayLogInfoRepositoryValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// report_host
	serverName := de.operationInfo.MySQLServer.GetServerName()
	if variableMap[dbConfigReportHost] != serverName {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigReportHost,
			variableMap[dbConfigReportHost],
		})
		advice = "It is recommended that " + dbConfigReportHost + " be set to " + serverName + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// report_port
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	if variableMap[dbConfigReportPort] != strconv.Itoa(portNum) {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigReportPort,
			variableMap[dbConfigReportPort],
		})
		advice = "It is recommended that " + dbConfigReportPort + " be set to " + strconv.Itoa(portNum) + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}

	// innodb_flush_method
	if variableMap[dbConfigInnodbFlushMethod] != dbConfigInnodbFlushMethodValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbFlushMethod,
			variableMap[dbConfigInnodbFlushMethod],
		})
		advice = "It is recommended that " + dbConfigInnodbFlushMethod + " be set to " + dbConfigInnodbFlushMethodValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_monitor_enable
	if variableMap[dbConfigInnodbMonitorEnable] != dbConfigInnodbMonitorEnableValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbMonitorEnable,
			variableMap[dbConfigInnodbMonitorEnable],
		})
		advice = "It is recommended that " + dbConfigInnodbMonitorEnable + " be set to " + dbConfigInnodbMonitorEnableValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_print_all_deadlocks
	if variableMap[dbConfigInnodbPrintAllDeadlocks] != dbConfigInnodbPrintAllDeadlocksValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbPrintAllDeadlocks,
			variableMap[dbConfigInnodbPrintAllDeadlocks],
		})
		advice = "It is recommended that " + dbConfigInnodbPrintAllDeadlocks + " be set to " + dbConfigInnodbPrintAllDeadlocksValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slow_query_log
	if variableMap[dbConfigSlowQueryLog] != dbConfigSlowQueryLogValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlowQueryLog,
			variableMap[dbConfigSlowQueryLog],
		})
		advice = "It is recommended that " + dbConfigSlowQueryLog + " be set to " + dbConfigSlowQueryLogValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// performance_schema
	if variableMap[dbConfigPerformanceSchema] != dbConfigPerformanceSchemaValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigPerformanceSchema,
			variableMap[dbConfigPerformanceSchema],
		})
		advice = "It is recommended that " + dbConfigPerformanceSchema + " be set to " + dbConfigPerformanceSchemaValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// database config data
	jsonBytesTotal, err := json.Marshal(dbConfigInvalid)
	if err != nil {
		return nil
	}
	de.result.DBConfigData = string(jsonBytesTotal)
	// database config advice
	jsonBytesAdvice, err := json.Marshal(dbConfigAdvice)
	if err != nil {
		return nil
	}
	de.result.DBConfigAdvice = string(jsonBytesAdvice)

	// database config score deduction
	dbConfigScoreDeduction := float64(dbConfigCount) * dbConfigConfig.ScoreDeductionPerUnitHigh
	if dbConfigScoreDeduction > dbConfigConfig.MaxScoreDeductionHigh {
		dbConfigScoreDeduction = dbConfigConfig.MaxScoreDeductionHigh
	}
	de.result.DBConfigScore = int(defaultMaxScore - dbConfigScoreDeduction)
	if de.result.DBConfigScore < constant.ZeroInt {
		de.result.DBConfigScore = constant.ZeroInt
	}
	return nil
}

// checkCPUUsage checks cpu usage
func (de *DefaultEngine) checkCPUUsage() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		sum(((avg by (mode) ( (clamp_max(rate(node_cpu{instance=~"%s",mode!="idle"}[$interval]),1)) 
		or (clamp_max(irate(node_cpu{instance=~"%s",mode!="idle"}[5m]),1)) ))*100 or 
		(avg_over_time(node_cpu_average{instance=~"%s", mode!="total", mode!="idle"}[$interval]) or 
		avg_over_time(node_cpu_average{instance=~"%s", mode!="total", mode!="idle"}[5m]))))
	`, host, host, host, host)
	case 2:
		query = fmt.Sprintf(`
		sum(avg by (node_name,mode) (clamp_max(((avg by (mode,node_name) ((
		clamp_max(rate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[20s]),1)) or
		(clamp_max(irate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[5m]),1)) ))*100 or
		(avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[20s]) or
		avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[5m]))),100)))
	`, serverName, serverName, serverName, serverName)
	}
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
			cpuUsageHighSum += cpuUsage
			cpuUsageHighCount++
		case cpuUsage >= cpuUsageConfig.LowWatermark:
			cpuUsageMediumSum += cpuUsage
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
	de.result.CPUUsageHigh = string(jsonBytesHigh)

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
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		rate(node_disk_io_time_ms{device=~"(sda|sdb|sdc|sr0)", instance=~"%s"}[$interval])/1000 or 
		irate(node_disk_io_time_ms{device=~"(sda|sdb|sdc|sr0)", instance=~"%s"}[5m])/1000
	`, host, host)
	case 2:
		query = fmt.Sprintf(`
		sum by (node_name) (rate(node_disk_io_time_seconds_total{device=~"(sda|sdb|sdc|sr0)",node_name=~"%s"}[5m]) or 
		irate(node_disk_io_time_seconds_total{device=~"(sda|sdb|sdc|sr0)",node_name=~"%s"}[5m]) or
		(max_over_time(rdsosmetrics_diskIO_util{device=~"(sda|sdb|sdc|sr0)",node_name=~"%s"}[5m]) or 
		max_over_time(rdsosmetrics_diskIO_util{device=~"(sda|sdb|sdc|sr0)",node_name=~"%s"}[5m]))/100)
	`, serverName, serverName, serverName, serverName)
	}
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	ioUtilConfig := de.getItemConfig(defaultIOUtilItemName)

	var (
		ioUtil            float64
		ioUtilHighSum     float64
		ioUtilHighCount   int
		ioUtilMediumSum   float64
		ioUtilMediumCount int

		ioUtilHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		ioUtil, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case ioUtil >= ioUtilConfig.HighWatermark:
			ioUtilHigh = append(ioUtilHigh, rowData)
			ioUtilHighSum += ioUtil
			ioUtilHighCount++
		case ioUtil >= ioUtilConfig.LowWatermark:
			ioUtilMediumSum += ioUtil
			ioUtilMediumCount++
		}
	}

	// io utilization data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.IOUtilData = string(jsonBytesTotal)
	// io utilization high
	jsonBytesHigh, err := json.Marshal(ioUtilHigh)
	if err != nil {
		return nil
	}
	de.result.IOUtilHigh = string(jsonBytesHigh)

	// io utilization high score deduction
	ioUtilScoreDeductionHigh := (ioUtilHighSum/float64(ioUtilHighCount) - ioUtilConfig.HighWatermark) / ioUtilConfig.Unit * ioUtilConfig.ScoreDeductionPerUnitHigh
	if ioUtilScoreDeductionHigh > ioUtilConfig.MaxScoreDeductionHigh {
		ioUtilScoreDeductionHigh = ioUtilConfig.MaxScoreDeductionHigh
	}
	// io utilization medium score deduction
	ioUtilScoreDeductionMedium := (ioUtilMediumSum/float64(ioUtilMediumCount) - ioUtilConfig.LowWatermark) / ioUtilConfig.Unit * ioUtilConfig.ScoreDeductionPerUnitMedium
	if ioUtilScoreDeductionMedium > ioUtilConfig.MaxScoreDeductionMedium {
		ioUtilScoreDeductionMedium = ioUtilConfig.MaxScoreDeductionMedium
	}
	// io utilization score
	de.result.IOUtilScore = int(defaultMaxScore - ioUtilScoreDeductionHigh - ioUtilScoreDeductionMedium)
	if de.result.IOUtilScore < constant.ZeroInt {
		de.result.IOUtilScore = constant.ZeroInt
	}

	return nil
}

// checkDiskCapacityUsage checks disk capacity usage
func (de *DefaultEngine) checkDiskCapacityUsage() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		node_filesystem_size{instance=~"%s",mountpoint="/", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"} 
		- node_filesystem_free{instance=~"%s",mountpoint="/", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"}
	`, host, host)
	case 2:
		query = fmt.Sprintf(`
		sum(avg by (node_name,mountpoint) (1 - (max_over_time(node_filesystem_free_bytes{node_name=~"%s", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"}[5m]) or 
		max_over_time(node_filesystem_free_bytes{node_name=~"%s", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"}[5m]))  
		(max_over_time(node_filesystem_size_bytes{node_name=~"%s", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"}[5m]) or 
		max_over_time(node_filesystem_size_bytes{node_name=~"%s", fstype!~"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs"}[5m]))))
	`, serverName, serverName, serverName, serverName)
	}
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	diskCapacityUsageConfig := de.getItemConfig(defaultDiskCapacityUsageItemName)

	var (
		diskCapacityUsage            float64
		diskCapacityUsageHighSum     float64
		diskCapacityUsageHighCount   int
		diskCapacityUsageMediumSum   float64
		diskCapacityUsageMediumCount int

		diskCapacityUsageHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		diskCapacityUsage, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case diskCapacityUsage >= diskCapacityUsageConfig.HighWatermark:
			diskCapacityUsageHigh = append(diskCapacityUsageHigh, rowData)
			diskCapacityUsageHighSum += diskCapacityUsage
			diskCapacityUsageHighCount++
		case diskCapacityUsage >= diskCapacityUsageConfig.LowWatermark:
			diskCapacityUsageMediumSum += diskCapacityUsage
			diskCapacityUsageMediumCount++
		}
	}

	// disk capacity usage data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.DiskCapacityUsageData = string(jsonBytesTotal)
	// disk capacity usage high
	jsonBytesHigh, err := json.Marshal(diskCapacityUsageHigh)
	if err != nil {
		return nil
	}
	de.result.DiskCapacityUsageHigh = string(jsonBytesHigh)

	// disk capacity usage high score deduction
	diskCapacityUsageScoreDeductionHigh := (diskCapacityUsageHighSum/float64(diskCapacityUsageHighCount) - diskCapacityUsageConfig.HighWatermark) / diskCapacityUsageConfig.Unit * diskCapacityUsageConfig.ScoreDeductionPerUnitHigh
	if diskCapacityUsageScoreDeductionHigh > diskCapacityUsageConfig.MaxScoreDeductionHigh {
		diskCapacityUsageScoreDeductionHigh = diskCapacityUsageConfig.MaxScoreDeductionHigh
	}
	// disk capacity usage medium score deduction
	diskCapacityUsageScoreDeductionMedium := (diskCapacityUsageMediumSum/float64(diskCapacityUsageMediumCount) - diskCapacityUsageConfig.LowWatermark) / diskCapacityUsageConfig.Unit * diskCapacityUsageConfig.ScoreDeductionPerUnitMedium
	if diskCapacityUsageScoreDeductionMedium > diskCapacityUsageConfig.MaxScoreDeductionMedium {
		diskCapacityUsageScoreDeductionMedium = diskCapacityUsageConfig.MaxScoreDeductionMedium
	}
	// disk capacity score
	de.result.DiskCapacityUsageScore = int(defaultMaxScore - diskCapacityUsageScoreDeductionHigh - diskCapacityUsageScoreDeductionMedium)
	if de.result.DiskCapacityUsageScore < constant.ZeroInt {
		de.result.DiskCapacityUsageScore = constant.ZeroInt
	}

	return nil
}

// checkConnectionUsage checks connection usage
func (de *DefaultEngine) checkConnectionUsage() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		max(max_over_time(mysql_global_status_threads_connected{instance=~"%s"}[$interval]) or 
		mysql_global_status_threads_connected{instance=~"%s"} )
	`, host, host)
	case 2:
		query = fmt.Sprintf(`
		clamp_max((avg by (service_name) (max_over_time(mysql_global_status_max_used_connections{service_name=~"%s"}[5m]) or 
		max_over_time(mysql_global_status_max_used_connections{service_name=~"%s"}[5m])) / avg by (service_name) 
		(mysql_global_variables_max_connections{service_name=~"%s"})),1)
	`, serverName, serverName, serverName)
	}
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	connectionUsageConfig := de.getItemConfig(defaultConnectionUsageItemName)

	var (
		connectionUsage            float64
		connectionUsageHighSum     float64
		connectionUsageHighCount   int
		connectionUsageMediumSum   float64
		connectionUsageMediumCount int

		connectionUsageHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		connectionUsage, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case connectionUsage >= connectionUsageConfig.HighWatermark:
			connectionUsageHigh = append(connectionUsageHigh, rowData)
			connectionUsageHighSum += connectionUsage
			connectionUsageHighCount++
		case connectionUsage >= connectionUsageConfig.LowWatermark:
			connectionUsageMediumSum += connectionUsage
			connectionUsageMediumCount++
		}
	}

	// connection usage data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.ConnectionUsageData = string(jsonBytesTotal)
	// connection usage high
	jsonBytesHigh, err := json.Marshal(connectionUsageHigh)
	if err != nil {
		return nil
	}
	de.result.CacheMissRatioHigh = ByteToFloat64(jsonBytesHigh)

	// connection usage high score deduction
	connectionUsageScoreDeductionHigh := (connectionUsageHighSum/float64(connectionUsageHighCount) - connectionUsageConfig.HighWatermark) / connectionUsageConfig.Unit * connectionUsageConfig.ScoreDeductionPerUnitHigh
	if connectionUsageScoreDeductionHigh > connectionUsageConfig.MaxScoreDeductionHigh {
		connectionUsageScoreDeductionHigh = connectionUsageConfig.MaxScoreDeductionHigh
	}
	// connection usage medium score deduction
	connectionUsageScoreDeductionMedium := (connectionUsageMediumSum/float64(connectionUsageMediumCount) - connectionUsageConfig.LowWatermark) / connectionUsageConfig.Unit * connectionUsageConfig.ScoreDeductionPerUnitMedium
	if connectionUsageScoreDeductionMedium > connectionUsageConfig.MaxScoreDeductionMedium {
		connectionUsageScoreDeductionMedium = connectionUsageConfig.MaxScoreDeductionMedium
	}
	// connection usage score
	de.result.ConnectionUsageScore = int(defaultMaxScore - connectionUsageScoreDeductionHigh - connectionUsageScoreDeductionMedium)
	if de.result.ConnectionUsageScore < constant.ZeroInt {
		de.result.ConnectionUsageScore = constant.ZeroInt
	}

	return nil
}

// checkActiveSessionNum check active session number
func (de *DefaultEngine) checkActiveSessionNum() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		avg_over_time(mysql_global_status_threads_running{instance=~"%s"}[$interval]) or 
		avg_over_time(mysql_global_status_threads_running{instance=~"%s"}[5m])
	`, host, host)
	case 2:
		query = fmt.Sprintf(`
		avg by (service_name) (avg_over_time(mysql_global_status_threads_running{service_name=~"%s"}[5m]) or 
		avg_over_time(mysql_global_status_threads_running{service_name=~"%s"}[5m]))
	`, serverName, serverName)
	}
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	activeSessionNumConfig := de.getItemConfig(defaultAverageActiveSessionNumItemName)

	var (
		activeSessionNum            float64
		activeSessionNumHighSum     float64
		activeSessionNumHighCount   int
		activeSessionNumMediumSum   float64
		activeSessionNumMediumCount int

		activeSessionNumHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		activeSessionNum, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case activeSessionNum >= activeSessionNumConfig.HighWatermark:
			activeSessionNumHigh = append(activeSessionNumHigh, rowData)
			activeSessionNumHighSum += activeSessionNum
			activeSessionNumHighCount++
		case activeSessionNum >= activeSessionNumConfig.LowWatermark:
			activeSessionNumMediumSum += activeSessionNum
			activeSessionNumMediumCount++
		}
	}

	// active session number data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.AverageActiveSessionNumData = string(jsonBytesTotal)
	// active session number high
	jsonBytesHigh, err := json.Marshal(activeSessionNumHigh)
	if err != nil {
		return nil
	}
	de.result.AverageActiveSessionNumHigh = string(jsonBytesHigh)

	// active session number high score deduction
	activeSessionNumScoreDeductionHigh := (activeSessionNumHighSum/float64(activeSessionNumHighCount) - activeSessionNumConfig.HighWatermark) / activeSessionNumConfig.Unit * activeSessionNumConfig.ScoreDeductionPerUnitHigh
	if activeSessionNumScoreDeductionHigh > activeSessionNumConfig.MaxScoreDeductionHigh {
		activeSessionNumScoreDeductionHigh = activeSessionNumConfig.MaxScoreDeductionHigh
	}
	// active session number medium score deduction
	activeSessionNumScoreDeductionMedium := (activeSessionNumMediumSum/float64(activeSessionNum) - activeSessionNumConfig.LowWatermark) / activeSessionNumConfig.Unit * activeSessionNumConfig.ScoreDeductionPerUnitMedium
	if activeSessionNumScoreDeductionMedium > activeSessionNumConfig.MaxScoreDeductionMedium {
		activeSessionNumScoreDeductionMedium = activeSessionNumConfig.MaxScoreDeductionMedium
	}
	// active session number score
	de.result.AverageActiveSessionNumScore = int(defaultMaxScore - activeSessionNumScoreDeductionHigh - activeSessionNumScoreDeductionMedium)
	if de.result.AverageActiveSessionNumScore < constant.ZeroInt {
		de.result.AverageActiveSessionNumScore = constant.ZeroInt
	}

	return nil
}

// checkCacheMissRatio checks cache miss ratio
func (de *DefaultEngine) checkCacheMissRatio() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	host := serverName + strconv.Itoa(portNum)
	// get prometheus version
	prometheusVersion, err := de.getPmmVersion()
	if err != nil {
		return err
	}
	var query string
	switch prometheusVersion {
	case 1:
		query = fmt.Sprintf(`
		1- (rate(mysql_global_status_table_open_cache_hits{instance=~"%s"}[$interval]) or 
		irate(mysql_global_status_table_open_cache_hits{instance=~"%s"}[5m]))/
		((rate(mysql_global_status_table_open_cache_hits{instance=~"%s"}[$interval]) or 
		irate(mysql_global_status_table_open_cache_hits{instance=~"%s"}[5m]))+
		(rate(mysql_global_status_table_open_cache_misses{instance=~"%s"}[$interval]) or 
		irate(mysql_global_status_table_open_cache_misses{instance=~"%s"}[5m])))
	`, host, host, host, host, host, host)
	case 2:
		query = fmt.Sprintf(`
		clamp_max((1 - avg by (service_name)(rate(mysql_global_status_table_open_cache_hits{service_name=~"%s"}[5m]) or 
		irate(mysql_global_status_table_open_cache_hits{service_name=~"%s"}[5m]))/
		avg by (service_name)((rate(mysql_global_status_table_open_cache_hits{service_name=~"%s"}[5m]) or 
		irate(mysql_global_status_table_open_cache_hits{service_name=~"%s"}[5m]))+
		(rate(mysql_global_status_table_open_cache_misses{service_name=~"%s"}[5m]) or 
		irate(mysql_global_status_table_open_cache_misses{service_name=~"%s"}[5m])))),1)
	`, serverName, serverName, serverName, serverName, serverName, serverName)
	}
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	cacheMissRatioConfig := de.getItemConfig(defaultCacheMissRatioItemName)

	var (
		cacheMissRatio            float64
		cacheMissRatioHighSum     float64
		cacheMissRatioHighCount   int
		cacheMissRatioMediumSum   float64
		cacheMissRatioMediumCount int

		cacheMissRatioHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		cacheMissRatio, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case cacheMissRatio >= cacheMissRatioConfig.HighWatermark:
			cacheMissRatioHigh = append(cacheMissRatioHigh, rowData)
			cacheMissRatioHighSum += cacheMissRatio
			cacheMissRatioHighCount++
		case cacheMissRatio >= cacheMissRatioConfig.LowWatermark:
			cacheMissRatioMediumSum += cacheMissRatio
			cacheMissRatioMediumCount++
		}
	}

	// cache miss ratio data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.CacheMissRatioData = ByteToFloat64(jsonBytesTotal)
	// cache miss ratio high
	jsonBytesHigh, err := json.Marshal(cacheMissRatioHigh)
	if err != nil {
		return nil
	}
	de.result.CacheMissRatioHigh = ByteToFloat64(jsonBytesHigh)

	// cache miss ratio high score deduction
	cacheMissRatioScoreDeductionHigh := (cacheMissRatioHighSum/float64(cacheMissRatioHighCount) - cacheMissRatioConfig.HighWatermark) / cacheMissRatioConfig.Unit * cacheMissRatioConfig.ScoreDeductionPerUnitHigh
	if cacheMissRatioScoreDeductionHigh > cacheMissRatioConfig.MaxScoreDeductionHigh {
		cacheMissRatioScoreDeductionHigh = cacheMissRatioConfig.MaxScoreDeductionHigh
	}
	// cache miss ratio medium score deduction
	cacheMissRatioScoreDeductionMedium := (cacheMissRatioHighSum/float64(cacheMissRatioMediumCount) - cacheMissRatioConfig.LowWatermark) / cacheMissRatioConfig.Unit * cacheMissRatioConfig.ScoreDeductionPerUnitMedium
	if cacheMissRatioScoreDeductionMedium > cacheMissRatioConfig.MaxScoreDeductionMedium {
		cacheMissRatioScoreDeductionMedium = cacheMissRatioConfig.MaxScoreDeductionMedium
	}
	// cache miss ratio score
	de.result.CacheMissRatioScore = int(defaultMaxScore - cacheMissRatioScoreDeductionHigh - cacheMissRatioScoreDeductionMedium)
	if de.result.CacheMissRatioScore < constant.ZeroInt {
		de.result.CacheMissRatioScore = constant.ZeroInt
	}

	return nil
}

// checkTableSize checks table size by checking rows
func (de *DefaultEngine) checkTableSize() error {
	// check table rows
	// get data
	sql := `
		select TABLE_SCHEMA,TABLE_NAME,TABLE_ROWS,(DATA_LENGTH+INDEX_LENGTH)/1024/1024/1024
		as TABLE_SIZE from TABLES
		where TABLE_TYPE='BASE TABLE';
	`
	result, err := de.monitorMysqlConn.Execute(sql)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	tableRowsConfig := de.getItemConfig(defaultTableRowsItemName)

	var (
		tableRows            float64
		tableRowsHighSum     float64
		tableRowsHighCount   int
		tableRowsMediumSum   float64
		tableRowsMediumCount int

		tableRowsHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		tableRows, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case tableRows >= tableRowsConfig.HighWatermark:
			tableRowsHigh = append(tableRowsHigh, rowData)
			tableRowsHighSum += tableRows
			tableRowsHighCount++
		case tableRows >= tableRowsConfig.LowWatermark:
			tableRowsMediumSum += tableRows
			tableRowsMediumCount++
		}
	}

	// table rows data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.TableSizeData = string(jsonBytesTotal)
	// table rows high
	jsonBytesHigh, err := json.Marshal(tableRowsHigh)
	if err != nil {
		return nil
	}
	de.result.TableSizeHigh = string(jsonBytesHigh)

	// table rows high score deduction
	tableRowsScoreDeductionHigh := (tableRowsHighSum/float64(tableRowsHighCount) - tableRowsConfig.HighWatermark) / tableRowsConfig.Unit * tableRowsConfig.ScoreDeductionPerUnitHigh
	if tableRowsScoreDeductionHigh > tableRowsConfig.MaxScoreDeductionHigh {
		tableRowsScoreDeductionHigh = tableRowsConfig.MaxScoreDeductionHigh
	}
	// table rows medium score deduction
	tableRowsScoreDeductionMedium := (tableRowsMediumSum/float64(tableRowsMediumCount) - tableRowsConfig.LowWatermark) / tableRowsConfig.Unit * tableRowsConfig.ScoreDeductionPerUnitMedium
	if tableRowsScoreDeductionMedium > tableRowsConfig.MaxScoreDeductionMedium {
		tableRowsScoreDeductionMedium = tableRowsConfig.MaxScoreDeductionMedium
	}
	// table rows score
	de.result.TableSizeScore = int(defaultMaxScore - tableRowsScoreDeductionHigh - tableRowsScoreDeductionMedium)
	if de.result.TableSizeScore < constant.ZeroInt {
		de.result.TableSizeScore = constant.ZeroInt
	}

	return nil
}

// checkSlowQuery checks slow query
// TODO
func (de *DefaultEngine) checkSlowQuery() error {
	// check slow query execution time
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

	slowQueryExecutionTimeConfig := de.getItemConfig(defaultSlowQueryExecutionTimeItemName)

	var (
		slowQueryExecutionTime            float64
		slowQueryExecutionTimeHighSum     float64
		slowQueryExecutionTimeHighCount   int
		slowQueryExecutionTimeMediumSum   float64
		slowQueryExecutionTimeMediumCount int

		slowQueryExecutionTimeHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		slowQueryExecutionTime, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case slowQueryExecutionTime >= slowQueryExecutionTimeConfig.HighWatermark:
			slowQueryExecutionTimeHigh = append(slowQueryExecutionTimeHigh, rowData)
			slowQueryExecutionTimeHighSum += slowQueryExecutionTime
			slowQueryExecutionTimeHighCount++
		case slowQueryExecutionTime >= slowQueryExecutionTimeConfig.LowWatermark:
			slowQueryExecutionTimeMediumSum += slowQueryExecutionTime
			slowQueryExecutionTimeMediumCount++
		}
	}

	// slow query execution time high score deduction
	slowQueryExecutionTimeScoreDeductionHigh := (slowQueryExecutionTimeHighSum/float64(slowQueryExecutionTimeHighCount) - slowQueryExecutionTimeConfig.HighWatermark) / slowQueryExecutionTimeConfig.Unit * slowQueryExecutionTimeConfig.ScoreDeductionPerUnitHigh
	if slowQueryExecutionTimeScoreDeductionHigh > slowQueryExecutionTimeConfig.MaxScoreDeductionHigh {
		slowQueryExecutionTimeScoreDeductionHigh = slowQueryExecutionTimeConfig.MaxScoreDeductionHigh
	}
	// slow query execution time medium score deduction
	slowQueryExecutionTimeScoreDeductionMedium := (slowQueryExecutionTimeMediumSum/float64(slowQueryExecutionTimeMediumCount) - slowQueryExecutionTimeConfig.LowWatermark) / slowQueryExecutionTimeConfig.Unit * slowQueryExecutionTimeConfig.ScoreDeductionPerUnitMedium
	if slowQueryExecutionTimeScoreDeductionMedium > slowQueryExecutionTimeConfig.MaxScoreDeductionMedium {
		slowQueryExecutionTimeScoreDeductionMedium = slowQueryExecutionTimeConfig.MaxScoreDeductionMedium
	}

	// check slow query rows examined
	// TODO
	query = fmt.Sprintf(`
		sum(avg by (node_name,mode) (clamp_max(((avg by (mode,node_name) ((
		clamp_max(rate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[20s]),1)) or
		(clamp_max(irate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[5m]),1)) ))*100 or
		(avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[20s]) or
		avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[5m]))),100)))
	`, serverName, serverName, serverName, serverName)
	result, err = de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length = result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	slowQueryRowsExaminedConfig := de.getItemConfig(defaultSlowQueryRowsExaminedItemName)

	var (
		slowQueryRowsExamined            float64
		slowQueryRowsExaminedHighSum     float64
		slowQueryRowsExaminedHighCount   int
		slowQueryRowsExaminedMediumSum   float64
		slowQueryRowsExaminedMediumCount int

		slowQueryRowsExaminedHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		slowQueryRowsExamined, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case slowQueryRowsExamined >= slowQueryRowsExaminedConfig.HighWatermark:
			slowQueryRowsExaminedHigh = append(slowQueryRowsExaminedHigh, rowData)
			slowQueryRowsExaminedHighSum += slowQueryRowsExamined
			slowQueryRowsExaminedHighCount++
		case slowQueryRowsExamined >= slowQueryRowsExaminedConfig.LowWatermark:
			slowQueryRowsExaminedMediumSum += slowQueryRowsExamined
			slowQueryRowsExaminedMediumCount++
		}
	}

	// slow query rows examined data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.SlowQueryData = string(jsonBytesTotal)
	// slow query rows examined high
	jsonBytesHigh, err := json.Marshal(slowQueryRowsExaminedHigh)
	if err != nil {
		return nil
	}
	de.result.SlowQueryAdvice = string(jsonBytesHigh)

	// slow query rows examined high score deduction
	slowQueryRowsExaminedScoreDeductionHigh := (slowQueryRowsExaminedHighSum/float64(slowQueryRowsExaminedHighCount) - slowQueryRowsExaminedConfig.HighWatermark) / slowQueryRowsExaminedConfig.Unit * slowQueryRowsExaminedConfig.ScoreDeductionPerUnitHigh
	if slowQueryRowsExaminedScoreDeductionHigh > slowQueryRowsExaminedConfig.MaxScoreDeductionHigh {
		slowQueryRowsExaminedScoreDeductionHigh = slowQueryRowsExaminedConfig.MaxScoreDeductionHigh
	}
	// slow query rows examined medium score deduction
	slowQueryRowsExaminedScoreDeductionMedium := (slowQueryRowsExaminedMediumSum/float64(slowQueryRowsExaminedMediumCount) - slowQueryRowsExaminedConfig.LowWatermark) / slowQueryRowsExaminedConfig.Unit * slowQueryRowsExaminedConfig.ScoreDeductionPerUnitMedium
	if slowQueryRowsExaminedScoreDeductionMedium > slowQueryRowsExaminedConfig.MaxScoreDeductionMedium {
		slowQueryRowsExaminedScoreDeductionMedium = slowQueryRowsExaminedConfig.MaxScoreDeductionMedium
	}

	// slow query score
	de.result.SlowQueryScore = int(defaultMaxScore - slowQueryExecutionTimeScoreDeductionHigh - slowQueryExecutionTimeScoreDeductionMedium - slowQueryRowsExaminedScoreDeductionHigh - slowQueryRowsExaminedScoreDeductionMedium)
	if de.result.SlowQueryScore < constant.ZeroInt {
		de.result.SlowQueryScore = constant.ZeroInt
	}

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
		de.result.SlowQueryScore*(de.getItemConfig(defaultSlowQueryExecutionTimeItemName).ItemWeight+de.getItemConfig(defaultSlowQueryRowsExaminedItemName).ItemWeight)) /
		constant.MaxPercentage

	if de.result.WeightedAverageScore < defaultMinScore {
		de.result.WeightedAverageScore = defaultMinScore
	}
}

// postRun performs post-run actions, for now, it ony saves healthcheck result to the middleware
func (de *DefaultEngine) postRun() error {
	// save result
	return de.Repository.SaveResult(de.result)
}

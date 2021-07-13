package healthcheck

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/now"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	defaultEngineConfigAddr   = "localhost:3306"
	defaultEngineConfigDBName = "performance_schema"
	defaultEngineConfigDBUser = "root"
	defaultEngineConfigDBPass = "root"

	applicationMysqlAddr   = "192.168.10.210:3306"
	applicationMysqlDBName = "performance_schema"
	applicationMysqlDBUser = "root"
	applicationMysqlDBPass = "root"

	defaultEngineConfigID                          = 1
	defaultEngineConfigItemName                    = "test_item"
	defaultEngineConfigItemWeight                  = 5
	defaultEngineConfigLowWatermark                = 50.00
	defaultEngineConfigHighWatermark               = 70.00
	defaultEngineConfigUnit                        = 10.00
	defaultEngineConfigScoreDeductionPerUnitHigh   = 20.00
	defaultEngineConfigMaxScoreDeductionHigh       = 100.00
	defaultEngineConfigScoreDeductionPerUnitMedium = 10.00
	defaultEngineConfigMaxScoreDeductionMedium     = 50.00
	defaultEngineConfigDelFlag                     = 0
	defaultEngineConfigCreateTimeString            = "2021-01-21 10:00:00.000000"
	defaultEngineConfigLastUpdateTimeString        = "2021-01-21 13:00:00.000000"

	serviceID        = 1
	serviceStartTime = "2021-01-21 10:00:00.000000"
	serviceEndTime   = "2021-01-21 13:00:00.000000"
	serviceStep      = 1 * time.Millisecond
)

var defaultEngineConfigRepo = initDefaultEngineConfigRepo()
var mysqlServerRepo = initMySQLServerRepo()

func initDefaultEngineConfigRepo() *Repository {
	pool, err := mysql.NewPoolWithDefault(defaultEngineConfigAddr, defaultEngineConfigDBName, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareClusterRepo() failed", err))
		return nil
	}

	return NewRepository(pool)
}

func initMySQLServerRepo() *metadata.MySQLServerRepo {
	pool, err := mysql.NewPoolWithDefault(defaultEngineConfigAddr, defaultEngineConfigDBName, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMySQLServerRepo() failed", err))
		return nil
	}

	return metadata.NewMySQLServerRepo(pool)
}

func TestDefaultEngineConfig_Validate(t *testing.T) {
	asst := assert.New(t)
	// load config
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	result, err := defaultEngineConfigRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Validate() failed", err))
	defaultEngineConfigList := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultEngineConfigList {
		defaultEngineConfigList[i] = NewEmptyDefaultItemConfig()
	}
	err = result.MapToStructSlice(defaultEngineConfigList, constant.DefaultMiddlewareTag)
	asst.Nil(err, common.CombineMessageWithError("test Validate() failed", err))
	entityList := NewEmptyDefaultEngineConfig()
	for i := range defaultEngineConfigList {
		itemName := defaultEngineConfigList[i].ItemName
		entityList[itemName] = defaultEngineConfigList[i]
	}
	// validate config
	validate := entityList.Validate()
	asst.Equal(nil, validate, "test Validate() failed")
}

func TestDefaultEngine_Run(t *testing.T) {
	asst := assert.New(t)
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)
	startTime, _ := now.Parse(serviceStartTime)
	endTime, _ := now.Parse(serviceEndTime)

	id, err := defaultEngineConfigRepo.InitOperation(serviceID, startTime, endTime, serviceStep)
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))

	mysqlServerService := metadata.NewMySQLServerService(mysqlServerRepo)
	err = mysqlServerService.GetByID(1)
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))

	applicationMySQLConn, err := mysql.NewConn(applicationMysqlAddr, applicationMysqlDBName, applicationMysqlDBUser, applicationMysqlDBPass)
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))

	monitorSystem, err := mysqlServer.GetMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	var (
		monitorPrometheusConn *prometheus.Conn
		monitorClickhouseConn *clickhouse.Conn
		monitorMySQLConn      *mysql.Conn
	)

	monitorSystemType := monitorSystem.GetSystemType()
	switch monitorSystemType {
	case 1:
		// pmm 1.x
		// init prometheus connection
		prometheusAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNum())
		prometheusConfig := prometheus.NewConfig(prometheusAddr, prometheus.DefaultRoundTripper)
		monitorPrometheusConn, err = prometheus.NewConnWithConfig(prometheusConfig)
		asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))

		// init mysql connection
		mysqlAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		monitorMySQLConn, err = mysql.NewConn(mysqlAddr, defaultMonitorMySQLDBName, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
		asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	case 2:
		// pmm 2.x
		// init prometheus connection
		prometheusAddr := fmt.Sprintf("%s:%d%s", monitorSystem.GetHostIP(), monitorSystem.GetPortNum(), monitorSystem.GetBaseURL())
		prometheusConfig := prometheus.NewConfigWithBasicAuth(prometheusAddr, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
		monitorPrometheusConn, err = prometheus.NewConnWithConfig(prometheusConfig)
		asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
		// init clickhouse connection
		clickhouseAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		monitorClickhouseConn, err = clickhouse.NewConnWithDefault(clickhouseAddr, defaultMonitorClickhouseDBName, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
		asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))

		operationInfo := NewOperationInfo(id, mysqlServer, monitorSystem, startTime, endTime, serviceStep)
		defaultEngine := NewDefaultEngine(defaultEngineConfigRepo, operationInfo, applicationMySQLConn, monitorPrometheusConn, monitorClickhouseConn, monitorMySQLConn)
		err = defaultEngine.run()
		asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	}
}

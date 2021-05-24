package healthcheck

import (
	"fmt"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"testing"
)

const (
	defaultEngineConfigAddr   = "localhost:3306"
	defaultEngineConfigDBName = "das"
	defaultEngineConfigDBUser = "root"
	defaultEngineConfigDBPass = "rootroot"

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
)

var defaultEngineConfigRepo = initDefaultEngineConfigRepo()

func initDefaultEngineConfigRepo() *Repository {
	pool, err := mysql.NewPoolWithDefault(defaultEngineConfigAddr, defaultEngineConfigDBName, defaultEngineConfigDBUser, defaultEngineConfigDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMiddlewareClusterRepo() failed", err))
		return nil
	}

	return NewRepository(pool)
}

func TestMiddlewareClusterRepo_Execute(t *testing.T) {
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healcheck Repository.loadEngineConfig() sql: \n%s\nplaceholders: %s", sql)
	result, _ := defaultEngineConfigRepo.Execute(sql)

	// init []*DefaultItemConfig
	defaultEngineConfigList := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultEngineConfigList {
		defaultEngineConfigList[i] = NewEmptyDefaultItemConfig()
	}
	// map to struct
	result.MapToStructSlice(defaultEngineConfigList, constant.DefaultMiddlewareTag)
	for i := range defaultEngineConfigList {
		fmt.Println(defaultEngineConfigList[i])
	}
}

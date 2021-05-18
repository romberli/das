package healthcheck

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	defaultEngineConfigAddr   = "192.168.10.210:3306"
	defaultEngineConfigDBName = "performance_schema"
	defaultEngineConfigDBUser = "root"
	defaultEngineConfigDBPass = "root"

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
	asst := assert.New(t)

	sql := `select variable_name, variable_value
		from global_variables;`
	result, err := defaultEngineConfigRepo.Execute(sql)
	vlist := make([]*GlobalVariables, result.RowNumber())
	for i := range vlist {
		vlist[i] = NewEmptyGlobalVariables()
	}
	// map to struct
	err = result.MapToStructSlice(vlist, constant.DefaultMiddlewareTag)
	// init entity
	elist := make(map[string]string, result.RowNumber())
	for i := range vlist {
		variableName := vlist[i].VariableName
		elist[variableName] = vlist[i].VariableValue
	}

	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

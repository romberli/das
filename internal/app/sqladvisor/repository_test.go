package sqladvisor

import (
	"testing"

	"github.com/romberli/das/global"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

func init() {
	initDASMySQLPool()
}

const (
	defaultDASMySQLAddr = "192.168.137.11:3306"
	defaultDASMySQLName = "das"
	defaultDASMySQLUser = "root"
	defaultDASMySQLPass = "root"

	defaultDBID    = 1
	defaultSQLText = "select * from t_meta_db_info where create_time<'2021-01-01';"
	defaultAdvice  = "[\n {\n  \"ID\": \"B95017DB61875675\",\n  \"Fingerprint\": \"select * from t_meta_db_info where create_time\\u003c?\",\n  \"Score\": 95,\n  \"Sample\": \"select * from t_meta_db_info where create_time\\u003c'2021-01-01'\",\n  \"Explain\": null,\n  \"HeuristicRules\": [\n    {\n      \"Item\": \"COL.001\",\n      \"Severity\": \"L1\",\n      \"Summary\": \"不建议使用 SELECT * 类型查询\",\n      \"Content\": \"当表结构变更时，使用 * 通配符选择所有列将导致查询的含义和行为会发生更改，可能导致查询返回更多的数据。\",\n      \"Case\": \"select * from tbl where id=1\",\n      \"Position\": 0\n    }\n  ],\n  \"IndexRules\": null,\n  \"Tables\": [\n    \"`soar`.`t_meta_db_info`\"\n  ]\n}\n]"
	defaultMessage = ""
)

var repository = initRepository()

func initDASMySQLPool() *mysql.Pool {
	var err error

	global.DASMySQLPool, err = mysql.NewPoolWithDefault(defaultDASMySQLAddr, defaultDASMySQLName, defaultDASMySQLUser, defaultDASMySQLPass)
	log.Infof("pool: %v, error: %v", global.DASMySQLPool, err)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}

	return global.DASMySQLPool
}

func initRepository() *Repository {
	return NewRepository(global.DASMySQLPool)
}

func deleteResult() error {
	sql := `delete from t_sa_operation_info;`
	_, err := repository.Execute(sql)

	return err
}

func TestRepositoryAll(t *testing.T) {
	TestRepository_Execute(t)
	TestRepository_Save(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := "select 1;"
	result, err := repository.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestRepository_Save(t *testing.T) {
	asst := assert.New(t)

	err := deleteResult()
	err = repository.Save(defaultDBID, defaultSQLText, defaultAdvice, defaultMessage)
	asst.Nil(err, "test Save() failed")
	err = deleteResult()
}

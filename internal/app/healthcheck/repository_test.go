package healthcheck

import (
	"fmt"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	// modify these connection information
	dbAddr   = "127.0.0.1:3306"
	dbDBName = "das"
	dbDBUser = "root"
	dbDBPass = "rootroot"
)

var repository = initRepository()

func initRepository() *Repository {
	pool, err := mysql.NewPoolWithDefault(dbAddr, dbDBName, dbDBUser, dbDBPass)
	log.Infof("pool: %v, error: %v", pool, err)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}

	return NewRepository(pool)
}

func TestRepository_GetEngineConfig(t *testing.T) {
	asst := assert.New(t)

	entities, err := repository.GetEngineConfig()
	asst.Nil(err, common.CombineMessageWithError("test GetEngineConfig() failed", err))
	fmt.Println(entities)
	//systemName := entities[0].GetAppName()
	//asst.Nil(err, common.CombineMessageWithError("test GetEngineConfig() failed", err))
	//asst.Equal(onlineAppName, systemName, "test GetEngineConfig() failed")
}

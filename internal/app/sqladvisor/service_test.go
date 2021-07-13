package sqladvisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultSoarBin    = "/Users/romber/work/source_code/go/src/github.com/romberli/das/bin/soar"
	defaultConfigFile = "/Users/romber/work/source_code/go/src/github.com/romberli/das/config/soar.yaml"

	defaultFingerprint = "select * from t_meta_db_info where create_time<?"
	defaultSQLID       = "B95017DB61875675"
)

var service = createService()

func createService() *Service {
	return NewService(defaultSoarBin, defaultConfigFile)
}
func TestService_All(t *testing.T) {
	TestService_GetFingerprint(t)
	TestService_GetFingerprint(t)
	TestService_GetSQLID(t)
	TestService_Advise(t)
}

func TestService_GetFingerprint(t *testing.T) {
	asst := assert.New(t)

	fingerprint := service.GetFingerprint(defaultSQLText)
	asst.Equal(defaultFingerprint, fingerprint, "test GetFingerprint() failed")
}

func TestService_GetSQLID(t *testing.T) {
	asst := assert.New(t)

	sqlID := service.GetSQLID(defaultSQLText)
	asst.Equal(defaultSQLID, sqlID, "test GetSQLID() failed")
}

func TestService_Advise(t *testing.T) {
	asst := assert.New(t)

	advice, err := service.Advise(defaultDBID, defaultSQLText)
	asst.Nil(err, "test Advise() failed")
	asst.NotEmpty(advice, "test Advise() failed")
}

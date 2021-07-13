package sqladvisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultDBSoarMySQLUser = "root"
	defaultDBSoarMySQLPass = "root"
)

var advisor = NewDefaultAdvisor(defaultSoarBin, defaultConfigFile)

func TestDefaultAdvisor_All(t *testing.T) {
	TestDefaultAdvisor_GetFingerprint(t)
	TestDefaultAdvisor_GetSQLID(t)
	TestDefaultAdvisor_Advise(t)
}

func TestDefaultAdvisor_GetFingerprint(t *testing.T) {
	asst := assert.New(t)

	fingerprint := advisor.GetFingerprint(defaultSQLText)
	asst.Equal(defaultFingerprint, fingerprint, "test GetFingerprint() failed")
}

func TestDefaultAdvisor_GetSQLID(t *testing.T) {
	asst := assert.New(t)

	sqlID := advisor.GetSQLID(defaultSQLText)
	asst.Equal(defaultSQLID, sqlID, "test GetSQLID() failed")
}

func TestDefaultAdvisor_Advise(t *testing.T) {
	asst := assert.New(t)

	advice, message, err := advisor.advise(defaultDBID, defaultSQLText, defaultDBSoarMySQLUser, defaultDBSoarMySQLPass)
	asst.Nil(err, "test Advise() failed")
	asst.NotEmpty(advice, "test Advise() failed")
	t.Log(message)
	t.Log(advice)
}

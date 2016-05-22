package service

import (
	testconfig "parkme-api/tests/config"
	"testing"
)

func TestServiceBase(t *testing.T) {
	testconfig.InitTestsDatabase()
	InitDbService()

	sess, col := Connect("testCollection")
	if sess == nil || col == nil {
		t.Fatal("Cannot connect to the mongodb service")
	}
}

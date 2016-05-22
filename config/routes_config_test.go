package config

import (
	"testing"
)

const routesFilePath = "../parkme-api/config/routes.json"

func TestRoutesConfig(t *testing.T) {
	InitRoutes(routesFilePath)

	if Routes == nil {
		t.Fatal("Retrieving the routes from the configuration file has failed!")
	}
}

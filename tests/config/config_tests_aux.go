package config

import (
	"log"
	"os"
	"parkme-api/config"
)

// All the api files are 3 levels deep into the folder hierarchy (ex: ./api/app/someapi/someapi.go), hence the ../../../ prefix
const routesFilePath = "../../../tests/config/test_routes.json"

const (
	envApplicationName    = "PARKME_TESTAPP_NAME"
	envAPIInstance        = "PARKME_TESTAPP_INSTANCE"
	envHTTPServerAddress  = "PARKME_TESTAPP_HTTP"
	envDatabaseName       = "PARKME_TESTAPP_DB_NAME"
	envDatabaseConnection = "PARKME_TESTAPP_DB_CONN"
)

// InitTestsApp initializes the application used for testing
func InitTestsApp() {
	config.ApplicationName = os.Getenv(envApplicationName)
	config.APIInstance = os.Getenv(envAPIInstance)
	config.HTTPServerAddress = os.Getenv(envHTTPServerAddress)
}

// InitTestsDatabase initializes the connection parameters to the database used for testing
func InitTestsDatabase() {
	dbName := os.Getenv(envDatabaseName)
	dbConn := os.Getenv(envDatabaseConnection)

	if len(dbName) == 0 || len(dbConn) == 0 {
		log.Fatal("Environment variables for the test database are not set!")
	}

	config.DbName = dbName
	config.DbConnectionString = dbConn
}

// InitTestsRoutes initializez the routes used for testing the endpoints
func InitTestsRoutes() {
	config.InitRoutes(routesFilePath)
}

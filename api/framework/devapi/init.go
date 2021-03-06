package devapi

import (
	"parkme-api/config"
	"parkme-api/util/jsonutil"
	"log"
)

// Routes configuration file path
var devRoutes = `
[
    {
        "id": "DevApiRoute",
        "endpoint": "/dev",
        "actions": {
            "CreateAppUser": {
                "type": "POST",
                "allowAnonymous": true
            },
            "ActivateAppUser": {
                "type": "GET",
                "allowAnonymous": true
            }
        }
    }
]`

// InitDevRoutes initializes the routes used for development purposes only
func InitDevRoutes() {
	var route []config.Route

	var err = jsonutil.DeserializeJSON([]byte(devRoutes), &route)
	if err != nil {
		log.Fatal(err)
	}

	err = config.AddRoutes(false, route...)
	if err != nil {
		log.Fatalf("[InitDevRoutes] %v\n", err)
	}
}

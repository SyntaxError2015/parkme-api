package parkservice

import (
	"parkme-api/orm/dbmodels"
	"parkme-api/orm/service"

	"gopkg.in/mgo.v2/bson"
)

const collectionName = "parks"

// Create adds a new Park to the database
func Create(park *dbmodels.Park) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	if park.ID == "" {
		park.ID = bson.NewObjectId()
	}

	err := collection.Insert(park)

	return err
}

// Update adds a new Park to the database
func Update(id bson.ObjectId, park *dbmodels.Park) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	err := collection.UpdateId(id, park)

	return err
}

// Get retrieves an Park from the database, based on its ID
func Get(parkID bson.ObjectId) (*dbmodels.Park, error) {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	park := dbmodels.Park{}
	err := collection.FindId(parkID).One(&park)

	return &park, err
}

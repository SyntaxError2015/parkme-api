package slotservice

import (
	"apiGO/service"
	"parkme-api/orm/dbmodels"
	"parkme-api/orm/models"

	"gopkg.in/mgo.v2/bson"
)

const collectionName = "slots"

// GetAll retrieves a Slot slice from the database, based on a park's ID
func GetAll(parkID bson.ObjectId) ([]*dbmodels.Slot, error) {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	var slots []*dbmodels.Slot
	err := collection.Find(bson.M{"parkID": parkID}).All(&slots)

	return slots, err
}

// CreateMultiple adds new Slots to the database
func CreateMultiple(slots []*models.Slot) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	err := collection.Insert(slots...)

	return slots, err
}

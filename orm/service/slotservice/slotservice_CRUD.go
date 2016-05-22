package slotservice

import (
	"errors"
	"parkme-api/orm/dbmodels"
	"parkme-api/orm/service"

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
func CreateMultiple(slots []dbmodels.Slot) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	for i := 0; i < len(slots); i++ {
		if len(slots[i].ID) == 0 {
			slots[i].ID = bson.NewObjectId()
		}

		err := collection.Insert(slots[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateMultiple updates multiple Slot entities
func UpdateMultiple(slots []dbmodels.Slot) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	var err error
	for _, slot := range slots {
		if len(slot.ID) == 0 {
			return errors.New("Empty ID for update")
		}

		_, err = collection.UpsertId(slot.ID, slot)

		if err != nil {
			return err
		}
	}

	return nil
}

// Update updates a certain slot
func Update(slotID bson.ObjectId, slot *dbmodels.Slot) error {
	session, collection := service.Connect(collectionName)
	defer session.Close()

	err := collection.UpdateId(slotID, slot)

	return err
}

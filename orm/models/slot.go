package models

import (
	"parkme-api/orm/dbmodels"

	"gopkg.in/mgo.v2/bson"
)

// Slot reprents a square in a parking lot, where the car is parked
type Slot struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Park     Park          `bson:"park" json:"park"`
	Position Point         `bson:"position" json:"position"`
}

// Expand copies the dbmodels.Slot to a Slot expands all
// the components by fetching them from the database
func (slot *Slot) Expand(dbSlot dbmodels.Slot) {
	slot.ID = dbSlot.ID
	slot.Position.Expand(dbSlot.Position)

	park, _ = parkservice.Get(dbSlot.ParkID)

	slot.Park = park
}

// Collapse coppies the Slot to a dbmodels.Slot user and
// only keeps the unique identifiers from the inner components
func (slot *Slot) Collapse() *dbmodels.Slot {
	dbslot := dbmodels.Slot{
		ID:     slot.ID,
		ParkID: slot.Park.ID,
	}

	var position = slot.Position.Collapse()

	dbslot.Position = *position

	return &dbslot
}

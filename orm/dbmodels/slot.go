package dbmodels

import "gopkg.in/mgo.v2/bson"

// Slot reprents a square in a parking lot, where the car is parked
type Slot struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	ParkID   bson.ObjectId `bson:"parkID,omitempty" json:"parkID"`
	Position Point         `bson:"position,omitempty" json:"position"`
}

// Equal compares two Slot objects. Implements the Objecter interface
func (slot Slot) Equal(obj Objecter) bool {
	otherSlot, ok := obj.(Slot)
	if !ok {
		return false
	}

	switch {
	case slot.ID != otherSlot.ID:
		return false
	case slot.ParkID != otherSlot.ParkID:
		return false
	case !slot.Position.Equal(otherSlot.Position):
		return false
	}

	return true
}

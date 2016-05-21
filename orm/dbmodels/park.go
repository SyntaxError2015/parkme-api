package dbmodels

import "gopkg.in/mgo.v2/bson"

// Park represents an entire parking lot, which has one or more slots in which cars are parked
type Park struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Address  string        `bson:"address" json:"address"`
	Status   int           `bson:"status" json:"status"`
	Position Point         `bson:"position" json:"position"`
}

// Equal compares two Park objects. Implements the Objecter interface
func (park Park) Equal(obj Objecter) bool {
	otherPark, ok := obj.(Park)
	if !ok {
		return false
	}

	switch {
	case park.ID != otherPark.ID:
		return false
	case park.Address != otherPark.Address:
		return false
	case park.Status != otherPark.Status:
		return false
	case !park.Position.Equal(otherPark.Position):
		return false
	}

	return true
}

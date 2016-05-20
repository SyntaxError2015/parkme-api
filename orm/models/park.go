package models

import (
	"parkme-api/orm/dbmodels"

	"gopkg.in/mgo.v2/bson"
)

// Park represents an entire parking lot, which has one or more slots in which cars are parked
type Park struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Address  string        `bson:"address" json:"address"`
	Position Point         `bson:"position" json:"position"`
}

// Expand copies the dbmodels.Park to a Park expands all
// the components by fetching them from the database
func (park *Park) Expand(dbPark dbmodels.Park) {
	park.ID = dbPark.ID
	park.Address = dbPark.Address

	var position = &Point{}
	position.Expand(dbPark.Position)

	park.Position = *position
}

// Collapse coppies the Park to a dbmodels.Park user and
// only keeps the unique identifiers from the inner components
func (park *Park) Collapse() *dbmodels.Park {
	dbPark := dbmodels.Park{
		ID:      park.ID,
		Address: park.Address,
	}

	var position = park.Position.Collapse()

	dbPark.Position = *position

	return &dbPark
}

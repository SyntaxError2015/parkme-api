package models

import (
	"parkme-api/orm/dbmodels"
	"parkme-api/orm/service/slotservice"

	"gopkg.in/mgo.v2/bson"
)

// Constants describing the status of a parking lot
const (
	ParkStatusOffline = iota
	ParkStatusOnline  = iota
)

// Park represents an entire parking lot, which has one or more slots in which cars are parked
type Park struct {
	ID        bson.ObjectId `json:"id"`
	AppUserID bson.ObjectId `json:"appUserID"`
	Address   string        `json:"address"`
	Status    int           `json:"status"`
	Position  Point         `json:"position"`
	Slots     []Slot        `json:"slots"`
}

// Expand copies the dbmodels.Park to a Park expands all
// the components by fetching them from the database
func (park *Park) Expand(dbPark dbmodels.Park) {
	park.ID = dbPark.ID
	park.AppUserID = dbPark.AppUserID
	park.Address = dbPark.Address
	park.Status = dbPark.Status

	var position = &Point{}
	position.Expand(dbPark.Position)

	dbSlots, _ := slotservice.GetAll(park.ID)
	slots := make([]Slot, len(dbSlots))
	for i := 0; i < len(dbSlots); i++ {
		slots[i].Expand(*dbSlots[i])
	}

	park.Position = *position
	park.Slots = slots
}

// Collapse coppies the Park to a dbmodels.Park user and
// only keeps the unique identifiers from the inner components
func (park *Park) Collapse() *dbmodels.Park {
	dbPark := dbmodels.Park{
		ID:        park.ID,
		AppUserID: park.AppUserID,
		Address:   park.Address,
		Status:    park.Status,
	}

	var position = park.Position.Collapse()

	dbPark.Position = *position

	return &dbPark
}

package slotapi

import (
	"parkme-api/orm/models"

	"gopkg.in/mgo.v2/bson"
)

// SlotUpdateModel is a model for receiving an update for a certain slot
type SlotUpdateModel struct {
	ParkID bson.ObjectId `json:"parkID"`
	Slot   models.Slot   `json:"slot"`
}

package parkapi

import (
	"log"
	"net/http"
	"parkme-api/api"
	"parkme-api/orm/dbmodels"
	"parkme-api/orm/models"
	"parkme-api/orm/service/parkservice"
	"parkme-api/orm/service/slotservice"
	"parkme-api/util/jsonutil"
)

// ParkAPI defines the API for managing parking lots
type ParkAPI int

// Register is an endpoint used for creating application users
func (p *ParkAPI) Register(params *api.Request) api.Response {
	var model = &models.Park{}

	err := jsonutil.DeserializeJSON(params.Body, model)
	if err != nil {
		return api.BadRequest(api.ErrEntityFormat)
	}

	var parkingSlots = model.Slots
	if parkingSlots == nil || len(parkingSlots) == 0 {
		return api.BadRequest(api.ErrEntityIntegrity)
	}

	existingParkPlace, err := parkservice.Get(model.ID)
	if err != nil || existingParkPlace == nil {
		return handleNewParkRegistration(model)
	}

	// update all the slots with their data
	slots := make([]dbmodels.Slot, len(parkingSlots))
	for i := 0; i < len(slots); i++ {
		slots[i] = *parkingSlots[i].Collapse()
	}
	err = slotservice.UpdateMultiple(slots)
	if err != nil {
		return api.InternalServerError(err)
	}

	return api.StatusResponse(http.StatusOK)
}

// GetAll returns all the exiting parking places
func (p *ParkAPI) GetAll(params *api.Request) api.Response {
	parkingPlaces, err := parkservice.GetAll()
	if err != nil {
		log.Println(err)
		return api.InternalServerError(err)
	}

	// append(piPark.Slots, models.Slot
	// parks := make([]*models.Park, len(parkingPlaces))

	var parks []models.Park

	log.Println(*parkingPlaces[0])

	for i := 0; i < len(parkingPlaces); i++ {
		parks = append(parks, models.Park{})

		parks[i].Expand(*parkingPlaces[i])
	}

	return api.JSONResponse(http.StatusOK, parks)
}

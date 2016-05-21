package parkapi

import (
	"net/http"
	"parkme-api/api"
	"parkme-api/orm/models"
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

	if !params.Identity.IsAuthorized() {
		return handleNewParkRegistration(park)
	}

	// update all the slots with their data
	slots := make([]dbModels.Slot, len(parkingSlots))
	for i := 0; i < len(slots); i++ {
		slots[i] = parkingSlots[i].Collapse()
	}
	err = slotservice.UpdateMultiple(slots)
	if err != nil {
		return api.InternalServerError(err)
	}

	return api.StatusResponse(http.StatusOK)
}

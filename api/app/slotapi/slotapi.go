package slotapi

import (
	"net/http"
	"parkme-api/api"
	"parkme-api/orm/service/slotservice"
	"parkme-api/util/jsonutil"
)

// SlotAPI defines the API for managing parking slots
type SlotAPI int

// UpdateSlot updates the status of a certain parking slot of a parking place
func (s *SlotAPI) UpdateSlot(params *api.Request) api.Response {
	model = &SlotUpdateModel{}

	err := jsonutil.DeserializeJSON(params.Body, model)
	if err != nil {
		return api.BadRequest(api.ErrEntityFormat)
	}

	slot = model.Collapse()
	err = slotservice.Update(slot)
	if err != nil {
		return err
	}

	return api.StatusResponse(http.StatusOK)
}

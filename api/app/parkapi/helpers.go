package parkapi

import (
	"fmt"
	"net/http"
	"parkme-api/api"
	"parkme-api/auth"
	"parkme-api/auth/identity"
	"parkme-api/orm/models"
	"parkme-api/orm/service/parkservice"
	"parkme-api/orm/service/slotservice"
	"parkme-api/util"

	"gopkg.in/mgo.v2/bson"
)

func handleNewParkRegistration(model *models.Park) api.Response {
	err := createParkAppUser(model)
	if err != nil {
		return api.InternalServerError(err)
	}

	park := model.Collapse()
	err = parkservice.Create(park)
	if err != nil {
		return api.InternalServerError(err)
	}

	err = slotservice.CreateMultiple(park.slots)
	if err != nil {
		return api.InternalServerError(err)
	}

	return respondWithCreatedPark(park.ID)
}

func respondWithCreatedPark(id bson.ObjectId) api.Response {
	park, err := service.Get(id)
	if err != nil {
		return api.InternalServerError(err)
	}

	model := &models.Park{}
	model.Expand(park)

	return api.JSONResponse(http.StatusCreated, model)
}

func createParkAppUser(park *models.Park) error {
	uuid, err := util.GenerateUUID()
	if err != nil {
		return err
	}

	appUser, err = auth.CreateAppUser(fmt.Sprintf("park-%s@parkme.syntaxerror2016", uuid), uuid, identity.AccountTypeNormalUser, "")
	if err != nil {
		return err
	}

	park.AppUserID = appUser.ID
}

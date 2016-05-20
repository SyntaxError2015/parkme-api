package bll

import (
	"parkme-api/api"
	"parkme-api/filter/apifilter"
	"parkme-api/orm/models"
	"parkme-api/orm/service/transactionservice"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// GetTransaction retrieves an existing Transaction based on its ID
func GetTransaction(transactionID bson.ObjectId) api.Response {
	dbTransaction, err := transactionservice.GetTransaction(transactionID)
	if err != nil || dbTransaction == nil {
		return api.NotFound(api.ErrEntityNotFound)
	}

	transaction := &models.Transaction{}
	transaction.Expand(dbTransaction)

	return api.JSONResponse(http.StatusOK, transaction)
}

// CreateTransaction creates a new Transaction
func CreateTransaction(transaction *models.Transaction) api.Response {
	if !apifilter.CheckTransactionIntegrity(transaction) {
		return api.BadRequest(api.ErrEntityIntegrity)
	}

	dbTransaction := transaction.Collapse()
	if dbTransaction == nil {
		return api.InternalServerError(api.ErrEntityProcessing)
	}

	err := transactionservice.CreateTransaction(dbTransaction)
	if err != nil {
		return api.InternalServerError(api.ErrEntityProcessing)
	}
	transaction.ID = dbTransaction.ID

	return api.JSONResponse(http.StatusCreated, transaction)
}
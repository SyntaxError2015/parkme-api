package transactionapi

import (
	"parkme-api/api"
	"parkme-api/bll"
	"parkme-api/filter"
	"parkme-api/filter/apifilter"
	"parkme-api/orm/models"
	"parkme-api/util/jsonutil"
)

// TransactionsAPI defines the API endpoint for application transactions of any kind
type TransactionsAPI int

// GetTransaction endpoint retrieves a certain transaction based on its Id
func (t *TransactionsAPI) GetTransaction(params *api.Request) api.Response {
	transactionID, found, err := filter.GetIDParameter("transactionId", params.Form)

	if err != nil {
		return api.BadRequest(err)
	}

	if !found {
		return api.NotFound(err)
	}

	return bll.GetTransaction(transactionID)
}

// CreateTransaction endpoint creates a new transaction with the valid transfer tokens and data
func (t *TransactionsAPI) CreateTransaction(params *api.Request) api.Response {
	transaction := &models.Transaction{}

	err := jsonutil.DeserializeJSON(params.Body, transaction)
	if err != nil || !apifilter.CheckTransactionIntegrity(transaction) {
		return api.BadRequest(api.ErrEntityFormat)
	}

	return bll.CreateTransaction(transaction)
}

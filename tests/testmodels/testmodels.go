package testmodels

import (
	"github.com/udonetsm/investing/models"
)

func Transaction() *models.Transaction {
	return &models.Transaction{
		Tid:           "775884",
		Type:          models.INVESTING,
		Requested_sum: 300000,
		Payer:         models.User{Uid: "998"},
		Reciever:      models.User{Uid: "8948"},
	}
}

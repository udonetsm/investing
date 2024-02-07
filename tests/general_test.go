// Testing of general functions
package tests

import (
	"testing"

	"github.com/udonetsm/investing/actions/general"
	"github.com/udonetsm/investing/actions/transactions"
	"github.com/udonetsm/investing/models"
)

func TestTransaction(t *testing.T) {
	ibill := models.Bills{Bill_id: 654321}
	sbill := models.Bills{Bill_id: 567890}
	investor := general.Payer{User: models.Users{
		User_name: "RICHMAN",
		User_id:   1234567890,
		Role:      true,
	},
		Bill: ibill,
	}
	startuper := general.Payer{User: models.Users{
		User_name: "POORMAN",
		User_id:   9987654321,
		Role:      false,
	},
		Bill: sbill,
	}
	transactions.Transaction(investor, startuper, 100000)
	transactions.Transaction(startuper, investor, 200000)
}

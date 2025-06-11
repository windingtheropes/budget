package argent

import (
	"github.com/windingtheropes/budget/types"
	"github.com/windingtheropes/budget/based"
)

var TransactionTable = based.NewTable[types.TransactionEntry, types.TransactionEntryForm]("transaction_entry")
var TransactionTypeTable = based.NewTable[types.TransactionType, types.TransactionTypeForm]("transaction_type")

func UserOwnsTransaction(user_id int64, transaction_id int64) bool {
	transactions, err := TransactionTable.Get("(id=?)", transaction_id)
	if err != nil {
		return false
	}
	if len(transactions) == 0 {
		return false
	}
	if transactions[0].User_Id == user_id { 
		return true
	}
	return false
}

func TransactionExists(transaction_id int64) bool {
	transactions, err := TransactionTable.Get("(id=?)", transaction_id)
	if err != nil {
		return false
	}
	if len(transactions) == 0 {
		return false
	} else {
		return true
	}
}
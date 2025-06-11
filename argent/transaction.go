package argent

import "github.com/windingtheropes/budget/tables"

func UserOwnsTransaction(user_id int64, transaction_id int64) bool {
	transactions, err := tables.Transaction.Get("(id=?)", transaction_id)
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
	transactions, err := tables.Transaction.Get("(id=?)", transaction_id)
	if err != nil {
		return false
	}
	if len(transactions) == 0 {
		return false
	} else {
		return true
	}
}
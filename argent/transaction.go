package argent
import ()

func UserOwnsTransaction(user_id int, transaction_id int) bool {
	transactions, err := GetTransactionById(transaction_id)
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

func TransactionExists(transaction_id int) bool {
	transactions, err := GetTransactionById(transaction_id)
	if err != nil {
		return false
	}
	if len(transactions) == 0 {
		return false
	} else {
		return true
	}
}
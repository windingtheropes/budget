package argent

import (
	"github.com/windingtheropes/budget/types"
	// "github.com/windingtheropes/budget/json"
	// "fmt"
)

func HydrateTransactionsWithTags(transactions []types.TransactionEntry) ([]types.HydTransactionEntry, error) {
	var hydratedTransactions []types.HydTransactionEntry
	for i := 0; i < len(transactions); i++ {
		dbTrans := transactions[i]
		tags, err := GetTransactionTags(dbTrans.Id)
		if err != nil {
			return make([]types.HydTransactionEntry, 0), err
		}
		hydratedTransactions = append(hydratedTransactions, types.HydTransactionEntry{
			Id:             dbTrans.Id,
			User_Id:        dbTrans.User_Id,
			Amount:         dbTrans.Amount,
			Currency:       dbTrans.Currency,
			Tags:           tags,
			Unix_Timestamp: dbTrans.Unix_Timestamp,
		})
	}
	return hydratedTransactions, nil
}

func AddTagsById(transaction_id int64, tag_ids []int) error {
	// var assignment_ids []int64;
	for i := 0; i < len(tag_ids); i++ {
		_, err := NewTagAssignment(tag_ids[i], transaction_id)
		if err != nil {
			return err
		}
		// assignment_ids = append(assignment_ids, id);
	}
	return nil
}
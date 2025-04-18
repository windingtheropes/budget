package argent

import (
	"github.com/windingtheropes/budget/types"
	// "github.com/windingtheropes/budget/json"
	// "fmt"
)

func HydrateTransactionsWithTags(transactions []types.TransactionEntry) ([]types.HydTransactionEntry, error) {
	var hydratedTransactions []types.HydTransactionEntry
	for i := 0; i < len(transactions); i++ {
		d_trans := transactions[i]
		tags, err := GetTransactionTags(d_trans.Id)
		if err != nil {
			return make([]types.HydTransactionEntry, 0), err
		}
		hydratedTransactions = append(hydratedTransactions, types.HydTransactionEntry{
			Id:             d_trans.Id,
			User_Id:        d_trans.User_Id,
			Msg: 			d_trans.Msg,
			Amount:         d_trans.Amount,
			Currency:       d_trans.Currency,
			Tags:           tags,
			Unix_Timestamp: d_trans.Unix_Timestamp,
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
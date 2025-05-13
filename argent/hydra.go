package argent

import (
	"github.com/windingtheropes/budget/types"
)

func HydrateTransactionsWithTags(transactions []types.TransactionEntry) ([]types.HydTransactionEntry, error) {
	var hydratedTransactions []types.HydTransactionEntry
	for i := range transactions {
		d_trans := transactions[i]
		tags, err := GetTransactionTags(d_trans.Id)
		if err != nil {
			return make([]types.HydTransactionEntry, 0), err
		}
		hydratedTransactions = append(hydratedTransactions, types.HydTransactionEntry{
			Id:             d_trans.Id,
			User_Id:        d_trans.User_Id,
			Type_Id:        d_trans.Type_Id,
			Msg:            d_trans.Msg,
			Amount:         d_trans.Amount,
			Currency:       d_trans.Currency,
			Tags:           tags,
			Unix_Timestamp: d_trans.Unix_Timestamp,
			Vendor:         d_trans.Vendor,
		})
	}
	return hydratedTransactions, nil
}

func AddTagsById(transaction_id int64, tag_ids []int) error {
	// var assignment_ids []int64;
	for i := range tag_ids {
		_, err := NewTagAssignment(tag_ids[i], transaction_id)
		if err != nil {
			return err
		}
	}
	return nil
}

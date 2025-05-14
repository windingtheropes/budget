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
			return nil, err
		}
		budget_entries, err := GetBudgetEntries(d_trans.Id)
		if err != nil {
			return nil, err
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
			Budget_Entries: budget_entries,
		})
	}
	return hydratedTransactions, nil
}

func HydrateTagsWithTagBudgets(tags []types.Tag) ([]types.HydTag, error) {
	var hydratedTags []types.HydTag
	for i := range tags {
		tag := tags[i]
		tagBudgets, err := GetTagBudget(tag.Id)
		if err != nil {
			return nil, err
		}
		hydratedTags = append(hydratedTags, types.HydTag{
			Id:      tag.Id,
			Name:    tag.Name,
			Tag_Budgets: tagBudgets,
		})
	}
	return hydratedTags, nil
}

func AddTagsById(transaction_id int64, tag_ids []int64) error {
	// var assignment_ids []int64;
	for i := range tag_ids {
		_, err := NewTagAssignment(tag_ids[i], transaction_id)
		if err != nil {
			return err
		}
	}
	return nil
}

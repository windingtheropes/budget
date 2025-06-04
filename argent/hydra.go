package argent

import (
	"github.com/windingtheropes/budget/types"
)

func HydrateTransactions(transactions []types.TransactionEntry) ([]types.HydTransactionEntry, error) {
	var hydratedTransactions []types.HydTransactionEntry
	for i := range transactions {
		transaction := transactions[i]
		tags, err := GetTransactionTags(transaction.Id)
		if err != nil {
			return nil, err
		}
		budget_entries, err := GetBudgetEntries(transaction.Id)
		if err != nil {
			return nil, err
		}
		hydratedTransactions = append(hydratedTransactions, types.HydTransactionEntry{
			Id:             transaction.Id,
			User_Id:        transaction.User_Id,
			Type_Id:        transaction.Type_Id,
			Msg:            transaction.Msg,
			Amount:         transaction.Amount,
			Currency:       transaction.Currency,
			Tags:           tags,
			Unix_Timestamp: transaction.Unix_Timestamp,
			Vendor:         transaction.Vendor,
			Budget_Entries: budget_entries,
		})
	}
	return hydratedTransactions, nil
}

func HydrateTagsWithTagBudgets(tags []types.Tag) ([]types.HydTag, error) {
	var hydratedTags []types.HydTag
	for i := range tags {
		tag := tags[i]
		tagBudgets, err := GetTagBudgetsByTagId(tag.Id)
		if err != nil {
			return nil, err
		}
		hydratedTags = append(hydratedTags, types.HydTag{
			Id:          tag.Id,
			Name:        tag.Name,
			Tag_Budgets: tagBudgets,
		})
	}
	return hydratedTags, nil
}

func HydrateBudgetsWithTagBudgets(budgets []types.Budget) ([]types.HydBudget, error) {
	var hydratedBudgets []types.HydBudget
	for i := range budgets {
		budget := budgets[i]
		tagBudgets, err := GetTagBudgetsByBudgetId(budget.Id)
		if err != nil {
			return nil, err
		}
		hydratedBudgets = append(hydratedBudgets, types.HydBudget{
			Id:          budget.Id,
			Name:        budget.Name,
			Goal: 		 budget.Goal,
			Tag_Budgets: tagBudgets,
		})
	}
	return hydratedBudgets, nil
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

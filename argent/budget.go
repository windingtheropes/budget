package argent

import (
	"log"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

var TagBudgetTable = based.NewTable[types.TagBudget, types.TagBudgetForm]("tag_budget")
var BudgetTable = based.NewTable[types.Budget, types.BudgetForm]("budget")
var BudgetEntryTable = based.NewTable[types.BudgetEntry, types.BudgetEntryForm]("budget_entry")

func UserBudgetNameExists(budget_name string, user_id int64) bool {
	budgets, err := BudgetTable.Get("user_id=?", user_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(budgets) == 0 {
		return false
	}
	for i := range budgets {
		budget := budgets[i]
		if budget.Name == budget_name {
			return true
		}
	}
	return false
}
func BudgetExists(budget_id int64) bool {
	budgets, err := BudgetTable.Get("id=?", budget_id)
	if err != nil {
		return false
	}
	if len(budgets) == 0 {
		return false
	} else {
		return true
	}
}

func UserOwnsBudget(user_id int64, budget_id int64) bool {
	budgets, err := BudgetTable.Get("id=?", budget_id)
	if err != nil {
		return false
	}
	if len(budgets) == 0 {
		return false
	}
	if budgets[0].User_Id == user_id {
		return true
	}
	return false
}

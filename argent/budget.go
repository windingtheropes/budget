package argent

import (
	"log"
	"github.com/windingtheropes/budget/tables"
)

func UserBudgetNameExists(budget_name string, user_id int64) bool {
	budgets, err := tables.Budget.Get("user_id=?", user_id)
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
	budgets, err := tables.Budget.Get("id=?", budget_id)
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
	budgets, err := tables.Budget.Get("id=?", budget_id)
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

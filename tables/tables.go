package tables

import (
	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

var User = based.NewTable[types.User, types.UserForm]("usr")
var Session = based.NewTable[types.Session, types.SessionForm]("session")
var TagBudget = based.NewTable[types.TagBudget, types.TagBudgetForm]("tag_budget")
var Budget = based.NewTable[types.Budget, types.BudgetForm]("budget")
var BudgetEntry = based.NewTable[types.BudgetEntry, types.BudgetEntryForm]("budget_entry")
var Tag = based.NewTable[types.Tag, types.TagForm]("tag")
var TagOwnership = based.NewTable[types.TagOwnership, types.TagOwnershipForm]("tag_ownership")
var TagAssignment = based.NewTable[types.TagAssignment, types.TagAssignmentForm]("tag_assignment")
var Transaction = based.NewTable[types.TransactionEntry, types.TransactionEntryForm]("transaction_entry")
var TransactionType = based.NewTable[types.TransactionType, types.TransactionTypeForm]("transaction_type")

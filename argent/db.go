package argent

import (
	"database/sql"
	"fmt"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

// Create a new budget entry.
func NewTransaction(user_id int64, type_id int64, amount float64, currency string, msg string, unix_timestamp int64, vendor string) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO transaction_entry (user_id, type_id, msg, amount, currency, unix_timestamp, vendor) VALUES (?,?,?,?,?,?,?)", user_id, type_id, msg, amount, currency, unix_timestamp, vendor)
	if err != nil {
		return 0, fmt.Errorf("newTransaction: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTransaction: %v", err)
	}

	return id, nil
}

// Delete a budget entry
func DeleteTransaction(entry_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM transaction_entry WHERE id=?", entry_id)
	if err != nil {
		return false, fmt.Errorf("deleteTransaction: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteTransaction: %v", err)
	}

	return true, nil
}

// Create a new tag.
func NewTag(name string) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO tag (tag_name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("newTag: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTag: %v", err)
	}

	return id, nil
}

// Create a new budget
func NewBudget(name string, user_id int64, type_id int64, goal float64) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO budget (user_id, type_id, budget_name, goal) VALUES (?, ?, ?, ?)", user_id, type_id, name, goal)
	if err != nil {
		return 0, fmt.Errorf("newBudget: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newBudget: %v", err)
	}

	return id, nil
}

// Get all budgets created by a user
func GetUserBudgets(user_id int64) ([]types.Budget, error) {
	// store matching sessions in the slcie
	var budgets []types.Budget

	rows, err := based.DB().Query("SELECT * FROM budget WHERE user_id = ?", user_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getBudgets %q: user id %v", user_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var budget types.Budget
		if err := rows.Scan(&budget.Id, &budget.User_Id, &budget.Name, &budget.Goal); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getBudgets %q: user id %v", user_id, err)
		}
		budgets = append(budgets, budget)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getBudgets %q: user id %v", user_id, err)
	}
	return budgets, nil
}

// Get all budget entries  on transaction
func GetBudgetEntries(transaction_id int64) ([]types.BudgetEntry, error) {
	// store matching sessions in the slcie
	var entries []types.BudgetEntry

	rows, err := based.DB().Query("SELECT * FROM budget_entry WHERE transaction_id = ?", transaction_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getBudgetEntries %q: trans id %v", transaction_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var entry types.BudgetEntry
		if err := rows.Scan(&entry.Id, &entry.Transaction_Id, &entry.Budget_Id, &entry.Amount); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getBudgetEntries %q: trans id %v", transaction_id, err)
		}
		entries = append(entries, entry)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getBudgetEntries %q: trans id %v", transaction_id, err)
	}
	return entries, nil
}

func GetBudgetEntriesOnTransaction(transaction_id int64) ([]types.BudgetEntry, error) {
	// store matching sessions in the slcie
	var entries []types.BudgetEntry

	rows, err := based.DB().Query("SELECT * FROM budget_entry WHERE transaction_id = ?", transaction_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("GetBudgetEntriesOnTransaction %q: budget id %v", transaction_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var entry types.BudgetEntry
		if err := rows.Scan(&entry.Id, &entry.Transaction_Id, &entry.Budget_Id, &entry.Amount); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("GetBudgetEntriesOnTransaction %q: budget id %v", transaction_id, err)
		}
		entries = append(entries, entry)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBudgetEntriesOnTransaction %q: budget id %v", transaction_id, err)
	}
	return entries, nil
}

// Get budget on a tag
func GetTagBudget(tag_id int64) ([]types.TagBudget, error) {
	// store matching sessions in the slcie
	var tagBudgets []types.TagBudget

	rows, err := based.DB().Query("SELECT * FROM tag_budget WHERE tag_id = ?", tag_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTagBudget %q: tag id %v", tag_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var tagBudget types.TagBudget
		if err := rows.Scan(&tagBudget.Id, &tagBudget.Tag_Id, &tagBudget.Budget_Id, &tagBudget.Goal, &tagBudget.Type_Id); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTagBudget %q: tag id %v", tag_id, err)
		}
		tagBudgets = append(tagBudgets, tagBudget)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTagBudget %q: tag id %v", tag_id, err)
	}
	return tagBudgets, nil
}

// Create a new budget entry from a transaction
func NewBudgetEntry(transaction_id int64, budget_id int64, amount float64) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO budget_entry (transaction_id, budget_id, amount) VALUES (?, ?, ?)", transaction_id, budget_id, amount)
	if err != nil {
		return 0, fmt.Errorf("newBudgetEntry: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newBudgetEntry: %v", err)
	}

	return id, nil
}

// Create a budget on a tag
func NewTagBudget(tag_id int64, budget_id int64, goal float64, type_id int64) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO tag_budget (tag_id, budget_id, goal, type_id) VALUES (?, ?, ?, ?)", tag_id, budget_id, goal, type_id)
	if err != nil {
		return 0, fmt.Errorf("newTagBudget: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTagBudget: %v", err)
	}

	return id, nil
}

// Create a new ownership record of a tag
func NewTagOwnership(tag_id int64, user_id int64) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO tag_ownership (tag_id, user_id) VALUES (?,?)", tag_id, user_id)
	if err != nil {
		return 0, fmt.Errorf("newTagownership: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTagownership: %v", err)
	}

	return id, nil
}

// Create a new tag assignment; add a tag to an entry.
func NewTagAssignment(tag_id int64, entry_id int64) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO tag_assignment (tag_id, entry_id) VALUES (?,?)", tag_id, entry_id)
	if err != nil {
		return 0, fmt.Errorf("newTagAssignment: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTagAssignment: %v", err)
	}

	return id, nil
}

// Get a list of transaction types
func GetTransactionTypes() ([]types.TransactionType, error) {
	// store matching sessions in the slcie
	var transaction_types []types.TransactionType

	rows, err := based.DB().Query("SELECT * FROM transaction_type")
	// Catch error with query
	if err != nil {
		return nil, fmt.Errorf("getTransactionTypes %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction_type types.TransactionType
		if err := rows.Scan(&transaction_type.Id, &transaction_type.Name, &transaction_type.Positive); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransactionTypes %v", err)
		}
		transaction_types = append(transaction_types, transaction_type)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransactionTypes %v", err)
	}
	return transaction_types, nil
}

// Get all budget entries by a user identified by user_id.
func GetUserTransactions(user_id int64) ([]types.TransactionEntry, error) {
	// store matching sessions in the slcie
	var transactions []types.TransactionEntry

	rows, err := based.DB().Query("SELECT * FROM transaction_entry WHERE user_id = ?", user_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTransaction %q: user id %v", user_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction types.TransactionEntry
		if err := rows.Scan(&transaction.Id, &transaction.User_Id, &transaction.Type_Id, &transaction.Msg, &transaction.Amount, &transaction.Currency, &transaction.Unix_Timestamp, &transaction.Vendor); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransaction %q: user id %v", user_id, err)
		}
		transactions = append(transactions, transaction)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransaction %q: user id %v", user_id, err)
	}
	return transactions, nil
}

// Get budget entry identified by its entry id
func GetTransactionById(entry_id int64) ([]types.TransactionEntry, error) {
	// store matching sessions in the slcie
	var transactions []types.TransactionEntry

	rows, err := based.DB().Query("SELECT * FROM transaction_entry WHERE id = ?", entry_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTransaction %q: entry id %v", entry_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction types.TransactionEntry
		if err := rows.Scan(&transaction.Id, &transaction.User_Id, &transaction.Type_Id, &transaction.Msg, &transaction.Amount, &transaction.Currency, &transaction.Unix_Timestamp, &transaction.Vendor); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransaction %q: entry id %v", entry_id, err)
		}
		transactions = append(transactions, transaction)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransaction %q: entry id %v", entry_id, err)
	}
	return transactions, nil
}

func GetUserTagOwnerships(user_id int64) ([]types.TagOwnership, error) {
	rows, err := based.DB().Query("SELECT * FROM tag_ownership WHERE user_id = ?", user_id)
	var ownership_records []types.TagOwnership
	if err != nil {
		return nil, fmt.Errorf("get tag, get ownership %q: %v", user_id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var tag_ownership types.TagOwnership
		if err := rows.Scan(&tag_ownership.Id, &tag_ownership.Tag_Id, &tag_ownership.User_Id); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTag ownership %q: %v", user_id, err)
		}
		ownership_records = append(ownership_records, tag_ownership)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get tag ownership %q: %v", user_id, err)
	}
	return ownership_records, nil
}
// Get a budget from the database.
func GetBudgetById(budget_id int64) ([]types.Budget, error) {
	// store matching sessions in the slice
	var budgets []types.Budget

	var rows *sql.Rows
	var err error

	rows, err = based.DB().Query("SELECT * FROM budget WHERE id = ?", budget_id)

	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getbudget %q: %v", budget_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var budget types.Budget
		if err := rows.Scan(&budget.Id, &budget.User_Id, &budget.Name, &budget.Goal); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getbudget %q: %v", budget_id, err)
		}
		budgets = append(budgets, budget)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getbudget %q: %v", budget_id, err)
	}
	return budgets, nil
}
// Get a tag from the database. Identifier is of type TagIdentifier, which can be a TagID or UserID.
func GetTagById(tag_id int64) ([]types.Tag, error) {
	// store matching sessions in the slice
	var tags []types.Tag

	var rows *sql.Rows
	var err error

	rows, err = based.DB().Query("SELECT * FROM tag WHERE id = ?", tag_id)

	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTag %q: %v", tag_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var tag types.Tag
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTag %q: %v", tag_id, err)
		}
		tags = append(tags, tag)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTag %q: %v", tag_id, err)
	}
	return tags, nil
}

// Get tag assignments to a specified entry_id
func GetTagAssignments(entry_id int64) ([]types.TagAssignment, error) {
	// store matching sessions in the slice
	var assignments []types.TagAssignment

	rows, err := based.DB().Query("SELECT * FROM tag_assignment WHERE entry_id = ?", entry_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTransactionTags %q: %v", entry_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var assignment types.TagAssignment
		if err := rows.Scan(&assignment.Id, &assignment.Tag_Id, &assignment.Transaction_Id); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransactionTags %q: %v", entry_id, err)
		}
		assignments = append(assignments, assignment)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransactionTags %q: %v", entry_id, err)
	}

	return assignments, nil
}

// Delete a budget entry
func DeleteTagOwnership(tag_id int64, entry_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM tag_ownership WHERE (tag_id=? AND user_id=?)", tag_id, entry_id)
	if err != nil {
		return false, fmt.Errorf("deleteTagonentry: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteTagonentry: %v", err)
	}

	return true, nil
}
func DeleteTag(tag_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM tag WHERE (id=?)", tag_id)
	if err != nil {
		return false, fmt.Errorf("deletetag: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deletetag: %v", err)
	}

	return true, nil
}

func DeleteTagOnEntry(tag_id int64, entry_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM tag_assignment WHERE (tag_id=? AND entry_id=?)", tag_id, entry_id)
	if err != nil {
		return false, fmt.Errorf("deleteTagonentry: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteTagonentry: %v", err)
	}

	return true, nil
}

// Delete a budget entry
func DeleteBudgetEntry(entry_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM budget_entry WHERE (id=?)", entry_id)
	if err != nil {
		return false, fmt.Errorf("deleteBudgetEntry: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteBudgetEntry: %v", err)
	}

	return true, nil
}

// Delete a budget
func DeleteBudget(budget_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM budget WHERE (id=?)", budget_id)
	if err != nil {
		return false, fmt.Errorf("deleteBudget: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteBudget: %v", err)
	}

	return true, nil
}

// Delete a tag budget
func DeleteTagBudget(tb_id int64) (bool, error) {
	result, err := based.DB().Exec("DELETE FROM tag_budget WHERE (id=?)", tb_id)
	if err != nil {
		return false, fmt.Errorf("deleteTagBudget: %v", err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("deleteTagBudget: %v", err)
	}

	return true, nil
}
package argent

import (
	"fmt"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

// Insert a new entry
func NewEntry(user_id int, amount float64, currency string) (int64, error) {
    result, err := based.DB().Exec("INSERT INTO budget_entry (user_id, amount, currency) VALUES (?,?,?)", user_id, amount, currency)
	if err != nil {
        return 0,fmt.Errorf("newEntry: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("newEntry: %v", err)
    }

    return id, nil
}
// Get all entries by a user
func GetEntries(user_id int) ([]types.BudgetEntry, error) {
    // store matching sessions in the slcie
    var entries []types.BudgetEntry

    rows, err := based.DB().Query("SELECT * FROM budget_entry WHERE user_id = ?", user_id)
	// Catch error with query
    if err != nil {
		// token is sensitive
        return nil, fmt.Errorf("getEntries %q: %v", user_id, err)
    }
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var entry types.BudgetEntry
        if err := rows.Scan(&entry.Id, &entry.User_Id, &entry.Amount, &entry.Currency); err != nil {
			// Catch error casting to struct
            return nil, fmt.Errorf("getEntries %q: %v", user_id, err)
        }
        entries = append(entries, entry)
    }
	// Catch a row error
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getEntries %q: %v", user_id, err)
    }
    return entries, nil
}
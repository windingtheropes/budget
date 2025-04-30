package argent

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

// Create a new budget entry.
func NewTransaction(user_id int, type_id int, amount float64, currency string, msg string, unix_timestamp int, vendor string) (int64, error) {
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

// Create a new tag.
func NewTag(user_id int, name string) (int64, error) {
	result, err := based.DB().Exec("INSERT INTO tag (tag_name, user_id) VALUES (?,?)", name, user_id)
	if err != nil {
		return 0, fmt.Errorf("newTag: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newTag: %v", err)
	}

	return id, nil
}

// Create a new tag assignment; add a tag to an entry.
func NewTagAssignment(tag_id int, entry_id int64) (int64, error) {
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
		if err := rows.Scan(&transaction_type.Id, &transaction_type.Name); err != nil {
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
func GetTransactions(user_id int) ([]types.TransactionEntry, error) {
	// store matching sessions in the slcie
	var transactions []types.TransactionEntry

	rows, err := based.DB().Query("SELECT * FROM transaction_entry WHERE user_id = ?", user_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTransaction %q: %v", user_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var transaction types.TransactionEntry
		if err := rows.Scan(&transaction.Id, &transaction.User_Id, &transaction.Type_Id, &transaction.Msg, &transaction.Amount, &transaction.Currency, &transaction.Unix_Timestamp, &transaction.Vendor); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransaction %q: %v", user_id, err)
		}
		transactions = append(transactions, transaction)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransaction %q: %v", user_id, err)
	}
	return transactions, nil
}

// Get a tag from the database. Identifier is of type TagIdentifier, which can be a TagID or UserID.
func GetTag[T types.TagIdentifier](identifier T) ([]types.Tag, error) {
	// store matching sessions in the slice
	var tags []types.Tag

	var rows *sql.Rows
	var err error

	if reflect.TypeOf(identifier) == reflect.TypeOf(types.TagID(0)) {
		rows, err = based.DB().Query("SELECT * FROM tag WHERE id = ?", identifier)
	} else if reflect.TypeOf(identifier) == reflect.TypeOf(types.UserID(0)) {
		rows, err = based.DB().Query("SELECT * FROM tag WHERE user_id = ?", identifier)
	}

	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTag %q: %v", identifier, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var tag types.Tag
		if err := rows.Scan(&tag.Id, &tag.Name, &tag.User_Id); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTag %q: %v", identifier, err)
		}
		tags = append(tags, tag)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTag %q: %v", identifier, err)
	}
	return tags, nil
}

// Get all tags on a budget entry by its entry_id.
func GetTransactionTags(transaction_id int) ([]types.Tag, error) {
	// store matching sessions in the slice
	var assignments []types.TagAssignment

	rows, err := based.DB().Query("SELECT * FROM tag_assignment WHERE entry_id = ?", transaction_id)
	// Catch error with query
	if err != nil {
		// token is sensitive
		return nil, fmt.Errorf("getTransactionTags %q: %v", transaction_id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var assignment types.TagAssignment
		if err := rows.Scan(&assignment.Id, &assignment.Tag_Id, &assignment.Entry_Id); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("getTransactionTags %q: %v", transaction_id, err)
		}
		assignments = append(assignments, assignment)
	}
	// Catch a row error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTransactionTags %q: %v", transaction_id, err)
	}

	if len(assignments) == 0 {
		return make([]types.Tag, 0), nil
	}

	var tags []types.Tag

	for i := 0; i < len(assignments); i++ {
		tag, err := GetTag(types.TagID(assignments[i].Tag_Id))
		if err != nil {
			return nil, err
		}
		if len(tag) == 0 {
			continue
		}
		tags = append(tags, tag[0])
	}
	return tags, nil
}

package argent

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

// Create a new budget entry.
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
// Create a new tag.
func NewTag(user_id int, name string) (int64, error) {
    result, err := based.DB().Exec("INSERT INTO tag (name, user_id) VALUES (?,?)", user_id, name, user_id)
	if err != nil {
        return 0,fmt.Errorf("newTag: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("newTag: %v", err)
    }

    return id, nil
}
// Create a new tag assignment; add a tag to an entry.
func NewTagAssignment(tag_id int, entry_id string) (int64, error) {
    result, err := based.DB().Exec("INSERT INTO tag_assignment (tag_id, entry_id) VALUES (?,?)", tag_id, entry_id)
	if err != nil {
        return 0,fmt.Errorf("newTagAssignment: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("newTagAssignment: %v", err)
    }

    return id, nil
}

// Get all budget entries by a user identified by user_id.
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
func GetEntryTags(entry_id int) ([]types.Tag, error) {
    // store matching sessions in the slice
    var assignments []types.TagAssignment

    rows, err := based.DB().Query("SELECT * FROM tag_assignment WHERE entry_id = ?", entry_id)
	// Catch error with query
    if err != nil {
		// token is sensitive
        return nil, fmt.Errorf("getEntryTags %q: %v", entry_id, err)
    }
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var assignment types.TagAssignment
        if err := rows.Scan(&assignment.Id, &assignment.Tag_Id, &assignment.Entry_Id); err != nil {
			// Catch error casting to struct
            return nil, fmt.Errorf("getEntryTags %q: %v", entry_id, err)
        }
        assignments = append(assignments, assignment)
    }
	// Catch a row error
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getEntryTags %q: %v", entry_id, err)
    }
    
    if len(assignments) == 0 { 
        return make([]types.Tag, 0), nil 
    }

    var tags []types.Tag;

    for i := 0; i < len(assignments); i++ {
        tag, err := GetTag(types.TagID(assignments[i].Tag_Id))
        if err != nil { return nil, err }
        if len(tag) == 0 { continue }
        tags = append(tags, tag[0])
    }
    return tags, nil
}
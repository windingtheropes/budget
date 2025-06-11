package based

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Table[Data any, Form any] struct {
	name              string // Table name in the database
	insertion_snippet string // Snippet for inserting to table, generated from the fields of Form
	update_snippet   string // Snippet for updating row in table, generated from fields of Form
}

// Get an array of Data rows from the table by a sql_condition. 
// Will substitute `?` in order that args are passed.
// 
// ex. Get("user_id=?",user_id)
func (t *Table[Data, Form]) Get(sql_condition string, args... any) ([]Data, error) {
	var query string

	// Allow for a wildcard shortcut
	if(sql_condition == "*") {
		query = fmt.Sprintf("SELECT * FROM %v", t.name)
	} else {
		query = fmt.Sprintf("SELECT * FROM %v WHERE %v", t.name, sql_condition)
	}
	rows, err := DB().Query(query, args...)
	var data_rows []Data

	// Catch error with query
	if err != nil {
		return nil, fmt.Errorf("Get %v: %v", t.name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var data_row Data

		data_row_value := reflect.ValueOf(&data_row).Elem()
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		// Array of pointers, which will be all pointers contained in typed_row
		pointers := make([]interface{}, len(columns))

		// prioritize the length of columns, must be able to accomodate exact length
		if len(columns) != data_row_value.NumField() {
			return nil, fmt.Errorf("length of struct does not match that of column")
		}
		for i := range len(columns) {
			field := data_row_value.Field(i)
			if !field.CanAddr() {
				return nil, fmt.Errorf("cannot address field")
			}
			// Add the pointer in place to the pointer array
			pointers[i] = field.Addr().Interface()
		}

		// Destruture pointers to be scanned into in the order they're defined as a struct
		if err := rows.Scan(pointers...); err != nil {
			// Catch error casting to struct
			return nil, fmt.Errorf("Get %v: %v", t.name, err)
		}
		data_rows = append(data_rows, data_row)
	}
	return data_rows, nil
}
// Add a row from the table given a Form struct
func (t *Table[Data, Form]) New(row Form) (int64, error) {
	query_ph := fmt.Sprintf("INSERT INTO %v %v", t.name, t.insertion_snippet)
	form_vals := getFieldValues(row);

	result, err := DB().Exec(query_ph, form_vals...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
// Update a row in the table given a Form struct and an sql_condition to identify it.
// Returns number of rows affected.
// 
// ex. Update(UserForm{userdata}, "id = ?", usr.id)
func (t *Table[Data, Form]) Update(row Form, sql_condition string, args... any) (int64, error) {
	query := fmt.Sprintf(`UPDATE %v SET %v WHERE %v`, t.name, t.update_snippet, sql_condition)
	form_vals := getFieldValues(row);
	
	// The number of columns is a known constant for any Table, so the number of replacement args (form_vals) will be the same
	// We can then safely concatenate the unknown length args to the end of this
	concat_args := append(form_vals, args...)

	result, err := DB().Exec(query, concat_args...)
	if err != nil {
		return 0, err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return ra, nil
}

// Delete a row from the table by a sql_condition. Will substitute `?` in order that args are passed.
// 
// ex. Delete("id=?",entry_id)
func (t *Table[Data, Form]) Delete(sql_condition string, args... any) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %v WHERE %v", t.name, sql_condition)
	result, err := DB().Exec(query, args...)
	if err != nil {
		return false, fmt.Errorf("delete %v: %v", t.name, err)
	}

	if _, err := result.LastInsertId(); err != nil {
		return false, fmt.Errorf("delete %v: %v", t.name, err)
	}

	return true, nil
}

func getDbTags[T any](i T) []string {
	interf_type := reflect.TypeOf(&i).Elem()
	var db_vals []string;
	for i := range interf_type.NumField() {
		field := interf_type.Field(i)
		db_val := field.Tag.Get("db")
		if db_val == "" {
			log.Fatal("no db tag provided on struct")
		}
		db_vals = append(db_vals, db_val)
	}
	return db_vals
}
func getFieldValues[T any](i T) []interface{} {
	row_value := reflect.ValueOf(&i).Elem()
	form_vals := make([]interface{}, row_value.NumField())
	for i := range form_vals {
		val := row_value.Field(i).Interface()
		form_vals[i] = val
	}
	return form_vals
}

// Create a new Table representing table `name` in database
// 
// Data is the container struct for rows, 
// the amount of fields must match the amount of columns
//
// Form is the struct for creating new rows, should not include the primary key
// and must be annotated with `db:"{column_name}"` tags
func NewTable[Data any, Form any](name string) Table[Data, Form] {
	var f Form;

	// Insertion snippet (from Form) -> (name) VALUES (?)
	form_columns := getDbTags(f)
	insertion_snippet := fmt.Sprintf("(%v) VALUES (%v)", strings.Join(form_columns, ","), strings.Join(strings.Split(strings.Repeat("?",len(form_columns)),""), ","))
	
	// Update snippet (from Data) -> name=?,something=?
	update_snippet := strings.Join(form_columns, "=?,") + "=?"
	
	return Table[Data, Form]{
		name,
		insertion_snippet,
		update_snippet,
	}
}
package based

import (
	"fmt"
	"reflect"
	"strings"
)

type Table[Data any, Form any] struct {
	name              string
	insertion_snippet string
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

	row_value := reflect.ValueOf(&row).Elem()

	form_vals := make([]interface{}, row_value.NumField())
	for i := range form_vals {
		val := row_value.Field(i).Interface()
		form_vals[i] = val
	}
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

// Create a new Table
// 
// Data is the container struct for rows, 
// the amount of fields must match the amount of columns
//
// Form is the struct for creating new rows, should not include the primary key.
// Form must be annotated with `db:"{column_name}"` tags
// 
// `name` is the name of the table in the database
func NewTable[Data any, Form any](name string) Table[Data, Form] {
	var f Form

	// Generate the insertion snippet: `(name, password) VALUES (?,?)`
	// Reflect Form as a type
	var column_names []string
	form_type := reflect.TypeOf(&f).Elem()
	for i := range form_type.NumField() {
		field := form_type.Field(i)
		column_name := field.Tag.Get("db")
		if column_name == "" {
			panic(fmt.Errorf("no db tag provided on table form struct"))
		}
		column_names = append(column_names, column_name)
	}
	val_ph := strings.Split(strings.Repeat("?", form_type.NumField()), "")
	insertion_snippet := fmt.Sprintf("(%v) VALUES (%v)", strings.Join(column_names, ","), strings.Join(val_ph, ","))

	return Table[Data, Form]{
		name,
		insertion_snippet,
	}
}
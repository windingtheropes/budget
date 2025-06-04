package auth

import (
	"fmt"
	"reflect"
	"time"

	"database/sql"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

// Get a user by a UserIdentifier, UserID | Email
func GetUser[T types.UserIdentifier](identifier T) ([]types.User, error) {
    // A users slice to store criteria matching users
    var users []types.User

    var rows *sql.Rows;
    var err error;
    
    // Change query based on the type of k, string is email, int is id
    if reflect.TypeOf(identifier) == reflect.TypeOf(types.UserID(0)) {
        rows, err = based.DB().Query("SELECT * FROM usr WHERE id = ?", identifier)
    } else if reflect.TypeOf(identifier) == reflect.TypeOf(types.Email("")) {
        rows, err = based.DB().Query("SELECT * FROM usr WHERE email = ?", identifier)
    }

	// Catch error with query
    if err != nil {
        return nil, fmt.Errorf("getUser %q: %v", identifier, err)
    }
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var usr types.User
        if err := rows.Scan(&usr.Id, &usr.First_Name, &usr.Last_Name, &usr.Email, &usr.Password); err != nil {
			// Catch error casting to struct
            return nil, fmt.Errorf("getUser %q: %v", identifier, err)
        }
        users = append(users, usr)
    }
	// Catch a row error
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getUser %q: %v", identifier, err)
    }
    return users, nil
}

// Add a user, returning its id
func AddUser(first_name string, last_name string, email string, pass_hashed string) (int64, error) {
    result, err := based.DB().Exec("INSERT INTO usr (first_name, last_name, email, pass) VALUES (?,?,?,?)", first_name, last_name, email, pass_hashed)
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    return id, nil
}

// New session given a user_id, and a lifetime in seconds
func NewSession(user_id int64, lifetime int64) ([]string, int64, error) {
	var token string = GenToken(64);
    expiry := time.Now().Unix() + lifetime
    result, err := based.DB().Exec("INSERT INTO session (token, user_id, expiry) VALUES (?,?,?)", token, user_id, expiry)
	if err != nil {
        return nil, 0, fmt.Errorf("newSession: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return nil, 0, fmt.Errorf("newSession: %v", err)
    }
	
	var res []string;
	res = append(res, token);

    return res, id, nil
}
// Get a session in the database by the oken
func GetSession(token string) ([]types.Session, error) {
    // store matching sessions in the slcie
    var sessions []types.Session

    rows, err := based.DB().Query("SELECT * FROM session WHERE token = ?", token)
	// Catch error with query
    if err != nil {
		// token is sensitive
        return nil, fmt.Errorf("getSession %q: %v", "TOKEN", err)
    }
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var session types.Session
        if err := rows.Scan(&session.Id, &session.Token, &session.User_Id, &session.Expiry); err != nil {
			// Catch error casting to struct
            return nil, fmt.Errorf("getSession %q: %v", "TOKEN", err)
        }
        sessions = append(sessions, session)
    }
	// Catch a row error
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getSession %q: %v", "TOKEN", err)
    }
    return sessions, nil
}
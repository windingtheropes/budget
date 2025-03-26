package auth

import (
	"fmt"

	"github.com/windingtheropes/budget/based"
    "github.com/windingtheropes/budget/types"
)

func GetUser(email string) ([]types.User, error) {
    // A users slice to store criteria matching users
    var users []types.User

    rows, err := based.DB().Query("SELECT * FROM usr WHERE email = ?", email)
	// Catch error with query
    if err != nil {
        return nil, fmt.Errorf("getUser %q: %v", email, err)
    }
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var usr types.User
        if err := rows.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Password); err != nil {
			// Catch error casting to struct
            return nil, fmt.Errorf("getUser %q: %v", email, err)
        }
        users = append(users, usr)
    }
	// Catch a row error
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getUser %q: %v", email, err)
    }
    return users, nil
}

// add a user and return its id
func AddUser(full_name string, email string, pass_hashed string) (int64, error) {
    result, err := based.DB().Exec("INSERT INTO usr (full_name, email, pass) VALUES (?,?,?)", full_name, email, pass_hashed)
	if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    return id, nil
}

// new session, returns the token, id
func NewSession(user_id int) ([]string, int64, error) {
	var token string = GenToken(64);

    result, err := based.DB().Exec("INSERT INTO session (token, user_id) VALUES (?,?)", token, user_id)
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

// get a session
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
        if err := rows.Scan(&session.Id, &session.Token, &session.User_Id); err != nil {
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
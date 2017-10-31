// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package model

import "database/sql"

func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.ID,
		&s.Nickname,
		&s.Email,
		&s.Password,
		&s.Created,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

func ScanUsers(rs *sql.Rows) ([]User, error) {
	structs := make([]User, 0, 16)
	var err error
	for rs.Next() {
		var s User
		if err = rs.Scan(
			&s.ID,
			&s.Nickname,
			&s.Email,
			&s.Password,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPost(r *sql.Row) (Post, error) {
	var s Post
	if err := r.Scan(
		&s.ID,
		&s.Description,
		&s.UserId,
		&s.Created,
	); err != nil {
		return Post{}, err
	}
	return s, nil
}

func ScanPosts(rs *sql.Rows) ([]Post, error) {
	structs := make([]Post, 0, 16)
	var err error
	for rs.Next() {
		var s Post
		if err = rs.Scan(
			&s.ID,
			&s.Description,
			&s.UserId,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanNotification(r *sql.Row) (Notification, error) {
	var s Notification
	if err := r.Scan(
		&s.ID,
		&s.Time,
		&s.Set,
		&s.Created,
	); err != nil {
		return Notification{}, err
	}
	return s, nil
}

func ScanNotifications(rs *sql.Rows) ([]Notification, error) {
	structs := make([]Notification, 0, 16)
	var err error
	for rs.Next() {
		var s Notification
		if err = rs.Scan(
			&s.ID,
			&s.Time,
			&s.Set,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}


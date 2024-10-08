// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Event struct {
	EventID    int64
	EventType  int64
	EventLevel int64
}

type Score struct {
	ScoreID     int64
	UserID      int64
	EventID     int64
	Timems      int64
	DateEntered sql.NullTime
}

type User struct {
	UserID   int64
	Username string
	ApiKey   string
	Country  sql.NullString
}

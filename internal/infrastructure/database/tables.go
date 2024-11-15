package database

import (
	"fmt"
)

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

// schemas
const (
	// Schema
	SchemaUser         = "users"
	SchemaAuth         = "auth"
	SchemaNotification = "notification"
)

var (
	//User Schema Tables
	UserTableUser = format(SchemaUser, "users")
)

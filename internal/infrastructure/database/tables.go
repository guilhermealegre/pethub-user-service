package database

import (
	"fmt"
)

func format(schema, table string) string {
	return fmt.Sprintf("\"%s\".%s", schema, table)
}

var (
	//schema
	SchemaUser = "user"
	// tables
	UserTableUser = format(SchemaUser, "users")
)

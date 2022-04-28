package main

import (
	"database/sql"
	"fmt"
	"strings"
)

// DoInList demonstrates using bind variables with an in-list.
func DoInList(conn *sql.DB) error {

	fmt.Println("Using SQL with bind variables to find a variable number of accounts with an in-list.")

	argList := []any{1000, 1001, 1002}

	inListSql := generateInListSql(len(argList))

	rowCount, err := OpenCursorAndProcess(conn, inListSql, argList...)
	if err != nil {
		return err
	}

	if rowCount == len(argList) {
		fmt.Printf("Expected and read %d rows.\n\n", len(argList))
	} else {
		fmt.Printf("*** Incorrectly read %d rows when %d were expected.\n\n", rowCount, len(argList))
	}

	return nil
}

// generateInListSql constructs the SQL that is used to prepare the statement. Though it is dynamically
// generating SQL, it is not generating SQL with run-time parameters in the SQL itself. It is generating
// SQL with enough bind variables to handle a number of parameters defined at run-time rather than compile-time.
func generateInListSql(variableCount int) string {

	var sqlBuilder strings.Builder

	sqlBuilder.WriteString("select ACCOUNT_ID, ACCOUNT_NAME, ACCOUNT_BALANCE::money::numeric " +
		" from BIND_VARIABLES_EXAMPLE_SCHEMA.ACCOUNTS where ACCOUNT_ID in (")

	for i := 0; i < variableCount; i++ {
		if i != variableCount-1 {
			sqlBuilder.WriteString(fmt.Sprintf("$%d, ", i+1))
		} else {
			sqlBuilder.WriteString(fmt.Sprintf("$%d)", i+1))
		}
	}

	return sqlBuilder.String()
}

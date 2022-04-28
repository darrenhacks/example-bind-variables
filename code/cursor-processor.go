package main

import (
	"database/sql"
	"fmt"
)

// OpenCursorAndProcess runs the query by supplying bind variables and processes the results. It returns the number of
// account records processed.
func OpenCursorAndProcess(conn *sql.DB, sqlToRun string, args ...any) (int, error) {

	stmt, err := conn.Prepare(sqlToRun)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	return ScanCursorAndPrint(rows)
}

// ScanCursorAndPrint will read in each row and print out information about the account. It returns the number of
// account records processed.
func ScanCursorAndPrint(rows *sql.Rows) (int, error) {

	var accountId int
	var accountName string
	var accountBalance float32

	rowCount := 0
	for rows.Next() {
		rowCount++
		err := rows.Scan(&accountId, &accountName, &accountBalance)
		if err != nil {
			return rowCount, err
		}
		fmt.Printf("[%d]: %s {$%.2f)\n", accountId, accountName, accountBalance)
	}

	return rowCount, nil
}

package main

import (
	"database/sql"
	"fmt"
)

// DoTempTableDemo demonstrates using a temporary table to return process multiple rows as
// an alternative to in-lists. This is only peripherally related to using bind variables. The
// only one this uses is on the insert. I just included it as another way to dynamically determine
// the rows processed at run time.
func DoTempTableDemo(conn *sql.DB) error {

	fmt.Println("Using a temporary table to find a variable number of accounts.")

	expectedRowCount := 3

	transaction, err := conn.Begin()
	if err != nil {
		return err
	}

	err = createTempTable(transaction)
	if err != nil {
		transaction.Rollback()
		return err
	}

	populateTemp(transaction, expectedRowCount)
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	rowCount, err := doQuery(conn)
	if err != nil {
		return err
	}
	if rowCount == expectedRowCount {
		fmt.Printf("Expected and read %d rows.\n\n", expectedRowCount)
	} else {
		fmt.Printf("*** Incorrect number of rows were processed. Expeted %d but processed %d.\n\n",
			expectedRowCount, rowCount)
	}
	return nil
}

// createTempTable runs the SQL to create the temporary table.
func createTempTable(transaction *sql.Tx) error {

	tempTableSql := "CREATE TEMP TABLE tmp_accounts " +
		" (account_id INTEGER NOT NULL PRIMARY KEY)"
	_, err := transaction.Exec(tempTableSql)
	return err
}

// populateTemp loads the temporary table with the account numbers to
// query for.
func populateTemp(transaction *sql.Tx, count int) error {

	insertTempSql := "INSERT INTO tmp_accounts(account_id) values ($1)"

	stmt, err := transaction.Prepare(insertTempSql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		accountId := 1000 + i
		_, err = stmt.Exec(accountId)
		if err != nil {
			return err
		}
	}

	return nil
}

// doQuery constructs a query that uses the temp table to determine which accounts to process, reads in, and prints
// out those account details. It returns the number of accounts processed.
func doQuery(conn *sql.DB) (int, error) {

	stmt, err := conn.Prepare("SELECT act.ACCOUNT_ID, act.ACCOUNT_NAME, act.ACCOUNT_BALANCE::money::numeric " +
		"from BIND_VARIABLES_EXAMPLE_SCHEMA.ACCOUNTS act, tmp_accounts tmp " +
		"where act.ACCOUNT_ID = tmp.ACCOUNT_ID")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	return ScanCursorAndPrint(rows)
}

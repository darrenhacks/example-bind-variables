package main

import (
	"database/sql"
	"fmt"
)

// DoNotVulnerableDemo demonstrates how using bind variables can help prevent SQL injection attacks from
// malicious actors.
func DoNotVulnerableDemo(conn *sql.DB) error {

	// This shows what a programmer is expecting. The user supplies their account number and the query
	// will return only that user's account.
	fmt.Println("Using SQL with bind variables to find account 1000.")
	err := doNotVulnerableDemo(conn, "1000")
	if err != nil {
		return err
	}

	// This demonstrates how using bind variables can help prevent SQL injection attacks. Not only
	// will the application not processes accounts it should not, the application will receive an
	// error because the input cannot be translated by the DB into a number.
	fmt.Println("Using SQL with bind variables to find an account with injected SQL.")
	err = doNotVulnerableDemo(conn, "1 or ACCOUNT_ID > 0")
	if err != nil {
		fmt.Printf("Expected and recieved error when trying to inject malicious SQL: %s.\n", err)
	} else {
		fmt.Printf("*** Did not recieve error when trying to inject malicious SQL as expected.\n")
	}
	return nil
}

// doNotVulnerableDemo constructs a query based on user input and processes the result. If the expected
// number of rows were processed, it prints out a success message. If an unexpected amount was processed,
// it prints out an error message.
func doNotVulnerableDemo(conn *sql.DB, constraint string) error {

	sqlWithBindVariables := "select ACCOUNT_ID, ACCOUNT_NAME, ACCOUNT_BALANCE::money::numeric " +
		" from BIND_VARIABLES_EXAMPLE_SCHEMA.ACCOUNTS where ACCOUNT_ID = $1"

	rowCount, err := runBindsSql(conn, sqlWithBindVariables, constraint)
	if err != nil {
		return err
	}
	if rowCount == 1 {
		fmt.Printf("Correct number of rows were processed.\n\n")
	} else {
		fmt.Printf("*** Incorrect number of rows were processed: %d. This call was vulnerable.\n\n", rowCount)
	}
	return nil
}

// runBindsSql runs the query by supplying bind variables and processes the result set. It will print out
//information about the account and return the number of account records processed.
func runBindsSql(conn *sql.DB, sqlWithBindVariables string, accountNumber string) (int, error) {

	stmt, err := conn.Prepare(sqlWithBindVariables)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountNumber)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var accountId int
	var accountName string
	var accountBalance float32

	rowCount := 0
	for rows.Next() {
		rowCount++
		err = rows.Scan(&accountId, &accountName, &accountBalance)
		if err != nil {
			return rowCount, err
		}
		fmt.Printf("[%d]: %s {$%.2f)\n", accountId, accountName, accountBalance)
	}

	return rowCount, nil
}

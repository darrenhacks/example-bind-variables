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
	var accountNumber int32 = 1000
	err := doNotVulnerableDemo(conn, accountNumber)
	if err != nil {
		return err
	}

	// This demonstrates how using bind variables can help prevent SQL injection attacks. Not only
	// will the application not processes accounts it should not, the application will receive an
	// error because the input cannot be translated by the DB into a number.
	fmt.Println("Using SQL with bind variables to find an account with injected SQL.")
	err = doNotVulnerableDemo(conn, "1 or ACCOUNT_ID > 0")
	if err != nil {
		fmt.Printf("Expected and recieved error when trying to inject malicious SQL: %s.\n\n", err)
	} else {
		fmt.Printf("*** Did not recieve error when trying to inject malicious SQL as expected.\n\n")
	}
	return nil
}

// doNotVulnerableDemo constructs a query based on user input and processes the result. If the expected
// number of rows were processed, it prints out a success message. If an unexpected amount was processed,
// it prints out an error message.
func doNotVulnerableDemo(conn *sql.DB, constraint any) error {

	sqlWithBindVariables := "select ACCOUNT_ID, ACCOUNT_NAME, ACCOUNT_BALANCE::money::numeric " +
		" from BIND_VARIABLES_EXAMPLE_SCHEMA.ACCOUNTS where ACCOUNT_ID = $1"

	rowCount, err := OpenCursorAndProcess(conn, sqlWithBindVariables, constraint)
	if err != nil {
		return err
	}
	if rowCount == 1 {
		fmt.Printf("Expected and read 1 row.\n\n")
	} else {
		fmt.Printf("*** Incorrect number of rows were processed: %d. This call was vulnerable.\n\n", rowCount)
	}
	return nil
}

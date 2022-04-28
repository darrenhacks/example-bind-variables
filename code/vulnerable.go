package main

import (
	"database/sql"
	"fmt"
)

// DoVulnerableDemo demonstrates what can happen when you construct SQL based on input from a user.
func DoVulnerableDemo(conn *sql.DB) error {

	// This shows what a programmer is expecting. The user supplies their account number and the query
	// will return only that user's account.
	fmt.Println("Using constructed SQL to find account 1000.")
	err := doVulnerableDemo(conn, "1000")
	if err != nil {
		return err
	}

	// This demonstrates what a nefarious actor can do. Rather than providing an account number, they
	// provide something that, when added to the query the application will run, provides the actor
	// with more account information than they should have access to.
	fmt.Println("Using constructed SQL to find an account with injected SQL.")
	return doVulnerableDemo(conn, "1 or ACCOUNT_ID > 0")
}

// doVulnerableDemo constructs a query based on user input and processes the result. If the expected
// number of rows were processed, it prints out a success message. If an unexpected amount was processed,
// it prints out an error message.
func doVulnerableDemo(conn *sql.DB, constraint string) error {

	baseSql := "select ACCOUNT_ID, ACCOUNT_NAME, ACCOUNT_BALANCE::money::numeric " +
		" from BIND_VARIABLES_EXAMPLE_SCHEMA.ACCOUNTS where ACCOUNT_ID = %s"
	constructedSql := fmt.Sprintf(baseSql, constraint)
	rowCount, err := runConstructedSql(conn, constructedSql)
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

// runConstructedSql runs the query and processes the result set. It will print out information about
// the account and return the number of account records processed.
func runConstructedSql(conn *sql.DB, constructedSql string) (int, error) {

	stmt, err := conn.Prepare(constructedSql)
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

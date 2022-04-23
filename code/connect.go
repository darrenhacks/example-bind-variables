package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// DbConnectionError Error returned when a database connection cannot be made.
type DbConnectionError struct {
	Message string
}

// Converts a DbConnectionError to a string.
func (err DbConnectionError) Error() string {
	return err.Message
}

// Connect will try and connect to a database. Since this is meant to run in a container, it
// handles the case when the database container is up but not yet accepting connections. It will
// try and connect up to maxRetryCount times as defined in the supplied DbConnectionParams. It will run
// a query defined in the testQuery property of DbConnectionParams to ensure the connection is active.
// If, after maxRetryCount tries, the function cannot make a successful connection, the function will
// return an error.
func Connect(dbConnectionParams DbConnectionParams) (*sql.DB, error) {

	var lastError error

	for tryCount := 0; tryCount < dbConnectionParams.maxRetryCount; tryCount++ {

		if tryCount > 0 {
			fmt.Printf("Retrying...\n")
			time.Sleep(2 * time.Second)
		}
		db, err := connect(dbConnectionParams)
		if err != nil {
			fmt.Printf("Unable to connect to DB: %s\n", err)
			lastError = err
		} else {
			return db, nil
		}
	}

	return nil, lastError
}

// connect tries to make a connection to the database and confirm the connection.
// It returns the connection if successful and nill with an error code if not.
func connect(dbConnectionParams DbConnectionParams) (*sql.DB, error) {

	db, err := sql.Open("pgx", dbConnectionParams.dbUri)
	if err != nil {
		return db, err
	}

	err = runTestQuery(db, dbConnectionParams.testQuery)
	if err != nil {
		return db, err
	}

	return db, nil
}

// runTestQuery runs the test query and returns an error if it fails
// or nil if successful.
func runTestQuery(db *sql.DB, query string) error {

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return DbConnectionError{"Test query did not return any results"}
	}
	return nil
}

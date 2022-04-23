package main

import (
	"fmt"
	"os"
	"strconv"
)

// DbConnectionParams holds the parameters needed to connect to the database.
type DbConnectionParams struct {
	dbUri         string
	maxRetryCount int
	testQuery     string
}

// EnvironmentSetupError Error returned if the setup is not properly defined in the environment.
type EnvironmentSetupError struct {
	message string
}

// Converts a EnvironmentSetupError to a string.
func (err EnvironmentSetupError) Error() string {
	return err.message
}

// Creates a DbConnectionParams from environment variables. The following can be defined:
// DB_URI (required) The URI to use to connect to the database.
// MAX_CONNECT_RETRY_COUNT (optional) The number of times to try connecting before giving up. Defaults to 5.
// DB_TEST_QUERY (optional) A query to run to make sure the connection is live. Defaults to 'SELECT 1'.
func parseFromEnvironment() (DbConnectionParams, error) {

	dbUri, isFound := os.LookupEnv("DB_URI")
	if !isFound {
		return DbConnectionParams{}, EnvironmentSetupError{"DB_URI is required"}
	}

	maxRetryCount := 5
	var err error
	maxRetryCountAsString, isFound := os.LookupEnv("MAX_CONNECT_RETRY_COUNT")
	if isFound {
		maxRetryCount, err = strconv.Atoi(maxRetryCountAsString)
		if err != nil {
			return DbConnectionParams{},
				EnvironmentSetupError{fmt.Sprintf("Unable to convert '%s' to an integer", maxRetryCountAsString)}
		}
	}

	testQuery, isFound := os.LookupEnv("DB_TEST_QUERY")
	if !isFound {
		testQuery = "SELECT 1"
	}

	return DbConnectionParams{dbUri: dbUri, maxRetryCount: maxRetryCount, testQuery: testQuery}, nil
}

package main

import (
	"fmt"
	"os"
)

// Driver function.
func main() {

	// Grab the application variables from the environment.
	connectionParams, err := parseFromEnvironment()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Connecting...\n")

	// Try and get a connection to the database.
	conn, err := Connect(connectionParams)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
	defer conn.Close()
	fmt.Printf("Database connection successful...\n\n")

	// Do a demonstration of how dynamically building SQL at runtime
	// is vulnerable to SQL injection.
	err = DoVulnerableDemo(conn)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// Do a demonstration of how using bind variables helps eliminate
	// SQL injection vulnerabilities.
	err = DoNotVulnerableDemo(conn)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// Do a demonstration of using bind variables with an in-list.
	err = DoInList(conn)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// Do a demonstration of using a temporary table.
	err = DoTempTableDemo(conn)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
}

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
}

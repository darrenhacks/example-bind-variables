# Bind Variables Example

An example program showing how you can prevent some SQL injection 
vulnerabilities by using bind variables.

The project includes a PostgreSQL database and a Go application. The
database holds a single table with example data. The Go application
shows some examples of not using bind variables and possible consequences.
It further goes on to show how these consequences can be avoided by 
using bind variable.

## Running the Application

Both the database and application are configured to run inside Docker 
containers. If you have Docker installed, you can build and run both 
with the command `docker-copose up --build`. The application container will shut
down after the application completes, but the database will remain
running. You can shut off the database with the command `docker-compose down`.

## What to Look For

There are two files here that are the core of the demonstration.

[vulnerable.go](code/vulnerable.go) shows a common pattern of dynamically building SQL based
on a hard-coded base and using string formatting or concatenation to dynamically change the
SQL at runtime. It shows how carefully crafted inputs can produce query results the programmer 
did not intend.

[not-vulnerable.go](code/not-vulnerable.go) shows how to use bind variables to eliminate that
issue.

## Caveat

I am new to Go, and chose to write this application in Go as a learning
exercise. The Go code may or may not be following best practices, but the 
effect of using bind variables (or not) is still visible. 


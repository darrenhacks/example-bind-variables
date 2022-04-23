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
with the command `docker-copose up`. The application container will shut
down after the application completes, but the database will remain
running. You can shut off the database with `docker-compose down`.

## Caveat

I am new to Go, and chose to write this application in Go as a learning
exercise. The Go code may or may not be following best practices, but the 
effect of using bind variables (or not) is still visible. 


# Book: Let's Go Professional Package by Alex Edwards

A book dedicated to learning about the Go language for beginners and professionals.

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run ./cmd/web
```

To download the dependencies of the project, you can use the command:

```bash
go mod download
```

If you want to compile the project, you can use the command:

```bash
go build <project-name>.go
```

But if you want to upgrade the dependencies of the project, you can use the command: `u` to upgrade the dependencies.

```bash
go get -u github.com/foo/bar
```

And if you want to see the compiled project, you can use the command:

```bash
./<project-name>
```

## What I studied in this book?

- [x] Introduction
  - [x] conventions
  - [x] About the author
  - [x] Copyright and disclaimer
  - [x] Prerequisites
    - [x] Background knowledge
    - [x] Go 1.21
    - [x] Other Software
- [x] Foundations
  - [x] Project setup and creating a module
    - [x] Creating a module
    - [x] Hello World!
    - [x] Additional Information
      - [x] Module paths for downloadable packages
  - [x] Web Application basics
    - [x] Additional Information
      - [x] Network addresses
      - [x] Using `go run`
  - [x] Routing requests
    - [x] Fixed path and subtree patterns
    - [x] Restricting the root url pattern
    - [x] Additional informations
      - [x] Servermux features and quirks
      - [x] Host name matching
      - [x] The default servemux
      - [x] What about RESTful routing?
  - [x] Customizing HTTP headers
    - [x] HTTP status codes
    - [x] Customizing headers
    - [x] The `http.Error` shortcut
    - [x] The `net/http` constants
    - [x] Additional information
      - [x] System generated headers and content sniffing
      - [x] Manipulating the header map
      - [x] Header canonicalization
      - [x] Suppressing system-generated headers
  - [x] URL query strings
    - [x] The `io.writer` interface
    - [x] Additional information
      - [x] **[Golang Interfaces Explained](https://www.alexedwards.net/blog/interfaces-explained)**
  - [x] Project structure and organization
    - [x] Refactoring your existing code
    - [x] Additional information
      - [x] The internal directory
  - [x] HTML templating and inheritance
    - [x] Template composition
    - [x] Embedding partials
    - [x] Additional information
      - [x] The block action
  - [x] Serving static files
    - [x] The `http.FileServer` handler
    - [x] Using the static files
    - [x] Additional information
      - [x] File server features and functions
      - [x] Performance
      - [x] Serving single files
      - [x] Disabling directory listings
  - [x] The `http.Handler` interface 
    - [x] Handler functions
    - [x] Chaining handlers
    - [x] Requests are handled concurrently
- [ ] Configuration
  - [x] Managing configuration settings
    - [x] Command-line flags
    - [x] Default values
    - [x] Type conversions
    - [x] Automated help
    - [x] Additional information
      - [x] Environment variables
      - [x] Boolean flags
      - [x] Pre-existing variables
  - [x] Structured logging
    - [x] Creating a structured logger
    - [x] Using a structured logger
    - [x] Adding structured logging to our application
    - [x] Additional information
      - [x] Safer attributes
      - [x] JSON formatted logs
      - [x] Minimum log level
      - [x] Caller location
      - [x] Decoupled logging
      - [x] Concurrent logging
  - [x] Dependency injection
    - [x] Adding a deliberate error
    - [x] Additional information
      - [x] Closures for dependency injection
  - [x] Centralized error handling
    - [x] Revert the deliberate error
    - [x] Additional information
      - [x] Stack traces
  - [x] Isolating the application routes
- [ ] Database-driven responses
  - [x] Setting up MySQL
    - [x] Scaffolding the database
    - [x] Creating a new user
    - [x] Test the new user
  - [x] Installing a database driver
  - [x] Modules and reproducible builds
    - [x] Additional information
      - [x] Upgrading packages
      - [x] Removing unused packages
  - [x] Creating a database connection pool
    - [x] Using it in our web application
    - [x] Testing a connection
  - [x] Designing a database model
    - [x] Using the `SnippetModel`
    - [x] Additional Information
      - [x] Benefits of this structure
  - [x] Executing SQL statements
    - [x] Executing the query
    - [x] Using the model in our handlers
    - [x] Additional information
      - [x] Placeholder parameters
  - [x] Single-record SQL queries
    - [x] Using the model in our handlers
    - [x] Additional information
      - [x] Checking for specific errors
      - [x] Shorthand single-record queries
  - [x] Multiple-record SQL queries
    - [x] Using the model in our handlers
  - [x] Transactions and other details
    - [x] The `database/sql` package
    - [x] Verbosity
    - [x] Managing null values
    - [x] Working with transactions
    - [x] Prepared statements
- [ ] Dynamic HTML templates
  

# Book: Let's Go Professional Package by Alex Edwards

A book dedicated to learning about the Go language for beginners and professionals.

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run <project-name>.go
```

If you want to compile the project, you can use the command:

```bash
go build <project-name>.go
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
- [ ] Foundations
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
    - [ ] Additional information
      - [ ] **[Golang Interfaces Explained](https://www.alexedwards.net/blog/interfaces-explained)**
  - [ ] Project structure and organization
    - [ ] Refactoring your existing code
    - [ ] Additional information
      - [ ] The internal directory
  - [ ] HTML templating and inheritance
  - [ ] Serving static files
  - [ ] The `http.Handler` interface 
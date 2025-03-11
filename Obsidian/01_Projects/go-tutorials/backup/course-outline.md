# Go: Beginner to Advanced

## Basics

- Creating a Go project from scratch (Explain IDE options and Go modules from get-go)
- Your first Go program, what it means (Discuss Go playground as a learning Tool)
- Variables, constants and Operators (Include type inference and best practices)
- Conditionals
- Loops
- Functions (Variadic and anonymous functions)
- Error Handling (Basics) (Best practices and common Go idioms)
- Console read and write
- Record the solution so people can cross-check (Encourage community code reviews)

### Project

- CLI tool (maybe calculator or a task/todo list application to add the CRUD operations)

## Mid level

- Struct and Methods (Composition over inheritence)
- Pointers and Memory Management (example of pointer misuse and common pitfalls)
- Arrays and Slices (Discuss copy behavior and capacity vs length)
- Maps (Include working of complex key types)
- Introduction to Package and Visibilty (exported vs un-exported identifiers)
- Value and Map pointer semantics (Pointers in depth - Pass by value and reference)

### Project

- Contact Management System (Include with local file system for persistence)

### Testing (Learn Go with Tests)

- Different type of tests (Introduce Table driven tests)
- Writing tests using go testing package
- Mocks and Interfaces in testing

### Project

- Refactoring the Contact management system with TDD approach.

## Advanced

- Concurrency vs Parallelism
- Defer, Panic and Recover (Discuss use-cases and control flows)
- Understanding Goroutines (best practices and common mistakes)
- Channels
- Select statement and Timeouts (example with network I/O)
- Concurrency Patterns (Producer-consumer, worker pools)
- Record a solution of building it, so people can cross check (common concurrency issues and their solutions)
- Handling File and I/O.
- Interfaces (Empty interfaces and type assertions)
- Readers, Stringer and Writer (common implementations)
- Uses of Interfaces
- Using outside packages (go get, vendoring, module proxies)
- Reflection (Discuss when to use and avoid reflection)

### Advanced Project

- Concurrent Web Crawler
- Microservices in Go
  - Microservice backend server using net/http.
    - Things learnt
      - net/http package
      - Routing and Middleware
      - Connecting to database
      - Marshalling and Unmarshalling
  - Recreate the entire thing using Gin
    - Things learnt
      - patterns and strucutre used in modern applications
      - Middlewares and extended features
      - Advanced features of Gin
- DevOps Project
  - Creating images and deploying on kubernetes.
  - CI/CD pipline using GitOps for deployment.

# AI suggested curriculum

# Go: Beginner to Advanced - Final Curriculum

## Basics

1. **Introduction to Go and the Development Environment**

   - Introducing Go (History, Features, Usage)
   - Setting up the Go Development Environment
   - Exploring IDE options and Go modules
   - Discussion of the Go Playground as a learning tool

2. **Understanding Go Basics**

   - Creating a Go project from scratch
   - Your first Go program and its structure
   - Variables, Constants, and Operators (including type inference)
   - Conditionals and Loops (including `range` and common patterns)

3. **Fundamentals of Go Programming**

   - Functions (including variadic and anonymous functions)
   - Understanding Arrays, Slices, and Maps
   - Structs and Methods (composition over inheritance)
   - Introduction to Error Handling in Go (best practices)

4. **Working with the Console and Basic File I/O**

   - Console read and write operations
   - Basic File Handling (reading from and writing to files)
   - Error Handling (continued with file operations)

5. **Go Packages and Visibility**
   - Basics of Go packages
   - Exported vs. unexported identifiers
   - Organizing code for reusability and modularity

### Project 1: CLI Tool (Task/To-Do List Application)

## Mid-Level

6. **Deep Dive into Go Data Types**

   - Pointers and Memory Management (including misuse and pitfalls)
   - Deep understanding of Slices and Arrays (copy behavior, capacity vs. length)
   - Maps in-depth (working with complex key types)

7. **Introduction to Go Testing**
   - Different types of tests (unit, integration, end-to-end)
   - Writing tests using the Go testing package
   - Table-driven tests
   - Mocks and Interfaces in testing

### Project 2: Contact Management System (with TDD)

## Advanced

8. **Advanced Go Concepts and Concurrency**

   - Concurrency vs Parallelism
   - Understanding Goroutines (best practices)
   - Advanced Control Structures (Defer, Panic, Recover)
   - Channels and Concurrency Patterns (select statement, timeouts, producer-consumer, worker pools)

9. **Working with Interfaces and I/O in Go**

   - Interfaces in-depth (empty interfaces, type assertions, best practices)
   - Readers, Stringer, and Writer (custom implementations)
   - Advanced File and I/O operations
   - Using APIs and handling network I/O

10. **Exploring External Packages and Reflection**

- Managing dependencies (Go modules, vendoring, module proxies)
- Security best practices in Go applications
- Reflection (use cases and caution)

11. **Effective Go Project Layout and Design**

- Go project layout best practices
- Package design for maintainable code
- CI/CD pipelines for Go applications (overview)

### Advanced Project: Implement a Concurrent Web Crawler

## Microservices with Go

12. **Building a Microservice Backend with `net/http`**
    - Introduction and use cases for microservices
    - Working with the `net/http` package (routing, middleware)
    - Connecting to databases (including practical security practices)
    - JSON handling (marshalling/unmarshalling)

### Project 3: Microservices Backend Server using `net/http`

13. **Using a Framework: Gin for Rapid Development**
    - Transition to Gin for microservice development
    - Overview of Gin (routing, middleware, advanced features)
    - Enhancing the microservice with Gin's extended features (performance, error handling)

### Project 4: Recreate the Microservice using Gin

**Additional Topics and Community Engagement:**

- Working with GraphQL in Go (Bonus Session)
- Live Q&A Sessions / 'Office Hours' for viewer engagement
- Open Source Contributions in Go (How to find and contribute to projects)
- Keeping up with Go (Staying current with updates and community resources)

# Finalised Plan

# Go: Beginner to Advanced - Enhanced Project Outline

## Basics

**Project: Command-Line Todo List Application**

- Skills Covered:
  - Basic Go syntax
  - File I/O for data persistence
  - Error handling with file operations
  - Structuring a Go project using packages
  - Command-line parsing

## Mid-Level

**Project: RESTful API for a Bookstore**

- Skills Covered:
  - HTTP server using `net/http`
  - Structs and methods for data models
  - JSON handling for request/response
  - CRUD operations
  - Pointers and interfaces
  - Unit tests for API endpoints

## Testing

**Project: Advanced Testing for Bookstore API**

- Skills Covered:
  - Table-driven tests
  - Mock objects for database interactions
  - Integration testing
  - Performance testing and benchmarking

## Advanced

**Project: Multi-user Chat Server**

- Skills Covered:
  - Goroutines for handling multiple clients
  - Channels for client communication
  - Network programming (TCP/WebSocket)
  - Concurrency patterns
  - Error handling with defer, panic, and recover

## Advanced Projects

**Project: Distributed URL Shortening Service**

- Skills Covered:
  - Microservice development with `net/http`
  - Database integration
  - Middleware for analytics or authentication
  - Interface-based design

**DevOps Project: Deployment and Automation**

- Skills Covered:
  - Creating images for deployment on Kubernetes
  - CI/CD pipeline using GitOps

These projects will allow students to apply learned concepts from the course and gain practical experience in realistic scenarios.

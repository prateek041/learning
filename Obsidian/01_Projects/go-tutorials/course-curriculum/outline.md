# Go: Beginner to Advanced - Enhanced Project Outline

Below is the project list

## Project Lists
- Simple Calculator (functions, user input and error handling)
- File reader and writer (handling the os package)
- CLI to-do list (structs, slices, maps, basic CRUD operations)
- Concurrent downloader
- URL shortener
- Chat Application
- Weather CLI app
- Personal Blog engine
## Basics
**Project: Command-Line Todo List Application** (or a calculator)
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
**Project: Multi-user Chat Server** (or e-com website)
- Skills Covered:
  - Go-routines for handling multiple clients
  - Channels for client communication
  - Network programming (TCP/WebSocket)
  - Concurrency patterns
  - Error handling with defer, panic, and recover
  - Database integration
  - Middleware for analytics or authentication
  - Interface-based design

**DevOps Project: Deployment and Automation**
- Skills Covered:
  - Creating images for deployment on Kubernetes
  - CI/CD pipeline using GitOps

# Topics to be covered and their order
## Go Basics
### Introduction to the language
#### Functiones, variables and types
- Basic Syntax and Types
- Functions
- Variables with Initializers
- Short variable Declarations
- Basic Types
- Zero Values
- Type Conversions
- Type inference
- Constants
- Numeric Constants

#### Flow Control
- For
- For is Go's while
- Forever
- If
- If with a short statement
- If and Else
- Switch
- Switch Evaluation order
- Defer
- Stacking Defer

#### Packages
- Difference between modules and packages in Go
- How package imports and exports work in Go.
- Go modules
- using third party libraries

> [!IMPORTANT]
> A project here, maybe just implementing a pretty printing logger built on top of an external library

#### Types, Slices and Maps
- Arrays
- Slices are like references to arrays
- Slice literals
- Slice defaults
- Slice length and capacity
- Nil Slices
- Creating a Slice with make
- Slices of slices
- Appending to a Slice
- Range
- Maps
- Map Literal
- Structs (why to use structs)
- Struct Fields
- Struct literals

> [!IMPORTANT]
> A Basic project here that brings together all the above topics, a calculator or a Coffee machine should work.

#### Testing in Go

- Why to test your programs
- How testing works in Go.
- Writing tests from Scratch
- What is TDD.
- Introduction to how to use TDD.

#### Pointers and their importance

- What are pointers
- how they work in Go
- why to use them

#### Methods and Interfaces

- What are interfaces
- What is the use of interfaces
- implementing interfaces, what does it mean

#### Generics

- What are generics?
- What is the use of Generics in a Go program

#### Concurrency

- Concurrency vs Parallelism
- How Go handles Concurrency
- Channels
- Signals
- Go routines

> [!IMPORTANT]
> A project that ties all the above topics together, I am still thinking of a good usecase here.

### Leveling Up

#### Rest APIs

##### Basics

- Intruduction to in built packages like `net/http`
- Web server introduction
- How to read docs
- Explaining `net/http` package internals like Multiplexer, Handlers, Server, Middlewares etc.
- Building a Rest API only with `net/http`.

##### Intermediate

- Migrating to robust frameworks like Gorilla
- Adding Database (mongoDB)
- How to read docs and search for pakages
- Best Practices like Graceful shutdown, server timeout etc.
- Full Fledged Rest API supporting all the CRUD operations along with User handling using Auth middlewares.- Testing your Rest API.

##### Advanced

- Introduction to gRPC
- How it works
- A Toy project of showing how a gRPC based microservice works

### Capstone Project

- Distributed microservices app development (API spec given) eg: Chat Application, E-com.
- Frontend is shared complementry (User doesn't have to know to Frontend Development)

## Add-ons

- Dockerizing the application.
- Instrumenting monitoring using Prometheus Go Library.
- Deploying on Kubernetes

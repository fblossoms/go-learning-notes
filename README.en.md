# Go Learning Notes and Code Examples

[中文](README.md) | **English**

The content starts from Go fundamentals and gradually covers data structures, functions, structs, interfaces, error handling, modularization, logging, serialization, concurrency, network programming, database access, and common Go ecosystem libraries.

This repository is not a monolithic business project. It is a learning-oriented codebase for review, lookup, and long-term accumulation. Each topic is kept as independent as possible so it can be run, read, and extended separately.

## Tech Stack

- **Language**: Go
- **Core syntax**: variables, constants, types, control flow, functions, pointers, structs, interfaces, error handling
- **Data structures**: arrays, slices, sub-slices, hash tables, strings, linear data structures
- **Concurrency**: goroutine, channel, timer, GMP model, channel multiplexing
- **Network programming**: TCP programming, Web Server basics, HTTP request and response model
- **Databases**: MySQL, MongoDB
- **Database tools**: `database/sql` standard library, SQL Builder, GORM
- **Logging**: standard `log`, `zerolog`, file logging
- **Serialization**: struct serialization, JSON, binary serialization, struct tag
- **Supporting examples**: Python concurrency, socket, simple Web Server experiments

## Knowledge Map

### Go Fundamentals

- Comments, identifiers, variables, constants, `iota`
- Integer, floating-point, string, character, and formatted output
- Branching, loops, random numbers, input handling
- Type conversion, scope, recursion, nested functions, closures

### Data Structures and Memory Model

- Pointers and pointer basics
- Arrays, slices, sub-slices, `append`, capacity growth strategy, `copy`
- Linear data structures and hash tables
- Reference types, shallow copy, and deep copy

### Functions, Structs, and Interfaces

- Function definitions, function types, parameters
- `defer` and execution order
- Structs, struct pointers, anonymous structs
- Constructors, embedded structs, method receivers
- Interfaces, type assertions, formatting interfaces
- Error handling and basic testing examples

### Modularization, Logging, and Serialization

- Go module and local package organization
- Standard logging and `zerolog`
- Error logging and file handling
- Struct serialization, JSON, binary serialization

### Concurrency and Network Programming

- Basic concepts of concurrency and parallelism
- goroutine and GMP model
- TCP programming
- Web Server basic examples
- channel, timer, `struct{}` channel, channel multiplexing, channel concurrency
- Python concurrency and socket examples as supporting comparison material

### Database and Persistence

- Database fundamentals and SQL
- MySQL access with `database/sql`
- SQL Builder
- ORM concepts and GORM
- MongoDB CRUD examples
- `context` usage in database and network-style programs

## Learning Path

1. `Day01` to `Day09`: build Go fundamentals, basic types, branching, loops, and input/output concepts.
2. `Day10` to `Day16`: move into arrays, slices, strings, map, type conversion, hash tables, and other data structure topics.
3. `Day17` to `Day26`: learn functions, scope, recursion, closures, `defer`, structs, interfaces, copying, error handling, and tests.
4. `Day27` to `Day34`: organize object-oriented-style patterns, serialization, modularization, logging, and file handling.
5. `Day35` to `Day40`: focus on concurrency, goroutine, channel, TCP, and Web Server.
6. `Day41` to `Day44`: enter database-related topics, including SQL, MySQL, SQL Builder, GORM, MongoDB, and `context`.

## Repository Structure

Each `DayXX` directory is further split into topic subdirectories. Most topic directories contain:

- One `.go` example file
- One `.md` note file for the same topic

Example:

```txt
Day42/
  01_database_development/
    01_database_development.go
    01_database_development.md
  02_sql_builder/
    02_sql_builder.go
    02_sql_builder.md
```

This structure allows each `package main` example to compile independently and avoids `main redeclared` errors when multiple examples exist under the same day.

## Run Examples

Run a single topic example:

```bash
go run ./Day42/02_sql_builder
```

Check all Go code in the repository:

```bash
go test ./...
```

Database-related examples require local MySQL or MongoDB services and database configuration matching the connection strings in the code.

## Skills

- Writing and organizing small runnable examples in Go
- Understanding value types, reference types, and basic memory behavior
- Modeling with arrays, slices, map, structs, and interfaces
- Mastering functions, closures, `defer`, method receivers, and error handling
- Using the Go standard library for I/O, time, logging, SQL, networking, and related tasks
- Using goroutine and channel to handle concurrency problems
- Using MySQL, MongoDB, GORM, and SQL Builder for basic persistence operations


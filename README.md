# Golang Practice (`belajargolang`)

A collection of Go (Golang) programming exercises, assignments, and practical examples covering basic syntax, data structures, object-oriented concepts, and advanced concurrency patterns in Go.

---

## 📚 Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
  - [Module 1: Getting Started with Go](#module-1-getting-started-with-go)
  - [Module 2: Functions, Methods, and Interfaces](#module-2-functions-methods-and-interfaces)
  - [Module 3: Concurrency in Go](#module-3-concurrency-in-go)
- [Prerequisites](#prerequisites)
- [How to Run](#how-to-run)
- [Author](#author)

---

## 📖 Overview

This repository contains exercises and solutions structured into three main learning modules:
1. **Module 1**: Fundamental Go concepts including basic types, string manipulation, slices, maps, JSON marshaling/unmarshaling, and file I/O.
2. **Module 2**: Intermediate Go concepts focusing on pointers, functions as first-class values, closures, structs, methods, and interface polymorphism.
3. **Module 3**: Advanced concurrency in Go, exploring goroutines, race conditions, channels, parallel sorting, and classic synchronization problems like the Dining Philosophers.

---

## 🗂️ Project Structure

```text
golangpractice/
├── modul1/                # Basic Go syntax and core data structures
│   ├── findian.go         # String search and matching
│   ├── makejson.go        # Map creation and JSON marshaling
│   ├── read.go            # File reading and parsing into structs
│   ├── slice.go           # Dynamic slices and sorting
│   └── trunc.go           # Floating-point truncation
│
├── modul2/                # Functions, structs, methods, and interfaces
│   ├── bubbleSort.go      # Bubble sort implementation with pointers/slices
│   ├── activity3.go       # Physics kinematics calculator using closures
│   ├── activity4.go       # Animal simulation using interfaces and structs
│   └── main_test.go       # Unit tests
│
├── modul3/                # Concurrency, goroutines, and channels
│   ├── activity3-1.go     # Race condition demonstration
│   ├── activity3-2.go     # Concurrent 4-partition array sort using goroutines
│   ├── activity3-3.go     # Dining Philosophers problem simulation
│   ├── activity3-4.go     # Dining Philosophers with host/waiter synchronization
│   └── explanation.txt    # Concurrency and race condition writeup
│
├── .gitignore             # Ignored binaries, build artifacts, and IDE configs
└── README.md              # Project documentation
```

---

### Module 1: Getting Started with Go

Focuses on foundational Go syntax, input/output, and standard library data structures:

- **`trunc.go`**: Prompts the user for a floating-point number and prints it truncated to an integer.
- **`findian.go`**: Analyzes user-entered strings to check if they begin with `'i'`, contain `'a'`, and end with `'n'` (case-insensitive).
- **`slice.go`**: Creates an empty integer slice, repeatedly prompts the user to add integers, sorts the slice in ascending order after each input, and exits when `'X'` is entered.
- **`makejson.go`**: Prompts for a name and address, stores them in a map, and serializes the map into a JSON object.
- **`read.go`**: Reads a text file containing first and last names separated by spaces, parses them into a slice of structs, and prints the results.

---

### Module 2: Functions, Methods, and Interfaces

Explores functional programming, object-oriented principles, and modular design in Go:

- **`bubbleSort.go`**: Implements Bubble Sort with an interactive CLI, using helper functions to swap slice elements.
- **`activity3.go`**: Computes displacement based on acceleration, initial velocity, and initial displacement using functions as first-class values and closures (`GenDisplaceFn`).
- **`activity4.go`**: Simulates animals (`cow`, `bird`, `snake`) using structs, methods (`Eat`, `Move`, `Speak`), and Go interfaces to handle dynamic creation and querying.

---

### Module 3: Concurrency in Go

Demonstrates concurrent execution, synchronization primitives, and communication between goroutines:

- **`activity3-1.go`**: Explains and demonstrates a race condition where multiple concurrent goroutines access and modify shared memory without synchronization.
- **`activity3-2.go`**: Implements a concurrent sorting program that divides an array into 4 partitions, sorts each partition concurrently using goroutines, and merges the sorted partitions into a single sorted array.
- **`activity3-3.go` & `activity3-4.go`**: Solves the classic **Dining Philosophers** synchronization problem using mutexes, channels, and a host/arbitrator pattern ensuring no more than two philosophers eat simultaneously to prevent deadlocks.
- **`explanation.txt`**: Detailed theoretical analysis of race conditions, interleaved operations, and non-deterministic behavior.

---

## 🛠️ Prerequisites

- **Go**: Version 1.18 or higher installed on your system.
  Verify your installation:
  ```bash
  go version
  ```

---

## 🚀 How to Run

Navigate into any module directory or run files directly using `go run`:

### Running Module 1 Examples
```bash
# Truncate floating-point numbers
go run ./modul1/trunc.go

# Interactive slice sorting
go run ./modul1/slice.go

# JSON serialization
go run ./modul1/makejson.go

# Read file and parse structs
go run ./modul1/read.go
```

### Running Module 2 Examples
```bash
# Bubble Sort
go run ./modul2/bubbleSort.go

# Kinematics displacement calculator
go run ./modul2/activity3.go

# Animal simulation CLI
go run ./modul2/activity4.go
```

### Running Module 3 Examples
```bash
# Race Condition demo (run with race detector)
go run -race ./modul3/activity3-1.go

# Concurrent partition sort
go run ./modul3/activity3-2.go

# Dining Philosophers simulation
go run ./modul3/activity3-4.go
```

---

## 👤 Author

- **Kukuh Dwi S** ([@kukuhdwis](https://github.com/kukuhdwis))

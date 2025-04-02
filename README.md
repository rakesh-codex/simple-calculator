# Simple Calculator

## Project Overview
This is the second project in a **100 Projects Series to Super Master Golang**. It’s a command-line calculator written in Golang that performs basic arithmetic operations (addition, subtraction, multiplication, division) based on user-provided arguments. The project builds on CLI argument handling and introduces numeric parsing and error management.

### Purpose
- Learn to parse and validate numeric input in Golang.
- Practice control structures (`switch`) and error handling.
- Develop skills for building interactive CLI tools.

### Difficulty Level
- **Beginner**: Suitable for those familiar with basic Golang syntax.

---

## Tech Stack
- **Language**: Golang (using the standard library)
- **Packages**:
- `fmt`: For input/output formatting.
- `os`: For accessing command-line arguments.
- `strconv`: For converting strings to numbers.
- **Tools**:
- Go compiler (version 1.21 or later recommended).
- Any text editor or IDE (e.g., VS Code with Go extension).

---

## Prerequisites
- **Go Installed**: Verify with `go version`.
- **Basic Terminal Knowledge**: Ability to run commands in a terminal.

---

## Setup Instructions
1. **Create the Project Directory**:
```bash
mkdir simple-calculator
cd simple-calculator
2. **Initialize a Go Module**:
```bash
go mod init simple-calculator
3. **Create the Source File:**:
Create `main.go` and add the code from the section.
4. Run the Program:
See the section for examples.


Implementation :

package main

import (
"fmt"
"os"
"strconv"
)

func main() {
if len(os.Args) != 4 {
fmt.Println("Usage: go run main.go <number1> <operation> <number2>")
            fmt.Println("Operations: add, sub, mul, div")
            fmt.Println("Example: go run main.go 5 add 3")
            os.Exit(1)
            }

            num1Str := os.Args[1]
            operation := os.Args[2]
            num2Str := os.Args[3]

            num1, err1 := strconv.ParseFloat(num1Str, 64)
            if err1 != nil {
            fmt.Println("Error: First argument must be a number")
            os.Exit(1)
            }

            num2, err2 := strconv.ParseFloat(num2Str, 64)
            if err2 != nil {
            fmt.Println("Error: Second argument must be a number")
            os.Exit(1)
            }

            var result float64
            switch operation {
            case "add":
            result = num1 + num2
            case "sub":
            result = num1 - num2
            case "mul":
            result = num1 * num2
            case "div":
            if num2 == 0 {
            fmt.Println("Error: Division by zero is not allowed")
            os.Exit(1)
            }
            result = num1 / num2
            default:
            fmt.Println("Error: Invalid operation. Use add, sub, mul, or div")
            os.Exit(1)
            }

            fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operation, num2, result)
}
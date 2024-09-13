// Package is a collection of common source code files
// Types of packages - Executable(associated with doing something) and Reusable(code dependencies | libraries)
package main // Used if we want to create an executable file

// fmt is a standatd library package - https://pkg.go.dev/std
import "fmt" // Gives main package access to other functions and logic written in other packages

// Function declaration
func main() {
	fmt.Println("Hi there!")
}

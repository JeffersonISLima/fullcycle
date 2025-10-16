/*
Understanding the First Line
In Go, the first line of code in a package is used to define the package name. This is a convention that helps organize code and avoid naming conflicts.

Key concepts:
- File names can be different (main.go and hello.go) but both must declare the same package name (package main)
- Files in the same package can share variables, functions, and constants
- The Hello constant is defined in hello.go and accessed in main.go because they belong to the same package scope
- When running: use 'go run .' to compile all files in the package, not 'go run main.go' alone
*/

package main

import "fmt"

func main() {
	fmt.Println(Hello)
}

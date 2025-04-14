package greetings_test

import (
    "fmt"

    "github.com/linglingye001/greetings"
)

// ExampleHello demonstrates how to use the Hello function.
//
// Parameters:
//   - para1: xxxx
//   - para2: xxxx
//
// Returns:
//   - para3: xxxx
func ExampleHello() {
    // Get a greeting message for "Alice"
    message := greetings.Hello("Alice")
    fmt.Println(message)
    // Output: Hi, Alice. Welcome!
}
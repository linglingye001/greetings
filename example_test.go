package greetings_test

import (
    "fmt"

    "github.com/linglingye001/greetings"
)

// ExampleHello demonstrates how to use the Hello function.
func ExampleHello() {
    // Get a greeting message for "Alice"
    message := greetings.Hello("Alice")
    fmt.Println(message)
    // Output: Hi, Alice. Welcome!
}
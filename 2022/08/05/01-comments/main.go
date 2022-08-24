// Package comments demonstrates how to use the new doc rendering feature added
// in Go 1.19 that is similar to Markdown.
//
// # What is included in this dummy package?
//
// A simple function called [HelloWorld], that's it.
package comments

// HelloWorld returns a welcome message
//
// Depending on who you ask something important, other examples:
//
//  1. hola
//  2. welcome!
//
// You can also search for other welcome messages on [Google]
//
// [Google]: https://google.com/
func HelloWorld() string {
	return "hello world"
}

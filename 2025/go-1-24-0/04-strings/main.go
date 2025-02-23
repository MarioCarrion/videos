package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	lines := `one
two
three`

	//- strings.Lines
	fmt.Println("- strings.Lines")
	for v := range strings.Lines(lines) {
		fmt.Printf("%s", v) // Keeps `\n`
	}
	fmt.Printf("\n")
	//-

	//- strings.SplitSeq
	fmt.Println("- strings.SplitSeq")
	for v := range strings.SplitSeq(lines, "\n") {
		fmt.Printf("%s\n", v) // Does not keep "\n"
	}

	//- strings.SplitAfterSeq
	fmt.Println("- strings.SplitAfterSeq")
	for v := range strings.SplitAfterSeq(lines, "\n") {
		fmt.Printf("%s", v) // Keeps "\n"
	}
	fmt.Printf("\n")

	//- strings.FieldsSeq
	fmt.Println("- strings.FieldsSeq")
	for v := range strings.FieldsSeq(lines) {
		fmt.Printf("%s\n", v) // Does not keep "whitespace" character
	}

	//- strings.FieldsFuncSeq
	fmt.Println("- strings.FieldsFuncSeq")
	f := func(r rune) bool {
		return unicode.IsSpace(r)
	}
	for v := range strings.FieldsFuncSeq(lines, f) {
		fmt.Printf("%s\n", v)
	}
}

package main

import (
	"5-avoid-stuttering-types/file"
)

func main() {
	_ = file.File{}
	_ = file.New()
	_ = file.NewFancy()
	_ = file.Fancy{}
}

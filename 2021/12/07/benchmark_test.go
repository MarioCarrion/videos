package benchmark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concat("a", "b")
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%s%s", str, "a")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer

	for n := 0; n < b.N; n++ {
		buf.WriteString("a")
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder

	for n := 0; n < b.N; n++ {
		builder.WriteString("a")
	}
}

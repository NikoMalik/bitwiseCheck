package bitwisecheck

import (
	"fmt"
	"strings"
	"testing"

	lowlevelfunctions "github.com/NikoMalik/low-level-functions"
)

func BenchmarkFlags(b *testing.B) {
	b.ReportAllocs()
	var flags Flags
	for i := 0; i < b.N; i++ {
		flags.Set(0)
		flags.Clear(0)
		flags.IsSet(0)
		_ = flags.String()
	}
}

func BenchmarkFlagsIsSet(b *testing.B) {
	b.ReportAllocs()
	var flags Flags
	flags.Set(0)
	for i := 0; i < b.N; i++ {
		flags.IsSet(0)
	}
}

func BenchmarkTestBuffer(b *testing.B) {
	b.ReportAllocs()
	buffer := lowlevelfunctions.NewStringBuffer(8) // Create once

	for i := 0; i < b.N; i++ {
		var flags Flags
		flags.Set(0)
		flags.Set(1)
		flags.Clear(0)
		buffer.Reset() // Reset buffer for reuse
		buffer.WriteString("00000001")
	}
}

func BenchmarkTestAnotherBuffer(b *testing.B) {
	b.ReportAllocs()
	var builder strings.Builder // Create once

	for i := 0; i < b.N; i++ {
		var flags Flags
		flags.Set(0)
		flags.Set(1)
		flags.Clear(0)
		builder.Reset() // Reset builder for reuse
		builder.WriteString("00000001")
	}
}

func BenchmarkSprintf(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var flags Flags
		flags.Set(0)
		flags.Set(1)
		flags.Clear(0)
		_ = fmt.Sprintf("%08b", 1)
	}
}

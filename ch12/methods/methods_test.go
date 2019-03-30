package methods

import (
	"strings"
	"time"
)

func PrintDuration() {
	Print(time.Hour)
}

func ExamplePrintDuration() {
	PrintDuration()
	// Output:
	// type time.Duration
	// func (time.Duration) Hours() float64
	// func (time.Duration) Minutes() float64
	// func (time.Duration) Nanoseconds() int64
	// func (time.Duration) Round(time.Duration) time.Duration
	// func (time.Duration) Seconds() float64
	// func (time.Duration) String() string
	// func (time.Duration) Truncate(time.Duration) time.Duration
}

func PrintReplacer() {
	Print(new(strings.Replacer))
}

func ExamplePrintReplacer() {
	PrintReplacer()
	// Output:
	// type *strings.Replacer
	// func (*strings.Replacer) Replace(string) string
	// func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
}

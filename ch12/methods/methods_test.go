package methods_test

import (
	"strings"
	"time"

	"gopl/ch12/methods"
)

func ExamplePrintDuration() {
	methods.Print(time.Hour)
	// Output:
	// type time.Duration
	// func (time.Duration) Hours() float64
	// func (time.Duration) Minutes() float64
	// func (time.Duration) Nanoseconds() int64
	// func (time.Duration) Seconds() float64
	// func (time.Duration) String() string
}

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
	// Output:
	// type *strings.Replacer
	// func (*strings.Replacer) Replace(string) string
	// func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
}

/*
//!+output
methods.Print(time.Hour)
// Output:
// type time.Duration
// func (time.Duration) Hours() float64
// func (time.Duration) Minutes() float64
// func (time.Duration) Nanoseconds() int64
// func (time.Duration) Seconds() float64
// func (time.Duration) String() string
methods.Print(new(strings.Replacer))
// Output:
// type *strings.Replacer
// func (*strings.Replacer) Replace(string) string
// func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
//!-output
*/

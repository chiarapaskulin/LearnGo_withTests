package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat1(t *testing.T) {
	repeated := Repeat1("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

/*
BENCHMARKS
go test -bench=.
OUTPUT:
	goos: darwin
	goarch: amd64
	pkg: github.com/chiarapaskulin/Iteration
	cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
	BenchmarkRepeat-8        8300698               141.1 ns/op
	PASS
	ok      github.com/chiarapaskulin/Iteration     1.583s
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat1("a")
	}
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat2() {
	repeated := Repeat2("b", 4)
	fmt.Println(repeated)
	// Output: bbbb
}

// Write a test using a function from the strings package
func TestRepeat3(t *testing.T) {
	repeated := Repeat2("a", 10)
	expected := strings.Repeat("a", 10)

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

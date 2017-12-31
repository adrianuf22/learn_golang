package matematica

// from: https://www.golang-book.com/books/intro/12

import (
	"fmt"
	"testing" // Package for tests, provide many recources including struct testing.T for errors and notifications
)

/*
	Structs for test Soma
	Like a data provider
*/
type somaPair struct {
	values   []int // Values to evaluate
	expected int   // Result of the evaluation
}

var somaTests = []somaPair{ // An array of somaPair for many tests
	/* Scenario 1 */ {[]int{1, 2}, 3}, // values[], expected
	/* Scenario 2 */ {[]int{0, 15}, 15},
	/* Scenario 3 */ {[]int{46, -10}, 36},
	/* Scenario 4 */ {[]int{-15, -30}, -45},
}

// To create a test with many scenarios (Using struct somaTest)
func TestSomaWithSamples(t *testing.T) {
	for _, pair := range somaTests {
		result := Soma(pair.values[0], pair.values[1])
		if result != pair.expected {
			t.Error("Expected ", pair.expected, "got ", result)
		}
	}
}

// Example can be used for tests and also will document exemples of usage of method
// Test and Documentation
func ExampleSoma() {
	fmt.Println(Soma(3, 4))
	fmt.Println(Soma(34, 47))
	// output:
	// 7
	// 81
}

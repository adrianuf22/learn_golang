package matematica

import (
	"fmt"
	"testing"
)

// A simple unit test
func TestSubtrai(t *testing.T) {
	expectedResult := 3
	result := Subtrai(4, 1)

	if result != expectedResult {
		t.Error("Expected ", expectedResult, "got ", result)
	}
}

func ExampleSubtrai() {
	fmt.Println(Subtrai(1, 2))
	fmt.Println(Subtrai(100, 34))
	// output:
	// -1
	// 66
}

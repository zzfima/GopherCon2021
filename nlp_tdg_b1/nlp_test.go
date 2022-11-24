package gophercon2021

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStem(t *testing.T) {
	w := Stem("threading")
	require.Equal(t, "thread", w)
}

func TestTokenize(t *testing.T) {
	tokens := Tokenize("I am going see cars")
	require.Equal(t, []string{"i", "am", "go", "see", "car"}, tokens)
}

func ExampleTokenize() {
	tokens := Tokenize("I am going see cars")
	fmt.Println(tokens)

	//Output:
	//[i am go see car]
}

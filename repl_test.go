package main
import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "  hello world from go lang ",
			expected: []string{"hello", "world", "from", "go", "lang"},
		},


		{
			input:    "  Hi Alkesh, what's up?",
			expected: []string{"hi", "alkesh,", "what's", "up?"},
		},

	}

	for _, c := range cases {

	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Expected %d, but got %d", len(c.expected),len(actual))
	}

	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("Expected %s, but got %s", expectedWord, word)
		}
	}

   }
}

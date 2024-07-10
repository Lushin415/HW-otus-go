package wordcounter

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "Hello, world! Hello.",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			input: "Test our program. Test Fast. Test Good!",
			expected: map[string]int{
				"test":    3,
				"our":     1,
				"program": 1,
				"fast":    1,
				"good":    1,
			},
		},
		{
			input: "123 123 123!",
			expected: map[string]int{
				"123": 3,
			},
		},
		{
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := CountWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

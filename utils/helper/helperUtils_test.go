package helper

import "testing"

func TestContainsString(t *testing.T) {
	playerNames := []string{"Mason Mount", "Nathaniel Chalobah", "Jasper Cillessen", "Daniel Drinkwater"}
	pairs := map[string]bool{
		"Lionel Messi":       false,
		"Nathaniel Chalobah": true,
	}
	for input, expected := range pairs {
		actual := ContainsString(playerNames, input)
		if actual != expected {
			t.Error(
				"For", input,
				"expected", expected,
				"got", actual,
			)
		}
	}
}
func TestAppendIfMissingString(t *testing.T) {
	playerNames := []string{"Fred", "Sebastian Rudy", "Lionel Messi", "Marcelo", "Chumi"}
	pairs := map[string][]string{
		"Jasper Cillessen": {"Fred", "Sebastian Rudy", "Lionel Messi", "Marcelo", "Chumi", "Jasper Cillessen"},
		"Sebastian Rudy":   {"Fred", "Sebastian Rudy", "Lionel Messi", "Marcelo", "Chumi"},
	}
	for input, expected := range pairs {
		actual := AppendIfMissingString(playerNames, input)
		if len(actual) != len(expected) {
			t.Error(
				"For", input,
				"expected", expected,
				"got", actual,
			)
		}
		for i, playerName := range expected {
			if playerName != actual[i] {
				t.Error(
					"For", input,
					"expected", expected,
					"got", actual,
				)
			}
		}
	}
}

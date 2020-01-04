package main

import "testing"

func TestParseCode(t *testing.T) {
	actual := parseCode(3)
	expected := [4]int{0, 0, 0, 3}

	if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
	}

	actual = parseCode(1100)
	expected = [4]int{0, 1, 1, 0}

	if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
	}

	actual = parseCode(101)
	expected = [4]int{0, 0, 1, 1}

	if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
	}

	actual = parseCode(1001)
	expected = [4]int{0, 1, 0, 1}

	if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
	}
}
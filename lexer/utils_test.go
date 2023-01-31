package lexer

import "testing"

func TestIsLetter(t *testing.T) {
	tests := []struct {
		char     byte
		expected bool
	}{
		{'a', true},
		{'0', false},
	}

	for _, tt := range tests {
		got := isLetter(tt.char)
		if got != tt.expected {
			t.Errorf("expected %v,got %v", tt.expected, got)
		}
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		char     byte
		expected bool
	}{
		{'a', false},
		{'0', true},
	}

	for _, tt := range tests {
		got := isNumber(tt.char)
		if got != tt.expected {
			t.Errorf("expected %v,got %v", tt.expected, got)
		}
	}
}

func TestReadNumber(t *testing.T) {
	input := `256mb`
	l := New(input)

	want := "256"
	got := l.readNumber()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

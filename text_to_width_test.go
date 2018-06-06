package texttowidth

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		Text   string
		Width  int
		Expect string
	}{
		{"My name is John Doe", 30, "My name is John Doe"},
		{"This week in the news, MS just bought Github.", 20, "This week in the\nnews, MS just\nbought Github."},
		{"This Week in the longwordthatexceedsthelinelimit has just bought Github.", 20, "This week in the\nlongwordthatexceeds\nthelinelimit has\njust bought Github."},
		{"This Week in the news, longwordthatexceedsthelinelimit has just bought Github.", 20, "This week in the\nnews,\nlongwordthatexceeds\nthelinelimit has\njust bought Github."},
	}

	for _, test := range tests {
		result := Format(test.Text, test.Width)
		if !strings.EqualFold(result, test.Expect) {
			t.Errorf("Expected\n%s\n\nBut got\n\n%s\n[end]", test.Expect, result)
		}
	}
}

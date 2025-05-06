package sanitize

import (
	"html"
	"strings"
)

func Sanitize(input string) string {
	input = strings.TrimSpace(input)
	return html.EscapeString(input)
}

package hello

import (
	"fmt"
	"strings"
)

func Hello(names []string) string {
	if len(names) <= 1 {
		return "Hello... no one :("
	}

	return fmt.Sprintf("Hello - %s!", strings.Join(names, ", "))
}

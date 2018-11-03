package pretty

import (
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {
	s := struct {
		i int
		s string
		m map[string]string
	}{
		i: 100,
		s: "test",
		m: map[string]string{
			"a": "b",
		},
	}
	fmt.Print(Sprint(s))
}

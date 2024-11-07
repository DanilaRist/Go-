package sprint

import (
	"fmt"
)

func Pairs() string {
	var pairs string

	for i := 0; i < 99; i++ {
		for j := i + 1; j <= 99; j++ {
			pairs += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}

	if len(pairs) > 2 {
		pairs = pairs[:len(pairs)-2]
	}

	return pairs
}
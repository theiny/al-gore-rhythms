package binary

import (
	"fmt"
	"strconv"
)

func BinaryGap(n int) int {
	b := strconv.FormatInt(int64(n), 2)
	var highest int
	var gap int

	for i := 0; i < len(b); i++ {
		if string(b[i]) == "0" {
			gap++
			fmt.Printf("Incrementing gap count: %d\n", gap)
		}

		if string(b[i]) == "1" {
			fmt.Printf("Gap: %d Highest: %d\n", gap, highest)
			if gap > highest {
				highest = gap
			}
			gap = 0
		}
	}

	fmt.Printf("Highest: %d\n", highest)
	return highest
}

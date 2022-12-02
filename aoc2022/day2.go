package aoc2022

import (
	"fmt"
	"strings"
)

func DayTwo(input string) {
	s := 0
	ps := 0
	for _, l := range strings.Split(input, "\n") {
		spl := strings.Split(l, " ")
		e := int(spl[0][0]) - 65
		m := int(spl[1][0]) - 88
		//part one
		s += m + 1
		if (e+1)%3 == m {
			s += 6
		} else if e == m {
			s += 3
		}
		//part two stuff
		if m == 1 {
			ps += 4 + e
		} else if m == 2 {
			ps += 7 + (e+1)%3
		} else {
			if e == 0 {
				ps += 3
			} else {
				ps += (e - 1) + 1
			}
		}
	}
	fmt.Printf("Part One Solution: %d\n", s)
	fmt.Printf("Part Two Solution: %d", ps)
}

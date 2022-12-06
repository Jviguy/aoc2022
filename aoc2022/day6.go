package aoc2022

import (
	"fmt"
	"strings"
)

func DaySix(input string) {
	p1 := false
	p2 := false
	for i := range input {
		if i > 3 {
			s := i - 4
			l := input[s:i]
			m := true
			for _, c := range l {
				if strings.Count(l, string(c)) > 1 {
					m = false
					break
				}
			}
			if m && !p1 {
				fmt.Printf("Part One Solution: %d\n", i)
				p1 = true
			}
		}
		if i > 13 {
			s := i - 14
			l := input[s:i]
			m := true
			for _, c := range l {
				if strings.Count(l, string(c)) > 1 {
					m = false
					break
				}
			}
			if m && !p2 {
				fmt.Printf("Part Two Solution: %d\n", i)
				p2 = true
			}
		}
	}
}

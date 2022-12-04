package aoc2022

import (
	"fmt"
	"strings"
)

func DayThree(input string) {
	t := 0
	s := [3]string{}
	pt := 0
	for i, r := range strings.Split(input, "\r\n") {
		//part one
		first := r[:len(r)/2]
		second := r[len(r)/2:]
		c := rune(first[strings.IndexFunc(first, func(r rune) bool {
			return strings.ContainsRune(second, r)
		})])
		p := int(c) - 96
		if p <= 0 {
			p += 58
		}
		t += p
		// part two
		s[i%3] = r
		if (i+1)%3 == 0 && i != 0 {
			c := rune(s[0][strings.IndexFunc(s[0], func(r rune) bool {
				return strings.ContainsRune(s[1], r) && strings.ContainsRune(s[2], r)
			})])
			p := int(c) - 96
			if p <= 0 {
				p += 58
			}
			pt += p
		}
	}
	fmt.Printf("Part One Solution: %d\n", t)
	fmt.Printf("Part Two Solution: %d\n", pt)
}

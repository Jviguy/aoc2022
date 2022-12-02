package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func DayOne(input string) {
	t := [3]int{}
	c := 0
	for _, l := range strings.Split(input, "\r\n") {
		if l == "" {
			for j := 2; j >= 0; j-- {
				if c > t[j] {
					t[j] = c
					break
				}
			}
			c = 0
		} else {
			a, _ := strconv.Atoi(l)
			c += a
		}
	}
	m := 0
	for _, n := range t {
		m += n
	}
	fmt.Printf("Part One Solution: %d\n", t[len(t)-1])
	fmt.Printf("Part Two Solution: %d", m)
}

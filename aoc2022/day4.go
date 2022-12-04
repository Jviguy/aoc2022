package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
)

func DayFour(input string) {
	t := 0
	pt := 0
	for _, l := range strings.Split(input, "\r\n") {
		pairs := strings.Split(l, ",")
		start, end := parseRange(pairs[0])
		start1, end1 := parseRange(pairs[1])
		if (start >= start1 && start <= end1 && end <= end1) || (start1 >= start && start1 <= end && end1 <= end) {
			t++
		}
		if start <= end1 && start1 <= end {
			pt++
		}
	}
	fmt.Printf("Part One Solution: %d\n", t)
	fmt.Printf("Part Two Solution: %d\n", pt)
}

func parseRange(pair string) (int, int) {
	s := strings.Split(pair, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	return start, end
}

package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DayFive(input string) {
	lines := strings.Split(input, "\r\n")
	m := make(map[int]int, 0)
	stacks := make([][]string, 0, len(m))
	stacks2 := make([][]string, 0, len(m))
	for _, line := range lines {
		if !strings.Contains(line, "[") {
			j := 0
			for i, c := range line {
				if unicode.IsDigit(c) {
					m[i] = j
					stacks = append(stacks, []string{})
					stacks2 = append(stacks2, []string{})
					j++
				}
			}
			break
		}
	}
	i := 0
	for strings.Count(lines[i], "[") > 0 {
		line := lines[i]
		for x, c := range line {
			if unicode.IsLetter(c) {
				j := m[x]
				stacks[j] = append(stacks[j], string(c))
				stacks2[j] = append(stacks2[j], string(c))
			}
		}
		i++
	}
	for _, command := range lines[i+2:] {
		count, err := strconv.Atoi(command[5:strings.Index(command, " from")])
		if err != nil {
			panic(err)
		}
		o := 0
		if count > 9 {
			o += 1
		}
		from, err := strconv.Atoi(string(command[12+o]))
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(string(command[17+o]))
		if err != nil {
			panic(err)
		}
		l := stacks[from-1][:count]
		stacks[from-1] = stacks[from-1][count:]
		for i := 0; i < len(l); i++ {
			stacks[to-1] = append([]string{l[i]}, stacks[to-1]...)
		}
		l = stacks2[from-1][:count]
		l = append([]string{}, l...)
		stacks2[from-1] = stacks2[from-1][count:]
		stacks2[to-1] = append(l, stacks2[to-1]...)
	}
	x := ""
	for _, s := range stacks {
		x += s[0]
	}
	y := ""
	for _, s := range stacks2 {
		y += s[0]
	}
	fmt.Printf("Part One Solution: %s\n", x)
	fmt.Printf("Part Two Solution: %s\n", y)
}

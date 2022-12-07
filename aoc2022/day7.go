package aoc2022

import (
	"fmt"
	"strings"
)

func DaySeven(input string) {
	cd := ""
	for _, l := range strings.Split(input, "\n") {
		if l[0] == '$' {
			spl := strings.Split(l, " ")
			args := spl[2:]
			command := spl[1]
			if command == "cd" {
				cd = args[0]
			} else if command == "ls" {
				
			}
		} else {
			fmt.Println("DATA: " + l)
		}
	}
}
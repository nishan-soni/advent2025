package day1

import (
	"fmt"
	"strconv"

	"github.com/nishan-soni/advent2025/utils"
)

func Solve() {
	lines, err := utils.GetFileLines("inputs/day1.txt")

	if err != nil {
		panic(err)
	}

	curr := 50
	ans := 0

	for line := range lines {
		count, _ := strconv.Atoi(line[1:])
		direction := line[:1]

		if direction == "L" {
			curr = (curr - count) % 100
		} else {
			curr = (curr + count) % 100
		}

		if curr == 0 {
			ans++
		}

	}
	fmt.Println(ans)
}

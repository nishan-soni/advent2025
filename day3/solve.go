package day3

import (
	"fmt"
	"strconv"

	"github.com/nishan-soni/advent2025/utils"
)

func Solve() {

	lines, err := utils.GetFileLines("inputs/day3.txt")

	if err != nil {
		panic(err)
	}

	ans := 0
	for line := range lines {

		var maxVolt rune
		var maxVolt2 rune
		maxCharIdx := 0

		for idx, volt := range line {
			if idx == len(line)-1 {
				break
			}
			if volt > maxVolt {
				maxVolt = volt
				maxCharIdx = idx
			}
		}

		for _, volt := range line[maxCharIdx+1:] {
			if volt > maxVolt2 {
				maxVolt2 = volt
			}
		}

		converted, _ := strconv.Atoi(string(maxVolt) + string(maxVolt2))
		ans += converted

	}

	fmt.Println(ans)
}

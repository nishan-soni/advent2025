package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nishan-soni/advent2025/utils"
)

func Solve() {
	lines, err := utils.GetFileLines("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	line := <-lines

	ranges := strings.Split(line, ",")

	ans := 0

	for _, r := range ranges {
		r_split := strings.Split(r, "-")
		start, _ := strconv.Atoi(r_split[0])
		end, _ := strconv.Atoi(r_split[1])

		for num := start; num <= end; num++ {
			num_str := strconv.Itoa(num)
			if num_str[:len(num_str)/2] == num_str[len(num_str)/2:] {
				ans += num
			}
		}
	}

	fmt.Println(ans)
}

package day4

import (
	"fmt"

	"github.com/nishan-soni/advent2025/utils"
)

func Solve() {

	lines, err := utils.GetFileLines("inputs/day4.txt")

	if err != nil {
		panic(err)
	}

	var grid [][]rune

	for line := range lines {
		grid = append(grid, []rune(line))
	}

	ans := 0

	rows := len(grid)
	cols := len(grid[0])

	direct := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '.' {
				continue
			}
			nbr_count := 0
			for _, pair := range direct {
				dr := pair[0]
				dc := pair[1]
				nr := r + dr
				nc := c + dc

				if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
					continue
				}

				if grid[nr][nc] == '@' {
					nbr_count++
				}
			}
			if nbr_count < 4 {
				ans++
			}

		}
	}

	fmt.Println(ans)
}

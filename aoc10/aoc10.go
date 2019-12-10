package aoc10

import(
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
)

func getSpace() [][]bool {
	lines := fileio.GetLinesFromFile("aoc10")
	var field [][]bool
	for i,line := range(lines) {
		field = append(field, make([]bool,0))
		for _,char := range(line) {
			field[i] = append(field[i],char == '#')
		}
	}

	return field
}

func countDetectableAteroids(field [][]bool, x ,y int ) int {
	return 0
}

func countAsteroidsInRow(field [][]bool, dx, dy int ) int {
	return 0
}

func findMaxAsteroids(field [][]bool) int {
	return 0
}

func Aoc10Main() {
	field := getSpace()
	fmt.Println(field)
}
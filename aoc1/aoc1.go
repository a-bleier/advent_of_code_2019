package aoc1

import(
	"log"
	"strconv"
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
)

func getNumberList (lines []string ) []int {
	var numbers []int

	for _, line := range(lines){
		number,err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func Aoc1Main() {
	lines := fileio.GetLinesFromFile("aoc1")
	numbers := getNumberList(lines)
	sum := 0

	for _, num := range(numbers){
		sum += num/3-2
	}

	fmt.Println("First star: " + strconv.Itoa(sum))

	sum = 0

	for _, num := range(numbers){
		fuel := num/3-2
		for fuel >= 0 {
			sum += fuel
			fuel = fuel/3-2
			
		}
	}

	fmt.Println("Second star: " + strconv.Itoa(sum))
}
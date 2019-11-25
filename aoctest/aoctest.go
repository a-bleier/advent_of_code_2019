package aoctest

import(
	"github.com/a-bleier/aoc/fileio"
	"fmt"
	"log"
	"strconv"
	"sort"
)

//First of December of the prev year

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




func AoctestMain() {

	lines := fileio.GetLinesFromFile("aoctest")
	numbers := getNumberList(lines)
	sum:= 0
	for _,num := range(numbers) {
		sum += num
	}
	fmt.Println(sum)

	var freqs []int
	sum = 0
	freqs = append(freqs,0)
	for true {
		for _,num := range(numbers){
			sum += num
			i := sort.SearchInts(freqs,sum)
			if i < len(freqs) && freqs[i] == sum{
				fmt.Printf("first %d appears twice\n", sum)
				return
				
			} else {
				freqs = append(freqs,0)
				copy(freqs[i+1:],freqs[i:])
				freqs[i] = sum
			}
		}

	}
	


}
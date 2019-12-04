package aoc4

import(
	"github.com/a-bleier/aoc2019/fileio"
	"strings"
	"fmt"
	"strconv"
	"log"
)

func getBounds(line string) (int,int) {
	bounds := strings.Split(line,"-")
	lowerBound, errL := strconv.Atoi(bounds[0])

	if errL != nil {
		log.Fatal(errL)
	}

	upperBound, errU := strconv.Atoi(bounds[1])
	if errU != nil {
		log.Fatal(errU)
	}

	return lowerBound, upperBound
	
}

func countPasswords(l, u int) (int, int) {

	counter, counter2 := 0, 0

	for i := l; i < u+1; i++ {
		if isPassword(i){
			counter++
		}

		if isPassword2(i) {
			counter2++
			fmt.Println(i)
		}
	}

	return counter, counter2
}

func isPassword(num int) bool {
	var hasDouble bool = false
	//var doubleDigit int 

	prevDigit := 11

	for num > 0 {

		digit := num % 10

		if digit > prevDigit {
			return false
		}

		hasDouble = hasDouble || prevDigit == digit

		prevDigit = digit

		num = num / 10
	}

	return hasDouble
}

func isPassword2(num int) bool {
	var hasGlobalDouble bool
	var hasLocalDouble bool

	prevDigit := 11
	counter := 1

	for num > 0 {

		digit := num % 10

		if digit > prevDigit {
			return false
		}

		if prevDigit == digit {
			counter++
			hasLocalDouble = counter == 2
		} else {
			hasGlobalDouble = hasGlobalDouble || hasLocalDouble
			counter = 1
		}

		prevDigit = digit

		num = num / 10
	}

	hasGlobalDouble = hasGlobalDouble || hasLocalDouble

	return hasGlobalDouble
}

func Aoc4Main(){
	lines := fileio.GetLinesFromFile("aoc4")
	l, u := getBounds(lines[0])

	count, count2 := countPasswords(l,u)

	fmt.Println(l,u)
	fmt.Println(count, count2)




}


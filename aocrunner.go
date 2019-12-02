package main

/*

package structur
-aoc2019
	aoc1
		code
		inputfile
	aoc2
		code
		inputfile
	...
	aocrunner (main; can start all aoc solver)


*/

import(
	"github.com/a-bleier/aoc2019/fileio"
	"github.com/a-bleier/aoc2019/aoc1"
	"github.com/a-bleier/aoc2019/aoc2"
	"fmt"
)

func main(){
	runAoc2()
}

func runAoc1(){
	aoc1.Aoc1Main()
}

func runAoc2(){
	aoc2.Aoc2Main()
}

func testFileio() {
	lines := fileio.GetLinesFromFile("aoctest")

	for _,line := range(lines){
		fmt.Println(line)
	}
}
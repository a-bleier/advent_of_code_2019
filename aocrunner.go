package main

/*

package structur
-aoc
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
	"github.com/a-bleier/aoc/fileio"
	"github.com/a-bleier/aoc/aoctest"
	"fmt"
)

func main(){
	aoctest.AoctestMain()
}

func testFileio() {
	lines := fileio.GetLinesFromFile("aoctest")

	for _,line := range(lines){
		fmt.Println(line)
	}
}
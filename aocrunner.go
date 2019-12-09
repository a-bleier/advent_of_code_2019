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
	"github.com/a-bleier/aoc2019/aoc3"
	"github.com/a-bleier/aoc2019/aoc4"
	"github.com/a-bleier/aoc2019/aoc5"
	"github.com/a-bleier/aoc2019/aoc6"
	"github.com/a-bleier/aoc2019/aoc7"
	"fmt"
)

func main(){
	runAoc7()
}



func runAoc1(){
	aoc1.Aoc1Main()
}

func runAoc2(){
	aoc2.Aoc2Main()
}

func runAoc3(){
	aoc3.Aoc3Main()
}

func runAoc4() {
	aoc4.Aoc4Main()
}

func runAoc5() {
	aoc5.Aoc5Main()
}

func runAoc6() {
	aoc6.Aoc6Main()
}

func runAoc7() {
	aoc7.Aoc7Main()
}

func testFileio() {
	lines := fileio.GetLinesFromFile("aoctest")

	for _,line := range(lines){
		fmt.Println(line)
	}
}
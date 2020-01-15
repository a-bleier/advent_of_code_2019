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
	"github.com/a-bleier/aoc2019/aoc8"
	"github.com/a-bleier/aoc2019/aoc9"
	"github.com/a-bleier/aoc2019/aoc10"
	"github.com/a-bleier/aoc2019/aoc11"
	"fmt"
)

func main(){
	runAoc10()
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

func runAoc8() {
	aoc8.Aoc8Main()
}

func runAoc9() {
	aoc9.Aoc9Main()
}

func runAoc10() {
	aoc10.Aoc10Main()
}

func runAoc11() {
	aoc11.Aoc11Main()
}

func testFileio() {
	lines := fileio.GetLinesFromFile("aoctest")

	for _,line := range(lines){
		fmt.Println(line)
	}
}
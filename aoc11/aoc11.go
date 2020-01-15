package aoc11

import(
	"github.com/a-bleier/aoc2019/fileio"
	"github.com/a-bleier/aoc2019/aoc9"
	"fmt"
)

const(
	UP = 0
	DOWN = 2
	LEFT = 3
	RIGHT = 1
	BLACK = 0
	WHITE = 1

)

type paneltable map[panel]int

type robot struct {
	brain aoc9.Computer
}

func (r *robot) run (){
	r.brain.Run()
}

type panel struct {
	x,y int
}

func makeRobot(in func()int, out func(int)) robot{
	program := fileio.GetLinesFromFile("aoc11")

	fmt.Println(program)
	brain := aoc9.NewComputer(aoc9.GetCode(program[0]), in, out)

	return robot{brain : brain}
}

func isPaintedAgain(panels paneltable, p panel) (bool) {
	_,ok := panels[p]
	return ok
	
}

func modLikePython(d, m int) int {
	var res int = d % m
	if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
	   return res + m
	}
	return res
 }


func run (){

	panels := make(paneltable)
	firstPanels := make(paneltable)
	pan := panel{0,0}
	paintingTime := true
	orientation := UP
	counter := 0



	out := func(x int) {
		//isBlack,_ := panelIsBlack(blackPanels, pan)
		//fmt.Println("Panel",pan," is black", isBlack, "x", x)
		if paintingTime {
			panels[pan] = x
			fmt.Println("painting" ,x)
			if !isPaintedAgain(firstPanels,pan){
				firstPanels[pan] = 1
				counter++
				fmt.Println("is blacked the first time", counter)
				
			}
		} else {
			var d int
			if x == 0 {
				d = -1
			} else {
				d = 1
			}
			//fmt.Println(orientation, d)
			orientation = modLikePython((orientation + d) , 4)
			fmt.Println(orientation)
			//fmt.Println(orientation)
			if orientation == UP {
				pan.y++
			}  else if orientation == DOWN {
				pan.y--
			} else if orientation == LEFT {
				pan.x--
			} else if orientation == RIGHT {
				pan.x++
			}

			fmt.Println("moving to", pan)
		}
		
		paintingTime = !paintingTime
	}

	in := func()int {
		col, ok := panels[pan]
		if ok {
			fmt.Println("visiting", pan, "color", col)
			return col
			
		} else {
			fmt.Println("visiting", pan, "color", 0)
			return 0
		}
	}

	rob := makeRobot(in,out)

	rob.run()

	
	fmt.Println(counter)
	fmt.Println(panels)


}

func Aoc11Main(){
	run()
}
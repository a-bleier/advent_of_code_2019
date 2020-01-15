package aoc10

import(
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
	"math"
)

type coord struct {
	x,y int
}


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

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func countDetectableAteroids(field [][]bool, x ,y int ) int {
	counter := 0
	xfactor, yfactor := 1,1
	diffy, diffx := 0,0
	startwert := 0

	for i := 0; i < 4; i++ {

		if i == 0 {
			xfactor, yfactor = 1,1
			diffy = len(field)-1-y
			diffx = len(field[0])-1-x
			startwert = 1
		} else if i == 1{
			xfactor, yfactor = -1,1
			diffy = len(field)-1-y
			diffx = x
			startwert = 0
		} else if i == 2{
			xfactor, yfactor = 1,-1
			diffy = y
			diffx = len(field[0])-1-x
			startwert = 0
		} else if i == 3{
			xfactor, yfactor = -1,-1
			diffy = y
			diffx = x
			startwert = -1
		}

		for dy := startwert; abs(dy) <= diffy; dy += yfactor{ 
			for dx := startwert; abs(dx) <= diffx; dx += xfactor {
	
				//fmt.Println("me here")
				euc := euclid(abs(dx),abs(dy))
				if  !(euc == 1 || euc == 0) { //it is a scalar product
					continue
				}
				if  dx == 0 && dy == 0 {
					continue
				}
			

				xtemp, ytemp := x+dx,y+dy
	
				
				
				for xtemp >= 0 && xtemp < len(field[0]) && ytemp >= 0 && ytemp < len(field){
					fmt.Printf("looking up %d %d ", ytemp, xtemp)
					if field[ytemp][xtemp] {
						fmt.Printf("found one\n")
						counter++
						break
					}
					fmt.Printf("\n")
					xtemp, ytemp = xtemp+dx,ytemp+dy
				}
	
				
			
			}
	
		}

	}	
	return counter
}

func euclid(a, b int) int {
	for b != 0 {
		h := a % b
		a = b
		b = h
	}
	return a
}

func countAsteroidsInRow(field [][]bool, dx, dy int ) int {
	return 0
}

func findMaxAsteroids(field [][]bool) int {
	var maxNum int
	var asteroids []coord

	for y,row := range(field) {
		for x,cell := range(row) {
			if cell {
				asteroids = append(asteroids,coord{x,y})
			}
		}
	}
	x,y := 0,0
	for _,ast := range(asteroids) {
		fmt.Println("checking", ast.y, ast.x)
		num := countDetectableAteroids(field, ast.x, ast.y)
		fmt.Println("num: ", num)
		if num > maxNum {
			maxNum = num
			x = ast.x
			y = ast.y
			
		}
	}
	fmt.Println("x: ", x,"y: ", y)
	return maxNum
}

func Aoc10Main() {
	field := getSpace()
	fmt.Println(field)
	fmt.Println("Max number of asteroids detected on the best location: ")
	fmt.Println(findMaxAsteroids(field))
}
package aoc3

import(
	"github.com/a-bleier/aoc2019/fileio"
	"strings"
	"fmt"
	"strconv"
	"log"
)

type spot struct {
	x int
	y int
}

func getWire(line string) []string{
	return strings.Split(line,",")
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func calculateIntersections(wire1, wire2 []string) ([]spot, int){
	var intersections []spot
	var spotH1, spotH2, spotV1, spotV2 spot

	var steps int

	wirePath1 := doWiring(wire1)
	wirePath2 := doWiring(wire2)

	fmt.Println(wirePath1, wirePath2)
	var steps1 int
	var steps2 int
	for index1 := range(wirePath1) {

		if index1 != 0 {
			fmt.Println("1--", abs(wirePath1[index1].x - wirePath1[index1-1].x) + abs(wirePath1[index1].y - wirePath1[index1-1].y))
			steps1 += abs(wirePath1[index1].x - wirePath1[index1-1].x) + abs(wirePath1[index1].y - wirePath1[index1-1].y)
		}
		steps2 = 0

		for index2 := range(wirePath2){
			if index1 == 0 || index2 == 0 {
				continue
			}
			
			steps2 += abs(wirePath2[index2].x - wirePath2[index2-1].x) + abs(wirePath2[index2].y - wirePath2[index2-1].y)

			if wirePath1[index1].x == wirePath1[index1-1].x {
				spotV1, spotV2 = wirePath1[index1], wirePath1[index1-1]

				if wirePath2[index2].y == wirePath2[index2-1].y {
					spotH1, spotH2 = wirePath2[index2], wirePath2[index2-1]
				} else {
					continue
				}
			} else if wirePath1[index1].y == wirePath1[index1-1].y {
				spotH1, spotH2 = wirePath1[index1], wirePath1[index1-1]

				if wirePath2[index2].x == wirePath2[index2-1].x {
					spotV1, spotV2 = wirePath2[index2], wirePath2[index2-1]
				} else {
					continue
				}
			} else {
				continue
			}
			
			if areCrossing(spotH1, spotH2, spotV1, spotV2) {

				steps2 -= abs(wirePath2[index2].x - spotV2.x) + abs(wirePath2[index2].y - spotH2.y)
				steps1 -= abs(wirePath1[index1].x - spotV2.x) + abs(wirePath1[index1].y - spotH2.y)

				intersections = append(intersections,spot{spotV2.x, spotH2.y})
				if steps1 + steps2 < steps || steps == 0{
					
					steps = steps1 + steps2
				}
			} 


		}
	}

	return intersections, steps

}

func areCrossing(spotH1, spotH2, spotV1, spotV2 spot) bool {
	return ((spotH1.y < spotV1.y && spotH1.y > spotV2.y) || (spotH1.y < spotV2.y && spotH1.y > spotV1.y) )&& ((
		spotV1.x < spotH1.x && spotV1.x > spotH2.x) || (spotV1.x < spotH2.x && spotV1.x > spotH1.x ))
}

func doWiring(wire []string) []spot{
	var wirePath []spot
	currSpot := spot{0,0}

	wirePath = append(wirePath,spot{0,0})

	for _,s := range(wire) {
		value, err := strconv.Atoi(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		if s[0] == 'R' {
			currSpot.x += value
		} else if s[0] == 'L' {
			currSpot.x -= value
		} else if s[0] == 'U' {
			currSpot.y += value
		} else if s[0] == 'D' {
			currSpot.y -= value
		}

		wirePath = append(wirePath,spot{currSpot.x, currSpot.y})
	}

	return wirePath
}

func Aoc3Main(){
	lines := fileio.GetLinesFromFile("aoc3")

	wire1 := getWire(lines[0])
	wire2 := getWire(lines[1])

	var manhattanDist uint

	manhattanDist = 0

	intersections, minsteps := calculateIntersections(wire1,wire2)

	for _,inter := range(intersections) {
		sum := (uint) (inter.x) + (uint) (inter.y)
		//fmt.Println(sum)
		if sum == 0 {
			continue
		}
		if sum < manhattanDist || manhattanDist == 0 {
			manhattanDist = sum
		}
	}



	//fmt.Println("intersections at", intersections)

	fmt.Println("Manhattan Distance", manhattanDist)

	fmt.Println("minimal number of steps to next intersection", minsteps)




}


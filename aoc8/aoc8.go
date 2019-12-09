package aoc8

import(	
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
	"strconv"
	"log"
)

type layer []int

func getLayers(line string, layerSize int) []layer{
	var layers []layer
	j := -1
	for i := 0; i < len(line); i++ {
		if i % layerSize == 0 {
			layers = append(layers, make(layer, layerSize))
			j++
		}
		var err error
		layers[j][i%layerSize],err = strconv.Atoi((string) (line[i]))

		if err != nil {
			log.Fatal(err)
		}

	}
	return layers
}

func checkSum(layers []layer, layerSize int) int{
	min0Digits,index := layerSize, -1

	for i,l := range(layers){
		fmt.Println(l)
		counter := 0
		for _,x := range(l) {
			if x == 0 {
				counter++
			}
		}
		if counter < min0Digits {
			fmt.Println(counter)
			min0Digits = counter
			index = i
		}
	}
	count1 := 0
	count2 := 0
	for _,x := range(layers[index]){
		if x == 1 {
			count1++
		} else if x == 2 {
			count2++
		}
	}

	return count1*count2
}

func decode(layers []layer, layerSize int) layer {

	var lay layer
	for i:= 0; i < layerSize; i++ {
		j := 0
		for layers[j][i] == 2 {
			j++
		}
		lay = append(lay, layers[j][i])
	}

	return lay

}

func layerToString(lay layer) string {
	var out string

	for i, x := range(lay) {
		if i % 25 == 0 {
			out += "\n"
		}
		out += strconv.Itoa(x)
	}

	return out
}

func Aoc8Main(){
	layerSize := 25*6
	lines := fileio.GetLinesFromFile("aoc8")
	//fmt.Println(lines[0])
	layers := getLayers(lines[0],layerSize)
	x := checkSum(layers, layerSize)
	//fmt.Println(layers)
	fmt.Println(x)
	pic := decode(layers, layerSize)
	s := layerToString(pic)
	fmt.Printf(s)

}
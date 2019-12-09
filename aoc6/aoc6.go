package aoc6

import(
	"github.com/a-bleier/aoc2019/fileio"
	"fmt"
	"strings"
)

type satellite struct {
	name string
	flag bool
}

func getSystem(lines []string) (map[string]string, map[string][]string) {

	satSystem := make(map[string]string)
	for _,line := range(lines) {
		splittedLine := strings.Split(line,")")
		sat1Name := splittedLine[0]
		sat2Name := splittedLine[1]

		_, ok := satSystem[sat2Name]

		if !ok {
			satSystem[sat2Name] = sat1Name
		}
		
	}

	transferSystem := make(map[string][]string)
	for _,line := range(lines) {
		splittedLine := strings.Split(line,")")
		sat1Name := splittedLine[0]
		sat2Name := splittedLine[1]

		_, ok := transferSystem[sat2Name]

		if !ok{
			transferSystem[sat2Name] = make([]string,0)
		}
		transferSystem[sat2Name] = append(transferSystem[sat2Name], sat1Name)
		
		_, ok = transferSystem[sat1Name]

		if !ok{
			transferSystem[sat1Name] = make([]string,0)
		}
		transferSystem[sat1Name] = append(transferSystem[sat1Name], sat2Name)
		
	}

	return satSystem, transferSystem
}

func findOptimalTransfer(tSystem map[string][]string, prev string, start string, dest string) int {
	sats := tSystem[start]
	var length int = 1000000
	for _,sat := range(sats) {

		if prev == sat {
			continue
			
		}

		if sat == dest {
			return 1
		}
		path := findOptimalTransfer(tSystem,start,sat,dest)

		if path < length {
			length = path
		}
	}
	return length+1
}

func countOrbits(system map[string]string) int {
	counter := 0	
	for sat := range system {
		center, ok := system[sat]

		for ok {
			counter++
			center, ok = system[center]
		}
		
		
	}

	return counter
}

func Aoc6Main() {

	satSystem, transferSystem := getSystem(fileio.GetLinesFromFile("aoc6"))
	
	count := countOrbits(satSystem)
	
	fmt.Println("Number of (in-)direct orbitals")
	fmt.Println(count)

	hops := findOptimalTransfer(transferSystem,"YOU",satSystem["YOU"],satSystem["SAN"])

	fmt.Println("The smallest number of orbital hops is")
	fmt.Println(hops)


}
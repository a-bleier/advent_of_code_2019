package fileio

import (
	"os"
	"log"
	"bufio"
)
//buildPath builds a path to an input file for a specific aoc riddle
func buildPath(aocname string) string {

	return "/home/adrian/go/src/github.com/a-bleier/aoc/" + aocname + "/" + aocname + ".txt"
}

//GetLinesFromFile returns the lines from a file in an array
func GetLinesFromFile(aocname string) []string {

	var fileLines []string

	filename := buildPath(aocname)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fileLines = append(fileLines, s)

	}

	return fileLines

}
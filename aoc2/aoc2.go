package aoc2

import(
	"log"
	"strconv"
	"strings"
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
)

const HALT = 99
const ADD = 1
const MUL = 2
const OUTPUT = 19690720

func getCode (line string) []int {
	var codes []int
	codeStrings := strings.Split(line,",")

	for _,codeString := range(codeStrings){
		opcode, err := strconv.Atoi(codeString)
		if err != nil {
			log.Fatal(err)
		}
		codes = append(codes,opcode)
	}
	return codes
}

//returns the program state after execution
func runProgram(program []int) []int{
	pc := 0
	instruction := program[pc]

	for instruction != HALT {
		src1, src2, dest := program[pc+1], program[pc+2], program[pc+3]
		if instruction == ADD {
			program[dest] = program[src1] + program[src2]
		} else if instruction == MUL {
			program[dest] = program[src1] * program[src2]
		}

		pc += 4
		instruction = program[pc]
	}

	return program
}

func Aoc2Main() {
	programSaved := getCode(fileio.GetLinesFromFile("aoc2")[0])
	program := make([]int,len(programSaved))
	copy(program,programSaved)
	fmt.Println(program)
	program[1], program[2] = 12 , 2
	runProgram(program)
	fmt.Println("Star 1")
	fmt.Println("After execution: ")
	fmt.Println(program)

	for noun := 0; noun <= 99; noun++{
		for verb := 0; verb <= 99; verb++{
			program := make([]int,len(programSaved))
			copy(program,programSaved)
			program[1], program[2] = noun, verb
			runProgram(program)
			if program[0] == OUTPUT {
				fmt.Println("Second star ")
				fmt.Println("The noun-verb pair is ",noun,verb)
				break
			}
		}
	}

}
package aoc5



import(
	"log"
	"strconv"
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/a-bleier/aoc2019/fileio"
)

const HALT = 99
const ADD = 1
const MUL = 2
const IN = 3
const OUT = 4
const JT = 5
const JF = 6
const LT = 7
const EQ = 8
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
			
			instr := 0
			pm_instr := strconv.Itoa(instruction)

			size := len(pm_instr)

			

			for i := 0; i< (5 - size); i++{
				pm_instr = "0" + pm_instr
			}

			//fmt.Println(pm_instr)

			var op1, op2, dest int

			for index, char := range(pm_instr) {

				if index == 0 {
					if  (pc+3 < len(program)){
						dest = program[pc+3]
					}
				} else if index == 1 {
					if char == '0'{
						
						if  (program[pc+2] < len(program)){
							
							op2 = program[program[pc+2]]
						}
						
					} else if char == '1'{
						op2 = program[pc+2]
					}
				} else if index == 2 {
					if char == '0'{
						
						if (program[pc+1] < len(program)){
							op1 = program[program[pc+1]]
						}

					} else if char == '1'{
						op1 = program[pc+1]
					}

				} else if index == 3 {
					if pm_instr[index:] == "01" {
						instr = ADD
					} else if pm_instr[index:] == "02" {
						instr = MUL
					} else if pm_instr[index:] == "03" {
						instr = IN
					} else if pm_instr[index:] == "04" {
						instr = OUT
					} else if pm_instr[index:] == "05" {
						instr = JT
					} else if pm_instr[index:] == "06" {
						instr = JF
					} else if pm_instr[index:] == "07" {
						instr = LT
					} else if pm_instr[index:] == "08" {
						instr = EQ
					}
				}

			}
			
			//fmt.Println(op1,op2, dest)

			if instr == ADD {
				program[dest] = op1 + op2
				pc += 4
			} else if instr == MUL {
				program[dest] = op1 * op2
				pc += 4
			} else if instr == IN {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter number: ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimRight(text,"\n")
				num, err := strconv.Atoi(text)
	
				if err != nil {
					log.Fatal(err)
				}
				address := program[pc+1]
				program[address] = num
	
				pc += 2
			} else if instr == OUT {
				fmt.Println(op1)
				pc += 2
			} else if instr == JT {
				if op1 != 0 {
					pc = op2
				} else {pc += 3}
				
			} else if instr == JF {
				if op1 == 0 {
					pc = op2
				} else {pc += 3}

			} else if instr == LT {
				if op1 < op2 {
					program[dest] = 1
				} else {
					program[dest] = 0
				}
				pc += 4
			} else if instr == EQ {
				if op1 == op2 {
					program[dest] = 1
				} else {
					program[dest] = 0
				}
				pc += 4
			}

		instruction = program[pc]
	}

	return program
}

func Aoc5Main() {
	program := getCode(fileio.GetLinesFromFile("aoc5")[0])
	
	runProgram(program)

	//Answer for one star is 7566643 (remove all jump or compare cases)

}
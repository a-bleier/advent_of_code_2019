package aoc9



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
const BAAD = 9 //base address
const OUTPUT = 19690720

type inputFunc func() int
type outputFunc func(int)

type Computer struct {
	pc int
	mem map[int]int
	base int
	code int
	operands []int
	input inputFunc
	output outputFunc

}



func (comp *Computer) Run (){
	for comp.parseInstruction() {
		comp.exec()
		//fmt.Println(comp)
	}
}

func (comp *Computer) exec() {

	dest, op2, op1 := comp.operands[2], comp.operands[1], comp.operands[0]

	if comp.code == ADD {
		comp.mem[dest] = op1 + op2
		comp.pc += 4
	} else if comp.code == MUL {
		comp.mem[dest] = op1 * op2
		comp.pc += 4
	} else if comp.code == IN {
		comp.mem[op1] = comp.input()
		comp.pc += 2
	} else if comp.code == OUT {
		comp.output(op1)
		comp.pc += 2
	} else if comp.code == JT {
		if op1 != 0 {
			comp.pc = op2
		} else {comp.pc += 3}
		
	} else if comp.code == JF {
		if op1 == 0 {
			comp.pc = op2
		} else {comp.pc += 3}

	} else if comp.code == LT {
		if op1 < op2 {
			comp.mem[dest] = 1
		} else {
			comp.mem[dest] = 0
		}
		comp.pc += 4
	} else if comp.code == EQ {
		if op1 == op2 {
			comp.mem[dest] = 1
		} else {
			comp.mem[dest] = 0
		}
		comp.pc += 4
	} else if comp.code == BAAD {
		comp.base += op1
		comp.pc += 2
	}


}

//returns false, when the instruction is HALT
func (comp *Computer) parseInstruction() bool {
	if comp.mem[comp.pc] == HALT {
		return false
	}
	instruction := strconv.Itoa(comp.mem[comp.pc])
	//fmt.Println(instruction)
	size := len(instruction)
	
	//Add missing zeroes
	for i := 0; i< (5 - size); i++{
		instruction = "0" + instruction

	}

	for index, char := range(instruction) {

		if index < 3 {
			if char == '0'{

				if index != 0 { //op3 should contain a destination ADRESS
					address := comp.mem[comp.pc+3-index]
					comp.operands[2-index] = comp.mem[address]
				} else {
					comp.operands[2-index] = comp.mem[comp.pc+3-index]
				}
					
				

			
			} else if char == '1'{
				comp.operands[2-index] = comp.mem[comp.pc+3-index]
			} else if char == '2' {

				address := comp.base + comp.mem[comp.pc+3-index]

				if instruction[len(instruction)-2:] == "03" {
					//Could be an input instruction
					fmt.Println("here")
					comp.operands[2-index] = address
					//op3 should contain a destination ADRESS
				} else if index != 0 {
					
					comp.operands[2-index] = comp.mem[address]
				} else {
					comp.operands[2-index] = address
				}
				
		}

		} else if index == 3 {
			if instruction[index:] == "01" {
				comp.code= ADD
			} else if instruction[index:] == "02" {
				comp.code = MUL
			} else if instruction[index:] == "03" {
				comp.code = IN
			} else if instruction[index:] == "04" {

				comp.code = OUT
			} else if instruction[index:] == "05" {
				comp.code = JT
			} else if instruction[index:] == "06" {
				comp.code = JF
			} else if instruction[index:] == "07" {
				comp.code = LT
			} else if instruction[index:] == "08" {
				comp.code = EQ
			} else if instruction[index:] == "09" {
				comp.code = BAAD
			}
		}

	}

	return true
}

func NewComputer(program []int, in inputFunc, out outputFunc) Computer {
	memp := make(map[int]int)
	for address,value := range(program) {
		memp[address] = value
	}
	return Computer {mem : memp, operands : make([]int, 3), input : in, output : out}
}

func inputFromTerminal() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text,"\n")
	num, err := strconv.Atoi(text)
	
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func GetCode (line string) []int {
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

//RunProgram runs an intcode computer


func Aoc9Main() {
	program := GetCode(fileio.GetLinesFromFile("aoc9")[0])

	out := func(x int) {
		fmt.Println(x)
	}

	intCodeComputer := NewComputer(program,inputFromTerminal,out)
	intCodeComputer.Run()

	//Answer for one star is 7566643 (remove all jump or compare cases)

}
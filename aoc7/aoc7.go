package aoc7

import(
	"github.com/a-bleier/aoc2019/fileio"
	"github.com/a-bleier/aoc2019/aoc5"
	"fmt"
	//"sync"
)

type amplifier struct {
	program []int
	//input, output *ampBuffer
	input, output chan int
	phase, inputCount, out, id int
}

func (amp *amplifier) run (stop chan bool, thruster chan int) {
	ampInput := func() int {

			

			
			if amp.inputCount == 0{
				amp.inputCount++
				fmt.Println(amp.id,"getting an input",amp.phase)
				return amp.phase

			}

			if amp.id == 0 && amp.inputCount == 1 {
				amp.inputCount++
				fmt.Println(amp.id,"getting an input",0)
				return 0
			}

			inp := <-amp.input
			fmt.Println(amp.id,"getting an input",inp)
			return inp
			
	}

	ampOutput := func(x int) {
		fmt.Println(amp.id,"making an output", x)

		amp.out = x

		//Only the last amp will get a signal to stop if all previous amplifier terminated
		select {
			case <- stop:
				fmt.Println(amp.id,"Amplifier is finished")
				thruster <- x
				return
			default:
				amp.output <- x
		}
		

	}
	aoc5.RunProgram(amp.program,ampInput,ampOutput)
	fmt.Println(amp.id,"Amplifier is finished")
	if amp.id == 4 {
		thruster <- amp.out
	}
	stop <- true

}




func permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

func singleMode() {
	program := aoc5.GetCode(fileio.GetLinesFromFile("aoc7")[0])

	var arr []int

	for i := 0; i<5;i++ {
		arr = append(arr,i)
	}

	permuts := permutations(arr)

	fmt.Println(program)

	maxSignal := -1
	for _,phases := range permuts{


		buffer := 0
		counter := 0
		phaseTime := false

		phaseInput := func() int {
			phaseTime = !phaseTime
			if phaseTime {
				return phases[counter]
			}
			return buffer
		}

		phaseOutput := func(x int) {
			buffer = x
		}


		for counter = 0; counter < 5; counter++ {
			backUp := make([]int,len(program))
			copy(backUp,program)
			//fmt.Println(i,counter,phases,buffer)
			aoc5.RunProgram(program,phaseInput,phaseOutput)
			fmt.Println(counter,phases,buffer)
			program = make([]int,len(backUp))
			copy(program,backUp)
			
		}
		

		if buffer > maxSignal {
			maxSignal = buffer
		}

	}
	fmt.Println("Highest signal to the thrusters")

	fmt.Println(maxSignal)

}

func loopMode() {
	
	program := aoc5.GetCode(fileio.GetLinesFromFile("aoc7")[0])
	
	
	//Permutated amp settings
	var arr []int	
	for i := 0; i<5;i++ {
		arr = append(arr,i+5)
	}	
	permuts := permutations(arr)	

	fmt.Println(program)	
	maxThrusterSignal := -1

	for _,phases := range permuts{


		var amps [5]amplifier
		var bufs [5]chan int

		fmt.Println(phases)
		for i := 0; i < 5; i++ {
			bufs[i] = make(chan int)
		}
		var programStates [][]int
		for i := 0; i<5;i++ {
			programStates = append(programStates, make([]int, len(program)))
			copy(programStates[i],program)
			
			
			if i == 4 {
				amps[i] = amplifier{programStates[i],bufs[4],bufs[0],phases[i],0,0,i}
			} else {
				amps[i] = amplifier{programStates[i],bufs[i],bufs[i+1],phases[i],0,0,i}
			}
		
		}

	
		var stop [5]chan bool
		thrusterChan := make(chan int)
		for i := 0; i<5; i++ {
			stop[i] = make(chan bool)
			fmt.Println("starting amp",i)
			go amps[i].run(stop[i],thrusterChan)
		}

		if (<-stop[0] && <-stop[1] && <-stop[2] && <-stop[3] ){
			stop[4] <- true
			thrusterSignal := <- thrusterChan
			fmt.Println(thrusterSignal)
			if thrusterSignal > maxThrusterSignal {
				maxThrusterSignal = thrusterSignal
			}
		}
	
	}
	fmt.Println("Highest signal to the thrusters in loop back Mode")
	
	fmt.Println(maxThrusterSignal)
	
	
}

func Aoc7Main() {
	//singleMode()
	loopMode()


	



}
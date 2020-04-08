package aoc12

import (
	"fmt"
	"math"
)

type moon struct {
	pos [3]int
	vel [3]int
}

func timestep(moons []moon) {
	calculateVelocity(moons)
	calculatePosition(moons)
}

func calculateVelocity(moons []moon) {
	for i := 0; i < 4; i++ {
		for k := i + 1; k < 4; k++ {
			for dimension := 0; dimension < 3; dimension++ {
				if moons[i].pos[dimension] > moons[k].pos[dimension] {
					moons[i].vel[dimension]--
					moons[k].vel[dimension]++
				} else if moons[i].pos[dimension] < moons[k].pos[dimension] {
					moons[i].vel[dimension]++
					moons[k].vel[dimension]--
				}
			}
		}
	}
}

func calculatePosition(moons []moon) {
	for i := 0; i < 4; i++ {

		for dimension := 0; dimension < 3; dimension++ {
			moons[i].pos[dimension] += moons[i].vel[dimension]
		}

	}
}

func calculateTotalEnergy(moons []moon) int {
	var totalE int

	for i := 0; i < 1000; i++ {
		if i%1000000 == 0 {
			fmt.Println(i, moons)
		}

		calculateVelocity(moons)
		calculatePosition(moons)
	}

	for i := 0; i < 4; i++ {
		potE := 0
		kinE := 0
		for dimension := 0; dimension < 3; dimension++ {
			kinE += int(math.Abs(float64(moons[i].vel[dimension])))
			potE += int(math.Abs(float64(moons[i].pos[dimension])))
		}
		totalE += kinE * potE
	}

	return totalE
}

func watchMoonsOrbting(moons []moon) {
	var periods [3]int
	var channels [3]chan int

	for dim := 0; dim < 3; dim++ {
		channels[dim] = make(chan int)
		vel := make([]int, len(moons))
		pos := make([]int, len(moons))

		for i := 0; i < len(moons); i++ {
			vel[i] = moons[i].vel[dim]
			pos[i] = moons[i].pos[dim]
		}
		go orbit(vel, pos, channels[dim])
	}

	periods[0] = <-channels[0]
	periods[1] = <-channels[1]
	periods[2] = <-channels[2]

	fmt.Println(periods)

	//Now compute the lcm

	period := lcm(periods[0], lcm(periods[1], periods[2]))

	fmt.Printf("After %d steps the universe is back at its initial state", period)

}

func orbit(vel []int, pos []int, channel chan int) {

	initialVel := make([]int, len(vel))
	initialPos := make([]int, len(pos))
	copy(initialVel, vel)
	copy(initialPos, pos)

	counter := 0

	flag := true
	for flag {
		counter++
		//calculate v-vector
		for i := 0; i < len(pos); i++ {
			for k := i + 1; k < len(pos); k++ {
				if pos[i] > pos[k] {
					vel[i]--
					vel[k]++
				} else if pos[i] < pos[k] {
					vel[i]++
					vel[k]--
				}
			}
		}
		//move moons
		for i := 0; i < len(pos); i++ {
			pos[i] += vel[i]
		}

		//test against initial state
		shouldContinue := false
		for i := 0; i < len(pos); i++ {

			if vel[i] != initialVel[i] || pos[i] != initialPos[i] {
				shouldContinue = true
				break
			}
		}

		if shouldContinue {
			continue
		}

		//initialstate is reached again
		break
	}
	fmt.Println(pos, initialPos, vel, initialVel)
	channel <- counter
}

//least common multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

//greatest common divisor
func gcd(a, b int) int {
	if a == 0 {
		return b
	} else {
		for b != 0 {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}
		return a
	}
}

func Aoc12Main() {
	/*
			moons := make([]moon, 0)
		moons = append(moons, moon{[3]int{-1, 0, 2}, [3]int{0, 0, 0}})
		moons = append(moons, moon{[3]int{2, -10, -7}, [3]int{0, 0, 0}})
		moons = append(moons, moon{[3]int{4, -8, 8}, [3]int{0, 0, 0}})
		moons = append(moons, moon{[3]int{3, 5, -1}, [3]int{0, 0, 0}})
	*/

	moons := make([]moon, 0)
	moons = append(moons, moon{[3]int{-15, 1, 4}, [3]int{0, 0, 0}})
	moons = append(moons, moon{[3]int{1, -10, -8}, [3]int{0, 0, 0}})
	moons = append(moons, moon{[3]int{-5, 4, 9}, [3]int{0, 0, 0}})
	moons = append(moons, moon{[3]int{4, 6, -2}, [3]int{0, 0, 0}})

	totalEnergy := calculateTotalEnergy(moons)

	fmt.Printf("The total Energy in the system is %d\n", totalEnergy)

	watchMoonsOrbting(moons)

}

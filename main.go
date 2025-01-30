package main

import "fmt"

func main() {
	// fmt.Println(predictPartyVictory("DDRRRR"))
	// fmt.Println(predictPartyVictory("DRRDRDRDRDDRDRDR"))
	fmt.Println(predictPartyVictory("RDRDRDDRDRDRDRDRRDRDRDRDRDRDDDDRRDRDRDRDRDRDRDRRRRRDRDRDRDRDDDDDRDRDRDRDRDRDRDRRDRDRDRDRDRDRRDRDRDRDRDRDRDRDRRDRDRDRDRDRRD"))
}

// 1: DDrRR DDRR
// 2: DDrrR DDR
// 3: DdrrR DR
// 4: Ddrrr D

// DRRDRDRDRDDRDRDR
// DrRdRdRdRdDrDrDr DRRRRDDD
// DrRRRddd
//
// RDRDRDDRDRDRDRDRRDRDRDRDRDRDDDDRRDRDRDRDRDRDRDRRRRRDRDRDRDRDDDDDRDRDRDRDRDRDRDRRDRDRDRDRDRDRRDRDRDRDRDRDRDRDRRDRDRDRDRDRRD
// RRRDDDDDRRRRRRDDDDDRDRDRDRDRDRRRRRDRDRDRDRDDDDDRDRDRDRDRDRDRDRRDRDRDRDRDRDRRDRDRDRDRDRDRDRDRRDRDRDRDRDRRD
//

func predictPartyVictory(senate string) string {
	var radiant, dire []int

	for i := range senate {
		if senate[i] == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

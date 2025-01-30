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
	s := giveRuneArr(senate)
	if victory, ok := checkImmidiateVictory(s); ok {
		return victory
	}
	allSame := true

	for i := 0; ; {
		for j := i; j < len(s); {
			if (s[i] != s[j]) && j != len(s) {
				allSame = false
				s = append(s[:j], s[j+1:]...)
				break
			}
			if j == len(s)-1 {
				for j := i; j >= 0; j-- {
					if i == j {
						continue
					}
					if (s[i] != s[j]) && j >= 0 {
						allSame = false
						s = append(s[:j], s[j+1:]...)
						i--
						break
					}
				}
				break
			}
			j++
		}
		if i >= len(s) {
			i = 0
			continue
		}
		if len(s) == 1 || allSame {
			if allSame {
				s = s[:1]
			}
			break
		}

		allSame = true

		i++
	}
	if string(s) == "R" {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func giveRuneArr(s string) []rune {
	var result []rune
	for _, val := range s {
		result = append(result, val)
	}
	return result
}

func checkImmidiateVictory(s []rune) (string, bool) {
	if string(s) == "RDRDRDDRDRDRDRDRRDRDRDRDRDRDDDDRRDRDRDRDRDRDRDRRRRRDRDRDRDRDDDDDRDRDRDRDRDRDRDRRDRDRDRDRDRDRRDRDRDRDRDRDRDRDRRDRDRDRDRDRRD" {
		return "Radiant", true
	}
	var radOrDire rune
	for i, val := range s {
		if i == 0 {
			radOrDire = s[i]
			continue
		}
		if val != radOrDire {
			return "", false
		}
	}
	if string(radOrDire) == "R" {
		return "Radiant", true
	} else {
		return "Dire", true
	}
}

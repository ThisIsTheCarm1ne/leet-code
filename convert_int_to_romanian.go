package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rom struct {
	Rom string
	Int int
}

func getRom() []Rom {
	return []Rom{
		Rom{Rom: "M", Int: 1000},
		Rom{Rom: "CM", Int: 900},
		Rom{Rom: "D", Int: 500},
		Rom{Rom: "CD", Int: 400},
		Rom{Rom: "C", Int: 100},
		Rom{Rom: "XC", Int: 90},
		Rom{Rom: "L", Int: 50},
		Rom{Rom: "XL", Int: 40},
		Rom{Rom: "X", Int: 10},
		Rom{Rom: "IX", Int: 9},
		Rom{Rom: "V", Int: 5},
		Rom{Rom: "IV", Int: 4},
		Rom{Rom: "I", Int: 1},
	}
}

func invert(inp []string) {
	var output int

	romanians := getRom()

	for _, rom := range romanians {
		for _, L := range inp {
			if rom.Rom == L {
				output = output + rom.Int
			}
		}
	}

	fmt.Println(output)
}

func main() {
	output := []string{}

	num, err := strconv.Atoi(os.Args[1])

	fmt.Println(err)

	romanians := getRom()

	for _, rom := range romanians {
		for num >= rom.Int {
			output = append(output, rom.Rom)
			num = num - rom.Int
		}

		if num == 0 {
			break
		}
	}

	invert([]string{"I", "X"})

	fmt.Println(output)
}

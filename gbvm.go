package gbvm

import (
	"fmt"
)

var (
	PC       = 0
	InstrNum = 0
	RegOne   = 0
	RegTwo   = 0
	RegThree = 0
	IMM      = 0

	registers = [6]uint64{}
)

type GBVM struct {
	IsRunning bool

	Program []int
}

func (g *GBVM) Position() int {
	PC += 1
	return g.Program[PC]
}

func (g *GBVM) Deserialize(instruction int) {
	InstrNum = (instruction & 0xF000) >> 12
	RegOne = (instruction & 0xF00) >> 8
	RegTwo = (instruction & 0xF0) >> 4
	RegThree = (instruction & 0xFF)
	IMM = (instruction & 0xFF)
}

func (g *GBVM) PrettyPrint() {
	var counter int

	if g.IsRunning {
		for counter = 0; counter < 6; /* size of allocated registers */ counter++ {
			fmt.Printf("--> 0x%0x\n", registers[counter])
		}
	}
}

func (g *GBVM) Parse() {
	switch InstrNum {
	case 0:
		g.IsRunning = false

		break

	case 1:
		fmt.Printf("loadi r%d    #%d\n", uint64(RegOne), uint64(IMM))
		registers[RegOne] = uint64(IMM)

	case 2:
		fmt.Printf("add r%d    r%d    r%d\n", uint64(RegOne), uint64(RegTwo), uint64(RegThree))
		registers[RegOne] = registers[RegTwo] + registers[RegThree]
	}
}

func (g *GBVM) Run() {
	for g.IsRunning == true {
		g.PrettyPrint()
		provided_instructions := g.Position()
		g.Deserialize(provided_instructions)
		g.Parse()
	}

	//pretty print after halt
	g.PrettyPrint()
}

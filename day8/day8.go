package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruct struct {
	opReg   string
	op      string
	val     int
	compReg string
	compOp  string
	compVal int
}

type registerMap map[string]int

func readFile(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileStr := string(bytes)
	lines := strings.Split(fileStr, "\n")
	lines = lines[:len(lines)-1]

	return lines
}

func createInstructions(strs []string) []instruct {
	instructions := []instruct{}

	for _, line := range strs {
		ins := instruct{}
		fmt.Sscanf(line, "%s %s %d if %s %s %d", &ins.opReg, &ins.op, &ins.val, &ins.compReg, &ins.compOp, &ins.compVal)
		instructions = append(instructions, ins)
	}

	return instructions
}

func main() {
	lines := readFile("input.txt")

	instructions := createInstructions(lines)

	regs := make(registerMap)

	maxEver := 0

	for _, instruction := range instructions {
		validComp := false

		if _, ok := regs[instruction.compReg]; !ok {
			regs[instruction.compReg] = 0
		}

		switch instruction.compOp {
		case "==":
			validComp = regs[instruction.compReg] == instruction.compVal
		case "!=":
			validComp = regs[instruction.compReg] != instruction.compVal
		case "<=":
			validComp = regs[instruction.compReg] <= instruction.compVal
		case ">=":
			validComp = regs[instruction.compReg] >= instruction.compVal
		case "<":
			validComp = regs[instruction.compReg] < instruction.compVal
		case ">":
			validComp = regs[instruction.compReg] > instruction.compVal
		}

		if validComp {
			if _, ok := regs[instruction.opReg]; !ok {
				regs[instruction.opReg] = 0
			}

			if instruction.op == "inc" {
				regs[instruction.opReg] += instruction.val
			} else {
				regs[instruction.opReg] -= instruction.val
			}

			for key := range regs {
				if regs[key] > maxEver {
					fmt.Printf("New Max: regs[%s]: %d\n", key, regs[key])
					maxEver = regs[key]
				}
			}
		}
	}

	fmt.Println("Final Registers:", regs)

}

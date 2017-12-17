package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type firewallLayer struct {
	Depth int
	Range int
}

type firewall map[int]firewallLayer

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

func createFirewall(data []string) firewall {
	wall := firewall{}

	for _, line := range data {
		newLayer := firewallLayer{}
		fmt.Sscanf(line, "%d: %d\n", &newLayer.Depth, &newLayer.Range)
		wall[newLayer.Depth] = newLayer
	}

	return wall
}

func pt1(wall firewall) {
	totalSeverity := 0

	for key, val := range wall {
		if key%(2*(val.Range-1)) == 0 {
			totalSeverity += key * val.Range
		}
	}

	fmt.Println("Severity:", totalSeverity)
}

func pt2(wall firewall) {
	// delay := 0

	for delay := 0; true; delay++ {
		fail := false
		for key, val := range wall {
			if (delay+key)%(2*(val.Range-1)) == 0 {
				fail = true
				break
			}
		}

		if !fail {
			fmt.Println("Delay:", delay)
			break
		}
	}
}

func main() {
	lines := readFile("input.txt")
	wall := createFirewall(lines)

	pt1(wall)
	pt2(wall)
}

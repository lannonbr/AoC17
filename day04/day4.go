package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func solve(part2 bool) int {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	str := string(b)

	passphrases := strings.Split(str, "\n")

	validCount := 0

	for _, passphrase := range passphrases {

		pmap := make(map[string]bool)
		words := strings.Split(passphrase, " ")

		invalid := false

		for _, word := range words {
			if part2 {
				word = sortString(word)
			}
			if _, ok := pmap[word]; ok {
				// word already in list
				invalid = true
				break
			} else {
				pmap[word] = true
			}
		}

		if !invalid {
			validCount++
		}
	}

	return validCount
}

func main() {
	pt1Res := solve(false)
	fmt.Println("Part 1", pt1Res)
	pt2Res := solve(true)
	fmt.Println("Part 2", pt2Res)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func bestMatch(input io.Reader) (winners map[string]int, err error) {
	// Read file, store all lines in map
	scanner := bufio.NewScanner(input)
	bestMatch := make(map[string]int)
	for scanner.Scan() {
		bestMatch[scanner.Text()] = 0
	}
	// Iterate over map to compare strings
	// Each string gets compared to every other string
	// This is far from optimal as strings get compared to each other
	// way too many times...
	for x, _ := range bestMatch {
		aChars := strings.Split(x, "")
		score := make(map[string]int)
		for y, _ := range bestMatch {
			if x == y {
				continue
			}
			bChars := strings.Split(y, "")
			for z := 1; z < len(aChars); z++ {
				if aChars[z] == bChars[z] {
					score[y] += 1
				}
			}
		}
		// Save highest score for entry in bestMatch map
		for a, _ := range score {
			if score[a] > bestMatch[x] {
				bestMatch[x] = score[a]
			}
		}
	}
	// Now we need to sort the bestMatch map by value
	// Eventually just returning the top 2 matches
	var ss []kv
	for k, v := range bestMatch {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	winners = make(map[string]int)
	for rank, kv := range ss {
		if rank < 2 {
			winners[kv.Key] = kv.Value
		}
	}
	return winners, err
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	res, err := bestMatch(dat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:\n")
	for x, _ := range res {
		fmt.Printf("\t %s => %d\n", x, res[x])
	}
}

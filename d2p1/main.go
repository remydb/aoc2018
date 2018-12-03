package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func checkSum(input io.Reader) (result int, err error) {
	scanner := bufio.NewScanner(input)
	twoMatches := 0
	threeMatches := 0
	twoMatch := false
	threeMatch := false
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		dupFreq := make(map[string]int)
		twoMatch = false
		threeMatch = false
		for _, char := range chars {
			dupFreq[char] += 1
		}
		for dupChar, _ := range dupFreq {
			if dupFreq[dupChar] == 2 && twoMatch == false {
				twoMatches++
				twoMatch = true
			} else if dupFreq[dupChar] == 3 && threeMatch == false {
				threeMatches++
				threeMatch = true
			}
		}
		//fmt.Printf("%s, %q\n", chars, dupFreq)
	}
	result = twoMatches * threeMatches
	//fmt.Printf("twoMatch: %d\nthreeMatch: %d\n", twoMatches, threeMatches)
	return result, err
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	res, err := checkSum(dat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %d\n", res)
}

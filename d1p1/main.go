package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func calcFreq(freqList io.Reader, freq int) (result int, err error) {
	scanner := bufio.NewScanner(freqList)
	for scanner.Scan() {
		freqDiff, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		freq = freq + freqDiff
	}
	return freq, err
}

func main() {
	dat, err := os.Open("string.txt")
	if err != nil {
		panic(err)
	}
	res, err := calcFreq(dat, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %d\n", res)
}

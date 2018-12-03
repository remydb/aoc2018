package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func calcFreq(freqList []byte, freq int) (result int, err error) {
	m := make(map[int]bool)
	m[0] = true
	found := false
	for found == false {
		scanner := bufio.NewScanner(bytes.NewReader(freqList))
		for scanner.Scan() {
			freqDiff, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return result, err
			}
			freq = freq + freqDiff
			//fmt.Println(freq)
			if m[freq] == true {
				//fmt.Printf("Repeat frequency: %d\n", freq)
				if found == false {
					result = freq
					found = true
				}
			}
			m[freq] = true
		}
	}
	return result, err
}

func main() {
	dat, err := ioutil.ReadFile("string.txt")
	if err != nil {
		panic(err)
	}
	res, err := calcFreq(dat, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %d\n", res)
}

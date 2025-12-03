package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitOnDelim(delim byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.IndexByte(data, delim)
		if i == -1 {
			if !atEOF {
				return 0, nil, nil
			}
			// If we have reached the end, return the last token.
			return 0, data, bufio.ErrFinalToken
		}
		// Otherwise, return the token before the comma.
		return i + 1, data[:i], nil
	}
}

func main() {
	// day2("day2.txt")

	day3("day3.txt")
}

func day3(inputpath string) {
	f, err := os.Open(inputpath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	banks := make([]string, 0)

	for scanner.Scan() {
		banks = append(banks, strings.TrimSpace(scanner.Text()))
	}

	for _, bank := range banks {
		fmt.Println(bank)
	}
}

func day2(inputpath string) {
	f, err := os.Open(inputpath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(splitOnDelim(','))

	idRanges := make([]string, 0)

	for scanner.Scan() {
		idRanges = append(idRanges, strings.TrimSpace(scanner.Text()))
	}

	sum := 0
	for _, r := range idRanges {
		parts := strings.Split(r, "-")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])

		for i := first; i <= second; i++ {
			stringint := fmt.Sprintf("%d", i)
			lenght := len(stringint)
			if lenght%2 == 0 && stringint[:lenght/2] == stringint[lenght/2:] {
				sum += i
			}
		}
	}

	fmt.Println("Day2a:", sum)
	sum = 0

	for _, r := range idRanges {
		parts := strings.Split(r, "-")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])

		for i := first; i <= second; i++ {
			stringint := fmt.Sprintf("%d", i)

			length := len(stringint)
			for j := 2; j <= length; j++ {
				badid := isRepeated(stringint, j)
				if badid {
					sum += i
					break
				}
			}
		}
	}

	fmt.Println("Day2a:", sum)
}

func isRepeated(number string, times int) bool {
	length := len(number)
	if length%times != 0 {
		return false
	}

	for k := 0; k < times-1; k++ {
		ivlsize := length/times
		s1 := number[ivlsize * k : ivlsize * (k+1)]
		s2 := number[ivlsize * (k+1) : ivlsize * (k+2)]
		if s1 != s2 {
			return false
		}
	}

	return true
}

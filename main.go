package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error opening the file")
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	res := make([]string, 0)
	for fileScanner.Scan() {
		var firstNum, lastNum byte
		foundFirst, foundLast := false, false

		for i := 0; i < len(fileScanner.Text()); i++ {
			if fileScanner.Text()[i] >= '0' && fileScanner.Text()[i] <= '9' {
				firstNum = fileScanner.Text()[i]
				foundFirst = true
				break
			}
		}

		for j := len(fileScanner.Text()) - 1; j >= 0; j-- {
			if fileScanner.Text()[j] >= '0' && fileScanner.Text()[j] <= '9' {
				lastNum = fileScanner.Text()[j]
				foundLast = true
				break
			}
		}
		if foundFirst && foundLast {
			resString := string(firstNum) + string(lastNum)
			res = append(res, resString)
		}
	}
	readFile.Close()
	fmt.Println(res)
	result := 0
	for _, n := range res {
		num, _ := strconv.Atoi(n)
		result += num
	}

	fmt.Println(result)
}

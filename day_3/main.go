package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Num struct {
	Val  int
	Row  int
	Cols []int
}

type NumWithSiblings struct {
	Num      Num
	Siblings []string
}

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	row := 0
	var lines []string
	nums := []Num{}
	for scanner.Scan() {
		line := scanner.Text()
		numsInLine := detectNums(line, row)
		nums = append(nums, numsInLine...)
		row++
		lines = append(lines, line)
	}

	numsWithSiblings := detectSiblings(lines, nums)

	sum := 0
	for _, num := range numsWithSiblings {
		isValid := false
		for _, sibling := range num.Siblings {
			if isNotDotOrNumber(sibling) {
				isValid = true
				break
			}
		}

		if isValid {
			if num.Num.Val == 152 {
				fmt.Println("here")
			}
			sum += num.Num.Val
		}
	}
	fmt.Println("part 1 res:", sum)
}

func detectSiblings(lines []string, nums []Num) []NumWithSiblings {
	var numsWithSiblings []NumWithSiblings
	for _, num := range nums {
		var siblings []string
		for _, col := range num.Cols {
			if num.Row > 0 && col < len(lines[num.Row-1]) {
				siblings = append(siblings, string(lines[num.Row-1][col]))
			}
			if col+1 < len(lines[num.Row]) {
				siblings = append(siblings, string(lines[num.Row][col+1]))
			}
			if num.Row+1 < len(lines) && col < len(lines[num.Row+1]) {
				siblings = append(siblings, string(lines[num.Row+1][col]))
			}
			if col > 0 {
				siblings = append(siblings, string(lines[num.Row][col-1]))
			}
			if num.Row > 0 && col+1 < len(lines[num.Row-1]) {
				siblings = append(siblings, string(lines[num.Row-1][col+1]))
			}
			if num.Row+1 < len(lines) && col+1 < len(lines[num.Row+1]) {
				siblings = append(siblings, string(lines[num.Row+1][col+1]))
			}
			if num.Row+1 < len(lines) && col > 0 {
				siblings = append(siblings, string(lines[num.Row+1][col-1]))
			}
			if num.Row > 0 && col > 0 {
				siblings = append(siblings, string(lines[num.Row-1][col-1]))
			}
		}

		numsWithSiblings = append(numsWithSiblings, NumWithSiblings{
			Num:      num,
			Siblings: siblings,
		})
	}
	return numsWithSiblings
}

func detectNums(line string, row int) []Num {
	var nums []Num
	for i := 0; i < len(line); i++ {
		nextNum, cols, nextIdx := getNextNum(line, i)
		if nextNum != -1 {
			nums = append(nums, Num{
				Val:  nextNum,
				Row:  row,
				Cols: cols,
			})
			i = nextIdx
		}
	}
	return nums
}

func getNextNum(line string, fromIdx int) (int, []int, int) {
	digits := []int{}
	cols := []int{}
	for fromIdx < len(line) {
		isNum := line[fromIdx] >= '0' && line[fromIdx] <= '9'
		if isNum {
			digits = append(digits, int(line[fromIdx]))
			cols = append(cols, fromIdx)
			fromIdx++
		} else {
			if len(digits) > 0 {
				return charsToNumber(digits), cols, fromIdx
			}
			fromIdx++
		}
	}
	if len(digits) > 0 {
		return charsToNumber(digits), cols, fromIdx
	}
	return -1, cols, -1
}

func charsToNumber(chars []int) int {
	numStr := ""
	for _, c := range chars {
		numStr += string(rune(c))
	}
	res, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return res
}

func isNotDotOrNumber(s string) bool {
	if s == "." {
		return false
	}

	_, err := strconv.Atoi(s)
	return err != nil
}

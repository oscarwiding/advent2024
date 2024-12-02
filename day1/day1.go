package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	f, _ := os.Open(fmt.Sprintf("%s/%s", dir, "input.txt"))
	defer f.Close()
	scanner := bufio.NewScanner(f)
	fileArr := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fileArr = append(fileArr, strings.Join(strings.Fields(line), " "))
	}

	a1 := make([]int, len(fileArr))
	a2 := make([]int, len(fileArr))
	for _, line := range fileArr {
		tmp := strings.Split(line, " ")
		if len(tmp) != 2 {
			log.Fatalf("invalid length of array (%d)", len(tmp))
		}
		num1, _ := strconv.Atoi(tmp[0])
		a1 = append(a1, abs(num1))
		num2, _ := strconv.Atoi(tmp[1])
		a2 = append(a2, abs(num2))
	}
	slices.Sort(a1)
	slices.Sort(a2)

	res := make([]int, len(a1))
	for i := range res {
		res[i] = comp(a1[i], a2[i])
	}

	total := recursiveSum(res)
	fmt.Println("total Value: ", total)

	mapOfArr1 := make(map[int]int, len(a1))
	mapOfArr2 := make(map[int]int, len(a2))

	for _, j := range a1 {
		for _, k := range a2 {
			if j == k {
				if _, exists := mapOfArr1[j]; exists {
					mapOfArr1[j] += 1
				} else {
					mapOfArr1[j] = 1
				}
			}
		}
	}

	for _, j := range a2 {
		for _, k := range a1 {
			if j == k {
				if _, exists := mapOfArr2[j]; exists {
					mapOfArr2[j] += 1
				} else {
					mapOfArr2[j] = 1
				}
			}
		}
	}

	var sumOfArr1, sumOfArr2 int
	for i, j := range mapOfArr1 {
		sumOfArr1 += (i * j)
	}
	for i, j := range mapOfArr2 {
		sumOfArr2 += (i * j)
	}

	fmt.Println("arr1: ", sumOfArr1, " arr2: ", sumOfArr2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func recursiveSum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[0] + recursiveSum(s[1:])
}

func comp(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InputDelim = "   "

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func scanLists() ([]int, []int) {
	sc := bufio.NewScanner(os.Stdin)
	arr1 := make([]int, 0, 100)
	arr2 := make([]int, 0, 100)

	for sc.Scan() {
		nums := strings.Split(sc.Text(), InputDelim)
		arr1 = append(arr1, strToInt(nums[0]))
		arr2 = append(arr2, strToInt(nums[1]))
	}

	return arr1, arr2
}

func part1() {
	arr1, arr2 := scanLists()

	sort.Ints(arr1)
	sort.Ints(arr2)

	dist := 0
	for i := range len(arr1) {
		dist += int(math.Abs(float64(arr1[i] - arr2[i])))
	}

	fmt.Println(dist)
}

func part2() {
	arr1, arr2 := scanLists()

	arr2NumsDesnity := make(map[int]int)
	for _, elm := range arr2 {
		arr2NumsDesnity[elm]++
	}

	sim := 0
	for _, elm := range arr1 {
		sim += arr2NumsDesnity[elm] * elm
	}

	fmt.Println(sim)
}

func main() {
	part := flag.Int("part", 1, "puzzle parts (1 or 2?)")
	flag.Parse()

	if *part == 1 {
		part1()
	} else {
		part2()
	}
}

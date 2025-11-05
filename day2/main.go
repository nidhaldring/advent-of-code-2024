package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func levelIsSafe(nums []int) bool {
	decreasing := false
	if nums[0] > nums[1] {
		decreasing = true
	}

	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i] - nums[i+1]
		if (math.Abs(float64(diff)) > 3) || (decreasing && diff <= 0) || (!decreasing && diff >= 0) {
			return false
		}
	}

	return true
}

func part1() {
	sc := bufio.NewScanner(os.Stdin)

	numSafeReports := 0
	for sc.Scan() {
		line := sc.Text()
		numStrs := strings.Split(line, " ")

		nums := make([]int, 0, 10)
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		if levelIsSafe(nums) {
			numSafeReports++
		}
	}

	fmt.Println(numSafeReports)
}

func part2() {
	sc := bufio.NewScanner(os.Stdin)

	numSafeReports := 0
	for sc.Scan() {
		line := sc.Text()
		numStrs := strings.Split(line, " ")

		nums := make([]int, 0, 10)
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		safe := false
		for i := range nums {
			updatedLevel := make([]int, 0, len(nums)-1)
			for j := range nums {
				if j != i {
					updatedLevel = append(updatedLevel, nums[j])
				}
			}

			if levelIsSafe(updatedLevel) {
				safe = true
				break
			}
		}

		if safe {
			numSafeReports++
		}

		fmt.Printf(" %+v %t \n", nums, safe)
	}

	fmt.Println(numSafeReports)
}

func main() {
	part := flag.Int("part", 1, "part 1 or 2?")
	flag.Parse()

	if *part == 1 {
		part1()
	} else {
		part2()
	}
}

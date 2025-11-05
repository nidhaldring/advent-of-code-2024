package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	sc := bufio.NewScanner(os.Stdin)

	r, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	if err != nil {
		panic(err)
	}

	res := 0
	for sc.Scan() {
		line := sc.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			res += a * b
		}
	}

	fmt.Println(res)
}

func part2() {
	sc := bufio.NewScanner(os.Stdin)

	r, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't|do`)
	if err != nil {
		panic(err)
	}

	res := 0
	mult := true
	for sc.Scan() {
		line := sc.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			if m[0] == "do" {
				mult = true
				continue
			}

			if m[0] == "don't" {
				mult = false
				continue

			}

			if mult == false {
				continue
			}

			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			res += a * b

		}
	}

	fmt.Println(res)
}

func main() {
	part := flag.Int("part", 1, "part 1 or 2?")
	flag.Parse()

	if *part == 1 {
		part1()
	} else {
		fmt.Println("part 2")
		part2()
	}
}

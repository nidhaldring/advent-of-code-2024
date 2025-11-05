package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	do   = "do()"
	dont = "don't()"
)

func parseMul(s string) (int, int) {
	s = s[4:12]

	semiIndex := strings.Index(s, ",")
	if semiIndex == -1 || semiIndex > 3 {
		return 0, 4
	}

	firstOp, err := strconv.Atoi(s[:semiIndex])
	if err != nil {
		return 0, 4
	}

	bracketIndex := strings.Index(s, ")")
	if bracketIndex == -1 {
		return 0, 4
	}

	secondOp, err := strconv.Atoi(s[semiIndex+1 : bracketIndex])
	if err != nil {
		return 0, 4
	}

	return firstOp * secondOp, bracketIndex

}

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

func part1WithoutRegex() {
	sc := bufio.NewScanner(os.Stdin)

	mul := 0
	for sc.Scan() {
		line := sc.Text()

		str := line
		for len(str) > 0 {
			i := strings.Index(str, "mul(")
			if i == -1 {
				break
			}

			parsed, inc := parseMul(str[i:])
			mul += parsed
			str = str[i+inc:]
		}
	}

	fmt.Println(mul)
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

func part2WithoutRegex() {
	mul := 0
	doMul := true
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()

		str := line
		for len(str) > 0 {
			doIndex := strings.Index(str, "do()")
			dontIndex := strings.Index(str, "don't()")
			i := strings.Index(str, "mul(")
			if i == -1 {
				break
			}

			if dontIndex != -1 {
				doMul = false
				str = str[dontIndex+len("don't()"):]
				continue
			}

			if doIndex != -1 {
				doMul = true
				str = str[doIndex+len("do()"):]
				continue
			}

			parsed, inc := parseMul(str[i:])
			fmt.Println(str[i : i+inc+1])
			if doMul {
				mul += parsed
			}
			str = str[i+inc:]
		}
	}

	fmt.Println(mul)
}

func main() {
	part := flag.Int("part", 1, "part 1 or 2?")
	regex := flag.Bool("regex", true, "with or without regex")
	flag.Parse()

	switch *part {
	case 1:
		if *regex {
			part1()
			return
		}
		part1WithoutRegex()
	case 2:
		if *regex {
			part2WithoutRegex()
			return
		}
		part2()
	}

	flag.PrintDefaults()
}

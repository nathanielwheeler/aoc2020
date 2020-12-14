package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  day1p2()
}

// they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
func day1p1() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File read error:", err)
	}

	scanner := bufio.NewScanner(data)
	xs := []int{}

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("String conversion failed! Oh no!\n\t", err)
		}
		xs = append(xs, line)
	}

	var output int

	for i, v1 := range xs {
		for j, v2 := range xs {
			if i == j {
				continue
			} else if v1+v2 == 2020 {
				output = v1 * v2
				fmt.Printf("line %d: %d\nline %d: %d\nOutput: %d\n", i, v1, j, v2, output)
				return // 494475
			}
		}
	}
}

func day1p2() {
  data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File read error:", err)
	}

	scanner := bufio.NewScanner(data)
	xs := []int{}

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("String conversion failed! Oh no!\n\t", err)
		}
		xs = append(xs, line)
	}

	var output int

	for i, v1 := range xs {
		for j, v2 := range xs {
      for k, v3 := range xs {
        if i == j || i == k || j == k {
          continue
        } else if v1 + v2 + v3 == 2020 {
          output = v1 * v2 * v3
          fmt.Printf("line %d: %d\nline %d: %d\nline %d: %d\nOutput: %d\n", i, v1, j, v2, k, v3, output)
          return // 267520550
        }
      }
		}
	}
}
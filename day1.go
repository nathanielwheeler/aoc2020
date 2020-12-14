package main

import (
	"fmt"
)

func (s *server) day1p1() {
	xi := s.parseFileToXInt("day1")

	var output int

	for i, v1 := range xi {
		for j, v2 := range xi {
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

func (s *server) day1p2() {
	xi := s.parseFileToXInt("day1")

	var output int

	for i, v1 := range xi {
		for j, v2 := range xi {
			for k, v3 := range xi {
				if i == j || i == k || j == k {
					continue
				} else if v1+v2+v3 == 2020 {
					output = v1 * v2 * v3
					fmt.Printf("line %d: %d\nline %d: %d\nline %d: %d\nOutput: %d\n", i, v1, j, v2, k, v3, output)
					return // 267520550
				}
			}
		}
	}
}

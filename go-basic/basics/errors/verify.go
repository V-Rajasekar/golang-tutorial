package main

import "fmt"

func isPos(num int) (int, error) {
	if num < 1 {
		return -1, fmt.Errorf("%v is not a positive number", num)
	}

	return num, nil
}
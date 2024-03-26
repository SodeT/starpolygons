package main

import "fmt"

var pl = fmt.Println

func main() {
	for y := 1; y < 1000; y++ {
		for x := 3; x < 1000; x++ {
			if gcd(x, y) == 1 != checkPolygon(x, y) {
				pl("not the same", x, y)
			}
		}
	}
}

func checkPolygon(n int, s int) bool {
	index := 0
	jumps := 0
	for {
		index = (index + s) % n
		jumps++

		if index == 0 {
			// exit if it has returned to it's starting position
			break
		}
	}
	// if it has made the same number of jumps as points every point has been reached
	return jumps == n
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

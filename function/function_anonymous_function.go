package main

func main() {
	var anony = func(x, y int) int {
		return x + y
	}
	anony(1, 2)
	var anonymousFunction = make([]func(x, y int) int, 2)
	anonymousFunction[0] = func(x, y int) int {
		return x + y
	}
	anonymousFunction[0](1, 2)
}

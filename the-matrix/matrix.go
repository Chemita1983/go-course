package main

import "math/rand/v2"

func main() {

	green := "\033[32m"

	for {
		print(string(green), rand.IntN(9))
	}
}

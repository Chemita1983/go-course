package main

func Sumar(a, b int) int {
	return a + b
}

func Mayor(a, b int) int {
	if a == b {
		return a
	}

	if a > b {
		return a
	}

	return b
}

// Sumar los numeros ante
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a + b
	}
	return b
}
*/

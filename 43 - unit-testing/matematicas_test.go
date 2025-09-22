// Es necesario crear el archivo con _test para que go reconozca como un archivo de tipo test

package main

import "testing"

func TestSumar(t *testing.T) {
	/*
		r := Sumar(2, 3)
		e := 5
		if r != e {
			t.Errorf("Error al sumar, se esperaba %d, pero se obtuvo %d", e, r)
		}
	*/

	cases := []struct {
		a, b, e int
	}{
		{3, 2, 5},
		{3, 3, 6},
		{10, 10, 20},
		{10, -1, 9},
		{0, 0, 0},
		{-0, 0, 0},
		{-10, 5, -5},
		{-20, -10, -30},
	}

	for _, c := range cases {
		result := Sumar(c.a, c.b)
		if result != c.e {
			t.Errorf("Error al sumar, se esperaba %d, pero se obtuvo %d", c.e, result)
		}

	}
}

func TestMayor(t *testing.T) {

	cases := []struct {
		a, b, e int
	}{
		{3, 2, 3},
		{0, 5, 5},
		{3, 3, 3},
		{10, 10, 10},
		{10, -1, 10},
		{0, 0, 0},
		{-0, 0, 0},
		{-10, 5, 5},
		{-20, -10, -10},
	}

	for _, c := range cases {
		result := Mayor(c.a, c.b)
		if result != c.e {
			t.Errorf("Error al comprobar el numero mayor, se esperaba %d, pero se obtuvo %d", c.e, result)
		}

	}
}

func TestFibonacci(t *testing.T) {

	cases := []struct {
		a, r int
	}{
		{0, 1},
		{1, 1},
		{7, 13},
		{10, 55},
		{50, 12586269025},
	}

	for _, c := range cases {
		result := Fibonacci(c.a)
		if result != c.r {
			t.Errorf("Error al calcular Fibonacci, se esperaba %d, pero se obtuvo %d", c.a, result)
		}

	}
}

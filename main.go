package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "Suma dos num.")
	subs := flag.Bool("subs", false, "resta dos num.")
	mult := flag.Bool("mult", false, "multiplica dos num.")
	div := flag.Bool("div", false, "divide dos num.")

	flag.Parse()

	var a, b float64
	fmt.Println("Ingresa el 1er num.")
	fmt.Scan(&a)
	fmt.Println("Ingresa el 2do num.")
	fmt.Scan(&b)

	switch {
	case *add:
		fmt.Printf("Suma %g \n", addition(a, b))
	case *subs:
		fmt.Printf("Resta %g \n", subtract(a, b))
	case *mult:
		fmt.Printf("Multiplicar %g \n", multiply(a, b))
	case *div:
		fmt.Printf("Dividir %g \n", division(a, b))
	default:
		fmt.Fprintln(os.Stderr, "Error de operador, intenta con add, subs, mult, div.")
		os.Exit(1)
	}

}

func addition(a, b float64) float64 {
	return a + b
}
func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func division(a, b float64) float64 {
	return a / b
}

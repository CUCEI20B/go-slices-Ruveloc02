package main

import "fmt"

func main() {
	var n int
	var suma float64
	fmt.Scan(&n)
	numeros := make([]float64, n)
	for f := 0; f < len(numeros); f++ {
		fmt.Scan(&numeros[f])
		suma = suma + numeros[f]
	}
	fmt.Println(suma)
}

package main

import "fmt"

func main() {
	// Ponto de ebulição da água em Kelvin
	var kelvin float64 = 373.0

	// Converte para Celsius
	celsius := kelvin - 273

	// Resultado
	fmt.Printf("Temperatura em Kelvin: %.2f K\n", kelvin)
	fmt.Printf("Temperatura em Celsius: %.2f °C\n", celsius)
}
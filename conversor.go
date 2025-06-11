package main

import "fmt"

// A temperatura de ebulição da água em Kelvin
const ebulicaoK = 373.2

func main() {
	var tempK = ebulicaoK

	// Converte a temperatura de Kelvin para Celsius
	tempC := tempK - 273.2

	// Exibe a temperatura de ebulição da água em Kelvin e Celsius
	fmt.Printf("A temperatura de ebulição da água é %.2f K ou %.2f C\n", tempK, tempC)

}

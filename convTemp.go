package main

import "fmt"

const ebulicaoK = 373

func main() {
	var tempK = ebulicaoK
	var tempC = tempK - 273
	fmt.Printf("O ponto de ebulição da água em ºK = %v e o ponto de ebulição da águua em ºC = %v", tempK, tempC)
}

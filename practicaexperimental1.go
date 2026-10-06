package main

import "fmt"

func ejercicioArrayPrimos() {
	primos := [5]int{2, 3, 5, 7, 11}

	fmt.Println(primos[0])
	fmt.Println(primos[2])
	fmt.Println(primos[4])
}

func ejercicioArrayPromedio() {
	muestras := [3]float64{71.8, 56.2, 89.5}
	suma := 0.0

	for _, valor := range muestras {
		suma += valor
	}

	promedio := suma / float64(len(muestras))
	fmt.Printf("El promedio es igual a %.2f\n", promedio)
}

func slicePrimos() {
	primos := []int{2, 3, 5}
	primos = append(primos, 7, 11)

	fmt.Println("Primos", primos, "Corte", primos[1:4])
}

func ejercicioMap() {
	votos := []string{"Amber", "Bryan", "Bryan", "Amber"}
	conteo := make(map[string]int)

	for _, v := range votos {
		conteo[v]++
	}

	fmt.Println("Votos", conteo, votos)
}

func main() {
	ejercicioArrayPrimos()
	ejercicioArrayPromedio()
	slicePrimos()
	ejercicioMap()
}

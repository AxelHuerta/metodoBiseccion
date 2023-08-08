package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64 = 0
	var x2 float64 = 0
	var c float64 = 0
	var intervalos [][]float64

	// intervalo
	intervalo := []int{0, 100}
	x1 = float64(intervalo[0])
	x2 = float64(intervalo[1])

	// dividir el intervalo (si es grande) en intervalos mas pequenos
	for x2-x1 > 5 {
		c = x1 + 5
		var fila = []float64{x1, c}
		intervalos = append(intervalos, fila)
		x1 = c
	}

	var fila = []float64{x1, x2}
	intervalos = append(intervalos, fila)

	// iniciar metodo
	for i := 0; i < len(intervalos); i++ {
		iniciarMetodo(intervalos[i])
	}

}

// iniciar metodo
func iniciarMetodo(intervalos []float64) {
	// n = numero de iteraciones
	var n = 100
	var x1 = intervalos[0]
	var x2 = intervalos[1]

	buscarRaices(x1, x2, n)
}

func buscarRaices(x1, x2 float64, n int) {
	if n <= 0 {
		fmt.Printf("Hay una raíz en %f, %f \n", x1, x2)
		return
	}

	// valor medio
	var m = valormedio(x1, x2)
	// funciones
	var f1 = funcion(x1)
	var f2 = funcion(x2)
	var fm = funcion(m)

	if f1*fm == 0 || fm*f2 == 0 {
		fmt.Println("Hay una raíz en 0")
		if f1*fm == 0 && fm*f2 == 0 {
			return
		}

		// dado que ya sabemos que hay una raiz en cero, cambiamos el intervalo
		if x1 == 0 {
			x1 = x1 + math.SmallestNonzeroFloat64
			f1 = funcion(x1)
		}
	}

	if f1*fm < 0 && fm*f2 < 0 {
		// la raiz esta en ambos intervalos
		buscarRaices(x1, m, n)
		buscarRaices(m, x2, n)
	} else if f1*fm < 0 {
		buscarRaices(x1, m, n-1)
	} else if fm*f2 < 0 {
		buscarRaices(m, x2, n-1)
	}

}

// calcular valor medio
func valormedio(x1, x2 float64) float64 {
	return (x1 + x2) / 2
}

// aquí tenemos nuestra funcion
func funcion(x float64) float64 {
	return math.Sin(x)
}

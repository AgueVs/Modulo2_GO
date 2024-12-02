package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AnyOf devuelve verdadero si algún elemento de la lista satisface el predicado.
func AnyOf[T any](slice []T, pred func(T) bool) bool {

	fmt.Printf("el contenido de []T: %T %v - len:%v - cap:%v\n", slice, slice, len(slice), cap(slice))

	//Ponemos esto _ porque no lo necesitamos pero corresponde con la posicion.
	for _, v := range slice {
		println("Imprimimos la variable v", v)
		if pred(v) {
			println("Hemos encontrado un numero que cumple la condicion:", v)
			return true
		}
	}
	println("No hemos encontrado ningun numero que cumpla la condicion")
	return false
}

// FindIf devuelve el índice del primer elemento que satisface el predicado, o -1 si no hay ninguno.
func FindIf[T any](slice []T, pred func(T) bool) (int, bool) {

	fmt.Printf("el contenido de []T: %T %v - len:%v - cap:%v\n", slice, slice, len(slice), cap(slice))

	for i, v := range slice {
		if pred(v) {
			println("Imprimimos el valor que estamos buscando su posicion", v)
			println("Imprimimos la posicion que cumpla la condicion de la funcion (recordar que la posicion empieza en 0): ", i)
			return i, true
		}
	}
	return -1, false
}

/*
// AdjacentFind devuelve el índice del primer par adyacente de elementos que satisface el predicado, o -1 si no hay tal par.
func AdjacentFind[T any](slice []T, pred func(T, T) bool) int {
	for i := 0; i < len(slice)-1; i++ {
		if pred(slice[i], slice[i+1]) {
			return i
		}
	}
	return -1
}
*/

// Equal comprueba si dos listas son iguales.
func Equal[T comparable](slice1, slice2 []T) bool {

	fmt.Printf("el contenido de slice1 []T: %T %v - len:%v - cap:%v\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("el contenido de slide2 []T: %T %v - len:%v - cap:%v\n", slice2, slice2, len(slice2), cap(slice2))

	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
/*
// ReplaceIf reemplaza los elementos que satisfacen el predicado por un nuevo valor.
func ReplaceIf[T any](slice []T, newValue T, pred func(T) bool) int {
	replaced := 0
	for i, v := range slice {
		if pred(v) {
			slice[i] = newValue
			replaced++
		}
	}
	return replaced
}

// RemoveIf elimina los elementos que satisfacen el predicado y devuelve el nuevo tamaño de la lista.
func RemoveIf[T any](slice []T, pred func(T) bool) []T {
	result := slice[:0]
	for _, v := range slice {
		if !pred(v) {
			result = append(result, v)
		}
	}
	return result
}

*/
// IsSorted comprueba si la lista está ordenada. 
func IsSorted[T constraints.Ordered](slice []T) bool {

        fmt.Printf("el contenido de []T: %T %v - len:%v - cap:%v\n", slice, slice, len(slice), cap(slice))
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

// Merge combina dos listas ordenadas en una nueva lista ordenada.
func Merge[T constraints.Ordered](slice1, slice2 []T) []T {

	fmt.Printf("el contenido de slice1 []T: %T %v - len:%v - cap:%v\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("el contenido de slice2 []T: %T %v - len:%v - cap:%v\n", slice2, slice2, len(slice2), cap(slice2))

	result := make([]T, 0, len(slice1)+len(slice2))
	i, j := 0, 0
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] <= slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}
	result = append(result, slice1[i:]...)
	result = append(result, slice2[j:]...)
	return result
}


func main() {

	// se define aqui una funcion generica, "func(x int) bool {}"
	booleano := AnyOf([]int{2, 1, 4, 6}, func(x int) bool { return x > 4 })
	fmt.Println("Algun numero cumple la condicion de la funcion AnyOf que le estamos pasando? ", booleano) // true porque hay algun numero mayor de 2 (3,4)
	booleano2 := AnyOf([]int{1, 2, 3, 4}, func(x int) bool { return x > 5 })
	fmt.Println("Algun numero cumple la condicion de la funcion AnyOf que le estamos pasando? ", booleano2) // false porque no porque no hay ningun numero mayor de 5
	booleano3 := AnyOf([]string{"pepe", "juan", "lola", "cuadrado"}, func(x string) bool { return x == "pepe" })
	fmt.Println("Algun numero cumple la condicion de la funcion AnyOf que le estamos pasando? ", booleano3) // true porque contiene la palabra pepe en la lista.

	fmt.Println("-----------------------------------------------------------------------------------")
	number, ok := FindIf([]int{8, 3, 1, 2}, func(x int) bool { return x == 2 })
	if ok == true {
		fmt.Println("Usando la funcion FindIf que nos devuelve: ", number)
	} else {
		fmt.Println("No esta el numero en la lista y no puedo encontrar su posicion porque no existe el numero en la lista")
	}
	number2, ok2 := FindIf([]int{10, 9, 8, 4}, func(x int) bool { return x == 2 })
	if ok2 == true {
		fmt.Println("Usando la funcion FindIf que nos devuelve: ", number2)
	} else {
		fmt.Println("No esta el numero en la lista y no puedo encontrar su posicion porque no existe el numero en la lista")
	}

	number3, ok3 := FindIf([]string{"maria", "pepa", "TOMATE", "SABEMOS"}, func(x string) bool { return x == "TOMATE" })
	if ok3 == true {
		fmt.Println("Usando la funcion FindIf que nos devuelve: ", number3)
	} else {
		fmt.Println("No esta el numero en la lista y no puedo encontrar su posicion porque no existe el numero en la lista")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	/*
	   fmt.Println(AdjacentFind([]int{1, 2, 2, 4}, func(x, y int) bool { return x == y })) // 1
	*/
	fmt.Println("-----------------------------------------------------------------------------------")

	fmt.Println("Aqui la funcion Equal: ")
	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 3}
	fmt.Println(Equal(num1, num2)) // true

	a := []complex64{1 + 2i, 3 + 4i, 5 + 6i}
	b := []complex64{1 + 2i, 3 + 4i, 5 + 6i}
	c := []complex64{1 + 2i, 3 + 4i, 6 + 7i}

	fmt.Println(Equal(a, b)) // true
	fmt.Println(Equal(a, c)) // false

	fmt.Println("-----------------------------------------------------------------------------------")

/*	
	   slice := []int{1, 2, 3, 4}
	   fmt.Println(ReplaceIf(slice, 0, func(x int) bool { return x > 2 })) // 2
	   fmt.Println(slice) // [1, 2, 0, 0]
	   fmt.Println(RemoveIf([]int{1, 2, 3, 4}, func(x int) bool { return x > 2 })) // [1, 2]
*/	   
           fmt.Println("-----------------------------------------------------------------------------------")

           fmt.Println("La lista esta ordenada:")
           fmt.Println(IsSorted([]string{"CASA", "PEPE", "LOLA"})) //false
	   fmt.Println(IsSorted([]string{"CASA", "LOLA", "PEPA"})) //true
	   fmt.Println(IsSorted([]int{1, 2, 3, 4})) // true
	   // Para ordenar los numeros complejos debe ser primero la parte real y despues la parte imaginaria.
	   // Como vamos a aplicar para esta funcion IsSorted los numeros complejos?
//           order := []complex64{1 + 2i, 3 + 4i, 5 + 6i}
//           fmt.Println(IsSorted(order))

           fmt.Println("-----------------------------------------------------------------------------------")
	   
	   fmt.Println("Mergear la listas:" )
	   fmt.Println(Merge([]int{1, 3, 5}, []int{2, 4, 6})) // [1, 2, 3, 4, 5, 6]
           fmt.Println(Merge([]string{"HOLA","QUE"}, []string{"POR?","NOSE"}))
	   // Para ordenar los numeros complejos debe ser primero la parte real y despues la parte imaginaria.
           // Como vamos a aplicar para esta funcion Merge los numeros complejos?
//	   aa := []complex64{1 + 2i, 3 + 4i, 5 + 6i}
//         cc := []complex64{2 + 5i, 5 + 6i, 6 + 7i}
//	   fmt.Println(Merge(aa, cc))


}

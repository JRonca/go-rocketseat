package main

import "fmt"

func main() {
	m := make(map[string]string)
	mapMake := make(map[string]string, 10)
	mapLiteral := map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
		"chave3": "valor3",
	}
	fmt.Println(m, mapMake, mapLiteral)
	slicemap := map[int][]int{
		1: {1, 2, 3},
		2: {4, 5, 6},
		3: {7, 8, 9},
	}
	fmt.Println(slicemap)
	fmt.Println(slicemap[1])
	fmt.Println(mapLiteral["chave1"])
	valueTrue, ok := mapLiteral["chave1"]
	valueFalse, noOk := mapLiteral["chave5"]
	fmt.Println(valueTrue, ok)
	fmt.Println(valueFalse, noOk)
	fmt.Println(mapLiteral["chave3"])
	delete(mapLiteral, "chave3")
	fmt.Println(mapLiteral["chave3"])
	fmt.Println(mapLiteral)
	clear(mapLiteral)
	fmt.Println(mapLiteral)

	for key, value := range slicemap {
		fmt.Printf("Key: %d Value: %v\n", key, value)
	}
}

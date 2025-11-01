package main

import "fmt"

func main() {
	fmt.Println(ToRNA("GC"))
}

type RnaModel struct {
	key   string
	value string
}

// result
func ToRNA(dna string) string {
	var result string
	DnaStorage := []RnaModel{
		{"G", "C"},
		{"C", "G"},
		{"T", "A"},
		{"A", "U"},
	}
	// fmt.Println(DnaStorage)
	for _, char := range dna {
		fmt.Println("first", string(char))
		for index := range DnaStorage {

			// fmt.Println(key, value)
			// fmt.Println(DnaStorage[index].key)
			if string(char) == DnaStorage[index].key {
				result += DnaStorage[index].value
			}
		}
	}
	return result
}

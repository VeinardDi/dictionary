package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	mapDIC1 := map[string]string{"машина":"car", "тарелка":"dish", "ручка":"pen", "книга":"book", "лампа":"lamp", "хоккей":"hockey",
	"мяч":"bal", "ружье":"gun","мясо":"meat", "картошка":"potato"}
	mapDIC2, _ := json.Marshal(mapDIC1)
	fmt.Println(string(mapDIC2))

	file, err := os.Create("dictionary.json")
	if err != nil {
		return
	}
	file.Write(mapDIC2)
}


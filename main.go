package main

import ("encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	size, err := file.Stat()
	if err != nil {
		return
	}

	array := make([]byte, size.Size())
	_, err = file.Read(array)
	if err != nil {
		return
	}

	s := string(array)
	re := regexp.MustCompile("[a-яА-ЯёЁ]+")
	newSTRINGS := re.FindAllString(s, -1)

	f, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var dictionaryMAP map[string]string
	var inputARRAY []string

	err = json.Unmarshal(f, &dictionaryMAP)
	fmt.Println("ERROR parsing", err)

	for i := 0; i<10; i++ {
		if k, ok := dictionaryMAP[newSTRINGS[i]]; ok {
			inputARRAY = append(inputARRAY, k)
		}
	}

	outputSTRINGS := strings.Join(inputARRAY, " ")
	output, err:= os.Create("outout.txt")
	if err != nil {
		return
	}
	output.WriteString(outputSTRINGS)
	defer output.Close()

}
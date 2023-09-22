package paytable

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type PaytableType map[int][]int

var Paytable PaytableType

func ReadFromFile(filename string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	Paytable = make(map[int][]int)
	slice := strings.Split(string(fileBytes), "\n")

	for i := range slice {
		digitsSlice := strings.Split(slice[i], ", ")
		for j := range digitsSlice {
			val, err := strconv.Atoi(digitsSlice[j])
			if err != nil {
				log.Fatal(err)
			}

			Paytable[i] = append(Paytable[i], val)
		}
	}
}

package paytable

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Description struct {
	MultiplySymbol int     `json:"multiply_symbol"`
	WildSymbol     int     `json:"wild_symbol"`
	ScatterSymbol  int     `json:"scatter_symbol"`
	Paytable       [][]int `json:"paytable"`
	FreeGames      []int   `json:"free_games"`
}

var Desc Description

func ParseDescriptionJSON(filename string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileBytes, &Desc)
	if err != nil {
		log.Fatal(err)
	}
}

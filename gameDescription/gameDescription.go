package gameDescription

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Description struct {
	MultiplySymbol  int      `json:"multiply_symbol"`
	MultiplyValue   int      `json:"multiply_value"`
	MultiplyReelNum int      `json:"multiply_reel"`
	WildSymbol      int      `json:"wild_symbol"`
	ScatterSymbol   int      `json:"scatter_symbol"`
	Paytable        [][]int  `json:"paytable"`
	FreeGames       []int    `json:"free_games"`
	Lines           [][]int  `json:"lines"`
	ScatterType     string   `json:"scatter_type"`
	ScatterTypes    []string `json:"scatter_type_enum"`
}

var Desc Description

func ParseDescriptionJSON(filename string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(fileBytes, &Desc)
	if err != nil {
		log.Panic(err)
	}

	for i := range Desc.ScatterTypes {
		if Desc.ScatterType == Desc.ScatterTypes[i] {
			return
		}
	}

	log.Panicf("undefined scatter type: \"%v\" not found in scatter types\n", Desc.ScatterType)
}

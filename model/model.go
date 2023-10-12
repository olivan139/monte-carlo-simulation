package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Description struct {
	NumOfIterations int      `json:"num_of_iterations"`
	MultiplySymbol  int      `json:"multiply_symbol"`
	MultiplyValue   int      `json:"multiply_value"`
	MultiplyReelNum int      `json:"multiply_reel"`
	WildSymbol      int      `json:"wild_symbol"`
	ScatterSymbol   int      `json:"scatter_symbol"`
	Paytable        [][]int  `json:"paytable"`
	FreeGames       []int    `json:"free_games"`
	Lines           [][]int  `json:"lines"`
	Reels           [][]int  `json:"reels"`
	FreeGamesReels  [][]int  `json:"free_games_reels"`
	ScatterType     string   `json:"scatter_type"`
	ScatterTypes    []string `json:"scatter_type_enum"`
}

type ParrotSpriteSheet struct {
	ParrotAnimFrames [][]string `json:"frames"`
}

var Model *Description
var Parrot *ParrotSpriteSheet

func ParseJSON(filename string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(fileBytes, &Model)
	if err != nil {
		log.Panic(err)
	}

	for i := range Model.ScatterTypes {
		if Model.ScatterType == Model.ScatterTypes[i] {
			return
		}
	}

	log.Panicf("undefined scatter type: \"%v\" not found in scatter types\n", Model.ScatterType)
}

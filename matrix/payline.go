package matrix

import (
	"monte-carlo-simulation/helper"
	"monte-carlo-simulation/model"
)

func GetLinePayoff(winLine []int) int {

	frstSymbol := winLine[0]
	wildCount := 0
	wildCountAsLine := 0
	wildAsLine := false
	symbolCount := 0
	mainSymbol := -1

	if frstSymbol == model.Model.WildSymbol {
		wildCount++
		wildCountAsLine++
		wildAsLine = true
	} else {
		symbolCount++
		mainSymbol = frstSymbol
	}

	for i := 1; i < len(winLine); i++ {
		if winLine[i] != model.Model.WildSymbol {
			wildAsLine = false

			if mainSymbol == -1 {
				mainSymbol = winLine[i]
				symbolCount++
			} else if mainSymbol == winLine[i] {
				symbolCount++
			} else {
				break
			}
		} else {
			if wildAsLine {
				wildCountAsLine++
			}

			wildCount++
		}
	}

	if mainSymbol == -1 {
		return model.Model.Paytable[model.Model.WildSymbol][wildCountAsLine]
	}
	if mainSymbol == model.Model.ScatterSymbol {
		return model.Model.Paytable[model.Model.ScatterSymbol][symbolCount]
	}
	return helper.Max(model.Model.Paytable[mainSymbol][symbolCount+wildCount],
		model.Model.Paytable[model.Model.WildSymbol][wildCountAsLine])
}

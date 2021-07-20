package controller

import (
	"fmt"
	"os"
	"qt-test/helpers"
	"qt-test/state"
)

// returns data [ [ 'url', 'status', ... ], ... ]
func LoadFormattedPings() [][]string {
	pings := state.GetPingResults()

	var formattedPings [][]string

	for _, pingResult := range pings {
		formattedPings = append(formattedPings, []string{
			pingResult.Url,
			fmt.Sprintf("%d", pingResult.StatusCode),
			helpers.ConvertUnixToString(pingResult.SentAt),
			helpers.ConvertUnixToString(pingResult.FinishedAt),
		})
	}

	return formattedPings
}

func SavePingsToCSV(savePath string) {
	f, err := os.Create(savePath)

	if err != nil {
		f.Close()
		// todo: show message box!
		panic("No file exists!")
	}

	// get current state
	pings := state.GetPingResults()

	// put headings
	fmt.Fprintln(f, fmt.Sprintf("%s;%s;%s;%s;", "Url", "Http status", "Started at", "Finished at"))

	// put pings
	for _, pingResult := range pings {
		fmt.Fprintln(f, fmt.Sprintf(
			"%s;%d;%s;%s;",
			pingResult.Url,
			pingResult.StatusCode,
			helpers.ConvertUnixToString(pingResult.SentAt),
			helpers.ConvertUnixToString(pingResult.FinishedAt)))
	}

	err = f.Close()
	if err != nil {
		panic("Not saved!")
	}
}

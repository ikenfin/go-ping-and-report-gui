package controller

import (
    "qt-test/state"
    "fmt"
    "os"
    "time"
)

// returns data [ [ 'url', 'status', ... ], ... ]
func LoadFormattedPings() [][]string {
    pings := state.GetPingResults()

    var formattedPings [][]string

    for _, pingResult := range pings {
        formattedPings = append(formattedPings, []string{
            pingResult.Url,
            fmt.Sprintf("%d", pingResult.StatusCode),
            convertUnixToString(pingResult.SentAt),
            convertUnixToString(pingResult.FinishedAt),
        })
    }

    return formattedPings
}

func SavePingsToCSV (savePath string) {
    f, err := os.Create(savePath)

    if err != nil {
        f.Close()
        panic("No file exists!")
    }

    // get current state
    pings := state.GetPingResults()
    // put headings
    fmt.Fprintln(f, fmt.Sprintf("%s;%s;%s;%s;", "Url", "Http status", "Started at", "Finished at"))

    // put p
    for _, pingResult := range pings {
        fmt.Fprintln(f, fmt.Sprintf("%s;%d;%s;%s;", pingResult.Url, pingResult.StatusCode, convertUnixToString(pingResult.SentAt), convertUnixToString(pingResult.FinishedAt)))
    }

    err = f.Close()
    if err != nil {
        panic("Not saved!")
    }
}
/*
func GetFormattedPingResults () {
    pings := state.GetPingResults()

}
*/
func convertUnixToString (timestamp int64) string {
    return time.Unix(timestamp, 0).Format(time.RFC3339) 
}

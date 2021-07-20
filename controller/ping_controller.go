package controller

import (
    "math"
    "qt-test/state"
    "qt-test/logic"
)


// onProgress -> 0..100
func PingUrlsByList (urls []string, onProgress chan int) {
    total := len(urls)
    step := calculatePercentagePerUrl(total)

    // initial
    onProgress <- 0

    // signalizer for single parsed url
    onUrlPingComplete := make(chan *state.PingResult)

    // make tasks
    for _, url := range urls {
        go logic.Ping(url, onUrlPingComplete)
    }

    // processed counter
    processedUrls := 0

    for {
        urlPingResult := <-onUrlPingComplete

    state.StorePingResult(urlPingResult)

        processedUrls++

        // notify parent thread
        onProgress <- calculateProcessedPercentage(step, processedUrls)

        // seems we done
        if processedUrls >= total {
            onProgress <- 100
            close(onProgress)
            break
        }
    }

}

// 1 url in percents
func calculatePercentagePerUrl (totalUrls int) float64 {
    return float64(100 / totalUrls)
}

// convert total pinged to percentage
func calculateProcessedPercentage (step float64, total int) int {
    return int(math.Round(step * float64(total)))
}

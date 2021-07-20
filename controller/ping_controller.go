package controller

import (
    "qt-test/state"
    "qt-test/logic"
    "qt-test/helpers"
)


// onProgress -> 0..100
func PingUrlsByList (urls []string, onProgress chan int) {
    total := len(urls)
    step := helpers.CalculatePercentPerItem(total)

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
        onProgress <- helpers.CalculateProcessedPercentage(step, processedUrls)

        // seems we done
        if processedUrls >= total {
            onProgress <- 100
            close(onProgress)
            break
        }
    }

}



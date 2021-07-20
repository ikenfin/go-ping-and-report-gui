package logic

import (
    "time"
    "net/http"
    "qt-test/state"
)


func Ping (url string, res chan *state.PingResult) {
    response, err := http.Get(url)
    result := &state.PingResult{
        Url: url,
        SentAt: time.Now().Unix(),
    }

    if err != nil {
        result.StatusCode = -1
        res <- result
        return
    }

    result.StatusCode = response.StatusCode
    result.FinishedAt = time.Now().Unix()

    res <- result
}

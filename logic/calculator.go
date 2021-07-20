package logic

import (
	"fmt"
	"time"
)

func Calculate(response chan int) {
    for i := 0; i <= 100; i++ {
        time.Sleep(time.Millisecond)
        fmt.Println(i)
        response <- i
    }

    fmt.Println("close")
    close(response)
}



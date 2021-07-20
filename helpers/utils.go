/*
    Common helper functions
*/
package helpers

import (
	"math"
	"time"
)

// 1 item in percents
func CalculatePercentPerItem (total int) float64 {
    return float64(100 / total)
}

// convert total pinged to percentage
func CalculateProcessedPercentage (step float64, total int) int {
    return int(math.Round(step * float64(total)))
}

// convert unix timestamp to RFC date string
func FormatUnixToString (timestamp int64) string {
    return time.Unix(timestamp, 0).Format(time.RFC3339)
}

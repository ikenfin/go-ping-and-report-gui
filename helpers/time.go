package helpers

import "time"

func ConvertUnixToString (timestamp int64) string {
    return time.Unix(timestamp, 0).Format(time.RFC3339) 
}

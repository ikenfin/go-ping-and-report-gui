package state

type PingResult struct {
    Url string
    StatusCode int
    SentAt int64
    FinishedAt int64
}

type PingResultCache struct {
    Pings []PingResult
}

var db *PingResultCache

func Init () bool {
    db = &PingResultCache{}
    return true
}

func StorePingResult (r *PingResult) {
    db.Pings = append(db.Pings, *r)
}

func GetPingResults () []PingResult {
    return db.Pings
}

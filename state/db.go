package state

import (
    //"fmt"
    //"github.com/therecipe/qt/sql"
)

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


/*
var db *sql.QSqlDatabase

func Init () bool {
    db = sql.QSqlDatabase_AddDatabase("QSQLITE", "")

    db.SetDatabaseName(":memory:")
    if !db.Open() {
        widgets.QMessageBox_Critical(nil, "Cannot open database", "", widgets.QMessageBox__Cancel, 0)
        return false
    }

    query := sql.NewQSqlQuery3(db)

    query.Exec(`
        CREATE TABLE pings (
            id INT PRIMARY KEY, 
            url VARCHAR(2100) NOT NULL, 
            status_code INT NOT NULL, 
            sent_at TIMESTAMP NOT NULL, 
            finished_at TIMESTAMP NOT NULL
        )
    `)

    return true
}

func StorePingResult (p *PingResult) bool {
    query := sql.NewQSqlQuery3(db)

    sqlQ := `
        INSERT INTO pings (url, status, sent_at, finished_at) 
        VALUES ("%s", %d, %d, %d)
    `

    return query.Exec(fmt.Sprintf(sqlQ, p.Url, p.StatusCode, p.SentAt, p.FinishedAt))
}

func GetPingResults () {
    query := sql.NewQSqlQuery3(db)

    query.Exec(fmt.Sprintf(`
        SELECT url, status, sent_at, finished_at
        FROM pings
    `))
}
*/

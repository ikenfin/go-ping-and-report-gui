package ui
import (
    "fmt"
    "qt-test/controller"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func (w *PingResultModal) init () {
    w.ConnectOpen(func () {
        fmt.Println("open")
    })

    tableRows := controller.LoadFormattedPings()

    for rowNum, columns := range tableRows {
        for colNum, cellData := range columns {
            cell := widgets.NewQTableWidgetItem2(cellData, 0)
            w.PingResultTable.SetItem(rowNum, colNum, cell)
        }
    }
}

func (w *PingResultModal) ExportToCsv () {
    saveDialog := widgets.NewQFileDialog(w, core.Qt__Dialog)
    csvFileName := saveDialog.GetSaveFileName(nil, "Export to CSV", "pings.csv", "CSV files (*.csv *.tsv)", "", widgets.QFileDialog__ShowDirsOnly)
    saveDialog.Close()

    controller.SavePingsToCSV(csvFileName)
}

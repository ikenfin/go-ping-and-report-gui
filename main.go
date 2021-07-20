package main
import (
    "os"
    "github.com/therecipe/qt/widgets"
    "qt-test/ui"
    "qt-test/state"
)

func main() {
    widgets.NewQApplication(len(os.Args), os.Args)

    if !state.Init() {
	widgets.QMessageBox_Critical(nil, "Cannot init the database", "", widgets.QMessageBox__Cancel, 0)

	return
    }

    u := ui.NewMainWindow(nil)
    u.Show()
    widgets.QApplication_Exec()
}

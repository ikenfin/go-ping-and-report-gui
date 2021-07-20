package ui

import (
    "qt-test/controller"
    "strings"
)


func (w *MainWindow) init () {
    w.resetProgress()
    w.enableControls()
    // Connect actions
    w.ActionExit.ConnectTriggered(func (bool) { w.Close(); })
    w.ActionAbout.ConnectTriggered(func (bool) {
        NewAboutModal(w).Show() 
    })
}

func (w MainWindow) resetProgress () {
    w.Progress.SetValue(0)
}

func (w MainWindow) disableControls () {
    w.Ok.SetEnabled(false)
    w.Cancel.SetEnabled(false)
}
func (w MainWindow) enableControls () {
    w.Ok.SetEnabled(true)
    w.Cancel.SetEnabled(true)
}

func (w *MainWindow) SayHello () {
    urls := strings.Split(w.HostsInput.ToPlainText(), "\n")

    w.disableControls()
    w.resetProgress()

    onProgress := make(chan int)

    go controller.PingUrlsByList(urls, onProgress)

    for completedPercent := range onProgress {
        w.Progress.SetValue(completedPercent)
    }

    w.enableControls()
    // open results
    NewPingResultModal(w).Show()
}

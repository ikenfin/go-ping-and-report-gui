package helper

import (
    "github.com/therecipe/qt/core"
    "fmt"
)

type MainThreadHelper struct {
    core.QObject

    _ func(f func()) `signal:"runOnMain",auto`
}

func (*MainThreadHelper) runOnMain (f func()) { fmt.Println("runOnMainCalled"); f(); }


package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Press F10 to trigger the command")
	// 提示Ctrl+C退出
	fmt.Println("Press Ctrl+C to exit")
	for {
		if robotgo.AddEvent("f10") { // 监听F10
			robotgo.Click()                    // 鼠标单击
			time.Sleep(10 * time.Millisecond)  // 延时10ms
			robotgo.Click()                    // 鼠标单击
			time.Sleep(10 * time.Millisecond)  // 延时10ms
			robotgo.Toggle("left")             // 鼠标左键按下
			time.Sleep(500 * time.Millisecond) // 延时500ms
			robotgo.Toggle("left", "up")       // 鼠标左键弹起
		}
	}
}

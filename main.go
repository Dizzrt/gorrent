package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start")

	// debug
	a := app.New()
	window := a.NewWindow("gorrent")
	window.SetContent(widget.NewLabel("hello fyne"))
	window.ShowAndRun()

	// file, err := os.Open("C:/Users/Dizzrt/Desktop/s.torrent")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// bt, err := torrent.NewWithReader(file)
	// if err == nil {
	// 	fmt.Println(bt.Info.Name)
	// } else {
	// 	fmt.Println(err)
	// }
}

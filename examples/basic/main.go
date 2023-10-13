package main

import "github.com/itnxs/webview"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Thanks for using webview!")
	//w.Navigate("https://www.qq.com/")
	w.Run()
}

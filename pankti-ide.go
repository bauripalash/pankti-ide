package main

import (
	_ "embed"

	"github.com/webview/webview"
	"palash.bauri/pankti"
)

//go:embed webs/index.html
var index_html string

func main() {

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(480, 320, webview.HintNone)

	w.Bind("myfunc", func(text string) string {
		return pankti.Run(text)
	})
	w.SetHtml(index_html)
	w.Run()
	//fmt.Println(pankti.Run("1+2"))

}

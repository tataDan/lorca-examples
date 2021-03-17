/*
All the code in this application (except for main.go) is from the
https://github.com/jonasschmedtmann/complete-javascript-course/tree/master/07-Pig-Game/final
web site.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	go http.Serve(ln, http.FileServer(http.Dir(curDir)))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	<-ui.Done()
}

/*
func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Pig Game")
	w.SetSize(800, 600, webview.HintNone)

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	w.Navigate("file:///" + curDir + string(os.PathSeparator) + "index.html")

	w.Run()
}
*/

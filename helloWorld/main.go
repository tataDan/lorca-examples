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
	ui, err := lorca.New("", "", 400, 200)
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

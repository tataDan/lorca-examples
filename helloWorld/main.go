package main

import (
	"log"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ui, err := lorca.New("file:///"+curDir+string(os.PathSeparator)+"index.html", "",
		400, 200)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	<-ui.Done()
}

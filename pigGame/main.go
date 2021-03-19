/*
All the code in this application (except for main.go) is from the
https://github.com/jonasschmedtmann/complete-javascript-course/tree/master/07-Pig-Game/final
web site.
*/

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
		800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	<-ui.Done()
}

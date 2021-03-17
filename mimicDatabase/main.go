package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/zserge/lorca"
)

type country struct {
	name        string
	area        string
	nationalDay string
}

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

	countries := make([]country, 0, 256)

	ui.Bind("connect", func(password string) {
		if password != "password" {
			evalStr := fmt.Sprintf("alert(\"Closing application. Incorrect password was entered.\");")
			ui.Eval(evalStr)
			ui.Close()
		}
	})

	ui.Bind("query", func(matchType, queryValue string) {
		var rows []string
		var result string

		dashes := "-----------------------------------------------------------------------------"
		headings := fmt.Sprintf("%-50s%-15s%-15s\\n%s\\n", "Name", "Area", "National Day", dashes)

		for _, v := range countries {

			if matchType == "LIKE" {
				if strings.HasPrefix(v.name, queryValue) {
					row := fmt.Sprintf("%-50s%-15s%-15s\\n%s\\n", v.name, v.area, v.nationalDay, dashes)
					rows = append(rows, row)
				}
			} else {
				if v.name == queryValue {
					row := fmt.Sprintf("%-50s%-15s%-15s\\n%s\\n", v.name, v.area, v.nationalDay, dashes)
					rows = append(rows, row)
					break
				}
			}
		}
		joinedRows := strings.Join(rows, "")
		result = headings + joinedRows
		evalStr := fmt.Sprintf("updateTextArea(\"%s\");", result)
		ui.Eval(evalStr)
	})

	fileName := "countries.txt"

	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("Error is = ", e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var country country
		s := strings.Split(scanner.Text(), ",")
		country.name, country.area, country.nationalDay = s[0], s[1], s[2]
		countries = append(countries, country)
	}

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	go http.Serve(ln, http.FileServer(http.Dir(curDir)))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	<-ui.Done()
}

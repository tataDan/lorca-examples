// This application assumes that either mysql or mariadb is installed
// and running, and that the nation example database is installed.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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

	var db *sql.DB
	var dbErr error

	ui.Bind("connect", func(password string) {
		cs := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/nation", password)
		db, dbErr = sql.Open("mysql", cs)
		if dbErr != nil {
			log.Println("Failed to connect to database. Ensure that the database driver is 'mysql'")
			ui.Close()
		}
	})

	ui.Bind("query", func(matchType, queryValue string) {
		var country country
		var rows []string
		var result string

		var whereValue string
		if matchType == "LIKE" {
			whereValue = fmt.Sprintf("LIKE '%s%%'", queryValue)
		} else {
			whereValue = fmt.Sprintf(" = '%s'", queryValue)
		}
		query := fmt.Sprintf("SELECT name, area, IFNULL(national_day, 'N/A') FROM countries WHERE name %s",
			whereValue)
		countries, err := db.Query(query)
		if err != nil {
			evalStr := fmt.Sprintf("alert(\"Closing application. Failed to execute query. Possibly a incorrect password was used to connect to database.\");")
			ui.Eval(evalStr)
			ui.Close()
		}

		dashes := "-----------------------------------------------------------------------------"
		headings := fmt.Sprintf("%-50s%-15s%-15s\\n%s\\n", "Name", "Area", "National Day", dashes)

		for countries.Next() {
			err := countries.Scan(&country.name, &country.area, &country.nationalDay)
			if err != nil {
				log.Println("Failed to scan database results into country struct.")
				os.Exit(4)
			}
			row := fmt.Sprintf("%-50s%-15s%-15s\\n%s\\n", country.name, country.area, country.nationalDay, dashes)
			rows = append(rows, row)
		}

		joinedRows := strings.Join(rows, "")
		result = headings + joinedRows
		evalStr := fmt.Sprintf("updateTextArea(\"%s\");", result)
		ui.Eval(evalStr)
	})

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ui.Load("file:///" + curDir + string(os.PathSeparator) + "index.html")

	<-ui.Done()
}

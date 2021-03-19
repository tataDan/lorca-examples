package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/zserge/lorca"

	"github.com/ncruces/zenity"
)

type optionsType struct {
	CaseInsensitive bool
	WholeWord       bool
	WholeLine       bool
	FilenameOnly    bool
	Inverted        bool
}

func main() {
	ui, err := lorca.New("", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("search", func(pattern string, path string, options optionsType) {
		if options.WholeWord == true {
			pattern = `\b` + pattern + `\b`
		}
		if options.WholeLine == true {
			pattern = "^" + pattern + "$"
		}
		if options.CaseInsensitive == true {
			pattern = "(?i)" + pattern
		}

		var result string
		var err error

		evalStr := fmt.Sprintf("loadTextArea(\"%s\");", "")
		ui.Eval(evalStr)
		evalStr = fmt.Sprintf("updateStatus(\"%s\");", "Searching...")
		ui.Eval(evalStr)

		result, err = walkDir(path, pattern, options)
		if err != nil {
			errMsg := fmt.Sprintf("%s\n", err)
			zenity.Error(errMsg)
		}

		if result == "" {
			evalStr := fmt.Sprintf("loadTextArea(\"%s\");", "")
			ui.Eval(evalStr)
			evalStr = fmt.Sprintf("updateStatus(\"%s\");", "No results were found")
			ui.Eval(evalStr)
		} else {
			evalStr := fmt.Sprintf("loadTextArea(\"%s\");", result)
			ui.Eval(evalStr)
			evalStr = fmt.Sprintf("updateStatus(\"%s\");", "Success")
			ui.Eval(evalStr)
		}
	})

	ui.Bind("selectFolder", func() {
		pathSelected, err := zenity.SelectFile(zenity.Directory())
		if err != nil {
			errMsg := fmt.Sprintf("Error selecting a file. %s", err)
			zenity.Error(errMsg)
		}

		if pathSelected != "" {
			pathSelected = strings.ReplaceAll(pathSelected, `\`, `\\`)
			evalStr := fmt.Sprintf("setPath(\"%s\")", pathSelected)
			ui.Eval(evalStr)
		}
	})

	ui.Bind("selectFile", func() {
		pathSelected, err := zenity.SelectFile()
		if err != nil {
			errMsg := fmt.Sprintf("Error selecting a file. %s", err)
			zenity.Error(errMsg)
		}

		if pathSelected != "" {
			pathSelected = strings.ReplaceAll(pathSelected, `\`, `\\`)
			evalStr := fmt.Sprintf("setPath(\"%s\")", pathSelected)
			ui.Eval(evalStr)
		}
	})

	ui.Bind("error", func(msg string) {
		zenity.Error(msg)
		if err != nil {
			fmt.Println(err)
		}
	})

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ui.Load("file:///" + curDir + string(os.PathSeparator) + "index.html")

	<-ui.Done()
}

func walkDir(dirToWalk string, pattern string, options optionsType) (string, error) {
	var matches []string
	err := filepath.Walk(dirToWalk, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			matchesFromFile, err2 := checkFileForPattern(path, pattern, options)
			if err2 != nil {
				log.Printf("Failed opening file: %s", err2)
			} else {
				matches = append(matches, matchesFromFile...)
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	result := strings.Join(matches, "")
	return result, nil
}

func isBinary(fileToRead string) (bool, error) {
	data := make([]byte, 256)
	file, err := os.Open(fileToRead)
	if err != nil {
		return false, err
	}
	defer file.Close()
	count, err := file.Read(data)
	if err != nil {
		return false, err
	}
	for i := 0; i < count; i++ {
		if data[i] == 0 {
			return true, nil
		}
	}
	return false, nil
}

func checkFileForPattern(fileToRead string, pattern string, options optionsType) ([]string, error) {
	matches := make([]string, 0)
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Size() == 0 {
		return nil, nil
	}

	fileIsBinary, err := isBinary(fileToRead)
	if err != nil {
		return nil, err
	}
	if fileIsBinary {
		log.Printf("%s is binary\n", fileToRead)
		return nil, nil
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	if len(txtlines) == 0 {
		log.Printf("%s has no new line control characters.\n", fileToRead)
		return nil, nil
	}
	fileToRead = strings.ReplaceAll(fileToRead, `\`, `\\`)
	numNonMatches := 0
	numLinesMatched := 0
	for lineNum, line := range txtlines {
		if r.MatchString(line) {
			numLinesMatched++
			if options.Inverted == false {
				if options.FilenameOnly == true {
					if numLinesMatched == 1 {
						match := fmt.Sprintf("\\n%s\\n", fileToRead)
						matches = append(matches, match)
						break
					}
				}
				var printableLine string
				var sb strings.Builder
				for _, r := range line {
					if int(r) >= 32 && int(r) != 127 {
						if r == '\\' || r == '"' {
							sb.WriteRune('\\')
						}
						sb.WriteRune(r)
					}
				}
				printableLine = sb.String()
				if numLinesMatched == 1 {
					match := fmt.Sprintf("\\n%s\\n\\t%d: %s\\n", fileToRead, lineNum+1, printableLine)
					matches = append(matches, match)
				} else {
					match := fmt.Sprintf("\\t%d: %s\\n", lineNum+1, printableLine)
					matches = append(matches, match)
				}
			}
		} else {
			numNonMatches++
		}
	}

	if options.Inverted == true && numNonMatches == len(txtlines) {
		match := fmt.Sprintf("%s\\n\\n", fileToRead)
		matches = append(matches, match)
	}
	return matches, nil
}

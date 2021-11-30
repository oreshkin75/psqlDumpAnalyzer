package ui

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

var mu sync.Mutex

type exec struct {
	columns []string
	data    string
}

func (c *Creator) dumpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(c.config.HtmlDumpPage)
	if err != nil {
		http.Redirect(w, r, "/error", 301)
		c.logError.Print("Can't parse main page", err)
	}

	if c.columns == nil {
		err = tmpl.Execute(w, "Не получилось найти ваш запрос в дампе")
		if err != nil {
			http.Redirect(w, r, "/error", 301)
			c.logError.Print("Can't execute main page", err)
		}
		return
	}

	// TODO добавить чтение нескольких файлов
	err = c.dumpReader.OpenDumpFile(c.dumpFiles[0])
	defer c.dumpReader.CloseDumpFile()
	if err != nil {
		c.logError.Print(err)
		return
	}
	var allData []byte
	for {
		data, _, err := c.dumpReader.Read(1000)
		if err == io.EOF {
			break
		}
		go c.writeDumpToFile(data)
		if err != nil {
			panic(err)
		}
		if c.checkAllContains(string(data)) {
			allData = append(allData, data...)
		}
	}

	dataWithoutNulls := c.dumpReader.DeleteNulls(allData)
	clearData := c.dumpReader.DeleteUnprintableCharacters(dataWithoutNulls)

	var execute exec = exec{columns: c.columns, data: clearData}
	err = tmpl.Execute(w, execute)
	if err != nil {
		http.Redirect(w, r, "/error", 301)
		c.logError.Print("Can't execute main page", err)
	}
}

// проверка вхождения искомых данных
func (c *Creator) checkAllContains(data string) bool {
	var counter int
	for i, _ := range c.columns {
		if strings.Contains(data, c.columns[i]) {
			counter++
		}
	}

	if counter == len(c.columns) {
		return true
	} else {
		return false
	}
}

// запись дампа в файл
func (c *Creator) writeDumpToFile(data []byte) {
	mu.Lock()
	f, err := os.OpenFile(c.dumpFiles[0]+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		c.logError.Print(err)
	}
	mu.Unlock()
	dataWithouNulls := c.dumpReader.DeleteNulls(data)
	clearData := c.dumpReader.DeleteUnprintableCharacters(dataWithouNulls)

	mu.Lock()
	f.Write([]byte(clearData))
	mu.Unlock()
}

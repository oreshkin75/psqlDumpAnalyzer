package ui

import (
	"html/template"
	"net/http"
	"strings"
)

func (c *Creator) mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		query := r.FormValue("query")
		proc, err := c.platformWorker.FindAllPsqlProcesses()
		if err != nil {
			http.Redirect(w, r, "/error", 301)
		}
		c.logInfo.Print("all psql processes ", proc)
		if query != "" {
			//TODO выполнить запрос и сделать дамп
			c.psqlWorker.Connect()
			lowerQuery := strings.ToLower(query)
			if firstWord := strings.HasPrefix(lowerQuery, "select"); firstWord {
				c.columns, _ = c.psqlWorker.Select(lowerQuery)
			} else if firstWord := strings.HasPrefix(lowerQuery, "update"); firstWord {
				c.psqlWorker.Update(lowerQuery)
			} else if firstWord := strings.HasPrefix(lowerQuery, "insert"); firstWord {
				c.psqlWorker.Insert(lowerQuery)
			} else if firstWord := strings.HasPrefix(lowerQuery, "delete"); firstWord {
				c.psqlWorker.Delete(lowerQuery)
			}
			processes, err := c.platformWorker.FindQueryProcess()
			if err != nil {
				http.Redirect(w, r, "/error", 301)
			}
			c.logInfo.Print("query psql processes ", processes)
			files, _, _ := c.platformWorker.CreateDump()
			c.dumpFiles = nil
			c.dumpFiles = append(c.dumpFiles, files...)
			c.logInfo.Print("Create dump files ", files)

			http.Redirect(w, r, "/dump", 301)
		} else {
			// TODO сделать дамп
			http.Redirect(w, r, "/dump", 301)
		}
	} else {
		// вывод конфигурации
		tmpl, err := template.ParseFiles(c.config.HtmlMainPage)
		if err != nil {
			http.Redirect(w, r, "/error", 301)
			c.logError.Print("Can't parse main page", err)
		}
		err = tmpl.Execute(w, c.config)
		if err != nil {
			http.Redirect(w, r, "/error", 301)
			c.logError.Print("Can't execute main page", err)
		}
	}
}

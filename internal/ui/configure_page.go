package ui

import (
	"net/http"
)

func (c *Creator) configurePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", 301)
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			c.logError.Print("Can't parse form ", err)
			http.ServeFile(w, r, c.config.HtmlConfigurePage)
		}
		if tmp := r.FormValue("dllPath"); tmp != "" {
			c.config.DllPath = tmp
		}
		if tmp := r.FormValue("processName"); tmp != "" {
			c.config.NameProcess = tmp
		}
		if tmp := r.FormValue("ipPsql"); tmp != "" {
			c.config.PsqlIp = tmp
		}
		if tmp := r.FormValue("portPsql"); tmp != "" {
			c.config.PsqlPort = tmp
		}
		if tmp := r.FormValue("userPsql"); tmp != "" {
			c.config.PsqlUser = tmp
		}
		if tmp := r.FormValue("passwordPsql"); tmp != "" {
			c.config.PsqlPassword = tmp
		}
		if tmp := r.FormValue("dbasePsql"); tmp != "" {
			c.config.PsqlDBName = tmp
		}
		http.Redirect(w, r, "/main", 301)
	} else {
		http.ServeFile(w, r, c.config.HtmlConfigurePage)
	}
}

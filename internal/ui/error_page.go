package ui

import (
	"net/http"
)

func (c *Creator) errorPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, c.config.HtmlErrorPage)
}

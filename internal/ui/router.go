package ui

import (
	"fmt"
	"net/http"
)

func (c *Creator) StartWebUI(ip, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", c.configurePage)
	mux.HandleFunc("/main", c.mainPage)
	mux.HandleFunc("/error", c.errorPage)
	mux.HandleFunc("/dump", c.dumpPage)

	c.logInfo.Print("Web server start and listen on ", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		c.logError.Panic("Can't start web ui", err)
		return err
	}
	return nil
}

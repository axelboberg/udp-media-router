package http

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/axelboberg/go-rtp-repeater/internal/logger"
)

func ReturnOk (w http.ResponseWriter, r *http.Request) {
	logger.Info("HTTP", "Got request to /")
  fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

func Listen (port int) {
	logger.Info("HTTP", "Listening on port " + strconv.Itoa(port))
	http.HandleFunc("/", ReturnOk)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

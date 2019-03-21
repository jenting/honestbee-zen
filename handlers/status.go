package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/honestbee/Zen/config"
)

// StatusHandler handles status request
func StatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(struct {
		GoVersion  string `json:"go-version"`
		AppVersion string `json:"app-version"`
		ServerTime string `json:"server-time"`
	}{
		GoVersion:  runtime.Version(),
		AppVersion: config.Version,
		ServerTime: time.Now().UTC().String(),
	})
}

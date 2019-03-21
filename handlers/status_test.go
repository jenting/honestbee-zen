package handlers_test

import (
	"encoding/json"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/handlers"
)

func TestStatusHandler(t *testing.T) {
	testCases := [...]struct {
		description string
		expect      map[string]interface{}
	}{
		{
			description: "testing normal case",
			expect: map[string]interface{}{
				"go-version":  runtime.Version(),
				"app-version": config.Version,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			r := httptest.NewRequest("GET", "http://example.com/foo", nil)
			w := httptest.NewRecorder()
			handlers.StatusHandler(w, r, nil)

			actual := make(map[string]interface{})
			resp := w.Result()
			json.NewDecoder(resp.Body).Decode(&actual)

			if tt.expect["go-version"] != actual["go-version"] {
				t.Errorf("[%s] go version expect:%v != actual:%v",
					tt.description,
					tt.expect["go-version"],
					actual["go-version"],
				)
			}
			if tt.expect["app-version"] != actual["app-version"] {
				t.Errorf("[%s] app version expect:%v != actual:%v",
					tt.description,
					tt.expect["app-version"],
					actual["app-version"],
				)
			}
		})
	}
}

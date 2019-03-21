// +build integration

package integration

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestHandlersGraphiQL(t *testing.T) {
	ts := newTserver()
	defer ts.closeAll()
	testCases := []struct {
		description  string
		method       string
		addr         string
		args         url.Values
		expectStatus int
	}{
		{
			description:  "testing GET method case",
			method:       http.MethodGet,
			addr:         "/graphiql",
			args:         url.Values{},
			expectStatus: http.StatusOK,
		},
		{
			description:  "testing DELETE method failed case",
			method:       http.MethodDelete,
			addr:         "/graphiql",
			args:         url.Values{},
			expectStatus: http.StatusMethodNotAllowed,
		},
		{
			description:  "testing POST method failed case",
			method:       http.MethodPost,
			addr:         "/graphiql",
			args:         url.Values{},
			expectStatus: http.StatusMethodNotAllowed,
		},
		{
			description:  "testing PUT method failed case",
			method:       http.MethodPut,
			addr:         "/graphiql",
			args:         url.Values{},
			expectStatus: http.StatusMethodNotAllowed,
		},
		{
			description:  "testing patch method failed case",
			method:       http.MethodPatch,
			addr:         "/graphiql",
			args:         url.Values{},
			expectStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, ts.URL+tt.addr+"?"+tt.args.Encode(), nil)
			if err != nil {
				t.Fatalf("[%s] http new request failed:%v", tt.description, err)
			}

			resp, err := ts.Client().Do(req)
			if err != nil {
				t.Fatalf("[%s] http client do failed:%v\n", tt.description, err)
			}

			if tt.expectStatus != resp.StatusCode {
				t.Fatalf("[%s] http status expect:%v != actual:%v", tt.description, tt.expectStatus, resp.StatusCode)
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("[%s] io readall failed:%v", tt.description, err)
			}

			if len(b) == 0 {
				t.Fatalf("[%s] empty response body:%v", tt.description, err)
			}
		})
	}
}

package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	t.Run("", func(t *testing.T) {
		a := assert.New(t)

		// testing most used methods
		for _, method := range []string{http.MethodGet, http.MethodPut, http.MethodDelete} {
			t.Run(method, func(t *testing.T) {
				a.HTTPRedirect(countHandler, method, "/count", nil)
			})
		}
	})

	t.Run("json play", func(t *testing.T) {
		a := assert.New(t)

		var (
			payload = []byte(`
                  {
                        "startDate": "2024-01-01",
                        "endDate": "2024-02-01",
                        "maxCount": 200,
                        "minCount": 50
                  }
                  `)
			data map[string]any
		)

		a.NoError(json.Unmarshal(payload, &data))
		a.Equal("2024-01-01", data["startDate"])
		a.HTTPStatusCode(countHandler, "POST", "/count", nil, http.StatusAccepted)
	})

	t.Run("count docs", func(t *testing.T) {
		a := assert.New(t)

		a.HTTPSuccess(countDocHandler, http.MethodGet, "/", nil)
	})
}

func TestMainHandler(t *testing.T) {
	t.Run("main", func(t *testing.T) {
		a := assert.New(t)

		a.HTTPSuccess(mainHandler, http.MethodGet, "/", nil)
	})
}

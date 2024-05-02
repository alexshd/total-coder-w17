package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// firt endpoint -
//
//	JSON payload with 4 fields
//	- startDate (YYYY-MM-DD)
//	- endDate
//	- minCount
//	- maxCount

func TestCount(t *testing.T) {
	t.Run("", func(t *testing.T) {
		a := assert.New(t)

		for _, method := range []string{http.MethodGet, http.MethodPut, http.MethodDelete} {
			t.Run(method, func(t *testing.T) {
				a.HTTPRedirect(countHandler, method, "/count", nil)
			})
		}
	})

	t.Run("json play", func(t *testing.T) {
		a := assert.New(t)

		payload := []byte(`{"startDate": "2024-01-01", "endDate": "2024-02-01", "maxCount": 200, "minCount": 50}`)
		var dat map[string]any
		a.NoError(json.Unmarshal(payload, &dat))
		a.Equal("2024-01-01", dat["startDate"])
		a.HTTPStatusCode(countHandler, "POST", "/count", nil, http.StatusAccepted)
	})
}

type ServerSuite struct {
	suite.Suite
	str string
}

func (s *ServerSuite) SetupTest() {
	s.str = "string"
}

func (s *ServerSuite) TestServerSuite() {
	s.Equal(s.str, "string", "come and fail")
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

// func string2Time(timeStr string) time.Time {
// 	theTime, err := time.Parse("2006-01-02", timeStr)
// 	if err != nil {
// 		slog.Error("time parse failed", err)
// 	}
//
// 	return theTime
// }

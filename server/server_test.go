package server

import (
	"log/slog"
	"testing"
	"time"

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
type Counters struct {
	startDate string `json:"start_date"`
	endDate   string `json:"end_date"`
	maxCount  int64  `json:"max_count"`
	minCount  int64  `json:"min_count"`
}

func TestGETCount(t *testing.T) {
	t.Run(" /count no payload", func(t *testing.T) {
		a := assert.New(t)

		want := "la"

		a.HTTPBodyContains(countHandler, "GET", "/count", nil, want)
	})

	t.Run("json play", func(t *testing.T) {
		a := assert.New(t)
		counter := Counters{
			startDate: "2024-01-01",
			endDate:   "2024-02-01",
		}

		payload := `{"start_date": "2024-01-01", "end_date": "2024-02-01"}`

		a.JSONEq(`{"main": "line"}`, string(payload))
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

func string2Time(timeStr string) time.Time {
	theTime, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		slog.Error("time parse failed", err)
	}

	return theTime
}

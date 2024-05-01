package server

import (
	"testing"

	"github.com/matryer/is"
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
func TestGETCount(t *testing.T) {
	t.Run(" /count no payload", func(t *testing.T) {
		a := assert.New(t)

		want := "la"

		a.HTTPBodyContains(countHandler, "GET", "/count", nil, want)
	})
}

func TestWithIs(t *testing.T) {
	is := is.New(t)

	is.Equal(1, 1)
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

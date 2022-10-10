package e2e_test

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPingHealthcheck() {
	c := http.Client{}

	r, _ := c.Get("localhost:8083/healthcheck")

	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := ioutil.ReadAll(r.Body)

	s.JSONEq(`{"status": "OK", "messages": []}`, string(b))
}
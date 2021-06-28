package test

import (
	"embed"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Server *httptest.Server
	Data   *embed.FS
}

func (suite *Suite) TearDownSuite() {
	if suite.Server != nil {
		suite.Server.Close()
	}
}

func (suite *Suite) Initialize(router *Router) {
}

func (suite *Suite) Skip() {
	suite.T().SkipNow()
}

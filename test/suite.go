package test

import (
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
	router.inner().HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		suite.Require().Fail("Unexpected path requested")
	})
	suite.Server = httptest.NewServer(router.inner())
}

func (suite *Suite) Skip() {
	suite.T().SkipNow()
}

func (suite *Suite) FileResponseHandler(filename string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, error := suite.Data.ReadFile(filename)
		suite.Require().NoError(error)
		writer.Write(data)
	}
}

func (suite *Suite) Template(patterns ...string) *template.Template {
	template, error := template.ParseFS(suite.Data, patterns...)
	suite.Require().NoError(error)
	return template
}

func (suite *Suite) Render(writer io.Writer, path string, data interface{}) {
	template, error := template.ParseFS(suite.Data, path)
	suite.Require().NoError(error)
	error = template.Execute(writer, data)
	suite.Require().NoError(error)
}

func (suite *Suite) Parameters(request *http.Request) map[string]string {
	var parameters map[string]string
	decoder := json.NewDecoder(request.Body)
	suite.Require().NoError(decoder.Decode(&parameters))
	return parameters
}

package config

import (
	"testing"

	m "github.com/g4s8/go-matchers"
)

func Test_FromFiles(t *testing.T) {
	test1 := &CtlContext{Auth: &AuthToken{Token: "testtoken"}, Endpoint: "https:/central.artipie.com"}
	test2 := &CtlContext{Auth: &AuthBasic{UserName: "test", Password: "testpw"}, Endpoint: "http://artipie.local/artipie"}
	assert := m.Assert(t)
	config := new(ArtiCtlConfig)
	err := config.FromFiles("testdata/config.yaml")
	assert.That("config parsed without error", err, m.Nil())
	assert.That("currentContext parsed", config.CurrentContext, m.Eq("test1"))
	assert.That("test1 context parsed correctly", config.Contexts["test1"], m.Eq(test1))
	assert.That("test2 context parsed correctly", config.Contexts["test2"], m.Eq(test2))
}

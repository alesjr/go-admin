package gf

import (
	"net/http"
	"testing"

	"alesjr/go-admin/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestGf2(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}

package gofiber

import (
	"net/http"
	"testing"

	"alesjr/go-admin/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestGofiber(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}

package chi

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"go-admin/go-admin/tests/common"
)

func TestChi(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}

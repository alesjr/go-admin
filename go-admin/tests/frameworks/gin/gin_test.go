package gin

import (
	"net/http"
	"testing"

	"github.com/alesjr/go-admin/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestGin(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}

package common

import (
	"fmt"
	"net/http"

	"github.com/alesjr/go-admin/go-admin/modules/config"
	"github.com/alesjr/go-admin/go-admin/modules/language"
	"github.com/gavv/httpexpect"
)

func operationLogTest(e *httpexpect.Expect, sesID *http.Cookie) {

	fmt.Println()
	printlnWithColor("Operation Log", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Url("/info/op")).
		WithCookie(sesID.Name, sesID.Value).
		Expect().
		Status(200).
		Body().Contains(language.Get("operation log"))
}

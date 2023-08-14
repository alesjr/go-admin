package paginator

import (
	"testing"

	"alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/plugins/admin/modules/parameter"
	_ "alesjr/go-admin/themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}

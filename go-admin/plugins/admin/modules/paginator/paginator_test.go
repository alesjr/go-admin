package paginator

import (
	"testing"

	"go-admin/go-admin/modules/config"
	"go-admin/go-admin/plugins/admin/modules/parameter"
	_ "go-admin/themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}

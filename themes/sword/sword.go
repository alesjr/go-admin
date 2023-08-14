package sword

import (
	"strings"

	adminTemplate "github.com/alesjr/go-admin/go-admin/template"
	"github.com/alesjr/go-admin/go-admin/template/components"
	"github.com/alesjr/go-admin/go-admin/template/types"
	"github.com/alesjr/go-admin/themes/common"
	"github.com/alesjr/go-admin/themes/sword/resource"

	"github.com/gobuffalo/packr/v2"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: TemplateList,
	},
}

func init() {
	adminTemplate.Add("sword", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	path = strings.Replace(path, "/assets/dist", "", -1)
	box := packr.New("sword", "./resource/assets/dist")
	return box.Find(path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}

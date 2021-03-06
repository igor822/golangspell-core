package template

import (
	"github.com/danilovalente/golangspell-core/appcontext"
	"github.com/danilovalente/golangspell-core/config"
	"github.com/danilovalente/golangspell/gateway/template"
)

//GetRenderer lazy loads a Renderer
func GetRenderer() appcontext.Component {
	return &template.Renderer{}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Renderer, GetRenderer)
}

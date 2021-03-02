package tappx

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewTappxSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("tappx", 628, temp, adapters.SyncTypeIframe)
}

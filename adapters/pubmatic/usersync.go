package pubmatic

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewPubmaticSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("pubmatic", 76, temp, adapters.SyncTypeIframe)
}

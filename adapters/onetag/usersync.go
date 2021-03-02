package onetag

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSyncer(template *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("onetag", 241, template, adapters.SyncTypeIframe)
}

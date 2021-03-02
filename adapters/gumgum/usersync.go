package gumgum

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewGumGumSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("gumgum", 61, temp, adapters.SyncTypeIframe)
}

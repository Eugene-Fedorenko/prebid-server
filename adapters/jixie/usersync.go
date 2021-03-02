package jixie

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewJixieSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("jixie", 0, temp, adapters.SyncTypeRedirect)
}

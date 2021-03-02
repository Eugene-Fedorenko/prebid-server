package aja

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAJASyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("aja", 0, temp, adapters.SyncTypeRedirect)
}

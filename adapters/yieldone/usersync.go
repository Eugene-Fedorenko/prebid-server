package yieldone

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewYieldoneSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("yieldone", 0, temp, adapters.SyncTypeRedirect)
}

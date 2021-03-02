package nanointeractive

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewNanoInteractiveSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("nanointeractive", 72, temp, adapters.SyncTypeRedirect)
}

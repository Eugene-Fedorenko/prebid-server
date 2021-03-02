package acuityads

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAcuityAdsSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("acuityads", 231, temp, adapters.SyncTypeRedirect)
}

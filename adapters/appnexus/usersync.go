package appnexus

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAppnexusSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adnxs", 32, temp, adapters.SyncTypeRedirect)
}

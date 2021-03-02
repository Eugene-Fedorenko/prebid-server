package invibes

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewInvibesSyncer(urlTemplate *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer(
		"invibes",
		436,
		urlTemplate,
		adapters.SyncTypeIframe,
	)
}

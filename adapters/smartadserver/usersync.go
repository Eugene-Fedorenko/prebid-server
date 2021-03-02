package smartadserver

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSmartadserverSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("smartadserver", 45, temp, adapters.SyncTypeRedirect)
}

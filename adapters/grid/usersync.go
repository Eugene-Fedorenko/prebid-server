package grid

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewGridSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("grid", 686, temp, adapters.SyncTypeRedirect)
}

package adtelligent

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAdtelligentSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adtelligent", 0, temp, adapters.SyncTypeRedirect)
}

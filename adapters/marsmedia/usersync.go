package marsmedia

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewMarsmediaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("marsmedia", 0, temp, adapters.SyncTypeRedirect)
}

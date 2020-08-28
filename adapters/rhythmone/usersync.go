package rhythmone

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewRhythmoneSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("rhythmone", 36, temp, adapters.SyncTypeRedirect)
}

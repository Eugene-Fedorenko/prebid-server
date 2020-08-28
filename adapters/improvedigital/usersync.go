package improvedigital

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewImprovedigitalSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("improvedigital", 253, temp, adapters.SyncTypeRedirect)
}

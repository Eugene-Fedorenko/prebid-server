package triplelift

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewTripleliftSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("triplelift", 28, temp, adapters.SyncTypeRedirect)
}

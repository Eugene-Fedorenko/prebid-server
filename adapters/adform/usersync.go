package adform

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAdformSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adform", 50, temp, adapters.SyncTypeRedirect)
}

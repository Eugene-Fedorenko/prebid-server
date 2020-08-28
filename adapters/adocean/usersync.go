package adocean

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewAdOceanSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("adocean", 328, temp, adapters.SyncTypeRedirect)
}

package sovrn

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewSovrnSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("sovrn", 13, temp, adapters.SyncTypeRedirect)
}

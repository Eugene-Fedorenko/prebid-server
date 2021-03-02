package telaria

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewTelariaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("telaria", 202, temp, adapters.SyncTypeRedirect)
}

package visx

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewVisxSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("visx", 154, temp, adapters.SyncTypeRedirect)
}

package mgid

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewMgidSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("mgid", 358, temp, adapters.SyncTypeRedirect)
}

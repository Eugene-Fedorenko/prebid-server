package mediafuse

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewMediafuseSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("mediafuse", 411, temp, adapters.SyncTypeRedirect)
}

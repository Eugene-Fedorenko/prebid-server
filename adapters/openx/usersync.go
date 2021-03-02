package openx

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewOpenxSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("openx", 69, temp, adapters.SyncTypeRedirect)
}

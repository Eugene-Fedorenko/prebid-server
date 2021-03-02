package dmx

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewDmxSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("dmx", 144, temp, adapters.SyncTypeRedirect)
}
